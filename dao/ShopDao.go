package dao

import (
	"CloudRestaurant/model"
	"CloudRestaurant/tool"
)

type ShopDao struct {
	*tool.Orm
}

func NewShopDao() *ShopDao {
	return &ShopDao{tool.Dbengine}
}

const DEFAULT_RANGE = 5

//查询商铺的数据
func (sd *ShopDao) QueryShops(longitude, latitude float64, keyword string) []model.Shop {
	var shops []model.Shop
	if keyword == "" {
		err := sd.Where("longitude > ? and longitude < ? and latitude > ? and latitude < ? and name like ? and status = 1",
			longitude-DEFAULT_RANGE,
			longitude+DEFAULT_RANGE,
			latitude-DEFAULT_RANGE,
			latitude+DEFAULT_RANGE,
			keyword).Find(&shops)
		if err != nil {
			return nil
		}
	} else {
		err := sd.Where("longitude > ? and longitude < ? and latitude > ? and latitude < ? and name like ? and status = 1",
			longitude-DEFAULT_RANGE,
			longitude+DEFAULT_RANGE,
			latitude-DEFAULT_RANGE,
			latitude+DEFAULT_RANGE,
			keyword).Find(&shops)
		if err != nil {
			return nil
		}
	}

	return shops
}
