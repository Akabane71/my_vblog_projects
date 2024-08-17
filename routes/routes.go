package routes

import (
	"bluebell/controller"
	"bluebell/logger"
	"bluebell/middlewares"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// 注册路由

func Setup(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode) // gin设置为发布模式
	}

	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true),
		middlewares.RateLimitMiddleware(2*time.Second, 2), // 限流中间件
	)

	// 提供扩展
	v1 := r.Group("/api/v1")

	v1.GET("/ping", controller.Ping)

	// 注册业务路由
	v1.POST("/signup", controller.SignUpHandler)

	// 登录业务路由
	v1.POST("/login", controller.LoginHandler)

	// 应用上中间件
	v1.Use(middlewares.JWTAuthMiddleware()) // JWT中间件

	{
		v1.GET("/community", controller.CommunityHandler)
		v1.GET("/community/:id", controller.CommunityDetailHandler)
		v1.POST("/post/", controller.CreatePostHandler)
		v1.GET("/post/:id", controller.GetPostDetailHandler)
		v1.POST("/posts/", controller.GetPostListHandler)

		// 根据时间或者分数获取帖子列表
		v1.GET("post2", controller.GetPostListHandler2)

		// 投票功能
		v1.POST("vote", controller.PostVoteController)
	}

	r.NoRoute(func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}
