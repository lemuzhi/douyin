package controller

import (
	"douyin/internal/model/request"
	"douyin/internal/model/response"
	"douyin/internal/service"
	"douyin/pkg/errcode"
	"douyin/pkg/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CommentAction 评论操作
func CommentAction(c *gin.Context) {
	params := request.CommentRequest{}
	send := errcode.New(c)
	err := c.ShouldBindQuery(&params)
	if err != nil {
		send.RespFail(errcode.ErrInvalidParams)
		return
	}
	//获取当前请求的用户id
	svc := service.New(c)
	userId := c.GetUint("UserID")
	data, err := svc.CommentAction(&params, userId)
	if err != nil {
		send.RespFailDetail(errcode.Fail, err.Error())
		return
	}

	send.RespData(data)
}

func CommentListAction(c *gin.Context) {
	params := request.CommentListRequest{}
	send := errcode.New(c)
	err := c.ShouldBindQuery(&params)
	if err != nil {
		send.RespFail(errcode.ErrInvalidParams)
		return
	}
	svc := service.New(c)
	data := response.CommentListResponse{}
	fmt.Println(params.Token)
	// 检查是否携带token
	if params.Token == "" {
		data, err = svc.CommentListAction(&params, 0) //用户id为0，即未登录
	} else {
		claims, err := utils.ParseToken(params.Token)
		if err != nil {
			c.JSON(http.StatusOK, errcode.NewResponse(errcode.ErrAuthorized, err))
			c.Abort()
			return
		}
		data, err = svc.CommentListAction(&params, claims.UserID)
	}

	if err != nil {
		send.RespFailDetail(errcode.Fail, err.Error())
		return
	}

	send.RespData(data)
}
