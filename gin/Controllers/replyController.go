package controller

import (
	service "messageBoard/gin/Services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ReplyCreate(c *gin.Context) {
	var replyService service.Reply

	//parse json, validate the post json data
	err := c.ShouldBindJSON(&replyService)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, _ := c.Get("id")
	id2 := id.(uint)
	replyService.UserId = int(id2)
	replyService.CommentId, _ = strconv.Atoi(c.Param("comment_id"))
	err = replyService.CreateReply()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Create reply failed",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Create reply successfully",
	})
}
