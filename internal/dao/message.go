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
	message.CreateTime = time.Now()
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

func (dao *Dao) GetMessagebyIdAndTime(userId uint, toUserId uint, PreMsgTime int64) (message []model.Message, err error) {

	err = dao.db.Where("from_user_id= ?", userId).Where("to_user_id", toUserId).Where("create_time > ?", PreMsgTime).Find(&message).Error
	return
}
