package response

type RelationActionResponse struct {
	Response
}

type FollowListResponse struct {
	Response
	UserList []*User `json:"user_list"`
}

type FriendUser struct {
	User
	Avatar  string `json:"avatar"`
	Message string `json:"message"`
	MsgType uint   `json:"msgType"`
}

type FriendListResponse struct {
	Response
	UserList []*FriendUser `json:"user_list"`
}
