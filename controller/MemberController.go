package controller

import (
	"CloudRestaurant/param"
	"CloudRestaurant/service"
	"CloudRestaurant/tool"
	"fmt"

	"github.com/gin-gonic/gin"
)

type MemberController struct {
}

func (mc *MemberController) Router(engine *gin.Engine) {
	engine.GET("/api/sendcode", mc.Sendcode)
	engine.POST("/api/login_sms", mc.SmsLogin)
	engine.GET("/api/captcha", mc.Captcha)
	engine.GET("/api/vertifycha", mc.Vertifycha)
	engine.POST("/api/login_pwd", mc.NameLogin)
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
		tool.Success(context, "send success")
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
		tool.Success(context, member)
		return
	}
	tool.Failed(context, "member Login Failed")
}

//生成验证码,并返回客户端
func (mc *MemberController) Captcha(context *gin.Context) {
	tool.GenerateCaptcha(context)
}

//验证码是否正确
func (mc *MemberController) Vertifycha(context *gin.Context) {
	var captcha tool.CaptchaResult
	if err := tool.Decode(context.Request.Body, &captcha); err != nil {
		tool.Failed(context, "captcha decode error!")
		return
	}
	if res := tool.VertifyCaptcha(captcha.ID, captcha.VertifyValue); res {
		fmt.Println("验证通过")
	} else {
		fmt.Println("验证失败")
	}

}

func (mc *MemberController) NameLogin(context *gin.Context) {
	//1.解析用户登录传递参数
	var loginParam param.Loginparam
	if err := tool.Decode(context.Request.Body, &loginParam); err != nil {
		tool.Failed(context, "login param decode error")
		return
	}

	// 2.验证验证码
	if !(tool.VertifyCaptcha(loginParam.Id, loginParam.Value)) {
		tool.Failed(context, "captcha error,Please try again")
	}
	// 3.登录
	ms := service.MemberService{}

	member := ms.Login(loginParam.Name, loginParam.Password)
	if member.Id != 0 {
		tool.Success(context, &member)
		return
	}
	tool.Failed(context, "login error")
}
