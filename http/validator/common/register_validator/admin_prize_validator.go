package register_validator

import (
	"lottery_annual/config/globalConst"
	"lottery_annual/container"
	checkvalidator "lottery_annual/http/validator/common/check_validator"
)

// 各个业务模块验证器必须进行注册（初始化），程序启动时会自动加载到容器
func AdminPrizeValidator() {
	//创建容器
	containers := container.CreateContainersFactory()

	//  key 按照前缀+模块+验证动作 格式，将各个模块验证注册在容器
	var key string

	// 注册门户类表单参数验证器
	key = globalConst.ValidatorPrefix + "PrizeList"
	containers.Set(key, checkvalidator.PrizeList{})

	key = globalConst.ValidatorPrefix + "SavaOrUpdate"
	containers.Set(key, checkvalidator.PrizeAddAndUpdate{})

	key = globalConst.ValidatorPrefix + "PrizeDelete"
	containers.Set(key, checkvalidator.PrizeDelete{})

	key = globalConst.ValidatorPrefix + "PrizeListByAcId"
	containers.Set(key, checkvalidator.PrizeQueryByAcId{})

}
