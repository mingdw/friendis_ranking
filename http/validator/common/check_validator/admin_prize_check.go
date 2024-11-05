package checkvalidator

import (
	"fmt"
	"lottery_annual/config/globalConst"
	"strconv"

	"lottery_annual/http/controller"
	"lottery_annual/http/data_transfer"
	"lottery_annual/utils/response"

	"github.com/gin-gonic/gin"
)

type PrizeList struct {
	Code      string `form:"code" json:"code"`           // 必填、对于文本,表示它的长度>=1
	PrizeName string `form:"prizeName" json:"prizeName"` //  密码为 必填，长度>=6
	Status    int    `form:"status" json:"status"`
	Limit     int    `form:"limit" json:"limit"`
	PageSize  int    `form:"pageSize" json:"pageSize"`
}

type PrizeAddAndUpdate struct {
	Id        int    `form:"id" json:"id" errMsg:"id不合法"`
	AcId      int    `form:"acId" json:"acId" errMsg:"acid不合法"`
	PrizeName string `form:"prizeName" json:"prizeName" errMsg:"prizeName不合法"`
	ItemCode  string `form:"itemCode" json:"itemCode" errMsg:"itemCode不合法"`
	PrizeNum  int    `form:"prizeNum" json:"prizeNum" errMsg:"prizeNum不合法"`
	IsRepeat  int    `form:"isRepeat" json:"isRepeat" errMsg:"isRepeat不合法"`
}

type PrizeDelete struct {
	Ids []int `form:"ids" json:"ids"`
}

type PrizeQueryByAcId struct {
	AcId int `form:"acId" json:"acId" errMsg:"id不合法"`
}

// 验证器语法，参见 Register.go文件，有详细说明
func (ac PrizeList) CheckParams(context *gin.Context) {
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
		(&controller.PrizeController{}).QueryList(context, ac.Code, ac.PrizeName, ac.Status, ac.Limit, ac.PageSize)
	}

}

func (r PrizeAddAndUpdate) CheckParams(context *gin.Context) {
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
		itmCode, _ := strconv.Atoi(r.ItemCode)

		fmt.Println("saveOrUpdate请求参数, id:", r.Id, "; acId: ", r.AcId, "; itmCode: ", itmCode, "; isRepeat: ", r.IsRepeat, "; prizeNum: ", r.PrizeNum, "; prizeName: ", r.PrizeName)
		(&controller.PrizeController{}).PrizeAddAndUpDate(context, r.Id, r.AcId, itmCode, r.IsRepeat, r.PrizeNum, r.PrizeName)
	}

}

func (r PrizeDelete) CheckParams(context *gin.Context) {
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
		(&controller.PrizeController{}).Delete(context, r.Ids)
	}
}

func (r PrizeQueryByAcId) CheckParams(context *gin.Context) {
	acId := context.Param("acId")
	if acId == "" {
		response.ReturnFail(context, "活动编号不能为空")
	}
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
		id, err := strconv.Atoi(acId)
		if err == nil {
			(&controller.PrizeController{}).SelectByAcId(context, id)
		}

	}

}
