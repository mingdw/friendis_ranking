package register_validator

import (
	"friends_ranking/config/globalConst"
	"friends_ranking/container"
	checkvalidator "friends_ranking/http/validator/common/check_validator"
)

// 各个业务模块验证器必须进行注册（初始化），程序启动时会自动加载到容器
func WebRegisterValidator() {
	//创建容器
	containers := container.CreateContainersFactory()

	//  key 按照前缀+模块+验证动作 格式，将各个模块验证注册在容器
	var key string
	// Users 模块表单验证器按照 key => value 形式注册在容器，方便路由模块中调用
	key = globalConst.ValidatorPrefix + "Add"
	containers.Set(key, checkvalidator.Add{})

	key = globalConst.ValidatorPrefix + "Query"
	containers.Set(key, checkvalidator.Query{})

	key = globalConst.ValidatorPrefix + "Delete"
	containers.Set(key, checkvalidator.Delete{})
}
