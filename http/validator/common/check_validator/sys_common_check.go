package checkvalidator

import (
	"fmt"
	"lottery_annual/config/errorMsg"
	"lottery_annual/config/globalConst"
	"lottery_annual/http/controller"
	"lottery_annual/http/data_transfer"
	"lottery_annual/utils/response"

	"github.com/gin-gonic/gin"
)

type GetVerfyCode struct {
	Key string `form:"key" json:"key"  binding:"required" errMsg:"验证码key不能为空"` // 必填、对于文本,表示它的长度>=1
}

// 验证器语法，参见 Register.go文件，有详细说明

func (q GetVerfyCode) CheckParams(context *gin.Context) {
	q.Key = context.Query("key")
	fmt.Println("验证码获取key ： ", q.Key)
	//1.基本的验证规则没有通过
	if err := context.ShouldBind(&q); err != nil {
		response.ReturnFail(context, errorMsg.ErrorsValidatorNoPass)
		return
	}

	//  该函数主要是将本结构体的字段（成员）按照 consts.ValidatorPrefix+ json标签对应的 键 => 值 形式绑定在上下文，便于下一步（控制器）可以直接通过 context.Get(键) 获取相关值
	extraAddBindDataContext := data_transfer.DataAddContext(q, globalConst.ValidatorPrefix, context)
	if extraAddBindDataContext == nil {
		response.ReturnFail(context, "userLogin表单验证器json化失败")
		return
	} else {
		// 验证完成，调用控制器,并将验证器成员(字段)递给控制器，保持上下文数据一致性
		(&controller.SysCommonController{}).GetVerfyCode(extraAddBindDataContext)
	}

}
