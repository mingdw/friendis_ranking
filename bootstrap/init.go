package bootstrap

import (
	"log"
	"lottery_annual/config/errorMsg"
	"lottery_annual/config/variable"
	"lottery_annual/config/yamlConfig"
	"lottery_annual/http/validator/common/register_validator"
	"lottery_annual/service/sys_log_hook"
	"lottery_annual/utils/gorm_v2"
	"lottery_annual/utils/redis_factory"
	"lottery_annual/utils/zap_factory"
	"os"
)

func init() {
	log.Println("********************lottery_annual init **********************")
	checkRequiredFolders()

	//3.初始化表单参数验证器，注册在容器（Web、Api共用容器）
	register_validator.WebRegisterValidator()
	register_validator.ApiRegisterValidator()
	register_validator.CommonRegisterValidator()
	register_validator.AdminActivityValidator()

	variable.YamlConfig = yamlConfig.CreateYamlFactory()
	variable.YamlConfig.ConfigFileChangeListen()
	variable.RedisPool = redis_factory.InitRedisClientPool()

	// 5.初始化全局日志句柄，并载入日志钩子处理函数
	variable.ZapLog = zap_factory.CreateZapFactory(sys_log_hook.ZapLogHandler)

	if variable.YamlConfig.GetInt("Gormv2.Mysql.IsInitGlobalGormMysql") == 1 {
		if dbMysql, err := gorm_v2.GetOneMysqlClient(); err != nil {
			log.Fatal(errorMsg.ErrorsGormInitFail + err.Error())
		} else {
			variable.GormDbMysql = dbMysql
		}
	}
	log.Println("********************lottery_annual init end**********************")
}

// 检查项目必须的非编译目录是否存在，避免编译后调用的时候缺失相关目录
func checkRequiredFolders() {

	log.Println("check RequireFolders, current system profile: ", variable.ENV)
	//1.检查配置文件是否存在
	if _, err := os.Stat(variable.BasePath + "/resource/application_" + variable.ENV + ".yaml"); err != nil {
		log.Fatal(errorMsg.ErrorsConfigYamlNotExists + err.Error())
	}
	log.Println("application yaml folder check end!")

	//2.检查public目录是否存在
	if _, err := os.Stat(variable.BasePath + "/static/"); err != nil {
		log.Fatal(errorMsg.ErrorsPublicNotExists + err.Error())
	}
	log.Println("public  folder check end!")
	//3.检查storage/logs 目录是否存在
	if _, err := os.Stat(variable.BasePath + "/logs/"); err != nil {
		log.Fatal(errorMsg.ErrorsStorageLogsNotExists + err.Error())
	}
	log.Println("logs folder check end!")
}
