package internal

import (
	"douyin/internal/controller"
	"douyin/internal/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {

	//interactApi := controller.ControllerGroup.InteractController
	//socializeApi := controller.ControllerGroup.SocializeController

	//公开的，不需要鉴权的路由
	PublicRouter := r.Group("/douyin")
	{
		PublicRouter.GET("/feed/", controller.Feed)
		PublicRouter.POST("/user/login/", controller.Login)
	}

	//私密的，需要减权的路由
	PrivateRouter := r.Group("douyin")
	PrivateRouter.Use(middleware.JWT())
	{
		PrivateRouter.GET("/user/", controller.GetUserInfo)
	}
}
