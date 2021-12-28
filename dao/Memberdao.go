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

func (md *MemberDao) ValidateSmsCode(phone string, code string) *model.Smscode {
	var sms model.Smscode

	if _, err := md.Where("phone = ? and code = ?", phone, code).Get(&sms); err != nil {
		fmt.Println(err.Error())
	}

	return &sms
}

func (md *MemberDao) InsertCode(sms model.Smscode) int64 {
	result, err := md.InsertOne(&sms)
	if err != nil {
		logger.Error(err.Error())
	}
	return result
}
