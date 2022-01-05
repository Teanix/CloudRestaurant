package model

//商铺服务基础表 - 票 保 准 
type Service struct {
	Id          int64  `xorm:"pk autoincr" json:"id"`          //Id
	Name        string `xorm:"varchar(20)" json:"name"`        //服务名称
	Description string `xorm:"varchar(30)" json:"description"` //服务描述
	IconName    string `xorm:"varchar(3)" json:"icon_name"`    //服务图标名称
	IconColor   string `xorm:"varchar(6)" json:"icon_color"`   //服务图标颜色
}
