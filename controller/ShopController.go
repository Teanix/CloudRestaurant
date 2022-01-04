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
