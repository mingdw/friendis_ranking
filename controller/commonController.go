package controller

import "github.com/gin-gonic/gin"

type JsonStruct struct {
	Code  int         `json:"code"`  // 状态码
	Msg   interface{} `json:"msg"`   // 返回信息
	Data  interface{} `json:"data"`  // 返回值
	Count int64       `json:"count"` // 数量
}

type JsonErrorStruct struct {
	Code int         `json:"code"` // 状态码
	Msg  interface{} `json:"msg"`  // 返回信息
}

func ReturnSuccess(c *gin.Context, code int, msg interface{}, data interface{}, count int64) {
	json := &JsonStruct{Code: code, Msg: msg, Data: data, Count: count}
	c.JSON(200, json)
}

func ReturnError(c *gin.Context, msg interface{}) {
	json := &JsonErrorStruct{Code: -1, Msg: msg}
	c.JSON(200, json)
}
