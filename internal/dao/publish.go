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

	user, err := dao.GetUserInfo(userID)
	if err != nil {
		fmt.Println("publish list GetUserInfo error : ", err)
	}
	if beUserID > 0 {
		user.IsFollow, _ = dao.IsFollow(userID, beUserID)
	}

	for rows.Next() {
		var video model.Video
		// ScanRows 方法用于将一行记录扫描至结构体
		err = dao.db.ScanRows(rows, &video)
		if err != nil {
			fmt.Println("dao.db.ScanRows error: ", err)
		}

		// 业务逻辑...
		videoList = append(videoList, &response.VideoResponse{
			ID:            video.ID,
			Title:         video.Title,
			Author:        user,
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			FavoriteCount: dao.FavoriteCount(video.ID),
			CommentCount:  dao.CommentCount(video.ID),
			IsFavorite:    dao.IsFavorite(user.ID, video.ID),
		})
	}
	return videoList, nil
}

// WorkCount 获取已发布的作品数量
func (dao *Dao) WorkCount(uid uint) (count int64, videoList *[]model.Video) {
	dao.db.Model(&model.Video{}).Select("id").Where("user_id = ?", uid).Count(&count).Find(&videoList)
	return count, videoList
}

// UserFavoriteCount 获取点赞过的视频数量
func (dao *Dao) UserFavoriteCount(uid uint) int64 {
	var count int64
	dao.db.Model(&model.Favorite{}).Where("user_id = ?", uid).Count(&count)
	return count
}

// TotalFavorited 所有作品获得点赞的总数
func (dao *Dao) TotalFavorited(videoIdList *[]model.Video) int64 {
	var count int64
	var num int64
	for i := 0; i < len(*videoIdList); i++ {
		dao.db.Model(&model.Favorite{}).Where("video_id = ?", (*videoIdList)[i].ID).Count(&count)
		num += count
	}
	return num
}
