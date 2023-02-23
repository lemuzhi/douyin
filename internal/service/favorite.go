package service

import (
	"douyin/internal/model/request"
	"douyin/internal/model/response"
	"douyin/pkg/errcode"
	"fmt"
)

func (svc *Service) FavoriteAction(params *request.FavoriteRequest, userId uint) (response.FavoriteResponse, error) {

	response := response.FavoriteResponse{
		StatusCode: errcode.OK.Code,
		StatusMsg:  errcode.OK.Msg,
	}

	err := svc.dao.FavoriteAction(userId, params.VideoId, uint8(params.ActionType))
	return response, err
}

func (svc *Service) FavoriteListAction(params *request.FavoriteListRequest) (resp response.FavoriteListResponse, err error) {

	videos, err := svc.dao.FavoriteListAction(params.UserId)
	if err != nil {
		return resp, err
	}

	var videosRsp []response.VideoResponse

	//存储所有作者的id
	idList := make([]uint, len(videos))
	for i := 0; i < len(videos); i++ {
		idList = append(idList, videos[i].UserID)
	}

	//封装查询到的结果
	// 通过 in 查询 获取视频作者信息
	authorMap, err := svc.dao.GetUserInfoList(idList)
	if err != nil {
		fmt.Println("FavoriteListAction GetUserInfoList error: ", err)
	}
	//查询当前请求用户的所有关注

	//初始化返回参数
	for i := 0; i < len(videos); i++ {
		//查看当前用户是否关注了此用户isFollow
		isFollow, err := svc.dao.IsFollow(params.UserId, videos[i].UserID)
		authorMap[videos[i].UserID].IsFollow = isFollow

		if err != nil {
			return response.FavoriteListResponse{}, err
		}

		vRep := response.VideoResponse{
			ID:            videos[i].ID,
			Author:        authorMap[videos[i].UserID],
			PlayUrl:       videos[i].PlayUrl,
			CoverUrl:      videos[i].CoverUrl,
			FavoriteCount: svc.dao.FavoriteCount(videos[i].ID),
			CommentCount:  svc.dao.CommentCount(videos[i].ID),
			IsFavorite:    true,
			Title:         videos[i].Title,
		}
		videosRsp = append(videosRsp, vRep)
	}

	return response.FavoriteListResponse{
		Response:  errcode.NewResponse(errcode.OK),
		VideoList: videosRsp,
	}, err
}
