package controller

import (
	"CloudRestaurant/param"
	"CloudRestaurant/service"
	"CloudRestaurant/tool"

	"github.com/gin-gonic/gin"
)

type MemberController struct {
}

func (mc *MemberController) Router(engine *gin.Engine) {
	engine.GET("/api/sendcode", mc.Sendcode)
	engine.OPTIONS("/api/login_sms", mc.SmsLogin)
}

func (mc *MemberController) Sendcode(context *gin.Context) {
	//处理发送验证码Sendcode
	phone, exist := context.GetQuery("phone")
	if !exist {
		tool.Failed(context, "phone not exist")
		return
	}

	ms := service.MemberService{}
	isSend := ms.Sendcode(phone)
	if isSend {
		tool.Success(context, "send success", "success")
	} else {
		tool.Failed(context, "send error")
	}
}

func (mc *MemberController) SmsLogin(context *gin.Context) {
	var smsLoginParam param.SmsLoginParam

	if err := tool.Decode(context.Request.Body, &smsLoginParam); err != nil {
		tool.Failed(context, "json decode error")
		return
	}
	//完成手机+验证码登陆的逻辑
	us := service.MemberService{}
	if member := us.SmsLogin(smsLoginParam); member != nil {
		tool.Success(context, "member success", member)
		return
	}
	tool.Failed(context, "member Failed")
}
