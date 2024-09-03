package response

import (
	"friends_ranking/config/globalConst"

	"github.com/gin-gonic/gin"
)

type JsonStruct struct {
	Code  int         `json:"code"`  // 状态码
	Msg   interface{} `json:"msg"`   // 返回信息
	Data  interface{} `json:"data"`  // 返回值
	Count int64       `json:"count"` // 数量
}

func ReturnJson(c *gin.Context, httpCode int, dataCode int, msg string, data interface{}, count int64) {
	json := &JsonStruct{Code: dataCode, Msg: msg, Data: data, Count: count}
	//Context.Header("key2020","value2020")  	//可以根据实际情况在头部添加额外的其他信息
	c.JSON(httpCode, json)
}

func ReturnDefaultOk(c *gin.Context) {
	json := &JsonStruct{
		Code: globalConst.Response_Code_Ok,
		Msg:  globalConst.Response_Code_Ok,
		Data: "",
	}
	//Context.Header("key2020","value2020")  	//可以根据实际情况在头部添加额外的其他信息
	c.JSON(200, json)
}

func ReturnSuccess(c *gin.Context, data interface{}, count int64) {
	json := &JsonStruct{
		Code:  globalConst.Response_Code_Ok,
		Msg:   globalConst.Response_Code_Ok,
		Data:  data,
		Count: count,
	}
	c.JSON(200, json)
}

func ReturnDeFaultFail(c *gin.Context) {
	json := &JsonStruct{
		Code: globalConst.Response_Code_Fail,
		Msg:  globalConst.Response_Msg_Fail,
	}
	//Context.Header("key2020","value2020")  	//可以根据实际情况在头部添加额外的其他信息
	c.JSON(200, json)
}

func ReturnFail(c *gin.Context, erroMsg string) {
	json := &JsonStruct{
		Code: globalConst.Response_Code_Fail,
		Msg:  erroMsg,
	}
	//Context.Header("key2020","value2020")  	//可以根据实际情况在头部添加额外的其他信息
	c.JSON(200, json)
}
