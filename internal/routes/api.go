package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zs368/gin-example/configs"
	"github.com/zs368/gin-example/internal/app/controllers"
	"github.com/zs368/gin-example/internal/app/controllers/common_ctl"
)

func SetApiRouter(r *gin.Engine) {
	c := r.Group("/c")
	{
		upload := common_ctl.NewUpload()
		c.POST("/upload/file", upload.UploadFile)
		c.StaticFS("/static", http.Dir(configs.App.UploadSavePath))
	}

	apiV1 := r.Group("/api/v1")
	{
		user := controllers.NewUser()
		apiV1.GET("/user/:id", user.Get)

		article := controllers.NewArticle()
		apiV1.GET("/article/:id", article.Get)
		apiV1.POST("/article", article.Create)
		apiV1.PUT("/article/:id", article.Update)
		apiV1.DELETE("/article/:id", article.Delete)

		tag := controllers.NewTag()
		apiV1.GET("/tag/:id", tag.Get)
		apiV1.POST("/tag", tag.Create)
		apiV1.PUT("/tag/:id", tag.Update)
		apiV1.DELETE("/tag/:id", tag.Delete)
	}
}
