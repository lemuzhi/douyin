package model

import (
	"time"
)

type MessageList struct {
	MessageList []Message `json:"message_list"` // 用户列表
	StatusCode  string    `json:"status_code"`  // 状态码，0-成功，其他值-失败
	StatusMsg   *string   `json:"status_msg"`   // 返回状态描述
}

// Message
type Message struct {
	FromUserID uint      `json:"from_user_id"` // 消息发送者id
	ID         uint64    `json:"id"`           // 消息id
	ToUserID   uint      `json:"to_user_id"`   // 消息接收者id
	Content    string    `json:"content"`      // 消息内容
	CreateTime time.Time `json:"create_time"`  // 消息发送时间 yyyy-MM-dd HH:MM:ss

}
