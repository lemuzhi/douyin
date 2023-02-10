package response

type FavoriteResponse struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

type FavoriteListResponse struct {
	Response
	VideoList []VideoResponse `json:"video_list"`
}
