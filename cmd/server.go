package cmd

import (
	"douyin/initialize"
	"douyin/internal"
	"github.com/gin-gonic/gin"
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
	config := initialize.InitConfig("./config/config.toml")

	//初始化mysql
	initialize.InitMysql(config)

	//初始化redis
	initialize.InitRedis(config)

	//设置gin的启动模式
	gin.SetMode(config.GetString("gin.mode"))

	r := gin.Default()

	internal.InitRouter(r)

	return r.Run(config.GetString("gin.addr"))
}
