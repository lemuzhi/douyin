package dao

import "douyin/internal/model"

func (dao *Dao) AddComment(userId, videoId uint, commentText string) (comment model.Comment, user model.User, err error) {
	/*
	   添加评论并返回此条评论和评论者相关信息
	*/

	//插入评论
	comment.UserID = userId
	comment.VideoID = videoId
	comment.Content = commentText
	err = dao.db.Create(&comment).Error
	if err != nil {
		return
	}

	//用户信息
	err = dao.db.Where("id = ?", uint(userId)).First(&user).Error
	if err != nil {
		return
	}

	return
}

func (dao *Dao) DeleteComment(commentId uint) (err error) {

	comment := model.Comment{}
	comment.ID = commentId

	err = dao.db.Delete(&comment).Error

	return
}

func (dao *Dao) GetCommentsByVideoId(videoId uint) (comments []model.Comment, err error) {

	err = dao.db.Where("video_id = ?", videoId).Find(&comments).Error

	return
}

// CommentCount 获取视频评论数
func (dao *Dao) CommentCount(vid uint) (count int64) {
	dao.db.Model(&model.Comment{}).Where("video_id = ?", vid).Count(&count)
	return
}
