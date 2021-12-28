package service

import (
	"CloudRestaurant/dao"
	"CloudRestaurant/model"
	"CloudRestaurant/param"
	"CloudRestaurant/tool"
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"github.com/wonderivan/logger"
)

type MemberService struct {
}

func (ms *MemberService) Sendcode(phone string) bool {

	//1.生成验证码
	code := fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))

	//2.调用SDK发送
	config := tool.GetConfig().Sms
	client, err := dysmsapi.NewClientWithAccessKey(config.RegionId, config.AppKey, config.Appsecret)
	if err != nil {
		logger.Error(err.Error())
		return false
	}
	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"
	request.SignName = config.SignName
	request.TemplateCode = config.TemplateCode
	request.PhoneNumbers = phone
	//组装待发送信息
	par, err := json.Marshal(map[string]interface{}{
		"name": code,
	})
	if err != nil {
		logger.Error(err.Error())
	}
	request.TemplateParam = string(par)
	//接收返回结果
	response, err := client.SendSms(request)
	if err != nil {
		logger.Error(err.Error())
		fmt.Print(err.Error())
		return false
	}
	fmt.Println(response)

	//3.判断发送状态
	if response.Code == "OK" {
		// 将验证码保存到数据库中
		smscode := model.Smscode{Phone: phone, BizID: response.BizId, Code: code, CreateTime: time.Now().Unix()}
		MemberDao := dao.MemberDao{tool.Dbengine}
		result := MemberDao.InsertCode(smscode)

		return result > 0

	}
	return false
}

//完成手机+验证码登陆的实体操作
func (ms *MemberService) SmsLogin(loginparam param.SmsLoginParam) *model.Member {
	//todo
	//1.获取到手机号+验证码

	//2.验证是否正确
	//3.根据手机号查询记录
	// 若不存在则新创建且保存
	return nil
}
