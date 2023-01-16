package initialize

import (
	"fmt"
	"github.com/spf13/viper"
)

func InitConfig(file string) *viper.Viper {
	vp := viper.New()
	vp.SetConfigFile(file)
	vp.SetConfigType("toml") // 配置文件的类型
	err := vp.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	return vp
}
