package dao

import (
	"douyin/internal/model"
	"log"
)

func (dao *Dao) Login(username string) (user *model.User, err error) {
	return user, dao.db.Where("username = ?", username).First(&user).Error
}

func (dao *Dao) GetUserInfo(id int64) (model.User, error) {
	var user model.User
	err := dao.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		log.Println("GetUserInfo function query error: ", err)
		return user, err
	}
	return user, nil
}
