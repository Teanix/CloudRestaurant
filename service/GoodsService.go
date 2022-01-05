package service

import (
	"CloudRestaurant/dao"
	"CloudRestaurant/model"
)

type GoodsService struct {
}

func NewGoodsService() *GoodsService {
	return &GoodsService{}
}

//获取商家的食品列表
func (gs *GoodsService) GetFoods(shopid int64) []model.Goods {
	goodsDao := dao.NewGoodsDao()

	var goods []model.Goods
	goods, err := goodsDao.QueryFoods(shopid)
	if err != nil {
		return nil
	}
	return goods
}
