package service

import (
	// "fmt"

	model "messageBoard/gin/Models"
	"time"
)

type Comment struct {
	Id       int       `json:"id" gorm:"primary_key;auto_increase'"`
	Content  string    `json:"content"`
	Date     time.Time `gorm:"type:timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP" json:"date"`
	IsPublic bool      `json:"is_public"`
	IsReply  bool      `json:"is_reply"`
	UserId   int       `json:"user_id"`
}

func (this *Comment) CreateComment() (err error) {
	var userModel model.Comment
	userModel.Content = this.Content
	userModel.Date = time.Now()
	userModel.IsPublic = this.IsPublic
	userModel.IsReply = this.IsReply
	userModel.UserId = this.UserId
	err = userModel.CreateComment()
	return err
}

func (this *Comment) SelectComment() (*[]model.Comment, error) {
	var commentModel model.Comment
	comment, err := commentModel.SelectComment()
	return comment, err
}

func (this *Comment) SelectCommentWithLike(search string) (*[]model.Comment, error) {
	var commentModel model.Comment
	comment, err := commentModel.SelectCommentWithLike(search)
	return comment, err
}

func (this *Comment) UpdateComment() (err error) {
	var commentModel model.Comment
	commentModel.Id = this.Id
	commentModel.IsPublic = this.IsPublic
	commentModel.IsReply = this.IsReply
	err = commentModel.UpdateComment()
	return err
}
