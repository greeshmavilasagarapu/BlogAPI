package routes

import (
	"api/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterBlogRoutes(router *gin.Engine, controller *controllers.Controller) {
	blogRoutes := router.Group("/blogs")
	{
		blogRoutes.GET("/author/:id", controller.GetBlogsByAuthor) 
		blogRoutes.POST("/", controller.CreateBlog)
	}
}
