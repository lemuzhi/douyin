package controller

import (
	"douyin/internal/model/request"
	"douyin/internal/service"
	"douyin/pkg/errcode"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	params := request.RegisterReq{}
	send := errcode.New(c)
	err := c.ShouldBindQuery(&params)
	if err != nil {
		send.RespFailDetail(errcode.ErrInvalidParams, "账号或密码长度不能小于6位")
		return
	}

	svc := service.New(c)
	data, err := svc.Register(params)
	if err != nil {
		send.RespFailDetail(errcode.Fail, err.Error())
		return
	}

	send.RespData(data)
}

func Login(c *gin.Context) {
	params := request.LoginReq{}
	send := errcode.New(c)
	err := c.ShouldBindQuery(&params)
	if err != nil {
		send.RespFail(errcode.ErrInvalidParams)
		return
	}

	svc := service.New(c)
	data, err := svc.Login(params)
	if err != nil {
		send.RespFailDetail(errcode.Fail, err.Error())
		return
	}
	send.RespData(data)
}

func GetUserInfo(c *gin.Context) {
	params := request.UserReq{}
	send := errcode.New(c)
	err := c.ShouldBindQuery(&params)
	if err != nil {
		send.RespFail(errcode.ErrInvalidParams)
		return
	}
	svc := service.New(c)
	data := svc.GetUserInfo(params)

	send.RespData(data)
}
