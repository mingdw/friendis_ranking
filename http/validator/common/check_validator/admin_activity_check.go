package checkvalidator

import (
	"bytes"
	"fmt"
	"friends_ranking/config/globalConst"

	"friends_ranking/http/controller"
	"friends_ranking/http/data_transfer"
	"friends_ranking/utils/response"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

type ActivityList struct {
	Code      string `form:"code" json:"code"`           // 必填、对于文本,表示它的长度>=1
	StartTime string `form:"startTime" json:"startTime"` //  密码为 必填，长度>=6
	EndTime   string `form:"endTime" json:"endTime"`     //  密码为 必填，长度>=6
	Status    int    `form:"status" json:"status"`
	Limit     int    `form:"limit" json:"limit"`
	PageSize  int    `form:"pageSize" json:"pageSize"`
}

type ActivityAdd struct {
	Title     string `form:"title" json:"title"`         // 必填、对于文本,表示它的长度>=1
	Desc      string `form:"Desc" json:"Desc"`           // 必填、对于文本,表示它的长度>=1
	StartTime string `form:"startTime" json:"startTime"` //  密码为 必填，长度>=6
	EndTime   string `form:"endTime" json:"endTime"`     //  密码为 必填，长度>=6
}

type ActivityUpdate struct {
	Id        int    `form:"id" json:"id"`
	Title     string `form:"title" json:"title"`         // 必填、对于文本,表示它的长度>=1
	Desc      string `form:"Desc" json:"Desc"`           // 必填、对于文本,表示它的长度>=1
	StartTime string `form:"startTime" json:"startTime"` //  密码为 必填，长度>=6
	EndTime   string `form:"endTime" json:"endTime"`     //  密码为 必填，长度>=6
}

type ActivityDelete struct {
	Ids []int `form:"ids" json:"ids"`
}

// 验证器语法，参见 Register.go文件，有详细说明
func (ac ActivityList) CheckParams(context *gin.Context) {
	body, err := context.GetRawData()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("请求参数", string(body))
	context.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	//1.基本的验证规则没有通过
	if err := context.ShouldBindJSON(&ac); err != nil {
		response.ReturnCheckFail(context, err, &ac)
		return
	}

	//  该函数主要是将本结构体的字段（成员）按照 consts.ValidatorPrefix+ json标签对应的 键 => 值 形式绑定在上下文，便于下一步（控制器）可以直接通过 context.Get(键) 获取相关值
	extraAddBindDataContext := data_transfer.DataAddContext(ac, globalConst.ValidatorPrefix, context)
	if extraAddBindDataContext == nil {
		response.ReturnFail(context, "userLogin表单验证器json化失败")
		return
	} else {
		// 验证完成，调用控制器,并将验证器成员(字段)递给控制器，保持上下文数据一致性
		(&controller.ActivityController{}).QueryList(context, ac.Code, ac.StartTime, ac.EndTime, ac.Status, ac.Limit, ac.PageSize)
	}

}

func (r ActivityAdd) CheckParams(context *gin.Context) {
	if err := context.ShouldBind(&r); err != nil {
		response.ReturnCheckFail(context, err, &r)
		return
	}

	//  该函数主要是将本结构体的字段（成员）按照 consts.ValidatorPrefix+ json标签对应的 键 => 值 形式绑定在上下文，便于下一步（控制器）可以直接通过 context.Get(键) 获取相关值
	extraAddBindDataContext := data_transfer.DataAddContext(r, globalConst.ValidatorPrefix, context)
	if extraAddBindDataContext == nil {
		response.ReturnFail(context, "userLogin表单验证器json化失败")
		return
	} else {
		// 验证完成，调用控制器,并将验证器成员(字段)递给控制器，保持上下文数据一致性
		(&controller.ActivityController{}).Add(context, r.Title, r.Desc, r.StartTime, r.EndTime)
	}

}

func (r ActivityUpdate) CheckParams(context *gin.Context) {
	if err := context.ShouldBind(&r); err != nil {
		response.ReturnCheckFail(context, err, &r)
		return
	}

	//  该函数主要是将本结构体的字段（成员）按照 consts.ValidatorPrefix+ json标签对应的 键 => 值 形式绑定在上下文，便于下一步（控制器）可以直接通过 context.Get(键) 获取相关值
	extraAddBindDataContext := data_transfer.DataAddContext(r, globalConst.ValidatorPrefix, context)
	if extraAddBindDataContext == nil {
		response.ReturnFail(context, "userLogin表单验证器json化失败")
		return
	} else {
		// 验证完成，调用控制器,并将验证器成员(字段)递给控制器，保持上下文数据一致性
		(&controller.ActivityController{}).Update(context, r.Title, r.Desc, r.StartTime, r.EndTime, r.Id)
	}
}

func (r ActivityDelete) CheckParams(context *gin.Context) {
	if err := context.ShouldBind(&r); err != nil {
		response.ReturnCheckFail(context, err, &r)
		return
	}

	//  该函数主要是将本结构体的字段（成员）按照 consts.ValidatorPrefix+ json标签对应的 键 => 值 形式绑定在上下文，便于下一步（控制器）可以直接通过 context.Get(键) 获取相关值
	extraAddBindDataContext := data_transfer.DataAddContext(r, globalConst.ValidatorPrefix, context)
	if extraAddBindDataContext == nil {
		response.ReturnFail(context, "userLogin表单验证器json化失败")
		return
	} else {
		// 验证完成，调用控制器,并将验证器成员(字段)递给控制器，保持上下文数据一致性
		(&controller.ActivityController{}).Delete(context, r.Ids)
	}
}
