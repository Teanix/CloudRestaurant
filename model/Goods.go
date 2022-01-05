package model

//食品结构体的定义
type Goods struct {
	Id          int64   `xorm:"pk autoincr" json:"id"`
	Name        string  `xorm:"varchar(12)" json:"name"`        //名称
	Description string  `xorm:"varchar(32)" json:"description"` //商品描述
	Icon        string  `xorm:"varchar(255)" json:"icon"`       //商品图标
	SellCount   int64   `xorm:"int" json:"sell_count"`          //销量
	Price       float32 `xorm:"float" json:"price"`             //优惠价格
	OldPrice    float32 `xorm:"float" json:"old_price"`         //初始价格
	ShopId      int64   `xorm:"int" json:"shop_id"`             //对应商家
}
