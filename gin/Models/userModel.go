package model

import (
	database "messageBoard/gin/Databases"
)

type User struct {
	Id           int    `json:"id" gorm:"primary_key;auto_increase'"`
	Account      string `json:"account"`
	Password     string `json:"password"`
	IsSuspension bool   `json:"is_suspension"`
	IsSuperuser  bool   `json:"is_superuser"`
}

func (user *User) TableName() string {
	return "user"
}

// User signs up a new account.
func (this *User) CreateUser() (account string, err error) {
	result := database.DB.Create(&this)
	if result.Error != nil {
		return this.Account, result.Error
	}
	return this.Account, result.Error
}

// Select a user's infp.
func (this *User) SelectUser() (*User, error) {
	result := database.DB.Where("account = ?", this.Account).First(&this)
	if result.Error != nil {
		return this, result.Error
	}
	return this, result.Error
}

// Superuser sets the user's is_suspension.
func (this *User) UpdateUserIsSuspension() (err error) {
	result := database.DB.Model(&this).Where("id = ?", this.Id).Update("is_suspension", this.IsSuspension)
	if result.Error != nil {
		return result.Error
	}
	return result.Error
}
