package dao

import (
	"CloudRestaurant/model"
	"CloudRestaurant/tool"
	"fmt"

	"github.com/wonderivan/logger"
)

type MemberDao struct {
	*tool.Orm
}

//根据电话和密码查询用户
func (md *MemberDao) Query(name string, password string) *model.Member {
	var member model.Member
	password = tool.EncoderSha256(password)
	if _, err := md.Where("user_name = ? and password = ?", name, password).Get(&member); err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return &member
}

//根据电话查询用户
func (md *MemberDao) QueryByPhone(phone string) *model.Member {
	var member model.Member

	if _, err := md.Where("mobile = ?").Get(&member); err != nil {
		fmt.Println(err.Error())
	}

	return &member
}

//插入新用户
func (md *MemberDao) InsertMember(member model.Member) int64 {
	res, err := md.InsertOne(&member)
	if err != nil {
		fmt.Println(err.Error())
		return 0
	}
	return res
}

//根据手机及验证码查询数据库
func (md *MemberDao) ValidateSmsCode(phone string, code string) *model.Smscode {
	var sms model.Smscode

	if _, err := md.Where("phone = ? and code = ?", phone, code).Get(&sms); err != nil {
		fmt.Println(err.Error())
	}

	return &sms
}

//插入电话及验证码
func (md *MemberDao) InsertCode(sms model.Smscode) int64 {
	result, err := md.InsertOne(&sms)
	if err != nil {
		logger.Error(err.Error())
	}
	return result
}
