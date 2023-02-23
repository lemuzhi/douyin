package dao

import (
	"douyin/internal/model"
	"douyin/internal/model/response"
	"fmt"
	"time"
)

func (dao *Dao) RelationAction(userID uint, beUserID uint, actionType uint8) (err error) {
	// 先查看一下数据库中是不是已经有 同样的 user_id, be_user_id 的row
	// TODO: 注意还要修改数据库中的follow_count和follower_count字段（如果有的话）; 有可能自己关注自己吗？
	follow := model.Follow{}
	result := dao.db.Where("user_id = ? AND be_user_id = ?", userID, beUserID).Find(&follow)
	err = result.Error
	if err != nil {
		return
	}

	if result.RowsAffected > 0 {
		err = dao.db.Model(&follow).Updates(map[string]interface{}{"status": actionType, "updated_at": time.Now()}).Error
	} else {
		follow.UserID = userID
		follow.BeUserID = beUserID
		follow.Status = actionType
		err = dao.db.Create(&follow).Error
	}

	return
}

func (dao *Dao) GetFollowList(userID uint) (userList []*response.User, err error) {
	// reference from https://gorm.io/zh_CN/docs/query.html#Joins
	// TODO: 查的时候应该只查 status = 1 的 ? 但是文档的示例图片是有查出取消了关注的
	rows, err := dao.db.Model(&model.Follow{}).Select("user.id, user.username, follow.status").Where("follow.user_id = ? AND follow.status = 1", userID).Joins("left join user on user.id = follow.be_user_id").Rows()
	if err != nil {
		fmt.Println("GetFollowList Rows() error: ", err)
		return nil, err
	}

	defer rows.Close()

	// 需要用tag指定列名，否则会只写入4个int
	type FollowCap struct {
		ID            uint   `gorm:"column:id"`
		Name          string `gorm:"column:username"`
		FollowCount   int64  `gorm:"column:follow_count"`
		FollowerCount int64  `gorm:"column:follower_count"`
		Status        uint   `gorm:"column:status"`
	}

	for rows.Next() {
		var followCap FollowCap

		err = dao.db.ScanRows(rows, &followCap)
		if err != nil {
			fmt.Println("dao.db.ScanRows error: ", err)
		}

		var isFollow bool = true
		if followCap.Status == 2 {
			isFollow = false
		}

		userList = append(userList, &response.User{
			ID:            followCap.ID,
			Name:          followCap.Name,
			FollowCount:   dao.FollowCount(followCap.ID),
			FollowerCount: dao.FollowerCount(followCap.ID),
			IsFollow:      isFollow,
		})
	}

	return userList, nil
}

func (dao *Dao) GetFollowerList(userID uint) (userList []*response.User, err error) {
	rows, err := dao.db.Model(&model.Follow{}).Select("user.id, user.username").Where("follow.be_user_id = ? AND follow.status = 1", userID).Joins("left join user on user.id = follow.user_id").Rows()
	if err != nil {
		fmt.Println("GetFollowList Rows() error: ", err)
		return nil, err
	}

	defer rows.Close()

	// 需要用tag指定列名，否则会只写入4个int
	type FollowCap struct {
		ID            uint   `gorm:"column:id"`
		Name          string `gorm:"column:username"`
		FollowCount   int64  `gorm:"column:follow_count"`
		FollowerCount int64  `gorm:"column:follower_count"`
		Status        uint   `gorm:"column:status"`
	}

	for rows.Next() {
		var followCap FollowCap

		err = dao.db.ScanRows(rows, &followCap)
		if err != nil {
			fmt.Println("dao.db.ScanRows error: ", err)
		}

		// 这里的IsFollow应该是当前登录用户本人是否关注了这个粉丝
		isFollow, err := dao.IsFollow(userID, followCap.ID)
		if err != nil {
			fmt.Println("dao.IsFollow userID=", userID, ", beUserId=", followCap.ID, " error: ", err)
		}

		userList = append(userList, &response.User{
			ID:            followCap.ID,
			Name:          followCap.Name,
			FollowCount:   dao.FollowCount(followCap.ID),
			FollowerCount: dao.FollowerCount(followCap.ID),
			IsFollow:      isFollow,
		})
	}

	return userList, nil
}

func (dao *Dao) GetFriendList(userID uint) (userList []*response.FriendUser, err error) {

	rows, err := dao.db.Raw("SELECT u.id, u.username, u.avatar FROM follow f1 JOIN follow f2 ON f1.status = 1 AND f2.status = 1 AND f1.user_id = ? AND f1.be_user_id = f2.user_id AND f2.be_user_id = ? JOIN user u WHERE u.id = f1.be_user_id", userID, userID).Rows()

	if err != nil {
		fmt.Println("GetFriendList Rows() error: ", err)
		return nil, err
	}

	defer rows.Close()

	type FriendCap struct {
		ID     uint   `gorm:"column:id"`
		Name   string `gorm:"column:username"`
		Avatar string `gorm:"column:avatar"`
	}

	var msg model.Message
	for rows.Next() {
		var friendCap FriendCap

		err = dao.db.ScanRows(rows, &friendCap)
		if err != nil {
			fmt.Println("dao.db.ScanRows error: ", err)
		}
		_ = dao.db.Debug().Where("from_user_id = ? AND to_user_id = ?", friendCap.ID, userID).Order("create_time DESC").Take(&msg).Error
		userList = append(userList, &response.FriendUser{
			User: response.User{
				ID:            friendCap.ID,
				Name:          friendCap.Name,
				FollowCount:   dao.FollowCount(friendCap.ID),
				FollowerCount: dao.FollowerCount(friendCap.ID),
				IsFollow:      true,
			},
			Avatar:  friendCap.Avatar,
			Message: msg.Content,
			MsgType: 0,
		})
	}

	return userList, nil
}
