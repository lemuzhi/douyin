package request

type FeedRequest struct {
	LatestTime int64  `json:"latest_time" binding:"omitempty"`
	Token      string `json:"token" binding:"omitempty"`
}
