package errcode

//type code [2]interface{}

type ResponseCode struct {
	Code int32
	Msg  string
}

var (
	OK               = ResponseCode{0, "成功"}
	Fail             = ResponseCode{1, "失败"}
	ErrService       = ResponseCode{100000, "服务错误"}
	ErrInvalidParams = ResponseCode{100001, "无效参数"}
	ErrUserExists    = ResponseCode{101001, "用户已存在"}
	ErrUserNotExist  = ResponseCode{101002, "用户不存在"}
	ErrUserOrPwd     = ResponseCode{101003, "用户名或密码错误"}

	ErrToken         = ResponseCode{102001, "token错误"}
	ErrTokenNotExist = ResponseCode{102002, "token不存在"}
	ErrTokenExpired  = ResponseCode{102003, "token已过期"}
	ErrTokenFormat   = ResponseCode{102004, "token格式错误"}
	ErrAuthorized    = ResponseCode{102005, "鉴权失败"}

	ErrVideoNotExist = ResponseCode{103001, "视频不存在"}
)
