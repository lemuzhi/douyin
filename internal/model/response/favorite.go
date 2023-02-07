package response

/*
   author: wubo
   互动接口 返回参数结构体
*/

// FavoriteResponse  点赞操作
type FavoriteResponse struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

type FavoriteListResponse struct {
	Response
	VideoList []Video `json:"video_list"`
}

type CommentResponse struct {
	Response
	Comment Comment `json:"comment"`
}

type CommentListResponse struct {
	Response
	CommentList []Comment `json:"comment_list"`
}
