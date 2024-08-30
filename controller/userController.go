package controller

import (
	"friends_ranking/models"

	"github.com/gin-gonic/gin"
)

type UserController struct {
}

func (u UserController) GetUserInfo(c *gin.Context) {
	//id := c.GetString("id")
	userInfo := models.GetUserInfoById(12)
	ReturnSuccess(c, 200, "成功", userInfo, 1)
}

func (u UserController) Query(c *gin.Context) {
	//id := c.GetString("id")
	//userInfo := models.GetUserInfoById(12)
	ReturnError(c, "查询有误")
}
