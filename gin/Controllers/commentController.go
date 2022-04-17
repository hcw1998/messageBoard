package controller

import (
	"fmt"
	service "messageBoard/gin/Services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CommentCreate(c *gin.Context) {
	var commentService service.Comment

	//parse json, validate the post json data
	err := c.ShouldBindJSON(&commentService)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, _ := c.Get("id")
	id2 := id.(uint)
	commentService.UserId = int(id2)
	err = commentService.CreateComment()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Create comment failed",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Create comment successfully",
	})
}

func CommentSelect(c *gin.Context) {
	var commentService service.Comment

	// err := c.ShouldBindJSON(&userService)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }
	comment, err := commentService.SelectComment()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}
	c.JSON(http.StatusOK, comment)
}

func CommentSelectWithLike(c *gin.Context) {
	var commentService service.Comment
	// err := c.ShouldBindJSON(&commentService)
	search := fmt.Sprintf("%%%s%%", c.Query("search"))
	fmt.Printf("%s", search)
	comment, err := commentService.SelectCommentWithLike(search)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}
	c.JSON(http.StatusOK, comment)
}

func CommentUpdate(c *gin.Context) {
	var commentService service.Comment
	err := c.ShouldBindJSON(&commentService)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = commentService.UpdateComment()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Update comment failed",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Update comment successfully",
	})
}
