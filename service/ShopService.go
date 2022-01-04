package service

import (
	"CloudRestaurant/dao"
	"CloudRestaurant/model"
	"strconv"
)

type ShopService struct {
}

//查询商铺列表数据
func (ss *ShopService) ShopList(long, lat string) []model.Shop {
	longitude, err := strconv.ParseFloat(long, 32)
	if err != nil {
		return nil
	}
	latitude, err := strconv.ParseFloat(lat, 32)
	if err != nil {
		return nil
	}
	shopDao := dao.NewShopDao().QueryShops(longitude, latitude)
	return shopDao
}
