package routers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/huangchao308/blog/docs"
	"github.com/huangchao308/blog/internal/middleware"
	v1 "github.com/huangchao308/blog/internal/routers/api/v1"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery(), middleware.Translations())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	post := v1.NewPost()
	tag := v1.NewTag()

	apiV1 := r.Group("/api/v1")
	{
		apiV1.POST("/posts", post.Create)
		apiV1.GET("/posts", post.List)
		apiV1.DELETE("/posts/:id", post.Delete)
		apiV1.PUT("/posts/:id", post.Update)
		apiV1.PATCH("/posts/:id/status", post.Update)

		apiV1.GET("/tag/:id", tag.Get)
		apiV1.POST("/tags", tag.Create)
		apiV1.GET("/tags", tag.List)
		apiV1.DELETE("/tags/:id", tag.Delete)
		apiV1.PUT("/tags/:id", tag.Update)
		apiV1.PATCH("/tags/:id/status", tag.Update)
	}

	return r
}
