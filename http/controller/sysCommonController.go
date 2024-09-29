package controller

import (
	"fmt"
	"friends_ranking/config/globalConst"
	sysCommonService "friends_ranking/service/sys_common"
	"friends_ranking/utils/response"

	"github.com/gin-gonic/gin"
)

type SysCommonController struct {
}

// 生成验证码
func (u SysCommonController) GetVerfyCode(c *gin.Context) {
	key := c.GetString(globalConst.ValidatorPrefix + "key")
	fmt.Println("key --> ", key)
	code := sysCommonService.GetVerfyCode(key)
	if code != "" {
		response.ReturnSuccess(c, code, 1)
	}
}
