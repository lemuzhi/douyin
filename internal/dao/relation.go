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

func (dao *Dao) GetFollowList(userID uint) (userList []*model.Follow, err error) {
	//查询我关注的，并且关注状态是1的用户，1已关注，2已取消
	err = dao.db.Model(&model.Follow{}).Select("be_user_id").Where("user_id = ? AND status = ?", userID, 1).
		Order("created_at DESC").Find(&userList).Error
	return
}

func (dao *Dao) GetFollowerList(userID uint) (userList []*model.Follow, err error) {
	//查询关注我，并且关注状态是1的用户，1已关注，2已取消
	err = dao.db.Model(&model.Follow{}).Select("user_id").Where("be_user_id = ? AND status = ?", userID, 1).
		Order("created_at DESC").Find(&userList).Error
	return
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
	for rows.Next() {
		var friendCap FriendCap
		var msg model.Message
		err = dao.db.ScanRows(rows, &friendCap)
		if err != nil {
			fmt.Println("dao.db.ScanRows error: ", err)
		}

		fmt.Println("from_user_id =", friendCap.ID)
		fmt.Println("to_user_id =", userID)
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
