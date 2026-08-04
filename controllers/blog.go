package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"api/repositories"
)

func (c *Controller) GetBlogsByAuthor(ctx *gin.Context) {
	id := ctx.Param("id")

	authorUUID, err := uuid.Parse(id)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": "Invalid author id",
		})
		return
	}

	author := repositories.ReadAuthor(authorUUID)
	if author.Id == uuid.Nil {
		ctx.JSON(404, gin.H{
			"error": "Author not found",
		})
		return
	}

	blogs := repositories.ReadBlogsByAuthor(authorUUID)
	ctx.JSON(200, gin.H{
		"blogs": blogs,
	})
}
func (c *Controller) GetBlogById(ctx *gin.Context) {
	id := ctx.Param("id")

	blogUUID, err := uuid.Parse(id)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": "Invalid blog id",
		})
		return
	}

	blog := repositories.ReadBlog(blogUUID)
	if blog.Id == uuid.Nil {
		ctx.JSON(404, gin.H{
			"error": "Blog not found",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"blog": blog,
	})
}
func (c *Controller) CreateBlog(ctx *gin.Context) {
	var request struct {
		Name      string `json:"name"`
		Content   string `json:"content"`
		Author_id string `json:"author_id"`
	}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(400, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	authorUUID, err := uuid.Parse(request.Author_id)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": "Invalid author id",
		})
		return
	}

	author := repositories.ReadAuthor(authorUUID)
	if author.Id == uuid.Nil {
		ctx.JSON(404, gin.H{
			"error": "Author not found",
		})
		return
	}

	newBlog := repositories.CreateBlog(request.Name, request.Content, authorUUID)

	ctx.JSON(201, gin.H{
		"blog": newBlog,
	})
}
func (c *Controller) GetAllBlogs(ctx *gin.Context) {
	blogs := repositories.ReadAllBlogs()
	ctx.JSON(200, gin.H{
		"blogs": blogs,
	})
}
func (c *Controller) GetAllAuthors(ctx *gin.Context) {
	authors := repositories.ReadAllAuthors()
	ctx.JSON(200, gin.H{
		"authors": authors,
	})
}
func (c *Controller) CreateAuthor(ctx *gin.Context) {
	var request struct {
		Name string `json:"name"`
	}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(400, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	newAuthor := repositories.CreateAuthor(request.Name)

	ctx.JSON(201, gin.H{
		"author": newAuthor,
	})
}
func (c *Controller) GetAuthorById(ctx *gin.Context) {
	id := ctx.Param("id")

	authorUUID, err := uuid.Parse(id)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": "Invalid author id",
		})
		return
	}

	author := repositories.ReadAuthor(authorUUID)
	if author.Id == uuid.Nil {
		ctx.JSON(404, gin.H{
			"error": "Author not found",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"author": author,
	})
}
func (c *Controller) UpdateBlog(ctx *gin.Context) {
	id := ctx.Param("id")

	blogUUID, err := uuid.Parse(id)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": "Invalid blog id",
		})
		return
	}

	var request struct {
		Name    string `json:"name"`
		Content string `json:"content"`
	}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(400, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	updatedBlog := repositories.UpdateBlog(blogUUID, request.Name, request.Content)
	if updatedBlog.Id == uuid.Nil {
		ctx.JSON(404, gin.H{
			"error": "Blog not found",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"blog": updatedBlog,
	})
}
func (c *Controller) DeleteBlog(ctx *gin.Context) {
	id := ctx.Param("id")

	blogUUID, err := uuid.Parse(id)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": "Invalid blog id",
		})
		return
	}

	deleted := repositories.DeleteBlog(blogUUID)

	if !deleted {
		ctx.JSON(404, gin.H{
			"error": "Blog not found",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "Blog deleted successfully",
	})
}
func (c *Controller) DeleteAllBlogs(ctx *gin.Context) {
	repositories.FlushBlogs()
	ctx.JSON(200, gin.H{
		"message": "All blogs deleted",
	})
}
func (c *Controller) DeleteAuthor(ctx *gin.Context) {
	id := ctx.Param("id")

	authorUUID, err := uuid.Parse(id)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": "Invalid author id",
		})
		return
	}

	deleted := repositories.DeleteAuthor(authorUUID)

	if !deleted {
		ctx.JSON(404, gin.H{
			"error": "Author not found",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "Author deleted successfully",
	})
}
func (c *Controller) DeleteAllAuthors(ctx *gin.Context) {
	repositories.FlushAuthors()
	ctx.JSON(200, gin.H{
		"message": "All authors deleted",
	})
}
