// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/zag07/gin-example/internal/data/ent/article"
)

// Article is the model entity for the Article schema.
type Article struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// Desc holds the value of the "desc" field.
	Desc string `json:"desc,omitempty"`
	// CoverImageURL holds the value of the "cover_image_url" field.
	CoverImageURL string `json:"cover_image_url,omitempty"`
	// Content holds the value of the "content" field.
	Content string `json:"content,omitempty"`
	// Status holds the value of the "status" field.
	// 1：正常使用 2：删除
	Status int8 `json:"status,omitempty"`
	// CreatedBy holds the value of the "created_by" field.
	CreatedBy string `json:"created_by,omitempty"`
	// UpdatedBy holds the value of the "updated_by" field.
	UpdatedBy string `json:"updated_by,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Article) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case article.FieldID, article.FieldStatus:
			values[i] = new(sql.NullInt64)
		case article.FieldTitle, article.FieldDesc, article.FieldCoverImageURL, article.FieldContent, article.FieldCreatedBy, article.FieldUpdatedBy:
			values[i] = new(sql.NullString)
		case article.FieldCreatedAt, article.FieldUpdatedAt, article.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Article", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Article fields.
func (a *Article) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case article.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			a.ID = int64(value.Int64)
		case article.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				a.Title = value.String
			}
		case article.FieldDesc:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field desc", values[i])
			} else if value.Valid {
				a.Desc = value.String
			}
		case article.FieldCoverImageURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field cover_image_url", values[i])
			} else if value.Valid {
				a.CoverImageURL = value.String
			}
		case article.FieldContent:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field content", values[i])
			} else if value.Valid {
				a.Content = value.String
			}
		case article.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				a.Status = int8(value.Int64)
			}
		case article.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				a.CreatedBy = value.String
			}
		case article.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				a.UpdatedBy = value.String
			}
		case article.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				a.CreatedAt = value.Time
			}
		case article.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				a.UpdatedAt = value.Time
			}
		case article.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				a.DeletedAt = new(time.Time)
				*a.DeletedAt = value.Time
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Article.
// Note that you need to call Article.Unwrap() before calling this method if this Article
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Article) Update() *ArticleUpdateOne {
	return (&ArticleClient{config: a.config}).UpdateOne(a)
}

// Unwrap unwraps the Article entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *Article) Unwrap() *Article {
	tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Article is not a transactional entity")
	}
	a.config.driver = tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Article) String() string {
	var builder strings.Builder
	builder.WriteString("Article(")
	builder.WriteString(fmt.Sprintf("id=%v", a.ID))
	builder.WriteString(", title=")
	builder.WriteString(a.Title)
	builder.WriteString(", desc=")
	builder.WriteString(a.Desc)
	builder.WriteString(", cover_image_url=")
	builder.WriteString(a.CoverImageURL)
	builder.WriteString(", content=")
	builder.WriteString(a.Content)
	builder.WriteString(", status=")
	builder.WriteString(fmt.Sprintf("%v", a.Status))
	builder.WriteString(", created_by=")
	builder.WriteString(a.CreatedBy)
	builder.WriteString(", updated_by=")
	builder.WriteString(a.UpdatedBy)
	builder.WriteString(", created_at=")
	builder.WriteString(a.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(a.UpdatedAt.Format(time.ANSIC))
	if v := a.DeletedAt; v != nil {
		builder.WriteString(", deleted_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteByte(')')
	return builder.String()
}

// Articles is a parsable slice of Article.
type Articles []*Article

func (a Articles) config(cfg config) {
	for _i := range a {
		a[_i].config = cfg
	}
}
