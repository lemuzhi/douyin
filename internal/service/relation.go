package service

import (
	"douyin/internal/model/request"
	"douyin/internal/model/response"
	"douyin/pkg/errcode"

	"github.com/gin-gonic/gin"
)

func (svc *Service) RelationAction(c *gin.Context, params request.RelationActionReq) error {
	// fmt.Println("RelationAction", c.GetInt64("UserID"), c.GetUint("UserID"))
	// 使用 GetUint 而非 GetInt64
	return svc.dao.RelationAction(c.GetUint("UserID"), params.ToUserID, uint8(params.ActionType))
}

func (svc *Service) FollowList(params request.FollowListReq) (*response.FollowListResponse, error) {
	// TODO: cannot use id (variable of type string) as type int64 in argument to svc.dao.GetFollowList
	userList, err := svc.dao.GetFollowList(params.UserID)
	if err != nil {
		return nil, err
	}

	data := response.FollowListResponse{
		Response: errcode.NewResponse(errcode.OK),
		UserList: userList,
	}
	return &data, nil
}

func (svc *Service) FollowerList(params request.FollowListReq) (*response.FollowListResponse, error) {
	userList, err := svc.dao.GetFollowerList(params.UserID)
	if err != nil {
		return nil, err
	}

	data := response.FollowListResponse{
		Response: errcode.NewResponse(errcode.OK),
		UserList: userList,
	}
	return &data, nil
}
