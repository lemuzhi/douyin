package service

import (
	"douyin/internal/model"
	"douyin/internal/model/request"
	"douyin/internal/model/response"
	"douyin/pkg/errcode"
	"time"
)

func (svc *Service) GetFeedList(uid uint, params *request.FeedRequest) (resp response.FeedResponse, err error) {
	var lastTime time.Time
	if params.LatestTime == 0 {
		//如果未传入时间戳，使用当前时间
		lastTime = time.Now().Local()
	} else {
		//对传入的时间戳进行处理
		lastTime = time.UnixMilli(params.LatestTime)
	}
	videoList, err := svc.dao.GetFeedList(lastTime)

	//得到视频列表长度
	n := len(*videoList)
	if n == 0 {
		return resp, err
	}

	var isFollow, isFavorite bool
	var video *model.Video
	var videos []*response.VideoResponse
	for i := 0; i < n; i++ {
		video = (*videoList)[i]
		//判断用户用户是否点赞该视频，是否关注该视频的用户
		if uid > 0 {
			isFavorite = svc.dao.IsFavorite(uid, video.ID)
			isFollow, _ = svc.dao.IsFollow(uid, video.User.ID)
		}

		workCount, videoIdList := svc.dao.WorkCount(video.User.ID) //作品数量

		videos = append(videos, &response.VideoResponse{
			ID:    video.ID,
			Title: video.Title,
			Author: response.User{
				ID:              video.User.ID,
				Name:            video.User.Username,
				FollowCount:     svc.dao.FollowCount(video.User.ID),
				FollowerCount:   svc.dao.FollowerCount(video.User.ID),
				IsFollow:        isFollow,
				Avatar:          video.User.Avatar,
				BackgroundImage: video.User.BackgroundImage,
				Signature:       video.User.Signature,
				TotalFavorited:  svc.dao.TotalFavorited(videoIdList),
				WorkCount:       workCount,
				FavoriteCount:   svc.dao.UserFavoriteCount(video.User.ID), //点赞数量,
			},
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			FavoriteCount: int64(len(video.Favorites)),
			CommentCount:  int64(len(video.Comments)),
			IsFavorite:    isFavorite,
		})
	}
	return response.FeedResponse{
		Response:  errcode.NewResponse(errcode.OK),
		VideoList: videos,
		NextTime:  (*videoList)[n-1].CreatedAt.UnixMilli(),
	}, err
}
