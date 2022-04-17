package service

import (
	// "fmt"
	model "messageBoard/gin/Models"
	"time"
)

type Reply struct {
	Id        int       `json:"id" gorm:"primary_key;auto_increase'"`
	Content   string    `json:"content"`
	Date      time.Time `gorm:"type:timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP" json:"date"`
	CommentId int       `json:"comment_id"`
	UserId    int       `json:"user_id"`
}

func (this *Reply) CreateReply() (err error) {
	var replyModel model.Reply
	replyModel.Content = this.Content
	replyModel.Date = time.Now()
	replyModel.CommentId = this.CommentId
	replyModel.UserId = this.UserId
	err = replyModel.CreateReply()
	// fmt.Printf("after create%v, %s	\n", userModel, err)
	return err
}
