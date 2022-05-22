package main

import (
	"go_blog/models"
	"go_blog/views"

	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()
	models.InitDB()

	apiv1 := server.Group("/api/v1")
	{
		postRouter := apiv1.Group("/posts")
		postViews := views.PostViews{}
		postRouter.GET("/:id", postViews.Retrieve)
	}
	server.Run("localhost:8000")

}
