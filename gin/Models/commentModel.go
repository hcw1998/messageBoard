package model

import (
	"fmt"
	database "messageBoard/gin/Databases"
	"time"
)

type Comment struct {
	Id       int       `json:"id" gorm:"primary_key;auto_increase'"`
	Content  string    `json:"content"`
	Date     time.Time `gorm:"type:timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP" json:"date"`
	IsPublic bool      `json:"is_public"`
	IsReply  bool      `json:"is_reply"`
	UserId   int       `json:"user_id"`
	Reply    []Reply
}

func (comment *Comment) TableName() string {
	return "comment"
}

//create a new user
func (this *Comment) CreateComment() (err error) {
	result := database.DB.Create(&this)
	if result.Error != nil {
		return result.Error
	}
	return result.Error
}

// select all public comments
func (this *Comment) SelectComment() (*[]Comment, error) {
	var comments *[]Comment
	// result := database.DB.Where("is_public = ?", true).Find(&comments)
	result := database.DB.Preload("Reply").Find(&comments)

	if result.Error != nil {
		fmt.Printf("select error:%#v	\n", result.Error)
		return comments, result.Error
	}
	fmt.Printf("...%#v	\n", comments)
	return comments, result.Error
}

func (this *Comment) SelectCommentWithLike(search string) (*[]Comment, error) {
	var comments *[]Comment
	result := database.DB.Where("content LIKE ?", search).Find(&comments)
	if result.Error != nil {
		fmt.Printf("select error:%#v	\n", result.Error)
		return comments, result.Error
	}
	return comments, result.Error
}

func (this *Comment) UpdateComment() error {
	result := database.DB.Model(&this).Updates(map[string]interface{}{"is_public": this.IsPublic, "is_reply": this.IsReply})
	if result.Error != nil {
		return result.Error
	}
	return result.Error
}
