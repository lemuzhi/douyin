package internal

import (
	"douyin/internal/controller"
	"douyin/internal/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter(r *gin.Engine) {

	//interactApi := controller.ControllerGroup.InteractController
	//socializeApi := controller.ControllerGroup.SocializeController

	// 为 multipart forms 设置较低的内存限制 (默认是 32 MiB)
	//r.MaxMultipartMemory = 15 << 20 // 8 MiB

	//静态资源
	r.StaticFS("/public", http.Dir("public"))

	//公开的，不需要鉴权的路由
	PublicRouter := r.Group("/douyin")
	{
		PublicRouter.GET("/feed/", controller.Feed)
		PublicRouter.POST("/user/login/", controller.Login)
		PublicRouter.POST("/user/register/", controller.Register)
	}

	//私密的，需要减权的路由
	PrivateRouter := r.Group("/douyin")
	PrivateRouter.Use(middleware.JWT())
	{
		PrivateRouter.GET("/user/", controller.GetUserInfo)
		PrivateRouter.GET("/publish/list/", controller.PublishList)
		PrivateRouter.POST("/publish/action/", controller.PublishAction)
	}
}
