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
