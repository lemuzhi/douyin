package dao

import (
	"douyin/internal/model"
)

func (dao *Dao) Feed() (video model.Video, user model.User, err error) {
	err = dao.db.Preload("User").Where("id=?", 1).First(&video).Error
	if err != nil {
		return
	}
	err = dao.db.Where("id=?", video.UserID).First(&user).Error
	if err != nil {
		return
	}
	return
}
