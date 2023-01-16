package controller

import (
	"douyin/internal/model/request"
	"douyin/internal/service"
	errcode2 "douyin/pkg/errcode"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	params := request.LoginReq{}
	err := c.ShouldBindQuery(&params)
	if err != nil {
		c.JSON(http.StatusBadRequest, errcode2.NewResponse(errcode2.ErrInvalidParams))
		return
	}

	svc := service.New(c)
	data, err := svc.Login(params)
	if err != nil {
		c.JSON(http.StatusBadRequest, errcode2.NewResponse(errcode2.Fail, err))
		return
	}

	c.JSON(http.StatusOK, data)
}

func GetUserInfo(c *gin.Context) {
	params := request.UserReq{}
	err := c.ShouldBindQuery(&params)
	if err != nil {
		c.JSON(http.StatusOK, errcode2.NewResponse(errcode2.ErrInvalidParams))
		return
	}
	svc := service.New(c)
	data := svc.GetUserInfo(params)

	c.JSON(http.StatusOK, data)
}
