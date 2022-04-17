package service

import (
	// "fmt"
	model "messageBoard/gin/Models"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id           int    `json:"id" gorm:"primary_key;auto_increase'"`
	Account      string `json:"account"`
	Password     string `json:"password"`
	IsSuspension bool   `json:"is_suspension"`
	IsSuperuser  bool   `json:"is_superuser"`
}

func PasswordHash(pwd string) (pwdHash string, err error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	pwdHash = string(hash)
	return pwdHash, err
}

func PasswordVerify(pwd string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(pwd), []byte(hash))
	return err == nil
}

func (this *User) CreateUser() (account string, err error) {
	var userModel model.User
	userModel.Account = this.Account
	userModel.Password, _ = PasswordHash(this.Password)
	account, err = userModel.CreateUser()
	// fmt.Printf("after create%v, %s	\n", userModel, err)
	return account, err
}

func (this *User) SelectUser() (*model.User, error) {
	var userModel model.User
	userModel.Account = this.Account
	userModel.Password = this.Password
	// account, password, isSuspension, err = userModel.Select()
	user, err := userModel.SelectUser()
	// fmt.Printf("after create%v, %s	\n", userModel, err)
	return user, err
}

func (this *User) UpdateUserIsSuspension() (err error) {
	var userModel model.User
	userModel.Id = this.Id
	userModel.IsSuspension = this.IsSuspension
	err = userModel.UpdateUserIsSuspension()
	return err
}
