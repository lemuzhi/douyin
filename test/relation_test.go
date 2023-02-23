package test

import (
	"douyin/initialize"
	"douyin/internal/dao"
	"fmt"
	"testing"
	"time"
)

func initServer() {
	//设置时区
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		panic(err)
	}
	time.Local = loc

	//初始化配置
	initialize.InitConfig("../config/config.toml")

	//初始化mysql
	initialize.InitMysql()

	//初始化redis
	initialize.InitRedis()
}

func TestRelationAction(t *testing.T) {
	// TODO: 在这里 initServer 似乎不太对
	initServer()
	// go test -v test/relation_test.go
	fmt.Println("TestRelationAction begin:")
	daoTest := dao.New()
	fmt.Println("TestRelationAction finish dao.New:", daoTest)
	var user_id, be_user_id, status = 2, 3, 1
	err := daoTest.RelationAction(uint(user_id), uint(be_user_id), uint8(status))
	if err != nil {
		fmt.Println("TestRelationAction err:", err.Error())
	}
}

func TestFollowList(t *testing.T) {
	// TODO
	initServer()
	// go test -v test/relation_test.go
	daoTest := dao.New()
	userList, err := daoTest.GetFollowList(2)
	if err != nil {
		fmt.Println("TestFollowList err:", err.Error())
	}

	for i, ele := range userList {
		fmt.Printf("index: %d, element: %s\n", i, string(ele.Name))
	}

}
