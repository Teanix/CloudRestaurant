package dao

import (
	"CloudRestaurant/model"
	"CloudRestaurant/tool"
)

type FoodCategoryDao struct {
	*tool.Orm
}

//实例化Dao对象
func NewFoodCategoryDao() *FoodCategoryDao {
	return &FoodCategoryDao{tool.Dbengine}
}

//从数据库查询食品种类并返回
func (fcd *FoodCategoryDao) QueryCategories() ([]model.FoodCategory, error) {
	var categories []model.FoodCategory

	if err := fcd.Engine.Find(&categories); err != nil {
		return nil, err
	}
	return categories, nil
}
