package dao

import (
	"douyin/internal/model"
	"douyin/pkg/snowflake"
	"time"
)

//// GenMessageID 合并生成消息id
//func GenMessageID(userIdA, userIdB uint) uint64 {
//	if userIdA > userIdB {
//		a := fmt.Sprintf("%d_%d", userIdB, userIdA)
//		key1, _ := strconv.Atoi(a)
//		return uint64(key1)
//	}
//	b := fmt.Sprintf("%d_%d", userIdA, userIdB)
//	key2, _ := strconv.Atoi(b)
//	return uint64(key2)
//
//}

func (dao *Dao) SendMessage(userId uint, toUserId uint, content string) error {
	message := model.Message{}

	message.Content = content
	message.CreateTime = time.Now().Unix()
	message.FromUserID = userId
	message.ToUserID = toUserId
	message.ID, _ = snowflake.GetID()

	err := dao.db.Create(&message).Error
	if err != nil {
		return err
	}
	err = dao.db.Where("from_user_id = ?", userId).First(&message).Error
	if err != nil {
		return err
	}
	return nil
}

func (dao *Dao) GetMessagebyIdAndTime(formUserId uint, toUserId uint, PreMsgTime int64) (message []model.Message, err error) {
	if PreMsgTime == 0 {
		//未传入上次最新的消息时间，从最早发布的消息开始获取
		err = dao.db.Where("(from_user_id = ? AND to_user_id = ?) OR (to_user_id = ? AND from_user_id = ?)", toUserId, formUserId, toUserId, formUserId).Order("create_time").Find(&message).Error
		return
	}
	err = dao.db.Where("from_user_id = ? AND to_user_id = ? AND create_time > ?", toUserId, formUserId, PreMsgTime).Find(&message).Error
	return
}
