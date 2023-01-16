package middleware

import (
	"douyin/pkg/errcode"
	"douyin/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
	"strings"
)

// JWT token验证中间件
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Query("token")
		if n := strings.Index(c.Request.Header.Get("content-type"), binding.MIMEMultipartPOSTForm); n != -1 {
			token = c.PostForm("token")
		}
		if token == "" {
			c.JSON(http.StatusOK, errcode.NewResponse(errcode.ErrTokenNotExist))
			c.Abort()
			return
		}
		_, err := utils.ParseToken(token)
		if err != nil {
			c.JSON(http.StatusOK, errcode.NewResponse(errcode.ErrAuthorized, err))
			c.Abort()
			return
		}
		c.Next()
	}
}
