package model

type ShopService struct {
	ShopId    int64 `xorm:"pk not null" json:"shop_id"`    //某一个商户的ID
	ServiceId int64 `xorm:"pk not null" json:"service_id"` //商户具备的服务ID
}
 