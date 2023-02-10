package cmd

import (
	"douyin/initialize"
	"douyin/internal"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"time"
)

func RunServer() error {
	//设置时区
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		panic(err)
	}
	time.Local = loc

	//初始化配置
	initialize.InitConfig("./config/config.toml")

	//初始化mysql
	initialize.InitMysql()

	//初始化redis
	initialize.InitRedis()

	//设置gin的启动模式
	gin.SetMode(viper.GetString("gin.mode"))

	r := gin.Default()

	internal.InitRouter(r)

	return r.Run(viper.GetString("gin.addr"))
}
