package register_validator

import (
	"lottery_annual/config/globalConst"
	"lottery_annual/container"
	checkvalidator "lottery_annual/http/validator/common/check_validator"
)

// 系统公共接口校验器
func CommonRegisterValidator() {
	//创建容器
	containers := container.CreateContainersFactory()

	//  key 按照前缀+模块+验证动作 格式，将各个模块验证注册在容器
	var key string
	// Users 模块表单验证器按照 key => value 形式注册在容器，方便路由模块中调用
	key = globalConst.ValidatorPrefix + "GetVerfyCode"
	containers.Set(key, checkvalidator.GetVerfyCode{})
}
