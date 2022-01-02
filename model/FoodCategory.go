package model

//食品种类
type FoodCategory struct {
	Id          int64  `xorm:"pk autoincr" json:"id"`          //类别ID
	Title       string `xorm:"varchar(20)" json:"title"`       //食品类别标题
	Description string `xorm:"varchar(30)" json:"description"` //食品描述
	ImageUrl    string `xorm:"varchar(255)" json:"image_url"`  //食品种类图片
	LinkUrl     string `xorm:"varchar(255)" json:"link_rl"`    //食品类别链接
	IsInServing bool   `json:"is_in_serving"`                  //是否处于服务状态
}
