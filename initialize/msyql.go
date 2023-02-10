package initialize

import (
	"database/sql"
	"douyin/internal/model"
	"fmt"
	mysql2 "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
	"strings"
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
func InitMysql() {
	//dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s",
		viper.GetString("mysql.username"),
		viper.GetString("mysql.password"),
		viper.GetString("mysql.addr"),
		viper.GetString("mysql.port"),
		viper.GetString("mysql.db"),
		viper.GetString("mysql.config"),
	)
	var err error
Label:
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: &schema.NamingStrategy{
			SingularTable: true,
		},
		//禁用事物
		//SkipDefaultTransaction: true,
	})
	if err != nil {
		mysqlErr, ok := err.(*mysql2.MySQLError)
		if !ok {
			panic(err)
		}

		if mysqlErr.Message == fmt.Sprintf("Unknown database '%s'", viper.GetString("mysql.db")) {
			db, errs := sql.Open("mysql", strings.SplitAfter(dsn, "/")[0])
			if errs != nil {
				log.Println("db Open error: ", errs)
			}
			if errs = db.Ping(); errs != nil {
				log.Println("db Ping error: ", errs)
			}
			createSql := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s DEFAULT CHARACTER SET utf8mb4 DEFAULT COLLATE utf8mb4_general_ci;", viper.GetString("mysql.db"))
			_, errs = db.Exec(createSql)
			if errs != nil {
				log.Panic("db Exec error: ", errs)
			}
			goto Label
		}
	}

	//迁移
	err = migrate(DB)
	if err != nil {
		log.Println("gorm mysql AutoMigrate error: ", err)
	}
}
