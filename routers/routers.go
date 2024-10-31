package routers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"luciana/controller"
	"luciana/errHandler"
	"luciana/errHandler/code"
	"luciana/middlewares"
	"net/http"
	"time"
)

func SetupRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	// 配置CORS中间件
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173"}        // 允许前端的地址
	config.AllowMethods = []string{"POST", "GET", "DELETE", "PUT"} // 允许的方法
	config.AllowHeaders = []string{"Authorization", "Content-Type"}
	config.ExposeHeaders = []string{"Content-Length"}
	config.AllowCredentials = true
	config.MaxAge = 12 * time.Hour

	// 使用zap
	r.Use(
		cors.New(config),
		//logger.GinLogger(),
		//logger.GinRecovery(false),
		//Recovery 中间件会 recover掉项目可能出现的panic，并使用zap记录相关日志
		//middlewares.RateLimitMiddleware(2*time.Second, 40), // 每两秒钟添加十个令牌  全局限流
	)

	v1 := r.Group("/api/v1")
	v1.POST("/register", controller.RegisterHandler)
	v1.POST("/login", controller.LoginHandler)
	v1.GET("/refresh_token", controller.RefreshTokenHandler)

	// 中间件
	v1.Use(middlewares.JWTAuthMiddleWare()) //JWT认证
	{
		v1.GET("/index", controller.ChatListHandler)
		v1.GET("/chat/:id", controller.GetChatHandler)
		v1.POST("/new_chat", controller.NewChat)
		v1.PUT("/rename", controller.RenameChatHandler)
		v1.DELETE("/delete/:id", controller.DeleteChatHandler)
		v1.POST("/prompt", controller.PromptHandler)
		v1.GET("/ping", func(c *gin.Context) {
			c.String(http.StatusOK, "pong")
		})

	}
	r.NoRoute(func(c *gin.Context) {
		errHandler.ResponseError(c, code.NotFound)
	})
	return r
}
