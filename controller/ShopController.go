package controller

import (
	"CloudRestaurant/service"
	"CloudRestaurant/tool"

	"github.com/gin-gonic/gin"
)

type ShopController struct {
}

func (sc *ShopController) Router(app *gin.Engine) {
	app.GET("/api/shops", sc.GetShopList)
	app.GET("/api/search_shops", sc.SearchShop)
}

func (sc *ShopController) GetShopList(context *gin.Context) {
	longitude := context.Query("longitud-e")
	latitude := context.Query("latitude")

	if longitude == "" || longitude == "undefined" || latitude == "" || latitude == "undefined" {
		longitude = "116.34"
		latitude = "40.34"
	}

	shopService := service.ShopService{}
	shops := shopService.ShopList(longitude, latitude)
	if len(shops) != 0 {
		tool.Success(context, shops)
		return
	}
	tool.Failed(context, "shoplist get error")
}

//关键词搜索商铺信息
func (sc *ShopController) SearchShop(context *gin.Context) {
	longitude := context.Query("longitude")
	latitude := context.Query("latitude")
	keyword := context.Query("keyword")

	if keyword == "" {
		tool.Failed(context, "error search try again")
	}
	if longitude == "" || longitude == "undefined" || latitude == "" || latitude == "undefined" {
		longitude = "116.34"
		latitude = "40.34"
	}
	//执行真正的搜索逻辑
	shopService := service.ShopService{}

	shopService.SearchShops(longitude, latitude, keyword)
}
