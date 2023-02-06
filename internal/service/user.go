package service

import (
	"douyin/internal/model"
	"douyin/internal/model/request"
	"douyin/internal/model/response"
	"douyin/pkg/errcode"
	"douyin/pkg/utils"
	"fmt"
	"gorm.io/gorm"
)

func (svc *Service) Register(params request.RegisterReq) (*response.LoginResponse, error) {
	pwd, err := utils.EncipherPassword(params.Password)
	if err != nil {
		return nil, err
	}
	fmt.Println("密码：", pwd)
	user := model.User{
		Username: params.Username,
		Password: pwd,
	}
	id, err := svc.dao.Register(&user)
	if err != nil {
		return nil, err
	}

	token, err := utils.GenerateToken(id)
	if err != nil {
		return nil, err
	}

	data := &response.LoginResponse{
		Response: errcode.NewResponse(errcode.OK),
		UserID:   int64(id),
		Token:    token,
	}
	return data, nil
}

func (svc *Service) Login(params request.LoginReq) (*response.LoginResponse, error) {

	user, err := svc.dao.Login(params.Username)
	if err != nil {
		return nil, err
	}

	if !utils.VerifyPassword(user.Password, params.Password) {
		return nil, fmt.Errorf("%v", errcode.ErrUserOrPwd)
	}

	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		return nil, err
	}

	data := &response.LoginResponse{
		Response: errcode.NewResponse(errcode.OK),
		UserID:   1,
		Token:    token,
	}
	return data, nil
}

func (svc *Service) GetUserInfo(params request.UserReq) *response.UserInfoResponse {
	user, err := svc.dao.GetUserInfo(params.UserID)

	data := &response.UserInfoResponse{
		User: response.User{
			ID:            user.ID,
			Name:          user.Nickname,
			FollowCount:   user.FollowCount,
			FollowerCount: user.FollowerCount,
			IsFollow:      false,
		},
	}

	if err != nil {
		//没有找到用户，用户不存在
		if err == gorm.ErrRecordNotFound {
			data.Response = errcode.NewResponse(errcode.ErrUserNotExist)
			return data
		}
		//其他错误
		data.Response = errcode.NewResponse(errcode.ErrService)
		return data
	}

	data.Response = errcode.NewResponse(errcode.OK)
	return data
}
