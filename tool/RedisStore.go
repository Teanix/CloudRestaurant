package tool

import (
	"log"
	"time"

	"github.com/go-redis/redis"
	"github.com/mojocn/base64Captcha"
)

type RedisStore struct {
	client *redis.Client
}

var Redis_Store RedisStore

func InitRedisStore() *RedisStore {
	config := GetConfig().RedisConfig
	client := redis.NewClient(&redis.Options{
		Addr:     config.Addr + ":" + config.Port,
		Password: config.Password,
		DB:       config.Db,
	})
	Redis_Store = RedisStore{client: client}
	base64Captcha.SetCustomStore(&Redis_Store)
	
	return &Redis_Store
}

//set
func (rs *RedisStore) Set(id string, value string) {
	if err := rs.client.Set(id, value, time.Minute*10).Err(); err != nil {
		log.Println(err)
	}
}

//get
func (rs *RedisStore) Get(id string, clear bool) string {
	val, err := rs.client.Get(id).Result()
	if err != nil {
		log.Println(err)
		return ""
	}
	if clear {
		if err := rs.client.Del(id).Err(); err != nil {
			log.Println(err)
			return ""
		}
	}

	return val
}
