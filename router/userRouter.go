package router

import (
	"friends_ranking/controller"

	"github.com/gin-gonic/gin"
)

func UserRouter() *gin.Engine {
	r := gin.Default()

	user := r.Group("/user")
	{
		user.GET("/getUserInfo", controller.UserController{}.GetUserInfo)

		user.POST("/query", controller.UserController{}.Query)
	}
	return r

}
