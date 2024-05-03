package routers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"luciana/controller"
	"luciana/middlewares"
	"net/http"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// 使用CORS中间件
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:8082"} // 允许前端的地址
	config.AllowMethods = []string{"POST", "GET", "DELETE"} // 允许的方法

	r.Use(cors.New(config))
	v1 := r.Group("/api/v1")
	v1.POST("/register", controller.RegisterHandler)
	v1.POST("/login", controller.LoginHandler)
	v1.GET("/refresh_token", controller.RefreshTokenHandler)

	// 中间件
	v1.Use(middlewares.JWTAuthMiddleWare()) //JWT认证
	{
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
