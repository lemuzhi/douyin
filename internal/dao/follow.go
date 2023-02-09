package dao

import "douyin/internal/model"

// IsFollow 查看关注情况
func (dao *Dao) IsFollow(userId int64, beUserId int64) (followFlag bool, err error) {
	/*
	   查看关注表
	   若没有数据，代表没有关注
	   若有数据，还要查看状态 1已经关注，2已经取消关注
	*/
	//使用联表查询
	follow := model.Follow{}
	result := dao.db.Where("user_id = ? AND be_user_id = ?", userId, beUserId).Limit(1).Find(&follow)
	//var favorites []model.Favorite
	err = result.Error
	if result.RowsAffected > 0 {
		if follow.Status == 1 {
			followFlag = true
		} else {
			followFlag = false
		}
	} else {
		followFlag = false
	}
	return
}
