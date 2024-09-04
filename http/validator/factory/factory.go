package factory

import (
	"friends_ranking/config/errorMsg"
	"friends_ranking/config/variable"
	"friends_ranking/container"
	"friends_ranking/http/validator/interf"

	"github.com/gin-gonic/gin"
)

// 表单参数验证器工厂（请勿修改）
func Create(key string) func(context *gin.Context) {

	if value := container.CreateContainersFactory().Get(key); value != nil {
		if val, isOk := value.(interf.ValidatorInterface); isOk {
			return val.CheckParams
		}
	}
	variable.ZapLog.Error(errorMsg.ErrorsValidatorNotExists + ", 验证器模块：" + key)
	return nil
}
