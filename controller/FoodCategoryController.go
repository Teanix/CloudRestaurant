package controller

import (
	"CloudRestaurant/service"
	"CloudRestaurant/tool"

	"github.com/gin-gonic/gin"
)

type FoodCategoryController struct {
}

func (fcc *FoodCategoryController) Router(engine *gin.Engine) {
	engine.GET("/api/food_category", fcc.foodCategory)
}

func (fcc *FoodCategoryController) foodCategory(ctx *gin.Context) {
	//调用service层获取食品种类信息
	foodCategoryService := service.FoodCategoryService{}
	categories, err := foodCategoryService.Categories()
	if err != nil {
		tool.Failed(ctx, "get foodCategory Failed")
		return
	}
 
	//转换格式
	for _, category := range categories {
		if category.ImageUrl != "" {
			category.ImageUrl = tool.FileServerAddr() + "/" + category.ImageUrl
		}
	}
	tool.Success(ctx, categories)
}
