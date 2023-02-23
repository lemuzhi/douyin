package dao

import (
	"douyin/internal/model"
	"time"
)

func (dao *Dao) FavoriteAction(userId uint, videoId uint, actionType uint8) (err error) {
	/*
	   点赞操作 首先检查表中是否已经有当前用户对此视频的点赞记录
	   若有 则修改此条记录 update
	   若无 则增加一条新记录 insert
	*/
	favorite := model.Favorite{}
	result := dao.db.Where("user_id = ? AND video_id = ?", userId, videoId).Find(&favorite)
	err = result.Error
	if err != nil {
		return
	}

	if result.RowsAffected > 0 { //表中已有记录
		err = dao.db.Model(&favorite).Updates(map[string]interface{}{"status": actionType, "updated_at": time.Now()}).Error
	} else {
		favorite.UserID = userId
		favorite.VideoID = videoId
		favorite.Status = actionType
		err = dao.db.Create(&favorite).Error
	}

	return
}

func (dao *Dao) FavoriteListAction(userId uint) (videos []model.Video, err error) {
	/*
	   获取用户所有点过赞的视频
	*/
	//使用联表查询
	err = dao.db.Raw("SELECT `video`.* from `favorite` JOIN `video` ON `favorite`.`video_id`=`video`.`id`  where `favorite`.`user_id`= ? ORDER BY `favorite`.created_at DESC ", userId).Scan(&videos).Error

	return
}

// FavoriteCount 获取视频点赞数
func (dao *Dao) FavoriteCount(vid uint) (count int64) {
	dao.db.Model(&model.Favorite{}).Where("video_id = ? AND status = ?", vid, 1).Count(&count)
	return
}

func (dao *Dao) IsFavorite(uid, vid uint) bool {
	var count int64

	dao.db.Model(&model.Favorite{}).Where("user_id = ? AND video_id = ? AND status = ?", uid, vid, 1).Count(&count)

	if count > 0 {
		return true
	}
	return false
}
