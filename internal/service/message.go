package service

import (
	"douyin/internal/model/request"
	"douyin/internal/model/response"
	"douyin/pkg/errcode"
)

func (svc *Service) MessageAction(params *request.MessageRequest, userId uint) (resp response.Response, err error) {
	if params.ActionType == 1 {

		resp := response.Response{
			StatusCode: errcode.OK.Code,
			StatusMsg:  errcode.OK.Msg,
		}
		err := svc.dao.SendMessage(userId, params.ToUserID, params.Content)

		return resp, err
	}
	return
}

func (svc *Service) MessageListAction(params *request.MessageListRequest, UserId uint) (resp response.MessageListResponse, err error) {

	res, err := svc.dao.GetMessagebyIdAndTime(UserId, params.ToUserID, params.PreMsgTime)

	resp = response.MessageListResponse{
		Response:    errcode.NewResponse(errcode.OK),
		MessageList: res,
	}
	return
}
