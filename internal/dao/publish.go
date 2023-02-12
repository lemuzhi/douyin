package dao

import (
	"douyin/internal/model"
	"douyin/internal/model/response"
	"fmt"
)

func (dao *Dao) PublishAction(userID uint, title, playUrl, coverUrl string) error {
	video := model.Video{
		Title:    title,
		PlayUrl:  playUrl,
		CoverUrl: coverUrl,
		UserID:   userID,
	}
	return dao.db.Create(&video).Error
}

func (dao *Dao) GetPublishList(userID, beUserID uint) (videoList []*response.VideoResponse, err error) {
	rows, err := dao.db.Model(&model.Video{}).Select("id", "title", "play_url", "cover_url").
		Where("user_id = ?", userID).Order("created_at DESC").Rows()
	if err != nil {
		fmt.Println("GetPublishList Rows() error: ", err)
		return nil, err
	}

	defer rows.Close()

	var user model.User
	err = dao.db.Select("id", "username").Where("id = ?", userID).Find(&user).Error
	if err != nil {
		return nil, err
	}

	var isFollow bool
	if beUserID > 0 {
		isFollow, _ = dao.IsFollow(userID, beUserID)
	}
	followCount := dao.FollowCount(userID)
	followerCount := dao.FollowerCount(beUserID)

	for rows.Next() {
		var video model.Video
		// ScanRows 方法用于将一行记录扫描至结构体
		err = dao.db.ScanRows(rows, &video)
		if err != nil {
			fmt.Println("dao.db.ScanRows error: ", err)
		}

		// 业务逻辑...
		videoList = append(videoList, &response.VideoResponse{
			ID:    video.ID,
			Title: video.Title,
			Author: response.User{
				ID:            user.ID,
				Name:          user.Username,
				FollowCount:   followCount,
				FollowerCount: followerCount,
				IsFollow:      isFollow,
			},
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			FavoriteCount: dao.FavoriteCount(video.ID),
			CommentCount:  dao.CommentCount(video.ID),
			IsFavorite:    dao.IsFavorite(user.ID, video.ID),
		})
	}

	return videoList, nil
}
