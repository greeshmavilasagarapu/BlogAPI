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
		blogRoutes.GET("/:id", controller.GetBlogById)
		blogRoutes.GET("/", controller.GetAllBlogs)
		blogRoutes.GET("/authors", controller.GetAllAuthors)
		blogRoutes.POST("/authors", controller.CreateAuthor)
		blogRoutes.GET("/authors/:id", controller.GetAuthorById)
		blogRoutes.PUT("/:id", controller.UpdateBlog)
		blogRoutes.DELETE("/:id", controller.DeleteBlog)
		blogRoutes.DELETE("/", controller.DeleteAllBlogs)
		blogRoutes.DELETE("/authors/:id", controller.DeleteAuthor)
		blogRoutes.DELETE("/authors", controller.DeleteAllAuthors)
	}

}
