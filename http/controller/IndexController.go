package controller

import (
	"friends_ranking/config/globalConst"
	"friends_ranking/models"
	"friends_ranking/utils/response"

	"github.com/gin-gonic/gin"
)

type IndexController struct {
}

func (index IndexController) Login(c *gin.Context) {
	account := c.GetString(globalConst.ValidatorPrefix + "account")
	userModels := models.CreateUserFactory("")
	user, err := userModels.SelectByAccount(account)
	token := "111"
	if err != nil {
		response.ReturnFail(c, "根据账号查询出错")
	} else {
		if user.Id == 0 {
			response.ReturnFail(c, "登陆用户不存在，请重新注册")
		} else {
			loginPass := c.GetString(globalConst.ValidatorPrefix + "pass1")
			if user.Password != loginPass {
				response.ReturnFail(c, "密码错误，请重新输入")
			}

		}
	}
	response.ReturnSuccess(c, token, 1)

}

func (index IndexController) Register(c *gin.Context) {

}
