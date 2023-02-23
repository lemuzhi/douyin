package dao

import (
	"douyin/internal/model"
	"douyin/internal/model/response"
	"log"
)

func (dao *Dao) Register(user *model.User) (uint, error) {
	return user.ID, dao.db.Create(&user).Error
}

func (dao *Dao) Login(username string) (user *model.User, err error) {
	return user, dao.db.Where("username = ?", username).First(&user).Error
}

// GetUserInfo 获取单个用户信息
func (dao *Dao) GetUserInfo(uid uint) (user *response.User, err error) {
	var u model.User
	err = dao.db.Select("id", "username", "avatar", "background_image", "signature").Where("id = ?", uid).First(&u).Error
	if err != nil {
		log.Println("GetUserInfo function query error: ", err)
		return user, err
	}

	workCount, videoIdList := dao.WorkCount(uid) //作品数量
	user = &response.User{
		ID:              u.ID,
		Name:            u.Username,
		FollowCount:     dao.FollowCount(u.ID),
		FollowerCount:   dao.FollowerCount(u.ID),
		Avatar:          u.Avatar,
		BackgroundImage: u.BackgroundImage,
		Signature:       u.Signature,
		TotalFavorited:  dao.TotalFavorited(videoIdList),
		WorkCount:       workCount,
		FavoriteCount:   dao.UserFavoriteCount(uid),
	}
	return user, nil
}

// GetUserInfoList 获取一个列表的用户信息
func (dao *Dao) GetUserInfoList(idList []uint) (map[uint]*response.User, error) {
	userList, err := dao.FindUserIdByIdList(idList)
	if err != nil {
		return nil, err
	}

	authorMap := make(map[uint]*response.User, len(userList))
	var user *response.User
	var workCount int64
	var videoIdList *[]model.Video

	for i := 0; i < len(userList); i++ {
		workCount, videoIdList = dao.WorkCount(userList[i].ID) //作品数量

		user = &response.User{
			ID:              userList[i].ID,
			Name:            userList[i].Username,
			FollowCount:     dao.FollowCount(userList[i].ID),
			FollowerCount:   dao.FollowerCount(userList[i].ID),
			Avatar:          userList[i].Avatar,
			BackgroundImage: userList[i].BackgroundImage,
			Signature:       userList[i].Signature,
			TotalFavorited:  dao.TotalFavorited(videoIdList),
			WorkCount:       workCount,
			FavoriteCount:   dao.UserFavoriteCount(userList[i].ID),
		}
		authorMap[userList[i].ID] = user
	}
	return authorMap, nil
}

func (dao *Dao) FindUserByID(userID uint) (user model.User, err error) {
	return user, dao.db.Select("id", "username").Where("id = ?", userID).Find(&user).Error
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
