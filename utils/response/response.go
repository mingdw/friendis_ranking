package response

import (
	"friends_ranking/config/globalConst"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type JsonStruct struct {
	Code  int         `json:"code"`  // 状态码
	Msg   interface{} `json:"msg"`   // 返回信息
	Data  interface{} `json:"data"`  // 返回值
	Count int64       `json:"count"` // 数量
}

func ReturnJson(c *gin.Context, httpCode int, dataCode int, msg string, data interface{}, count int64) {
	//Context.Header("key2020","value2020")  	//可以根据实际情况在头部添加额外的其他信息
	var resp = map[string]interface{}{
		"code":  dataCode,
		"msg":   msg,
		"data":  msg,
		"count": count,
	}
	c.JSON(httpCode, resp)
}

func ReturnDefaultOk(c *gin.Context) {
	var resp = map[string]interface{}{
		"code":  globalConst.Response_Code_Ok,
		"msg":   globalConst.Response_Msg_Ok,
		"data":  "",
		"count": 0,
	}
	//Context.Header("key2020","value2020")  	//可以根据实际情况在头部添加额外的其他信息
	c.JSON(200, resp)
}

func ReturnSuccess(c *gin.Context, data interface{}, count int64) {
	var resp = map[string]interface{}{
		"code":  globalConst.Response_Code_Ok,
		"msg":   globalConst.Response_Msg_Ok,
		"data":  data,
		"count": count,
	}
	c.JSON(200, resp)
}

func ReturnDeFaultFail(c *gin.Context) {
	var resp = map[string]interface{}{
		"code": globalConst.Response_Code_Fail,
		"msg":  globalConst.Response_Msg_Fail,
	}
	//Context.Header("key2020","value2020")  	//可以根据实际情况在头部添加额外的其他信息
	c.JSON(200, resp)
}

func ReturnFail(c *gin.Context, erroMsg string) {
	//Context.Header("key2020","value2020")  	//可以根据实际情况在头部添加额外的其他信息
	var resp = map[string]interface{}{
		"code": globalConst.Response_Code_Fail,
		"msg":  erroMsg,
	}
	//Context.Header("key2020","value2020")  	//可以根据实际情况在头部添加额外的其他信息
	c.JSON(200, resp)
}

func ReturnCheckFail(c *gin.Context, err error, obj any) {
	title := _GetValidMsg(err, obj)
	if title == "" {
		title = "请求参数不合法，请重新输入1"
	}

	var resp = map[string]interface{}{
		"code": globalConst.Response_Code_Fail,
		"msg":  title,
	}
	//Context.Header("key2020","value2020")  	//可以根据实际情况在头部添加额外的其他信息
	c.JSON(200, resp)
}

func _GetValidMsg(err error, obj any) string {
	// 使用的时候，需要传obj的指针
	getObj := reflect.TypeOf(obj)
	// 将err接口断言为具体类型
	if errs, ok := err.(validator.ValidationErrors); ok {
		// 断言成功
		for _, e := range errs {
			// 循环每一个错误信息
			// 根据报错字段名，获取结构体的具体字段
			if f, exits := getObj.Elem().FieldByName(e.Field()); exits {
				msg := f.Tag.Get("errMsg")
				return msg
			}
		}
	}
	return "请求参数不合法，请重新输入"
}
