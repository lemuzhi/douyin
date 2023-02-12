package dao

import (
	"douyin/internal/model"
	"log"
)

func (dao *Dao) Register(user *model.User) (uint, error) {
	return user.ID, dao.db.Create(&user).Error
}

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

func (dao *Dao) FindUserByID(userID uint) (user model.User, err error) {
	return user, dao.db.Select("id", "username", "follow_count", "follower_count").Where("id = ?", userID).Find(&user).Error
}

func (dao *Dao) FindUserByName(name string) (user model.User, err error) {
	return user, dao.db.Where("username = ?", name).Find(&user).Error
}

func (dao *Dao) FindUserIdByName(name string) (user model.User, err error) {
	return user, dao.db.Select("id").Where("username = ?", name).First(&user).Error
}

func (dao *Dao) FindUserIdByIdList(idList []uint) (users []model.User, err error) {

	return users, dao.db.Where("id IN ?", idList).Find(&users).Error
}
