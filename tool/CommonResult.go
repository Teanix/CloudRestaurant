package tool

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	SUCCESS int = 0
	FAILED  int = 1
)

//普通成功返回
func Success(ctx *gin.Context, msg interface{}, v interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": SUCCESS,
		"msg":  msg,
		"data": v,
	})
}

func Failed(ctx *gin.Context, msg interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": FAILED,
		"msg":  msg,
	})
}

type CommonResult struct {
}
