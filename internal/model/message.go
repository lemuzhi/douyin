package model

type MessageList struct {
	MessageList []Message `json:"message_list"` // 用户列表
	StatusCode  string    `json:"status_code"`  // 状态码，0-成功，其他值-失败
	StatusMsg   *string   `json:"status_msg"`   // 返回状态描述
}

type Message struct {
	ID         uint64 `json:"id"`           // 消息id
	FromUserID uint   `json:"from_user_id"` // 消息发送者id
	ToUserID   uint   `json:"to_user_id"`   // 消息接收者id
	Content    string `json:"content"`      // 消息内容
	CreateTime int64  `json:"create_time"`  // 消息发送时间 yyyy-MM-dd HH:MM:ss

}
