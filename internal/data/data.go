package data

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"github.com/go-redis/redis/extra/redisotel"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	// init mysql driver
	_ "github.com/go-sql-driver/mysql"

	"github.com/zag07/gin-example/internal/conf"
	"github.com/zag07/gin-example/internal/data/ent"
	"github.com/zag07/gin-example/internal/data/ent/migrate"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewBlogRepo)

// Data .
type Data struct {
	db  *ent.Client
	rdb *redis.Client
}

func NewData(conf *conf.Data, log *zap.Logger) (*Data, func(), error) {
	drv, err := sql.Open(
		conf.Database.Driver,
		conf.Database.Source,
	)
	if err != nil {
		log.Sugar().Errorf("failed opening connection to sqlite: %v", err)
		return nil, nil, err
	}
	sqlDrv := dialect.DebugWithContext(drv, func(ctx context.Context, i ...interface{}) {
		tracer := otel.Tracer("ent.")
		kind := trace.SpanKindServer
		_, span := tracer.Start(ctx,
			"Query",
			trace.WithAttributes(
				attribute.String("sql", fmt.Sprint(i...)),
			),
			trace.WithSpanKind(kind),
		)
		span.End()
	})
	client := ent.NewClient(ent.Driver(sqlDrv))
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background(), migrate.WithDropColumn(true)); err != nil {
		log.Sugar().Error("failed creating schema resources: %v", err)
		return nil, nil, err
	}
	rdb := redis.NewClient(&redis.Options{
		Addr:         conf.Redis.Addr,
		Password:     conf.Redis.Password,
		DB:           int(conf.Redis.Db),
		DialTimeout:  conf.Redis.DialTimeout.AsDuration(),
		WriteTimeout: conf.Redis.WriteTimeout.AsDuration(),
		ReadTimeout:  conf.Redis.ReadTimeout.AsDuration(),
	})
	rdb.AddHook(redisotel.TracingHook{})
	d := &Data{
		db:  client,
		rdb: rdb,
	}

	return d, func() {
		log.Info("message closing the data resources")
		if err := d.db.Close(); err != nil {
			log.Sugar().Errorf("err:%v", err)
		}
		if err := d.rdb.Close(); err != nil {
			log.Sugar().Errorf("err:%v", err)
		}
	}, nil
}
