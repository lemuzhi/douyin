package main

import (
	"douyin/cmd"
	"log"
)

func main() {
	//启动服务
	err := cmd.RunServer()
	if err != nil {
		log.Println("Run server arise error: ", err)
	}
}
