package response

import "douyin/internal/model"

type MessageListResponse struct {
	Response
	MessageList []model.Message `json:"message_list"`
}
