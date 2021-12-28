package controller

import (
	"github.com/gin-gonic/gin"
)

type HelloController struct { //方法
}

func (hello *HelloController) Router(engine *gin.Engine) {

	engine.GET("/hello", hello.Hello)
}

func (hello *HelloController) Hello(context *gin.Context) {
	context.JSON(200, gin.H{
		"name": "yang",
		"age":  "18",
	})
}
