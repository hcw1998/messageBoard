package model

import (
	database "messageBoard/gin/Databases"
	"time"
)

type Reply struct {
	Id        int       `json:"id" gorm:"primary_key;auto_increase'"`
	Content   string    `json:"content"`
	Date      time.Time `gorm:"type:timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP" json:"date"`
	CommentId int       `json:"comment_id"`
	UserId    int       `json:"user_id"`
}

func (reply *Reply) TableName() string {
	return "reply"
}

func (this *Reply) CreateReply() (err error) {
	result := database.DB.Create(&this)
	if result.Error != nil {
		return result.Error
	}
	return result.Error
}
