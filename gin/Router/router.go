package router

import (
	controller "messageBoard/gin/Controllers"
	helper "messageBoard/gin/Helper"

	middleware "messageBoard/gin/Middlewares"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	router := gin.Default()
	// router.GET("/", test)
	router.Use(middleware.Cors())
	v1 := router.Group("user")
	{
		v1.POST("/signup", controller.UserCreate)
		v1.POST("/authentication", controller.UserLogin)
		v1.POST("/comment", helper.VerifyToken, controller.CommentCreate)
		v1.POST("/comment/:comment_id/reply", helper.VerifyToken, controller.ReplyCreate)
		v1.GET("/comment", controller.CommentSelect)
	}
	v2 := router.Group("superuser")
	v2.Use(helper.VerifyToken, helper.VerifySuperuser)
	{
		v2.PUT("/comment", controller.CommentUpdate)
		v2.PUT("/user", controller.UserIsSuspensionUpdate)
		v2.GET("/comment", controller.CommentSelectWithLike)

	}

	router.Run(":3000")
}
