package initialize

import (
	"douyin/internal/model"
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
)

var DB *gorm.DB

func migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		model.User{},     //用户表
		model.Video{},    //视频表
		model.Comment{},  //评论表
		model.Favorite{}, //喜欢视频表
		model.Follow{},   //关注表
	)
}

// InitMysql 初始化连接mysql
func InitMysql(config *viper.Viper) {
	//dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s",
		config.GetString("mysql.username"),
		config.GetString("mysql.password"),
		config.GetString("mysql.addr"),
		config.GetString("mysql.port"),
		config.GetString("mysql.db"),
		config.GetString("mysql.config"),
	)
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: &schema.NamingStrategy{
			SingularTable: true,
		},
		//禁用事物
		//SkipDefaultTransaction: true,
	})
	if err != nil {
		panic(err)
	}

	//迁移
	err = migrate(DB)
	if err != nil {
		log.Println("gorm mysql AutoMigrate error: ", err)
	}
}
