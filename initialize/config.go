package initialize

import (
	"douyin/pkg/snowflake"
	"fmt"
	"github.com/spf13/viper"
)

func InitConfig(file string) {
	viper.New()
	viper.SetConfigFile(file)
	viper.SetConfigType("toml") // 配置文件的类型
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	if err := snowflake.Init(1); err != nil {
		fmt.Printf("init snowflake failed, err:%v\n", err)
		return
	}
}
