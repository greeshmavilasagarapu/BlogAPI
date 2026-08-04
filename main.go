package main

import (
	"api/controllers"
	"api/routes"
	"net/http"

	"github.com/gin-gonic/gin"
)

// var (
// 	authors []Author
// 	blogs   []Blog
// )

func main() {
	router := gin.Default()

	router.LoadHTMLFS(http.Dir("."), "index.html")
	router.LoadHTMLGlob("templates/*")

	routes.RegisterBlogRoutes(router, &controllers.Controller{})

	router.Run() // listens on 0.0.0.0:8080 by default
}
