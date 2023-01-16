package errcode

import (
	"douyin/internal/model/response"
	"fmt"
)

func NewResponse(code ResponseCode, err ...interface{}) response.Response {
	msg := code.Msg
	for _, arg := range err {
		fmt.Println(arg)
		switch arg.(type) {
		case string:
			msg += arg.(string)
		default:
			msg += fmt.Sprintf("%v", err)
		}
	}
	return response.Response{
		StatusCode: code.Code,
		StatusMsg:  msg,
	}
}
