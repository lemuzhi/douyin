package internal

import (
	"douyin/internal/controller"
	"douyin/internal/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
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
		//PublicRouter.GET("/feed/", controller.Feed)
		PublicRouter.GET("/feed/", controller.GetFeedList)
		PublicRouter.POST("/user/login/", controller.Login)
		PublicRouter.POST("/user/register/", controller.Register)
	}

	//私密的，需要鉴权的路由
	PrivateRouter := r.Group("/douyin")
	PrivateRouter.Use(middleware.JWT())
	{
		PrivateRouter.GET("/user/", controller.GetUserInfo)
		PrivateRouter.GET("/publish/list/", controller.PublishList)
		PrivateRouter.POST("/publish/action/", controller.PublishAction)

		// 互动-点赞相关
		PrivateRouter.POST("/favorite/action/", controller.FavoriteAction)
		PrivateRouter.GET("/favorite/list/", controller.FavoriteListAction)
		// 互动-评论相关
		PrivateRouter.POST("/comment/action/", controller.CommentAction)
		PrivateRouter.GET("/comment/list/", controller.CommentListAction)

		PrivateRouter.POST("/relation/action/", controller.RelationAction)
		PrivateRouter.GET("/relation/follow/list/", controller.FollowList)

	}
}
