package controller

import (
	"friends_ranking/config/globalConst"
	"friends_ranking/models"
	userService "friends_ranking/service/user/crud"
	"friends_ranking/utils/response"
	"time"

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
		return
	} else {
		if user.Id == 0 {
			response.ReturnFail(c, "登陆用户不存在，请注册新用户")
			return
		} else {
			loginPass := c.GetString(globalConst.ValidatorPrefix + "pass1")
			if user.Password != loginPass {
				response.ReturnFail(c, "密码错误，请重新输入")
				return
			}

		}
	}
	response.ReturnSuccess(c, token, 1)

}

func (index IndexController) Register(c *gin.Context) {
	userName := c.GetString(globalConst.ValidatorPrefix + "userName")
	nickName := c.GetString(globalConst.ValidatorPrefix + "nickName")
	account := c.GetString(globalConst.ValidatorPrefix + "account")
	idCard := c.GetString(globalConst.ValidatorPrefix + "idCard")
	password := c.GetString(globalConst.ValidatorPrefix + "password")
	email := c.GetString(globalConst.ValidatorPrefix + "email")
	age := c.GetInt(globalConst.ValidatorPrefix + "age")
	phone := c.GetString(globalConst.ValidatorPrefix + "phone")
	job := c.GetString(globalConst.ValidatorPrefix + "job")
	birthDay := c.GetString(globalConst.ValidatorPrefix + "birthDay")
	level := 0
	status := 0
	now := time.Now()
	nowTime := now.Format(globalConst.DateFormat)
	registerTime := nowTime
	creator := globalConst.SysAccount
	editTor := creator
	isDelete := 0
	userDetail := models.User{
		UserName:     userName,
		NickName:     nickName,
		Account:      account,
		IdCard:       idCard,
		Password:     password,
		Email:        email,
		Age:          age,
		Phone:        phone,
		Job:          job,
		BirthDay:     birthDay,
		Level:        level,
		Status:       status,
		RegisterTime: registerTime,

		Creator:  creator,
		Editor:   editTor,
		IsDelete: isDelete,
	}
	userService.Register(userDetail, c)
}
