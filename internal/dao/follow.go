package dao

import "douyin/internal/model"

// IsFollow 查看关注情况
func (dao *Dao) IsFollow(userId, beUserId uint) (followFlag bool, err error) {
	follow := model.Follow{}
	result := dao.db.Where("user_id = ? AND be_user_id = ?", userId, beUserId).Limit(1).Find(&follow)
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

// FollowCount 关注数量
func (dao *Dao) FollowCount(uid uint) (count int64) {
	dao.db.Model(&model.Follow{}).Where("user_id = ? AND status = ?", uid, 1).Count(&count)
	return count
}

// FollowerCount 粉丝数量
func (dao *Dao) FollowerCount(uid uint) (count int64) {
	dao.db.Model(&model.Follow{}).Where("be_user_id = ? AND status = ?", uid, 1).Count(&count)
	return count
}
