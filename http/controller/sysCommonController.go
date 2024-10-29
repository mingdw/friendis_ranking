package controller

import (
	"fmt"
	"lottery_annual/config/globalConst"
	sysCommonService "lottery_annual/service/sys_common"
	"lottery_annual/utils/response"

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
