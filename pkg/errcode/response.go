package errcode

import (
	"douyin/internal/model/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
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

type Resp struct {
	ctx *gin.Context
}

func New(c *gin.Context) *Resp {
	return &Resp{ctx: c}
}

func (r *Resp) RespOK() {
	r.ctx.JSON(http.StatusOK, response.Response{
		StatusCode: OK.Code,
		StatusMsg:  OK.Msg,
	})
}

func (r *Resp) RespFail(code ResponseCode) {
	if code == (ResponseCode{}) {
		code = Fail
	}
	r.ctx.JSON(http.StatusOK, response.Response{
		StatusCode: code.Code,
		StatusMsg:  code.Msg,
	})
}

func (r *Resp) RespFailDetail(code ResponseCode, err ...interface{}) {
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

	r.ctx.JSON(http.StatusOK, response.Response{
		StatusCode: code.Code,
		StatusMsg:  msg,
	})
}

func (r *Resp) RespData(data interface{}) {
	r.ctx.JSON(http.StatusOK, data)
}
