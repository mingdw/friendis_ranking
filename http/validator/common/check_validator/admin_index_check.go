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

type Login struct {
	UserName string `form:"userName" json:"userName"  binding:"required,min=2,max=32" errMsg:"登陆账户不合法，请重新输入"`  // 必填、对于文本,表示它的长度>=1
	Pass1    string `form:"pass1" json:"pass1" binding:"required,min=6,max=20" errMsg:"密码长度不得小于6或者大于20"`       //  密码为 必填，长度>=6
	Pass2    string `form:"pass2" json:"pass2" binding:"required,min=6,max=20,eqfield=Pass1" errMsg:"两次密码不一致"` //  密码为 必填，长度>=6
}

type Regist struct {
	UserName string `form:"userName" json:"userName"  binding:"required,min=2,max=32" errMsg:"用户名不能为空，长度不能小于2大于32" ` // 必填、对于文本,表示它的长度>=1
	NickName string `form:"nickName" json:"nickName"  binding:"max=32" errMsg:"昵称长度不得超过32"`                          // 必填、对于文本,表示它的长度>=1
	Account  string `form:"account" json:"account"  binding:"max=32" errMsg:"账号不得超过32位"`                             // 必填、对于文本,表示它的长度>=1
	IdCard   string `form:"idCard" json:"idCard"  errMsg:"身份证号不合法" `
	Password string `form:"password" json:"password" binding:"required,min=6" errMsg:"密码长度不得小于6位"`   //  密码为 必填，长度>=6
	Email    string `form:"email" json:"email" errMsg:"邮箱地址不合法"`                                     //  密码为 必填，长度>=6
	Age      int    `form:"age" json:"age" binding:"gte=1" errMsg:"年龄不合法"`                           //  密码为 必填，长度>=6
	Phone    string `form:"phone" json:"phone"  errMsg:"电话号码不合法"`                                    //  密码为 必填，长度>=6
	Job      string `form:"job" json:"job" binding:"max=32" errMsg:"职位长度不能超过32位"`                    //  密码为 必填，长度>=6
	BirthDay string `form:"birthDay" json:"birthDay" binding:"datetime=2006-01-02" errMsg:"日期格式不合法"` //  密码为 必填，长度>=6
	Sex      int    `form:"sex" json:"sex" errMsg:"性别不合法"`                                           //  密码为 必填，长度>=6

}

// 验证器语法，参见 Register.go文件，有详细说明

func (l Login) CheckParams(context *gin.Context) {
	//1.基本的验证规则没有通过
	if err := context.ShouldBind(&l); err != nil {
		response.ReturnCheckFail(context, err, &l)
		return
	}

	//  该函数主要是将本结构体的字段（成员）按照 consts.ValidatorPrefix+ json标签对应的 键 => 值 形式绑定在上下文，便于下一步（控制器）可以直接通过 context.Get(键) 获取相关值
	extraAddBindDataContext := data_transfer.DataAddContext(l, globalConst.ValidatorPrefix, context)
	if extraAddBindDataContext == nil {
		response.ReturnFail(context, "userLogin表单验证器json化失败")
		return
	} else {
		// 验证完成，调用控制器,并将验证器成员(字段)递给控制器，保持上下文数据一致性
		(&controller.IndexController{}).Login(extraAddBindDataContext)
	}

}

func (r Regist) CheckParams(context *gin.Context) {
	body, err := context.GetRawData()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("请求参数", string(body))
	context.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	//1.基本的验证规则没有通过
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
		(&controller.IndexController{}).Register(extraAddBindDataContext)
	}

}
