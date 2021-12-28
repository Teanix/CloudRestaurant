package main

import (
	"CloudRestaurant/controller"
	"CloudRestaurant/tool"

	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
)

func main() {

	cfg, err := tool.ParseConfig("./config/app.json")
	if err != nil {
		panic(err)
	}
	_, err = tool.OrmEngine(cfg)
	if err != nil {
		logger.Error(err.Error())
	}
	app := gin.Default()
	registerRouter(app)
	app.Run(cfg.AppHost + ":" + cfg.AppPort)
}

//路由设置
func registerRouter(router *gin.Engine) {
	new(controller.HelloController).Router(router)
	new(controller.MemberController).Router(router)
}
