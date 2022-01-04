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
func (sd *ShopDao) QueryShops(longitude, latitude float64) []model.Shop {
	var shops []model.Shop
	err := sd.Where("longitude > ? and longitude < ? and latitude > ? and latitude < ?",
		longitude-DEFAULT_RANGE,
		longitude+DEFAULT_RANGE,
		latitude-DEFAULT_RANGE,
		latitude+DEFAULT_RANGE).Find(&shops)
	if err != nil {
		return nil
	}
	return shops
}
