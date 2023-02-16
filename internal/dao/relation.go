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

	rows, err := dao.db.Model(&model.Follow{}).Select("user.id, user.username, follow.status").Where("follow.user_id = ?", userID).Joins("left join user on user.id = follow.be_user_id").Rows()
	// rows, err := dao.db.Table("follow").Select("user.id, user.username, follow_count, follower_count, follow.status").Where("follow.user_id = ?", userID).Joins("left join user on user.id = follow.be_user_id").Rows()
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
	// rows, err := dao.db.Model(&model.Follow{}).Select("user.id, user.username, follow_count, follower_count, follow.status").Where("follow.be_user_id = ?", userID).Joins("left join user on user.id = follow.user_id").Rows()
	rows, err := dao.db.Model(&model.Follow{}).Select("user.id, user.username, follow.status").Where("follow.be_user_id = ?", userID).Joins("left join user on user.id = follow.user_id").Rows()
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

		var isFollow bool
		isFollow, err = dao.IsFollow(userID, followCap.ID)
		if err != nil {
			fmt.Println("dao.IsFollow userID=", userID, ", beUserId=", followCap.ID, " error: ", err)
			return nil, err
		}
		// 这里的IsFollow应该是当前登录用户本人是否关注了这个粉丝
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
