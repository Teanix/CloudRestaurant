package controller

import (
	"CloudRestaurant/service"
	"CloudRestaurant/tool"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GoodsController struct {
}

func (gc *GoodsController) Router(app *gin.Engine) {
	app.GET("/api/foods", gc.GetGoods)
}

func (gc *GoodsController) GetGoods(context *gin.Context) {
	shopId, exist := context.GetQuery("shop_id")
	if !exist {
		tool.Failed(context, "get goods error")
		return
	}
	id, err := strconv.Atoi(shopId)
	if err != nil {
		tool.Failed(context, "shop-id to int  error")
		return
	}
	//service层
	goodservice := service.NewGoodsService()
	goods := goodservice.GetFoods(int64(id))
	if len(goods) == 0 {
		tool.Failed(context, "no data")
	}
	tool.Success(context, goods)
}
