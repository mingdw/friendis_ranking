package register_validator

import (
	"friends_ranking/config/globalConst"
	"friends_ranking/container"
	checkvalidator "friends_ranking/http/validator/common/check_validator"
)

// 各个业务模块验证器必须进行注册（初始化），程序启动时会自动加载到容器
func AdminActivityValidator() {
	//创建容器
	containers := container.CreateContainersFactory()

	//  key 按照前缀+模块+验证动作 格式，将各个模块验证注册在容器
	var key string

	// 注册门户类表单参数验证器
	key = globalConst.ValidatorPrefix + "ActivityList"
	containers.Set(key, checkvalidator.ActivityList{})

	key = globalConst.ValidatorPrefix + "ActivityAdd"
	containers.Set(key, checkvalidator.ActivityAdd{})

	key = globalConst.ValidatorPrefix + "ActivityUpdate"
	containers.Set(key, checkvalidator.ActivityUpdate{})

	key = globalConst.ValidatorPrefix + "ActivityDelete"
	containers.Set(key, checkvalidator.ActivityDelete{})

	key = globalConst.ValidatorPrefix + "ActivityUpdateStatus"
	containers.Set(key, checkvalidator.ActivityUpDateStatus{})

}
