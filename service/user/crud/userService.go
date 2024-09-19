package userService

import (
	"friends_ranking/models"
	"friends_ranking/utils/response"

	"github.com/gin-gonic/gin"
)

func Register(user *models.User, c *gin.Context) {
	//先检查注册的账号是否存在

	userModels := models.CreateUserFactory("")
	account := user.Account
	ruser, err := userModels.SelectByAccount(account)
	if err != nil {
		response.ReturnFail(c, "根据账号查询出错")
		return
	} else {
		if ruser.Id == 0 {
			isInsert := userModels.InsertUser(user)
			if !isInsert {
				response.ReturnFail(c, "新增用户出错")
				return
			}
		} else {
			response.ReturnFail(c, "账户已存在，请输入新账号")
			return
		}
	}
	response.ReturnDefaultOk(c)

}
