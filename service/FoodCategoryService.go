package service

import (
	"CloudRestaurant/dao"
	"CloudRestaurant/model"
)

type FoodCategoryService struct {
}

func (fcs *FoodCategoryService) Categories() ([]model.FoodCategory, error) {
	//调用数据库操作层
	foodCategorydao := dao.NewFoodCategoryDao()
	return foodCategorydao.QueryCategories()
}
