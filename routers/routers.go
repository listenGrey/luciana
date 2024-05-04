package routers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"luciana/controller"
	"luciana/middlewares"
	"net/http"
)

func SetupRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	// 配置CORS中间件
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:8082"} // 允许前端的地址
	config.AllowMethods = []string{"POST", "GET", "DELETE"} // 允许的方法

	// 使用zap
	r.Use(
		cors.New(config),
		//logger.GinLogger(),
		//logger.GinRecovery(false),
		// Recovery 中间件会 recover掉项目可能出现的panic，并使用zap记录相关日志
		//middlewares.RateLimitMiddleware(2*time.Second, 40), // 每两秒钟添加十个令牌  全局限流
	)

	v1 := r.Group("/api/v1")
	v1.POST("/register", controller.RegisterHandler)
	v1.POST("/login", controller.LoginHandler)
	v1.GET("/refresh_token", controller.RefreshTokenHandler)

	// 中间件
	v1.Use(middlewares.JWTAuthMiddleWare()) //JWT认证
	{
		v1.GET("/index", controller.IndexHandler)
		v1.POST("/new_chat", controller.NewChat)
		v1.PUT("/rename", controller.RenameHandler)
		v1.DELETE("/delete/:chat", controller.DeleteHandler)
		v1.POST("/chat", controller.ChatHandler)
		v1.POST("/request", controller.RequestHandler)

		v1.GET("/ping", func(c *gin.Context) {
			c.String(http.StatusOK, "pong")
		})
	}
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}
