package initialize

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"log"
)

var RDB *redis.Client

// InitRedis 初始化连接redis
func InitRedis(config *viper.Viper) {
	RDB = redis.NewClient(&redis.Options{
		Addr:     config.GetString("redis.addr"),
		Username: config.GetString("redis.username"),
		Password: config.GetString("redis.password"),
		DB:       config.GetInt("redis.db"),
	})
	_, err := RDB.Ping(context.TODO()).Result()
	if err != nil {
		log.Println("rdb.Ping(ctx).Result() = ", err)
	}
}
