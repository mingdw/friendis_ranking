package variable

import (
	"friends_ranking/config/errorMsg"
	"friends_ranking/config/yamlConfig/ymlConfigInterf"
	"log"
	"os"
	"strings"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

// ginskeleton 封装的全局变量全部支持并发安全，请放心使用即可
// 开发者自行封装的全局变量，请做好并发安全检查与确认

var (
	BasePath           string              // 定义项目的根目录
	EventDestroyPrefix        = "Destroy_" //  程序退出时需要销毁的事件前缀
	ConfigKeyPrefix           = "Config_"  //  配置文件键值缓存时，键的前缀
	// 表单验证器前缀
	ValidatorPrefix string = "Form_Validator_" //表单校验器前缀

	// 全局日志指针
	ZapLog *zap.Logger

	//gorm 数据库客户端，如果您操作数据库使用的是gorm，请取消以下注释，在 bootstrap>init 文件，进行初始化即可使用
	GormDbMysql      *gorm.DB // 全局gorm的客户端连接
	GormDbSqlserver  *gorm.DB // 全局gorm的客户端连接
	GormDbPostgreSql *gorm.DB // 全局gorm的客户端连接

	//websocket
	WebsocketHub              interface{}
	WebsocketHandshakeSuccess = `{"code":200,"msg":"ws连接成功","data":""}`
	WebsocketServerPingMsg    = "Server->Ping->Client"

	YamlConfig ymlConfigInterf.YmlConfigInterf // 全局配置文件指针
	ENV        string

	//  用户自行定义其他全局变量 ↓

)

func init() {
	log.Println("全局变量初始化!!!")
	// 1.初始化程序根目录
	if curPath, err := os.Getwd(); err == nil {
		// 路径进行处理，兼容单元测试程序程序启动时的奇怪路径
		if len(os.Args) > 1 && strings.HasPrefix(os.Args[1], "-test") {
			BasePath = strings.Replace(strings.Replace(curPath, `\test`, "", 1), `/test`, "", 1)
		} else {
			BasePath = curPath
		}
	} else {
		log.Println(errorMsg.ErrorsBasePath)
	}
	env := os.Getenv("ENV")
	if env == "" {
		env = "dev"
	}
	ENV = env
}
