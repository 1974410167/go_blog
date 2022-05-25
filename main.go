package main

import (
	"go_blog/middleware/jwt"
	"go_blog/models"
	"go_blog/views"

	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()
	models.InitDB()

	server.POST("/register", views.Register)
	server.POST("/login", views.Login)

	apiv1 := server.Group("/api/v1/")
	apiv1.Use(jwt.JWT())
	{
		postRouter := apiv1.Group("/posts/")
		{
			postViews := views.PostViews{}
			postRouter.GET(":id", postViews.Retrieve)
			postRouter.POST("", postViews.Create)
		}
		tagRouter := apiv1.Group("/tags/")
		{
			tagViews := views.TagViews{}
			tagRouter.POST("", tagViews.Create)
			tagRouter.GET("", tagViews.List)

		}
	}
	server.Run("localhost:8000")

}
