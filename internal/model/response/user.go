package response

type RegisterResponse struct {
	Response
	UserID int64  `json:"user_id"`
	Token  string `json:"token"`
}

type LoginResponse struct {
	Response
	UserID int64  `json:"user_id"`
	Token  string `json:"token"`
}

type UserInfoResponse struct {
	Response
	User User
}
