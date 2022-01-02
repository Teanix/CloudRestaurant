package tool

import (
	"fmt"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

func InitSession(engine *gin.Engine) {
	config := GetConfig().RedisConfig
	store, err := redis.NewStore(10, "tcp", config.Addr+":"+config.Port, "", []byte("secret"))
	if err != nil {
		fmt.Println(err.Error())
	}
	engine.Use(sessions.Sessions("mysession", store))
}

//set session
func SetSess(context *gin.Context, key interface{}, value interface{}) error {
	session := sessions.Default(context) //获取当前Session
	fmt.Println("key:", key, "value:", value)
	session.Set(key, value)
	fmt.Println("set session load")
	return session.Save()
}

//get session
func GetSess(context *gin.Context, key interface{}) interface{} {
	session := sessions.Default(context)
	return session.Get(key)
}
