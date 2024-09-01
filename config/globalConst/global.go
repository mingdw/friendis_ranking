package globalConst

const (
	Response_Code_Ok int    = 200
	Response_Msg_Ok  string = "Success"

	Response_Code_Fail int    = -400000
	Response_Msg_Fail  string = "系统错误，请联系管理员"

	Response_Code_JWTError int    = -400100
	Response_Msg_JWTError  string = "token校验失败"

	Response_Code_Bussniess_Error int    = -400200
	Response_Msg_Bussniess_Error  string = "系统也无能执行出错"

	DateFormat = "2006-01-02 15:04:05" //  设置全局日期时间格式

	JWTSecretKey = "friends_ranking" //默认jwt加解密密钥

	Profile = "dev" //dev/test/prod
)
