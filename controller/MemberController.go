package controller

import (
	"CloudRestaurant/model"
	"CloudRestaurant/param"
	"CloudRestaurant/service"
	"CloudRestaurant/tool"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type MemberController struct {
}

func (mc *MemberController) Router(engine *gin.Engine) {
	engine.GET("/api/sendcode", mc.Sendcode)
	engine.POST("/api/login_sms", mc.SmsLogin)
	engine.GET("/api/captcha", mc.Captcha)
	engine.POST("/api/vertifycha", mc.Vertifycha)
	engine.POST("/api/login_pwd", mc.NameLogin)
	engine.POST("/api/upload/avator", mc.UploadAvator)
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
		//将用户信息保存到session中
		sess, _ := json.Marshal(member)
		if err := tool.SetSess(context, "user_"+string(member.Id), sess); err != nil {
			tool.Failed(context, "session Set Failed")
			return
		}
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
		return
	}
	// 3.登录
	ms := service.MemberService{}

	member := ms.Login(loginParam.Name, loginParam.Password)
	if member.Id != 0 {
		//将用户信息保存到session中
		sess, err := json.Marshal(member)
		if err != nil {
			fmt.Println("json Marshal error!")
			return
		}
		// if err := tool.SetSess(context, "user_16"+string(member.Id), sess); err != nil {
		if err := tool.SetSess(context, "user_16", sess); err != nil {
			tool.Failed(context, "session set  error")
			return
		}

		tool.Success(context, &member)
		return
	}
	tool.Failed(context, "login error")
}

//头像文件上传
func (mc *MemberController) UploadAvator(context *gin.Context) {
	// 1.解析上传的参数：image-file , user-ID
	userID := context.PostForm("user_id")
	fmt.Println("userID >>", userID)
	file, err := context.FormFile("avatar")
	if err != nil {
		tool.Failed(context, "avator decode error")
		return
	}
	// 2.通过session判断用户是否已经登录
	sess := tool.GetSess(context, "user_"+userID)
	if sess == nil {
		tool.Failed(context, "get user session error")
		return
	}
	var member model.Member
	json.Unmarshal(sess.([]byte), &member) //将session的内容设置为member的内容
	// 3.将file保存到本地
	fileName := "./uploadfile/" + strconv.FormatInt(time.Now().Unix(), 10) + file.Filename
	if err := context.SaveUploadedFile(file, fileName); err != nil {
		tool.Failed(context, "save avator error")
		return
	}
	//3.1 将文件上传到fdfs
	fileID := tool.UploadFile(fileName)
	if fileID != "" {
		//删除本地的文件
		os.Remove(fileName)
		//将路径保存到用户表中的头像字段
		memberService := service.MemberService{}
		path := memberService.UploadAvator(member.Id, fileID)
		if path != "" {
			tool.Success(context, tool.FileServerAddr()+"/"+path)
		}
	}

	tool.Failed(context, "upload avator error")
	// 5.返回结果
}
