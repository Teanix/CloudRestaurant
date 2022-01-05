package dao

import (
	"CloudRestaurant/model"
	"CloudRestaurant/tool"
)

type GoodsDao struct {
	*tool.Orm
}

//实例化Dao对象
func NewGoodsDao() *GoodsDao {
	return &GoodsDao{tool.Dbengine}
}

//根据商户id查询拥有所有商品
func (gd *GoodsDao) QueryFoods(shop_id int64) ([]model.Goods, error) {
	var goods []model.Goods

	err := gd.Orm.Where("shop_id = ?", shop_id).Find(&goods)

	if err != nil {
		return nil, err
	}

	return goods, nil
}
