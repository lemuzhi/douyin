package service

import (
	"douyin/internal/model/request"
	"douyin/internal/model/response"
	"douyin/pkg/errcode"
)

func (svc *Service) FavoriteAction(params *request.FavoriteRequest, userId uint) (response.FavoriteResponse, error) {

	response := response.FavoriteResponse{
		StatusCode: errcode.OK.Code,
		StatusMsg:  errcode.OK.Msg,
	}

	err := svc.dao.FavoriteAction(userId, params.VideoId, uint8(params.ActionType))
	return response, err
}

func (svc *Service) FavoriteListAction(params *request.FavoriteListRequest) (response.FavoriteListResponse, error) {

	videos, err := svc.dao.FavoriteListAction(params.UserId)

	var videosRsp []response.VideoResponse

	if len(videos) == 0 {
		return response.FavoriteListResponse{
			Response:  errcode.NewResponse(errcode.OK),
			VideoList: videosRsp,
		}, err
	}

	//存储所有作者的id
	idList := make([]uint, len(videos))
	for i := 0; i < len(videos); i++ {
		idList = append(idList, videos[i].UserID)
	}

	//封装查询到的结果
	// 通过 in 查询 获取视频作者信息
	authorMap, err0 := svc.UsersMap(idList)

	if err0 != nil {
		return response.FavoriteListResponse{}, err0
	}
	//查询当前请求用户的所有关注

	//初始化返回参数
	for i := 0; i < len(videos); i++ {
		user := authorMap[videos[i].UserID]
		//查看当前用户是否关注了此作者isFollow
		followFlag, err2 := svc.dao.IsFollow(params.UserId, videos[i].UserID)

		if err2 != nil {
			return response.FavoriteListResponse{}, err2
		}

		// 获取作者的作品数量（发布的视频数量）和视频id列表
		wCount, videoList := svc.dao.WorkCount(user.ID)
		userRsp := response.User{
			ID:              user.ID,
			Name:            user.Username,
			FollowCount:     svc.dao.FollowCount(user.ID),
			FollowerCount:   svc.dao.FollowerCount(user.ID),
			IsFollow:        followFlag,
			Avatar:          user.Avatar,
			BackgroundImage: user.BackgroundImage,
			Signature:       user.Signature,
			TotalFavorited:  svc.dao.TotalFavorited(videoList),
			WorkCount:       wCount,
			FavoriteCount:   svc.dao.UserFavoriteCount(user.ID),
		}
		vRep := response.VideoResponse{
			ID:            videos[i].ID,
			Author:        userRsp,
			PlayUrl:       videos[i].PlayUrl,
			CoverUrl:      videos[i].CoverUrl,
			FavoriteCount: svc.dao.FavoriteCount(videos[i].ID),
			CommentCount:  svc.dao.CommentCount(videos[i].ID),
			IsFavorite:    true, // 必定是点赞了的视频
			Title:         videos[i].Title,
		}
		videosRsp = append(videosRsp, vRep)
	}

	return response.FavoriteListResponse{
		Response:  errcode.NewResponse(errcode.OK),
		VideoList: videosRsp,
	}, err
}
