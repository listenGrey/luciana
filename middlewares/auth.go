package middlewares

import (
	"github.com/gin-gonic/gin"
	"luciana/errHandler"
	"luciana/errHandler/code"
	"luciana/pkg/jwt"
	"strings"
)

// JWTAuthMiddleWare 基于 JWT 的认证中间件
func JWTAuthMiddleWare() func(c *gin.Context) {
	return func(c *gin.Context) {
		// Token放在Header的Authorization中，使用Bearer开头
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			errHandler.ResponseError(c, code.InvalidAuthFormat)
			c.Abort()
			return
		}
		// 按空格分割
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			errHandler.ResponseError(c, code.InvalidAuthFormat)
			c.Abort()
			return
		}
		// 解析Token
		user, err := jwt.ParseToken(parts[1])
		if err != nil {
			errHandler.ResponseError(c, code.InvalidAuthFormat)
			c.Abort()
			return
		}
		// 将user信息保存到context中
		c.Set(errHandler.ContextUserIDKey, user.UserID)
		c.Next()
		// 后续的处理函数可以用过GetCurrentUserID(c)来获取当前请求的用户信息
	}
}
