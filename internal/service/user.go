package service

import (
	"douyin/internal/model"
	"douyin/internal/model/request"
	"douyin/internal/model/response"
	"douyin/pkg/errcode"
	"douyin/pkg/utils"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"time"
)

func (svc *Service) Register(params request.RegisterReq) (*response.LoginResponse, error) {
	count := svc.dao.FindUserCount(params.Username)
	if count > 0 {
		return nil, errors.New(errcode.ErrUserExists.Msg)
	}
	pwd, err := utils.EncipherPassword(params.Password)
	if err != nil {
		return nil, err
	}
	now := time.Now()
	user := model.User{
		Username: params.Username,
		Password: pwd,
		//客户端用户注册暂无设置头像、背景图、简介功能，也无修改功能，所以暂时写死
		Avatar:          "https://c-ssl.duitang.com/uploads/blog/202102/08/20210208200511_45cb8.jpg",
		BackgroundImage: "https://article.autotimes.com.cn/wp-content/uploads/2022/04/95f35f8c40454bf1b4f18d7c79b5b948.jpg",
		Signature:       fmt.Sprintf("我于%d年%d月%d日%d时%d分%d秒注册了抖声，欢迎关注！", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second()),
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
		UserID:   id,
		Token:    token,
	}
	return data, nil
}

func (svc *Service) Login(params request.LoginReq) (*response.LoginResponse, error) {

	user, err := svc.dao.Login(params.Username)
	if err != nil {
		return nil, err
	}

	//验证密码
	if !utils.VerifyPassword(user.Password, params.Password) {
		return nil, fmt.Errorf("%v", errcode.ErrUserOrPwd)
	}

	//生成token
	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		return nil, err
	}

	data := &response.LoginResponse{
		Response: errcode.NewResponse(errcode.OK),
		UserID:   user.ID,
		Token:    token,
	}
	return data, nil
}

func (svc *Service) GetUserInfo(params request.UserReq) *response.UserInfoResponse {
	user, err := svc.dao.GetUserInfo(params.UserID)

	data := &response.UserInfoResponse{User: user}

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
