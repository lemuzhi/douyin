package initialize

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"log"
)

var RDB *redis.Client

// InitRedis 初始化连接redis
func InitRedis() {
	RDB = redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis.addr"),
		Username: viper.GetString("redis.username"),
		Password: viper.GetString("redis.password"),
		DB:       viper.GetInt("redis.db"),
	})
	_, err := RDB.Ping(context.TODO()).Result()
	if err != nil {
		log.Println("rdb.Ping(ctx).Result() = ", err)
	}
}
