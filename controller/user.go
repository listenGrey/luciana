package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"luciana/errHandler"
	"luciana/errHandler/code"
	"luciana/logic"
	"luciana/model"
	"luciana/pkg/jwt"
	"net/http"
	"strings"
)

// RegisterHandler 用户注册业务
func RegisterHandler(c *gin.Context) {
	//获取请求参数，校验数据
	var client *model.RegisterForm
	if err := c.ShouldBind(&client); err != nil {
		// 请求参数错误
		errHandler.ResponseError(c, code.InvalidParams)
		return
	}
	// 用户注册
	info, err := logic.Register(client)
	if err != nil {
		errHandler.ResponseError(c, code.Busy)
		return
	}
	if info == code.UserExist || info == code.InvalidInvitation {
		errHandler.ResponseError(c, info)
		return
	}

	errHandler.ResponseSuccess(c, nil)
}

// LoginHandler 用户登录业务
func LoginHandler(c *gin.Context) {
	//获取请求参数，校验参数
	var user *model.LoginForm
	if err := c.ShouldBind(&user); err != nil {
		// 请求参数错误
		errHandler.ResponseError(c, code.InvalidParams)
		return
	}

	// 用户登录
	curUser, info, err := logic.Login(user)
	if err != nil {
		errHandler.ResponseError(c, code.Busy)
		return
	}
	if info == code.UserNotExist || info == code.InvalidPwd {
		errHandler.ResponseError(c, info)
		return
	}

	//返回响应
	errHandler.ResponseSuccess(c, gin.H{
		"user_id":       fmt.Sprintf("%d", curUser.Uid),
		"user_name":     curUser.UserName,
		"access_token":  curUser.AccessToken,
		"refresh_token": curUser.RefreshToken,
	})
}

// RefreshTokenHandler 刷新accessToken
func RefreshTokenHandler(c *gin.Context) {
	rt := c.Query("refresh_token")

	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		errHandler.ResponseMsg(c, code.Unauthorized, "请求头缺少Token")
		c.Abort()
		return
	}
	// 按空格分割
	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		errHandler.ResponseMsg(c, code.Unauthorized, "Token格式错误")
		c.Abort()
		return
	}
	aToken, rToken, err := jwt.RefreshToken(parts[1], rt)
	if err != nil {
		errHandler.ResponseMsg(c, code.Unauthorized, "刷新Token错误")
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token":  aToken,
		"refresh_token": rToken,
	})
}
