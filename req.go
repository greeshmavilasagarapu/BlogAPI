package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var (
	authors []Author
	blogs   []Blog
)

func main() {
	router := gin.Default()

	router.LoadHTMLFS(http.Dir("."), "index.html")
	router.LoadHTMLGlob("templates/*")

	router.GET("/blogs", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"blogs": readallblogs(),
		})
	})

	router.GET("/authors", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"authors": readallauthors(),
		})
	})

	router.POST("/author", func(ctx *gin.Context) {
		name := ctx.PostForm("name")
		author := createAuthor(name)
		ctx.JSON(200, gin.H{
			"author": author,
		})
	})

	router.POST("/blog", func(ctx *gin.Context) {
		name := ctx.PostForm("name")
		content := ctx.PostForm("content")
		author_id := ctx.PostForm("author_id")
		author_uuid, err := uuid.Parse(author_id)
		if err != nil {
			ctx.JSON(400, gin.H{
				"error": "Invalid author_id",
			})
			return
		}
		blog := createBlog(name, content, author_uuid)
		ctx.JSON(200, gin.H{
			"blog": blog,
		})
	})

	router.GET("/author/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		author_uuid, err := uuid.Parse(id)
		if err != nil {
			ctx.JSON(400, gin.H{
				"error": "Invalid author id",
			})
			return
		}
		author := readauthor(author_uuid)
		if author.Id == uuid.Nil {
			ctx.JSON(404, gin.H{
				"error": "Author not found",
			})
			return
		}
		ctx.JSON(200, gin.H{
			"author": author,
		})
	})

	router.GET("/blog/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		blog_uuid, err := uuid.Parse(id)
		if err != nil {
			ctx.JSON(400, gin.H{
				"error": "Invalid blog id",
			})
			return
		}
		blog := readblog(blog_uuid)
		if blog.Id == uuid.Nil {
			ctx.JSON(404, gin.H{
				"error": "Blog not found",
			})
			return
		}
		ctx.JSON(200, gin.H{
			"blog": blog,
		})
	})
	router.GET("/author/:id/blogs", func(ctx *gin.Context) {
		id := ctx.Param("id")

		authorUUID, err := uuid.Parse(id)
		if err != nil {
			ctx.JSON(400, gin.H{
				"error": "Invalid author id",
			})
			return
		}

		author := readauthor(authorUUID)
		if author.Id == uuid.Nil {
			ctx.JSON(404, gin.H{
				"error": "Author not found",
			})
			return
		}

		blogs := readBlogsByAuthor(authorUUID)

		ctx.JSON(200, gin.H{
			"author": author,
			"blogs":  blogs,
		})
	})
	// PUT
	router.PUT("/blog/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		blog_uuid, err := uuid.Parse(id)
		if err != nil {
			ctx.JSON(400, gin.H{
				"error": "Invalid blog id",
			})
			return
		}
		name := ctx.PostForm("name")
		content := ctx.PostForm("content")
		for i, blog := range blogs {
			if blog.Id == blog_uuid {
				blogs[i].Name = name
				blogs[i].Content = content
				ctx.JSON(200, gin.H{
					"blog": blogs[i],
				})
				return
			}
		}
		ctx.JSON(404, gin.H{
			"error": "Blog not found",
		})
	})

	router.PUT("/author/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		author_uuid, err := uuid.Parse(id)
		if err != nil {
			ctx.JSON(400, gin.H{
				"error": "Invalid author id",
			})
			return
		}
		name := ctx.PostForm("name")
		for i, author := range authors {
			if author.Id == author_uuid {
				authors[i].Name = name
				ctx.JSON(200, gin.H{
					"author": authors[i],
				})
				return
			}
		}
		ctx.JSON(404, gin.H{
			"error": "Author not found",
		})
	})

	router.DELETE("/blog/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		blog_uuid, err := uuid.Parse(id)
		if err != nil {
			ctx.JSON(400, gin.H{
				"error": "Invalid blog id",
			})
			return
		}
		for i, blog := range blogs {
			if blog.Id == blog_uuid {
				blogs = append(blogs[:i], blogs[i+1:]...)
				ctx.JSON(200, gin.H{
					"message": "Blog deleted",
				})
				return
			}
		}
		ctx.JSON(404, gin.H{
			"error": "Blog not found",
		})
	})

	router.DELETE("/author/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		author_uuid, err := uuid.Parse(id)
		if err != nil {
			ctx.JSON(400, gin.H{
				"error": "Invalid author id",
			})
			return
		}
		for i, author := range authors {
			if author.Id == author_uuid {
				authors = append(authors[:i], authors[i+1:]...)
				ctx.JSON(200, gin.H{
					"message": "Author deleted",
				})
				return
			}
		}
		ctx.JSON(404, gin.H{
			"error": "Author not found",
		})
	})

	router.DELETE("/blogs", func(ctx *gin.Context) {
		flushBlogs()
		ctx.JSON(200, gin.H{
			"message": "All blogs deleted",
		})
	})

	router.DELETE("/authors", func(ctx *gin.Context) {
		flushAuthors()
		ctx.JSON(200, gin.H{
			"message": "All authors deleted",
		})
	})

	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "blogs.html", gin.H{
			"title": "Blog API",
			"Blogs": readallblogs(),
		})
	})
	router.Run() // listens on 0.0.0.0:8080 by default
}
