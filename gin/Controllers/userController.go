package controller

import (
	helper "messageBoard/gin/Helper"
	service "messageBoard/gin/Services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserCreate(c *gin.Context) {
	var userService service.User

	//parse json, validate the post json data
	err := c.ShouldBindJSON(&userService)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	account, err := userService.CreateUser()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Insert user failed!",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "sign up successfully",
		"data":    account,
	})
}

func UserLogin(c *gin.Context) {
	var userService service.User

	err := c.ShouldBindJSON(&userService)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := userService.SelectUser()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "SelectUser failed",
		})
		return
	}
	match := service.PasswordVerify(user.Password, userService.Password) // (hash,no hash)
	token, _ := helper.GenerateToken(user)
	if !match {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "login failed",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "login successfully",
			"token":   token,
		})
	}
}

func UserIsSuspensionUpdate(c *gin.Context) {
	var userService service.User
	err := c.ShouldBindJSON(&userService)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = userService.UpdateUserIsSuspension()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Update user's is_suspension failed",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Update user's is_suspension successfully",
	})
}
