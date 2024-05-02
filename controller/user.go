package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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
	var client *model.RegisterFrom
	if err := c.ShouldBindJSON(&client); err != nil {
		//判断 err 是否为 validator 类型
		errs, ok := err.(validator.ValidationErrors)
		if ok {
			// 翻译错误
			errHandler.ResponseMsg(c, code.StatusInvalidParams, errs.Translate(trans))
			return
		}
		// 请求参数错误
		errHandler.ResponseError(c, code.StatusInvalidParams)
		return
	}
	// 用户注册
	if info := logic.Register(client); info != code.StatusSuccess {
		if info == code.StatusUserExist {
			errHandler.ResponseError(c, code.StatusUserExist)
			return
		} else if info == code.StatusInvalidInvitation {
			errHandler.ResponseError(c, code.StatusInvalidInvitation)
			return
		}
		errHandler.ResponseError(c, code.StatusBusy)
		return
	}

	errHandler.ResponseSuccess(c, nil)
}

// LoginHandler 用户登录业务
func LoginHandler(c *gin.Context) {
	//获取请求参数，校验参数
	var user *model.LoginForm
	if err := c.ShouldBindJSON(&user); err != nil {
		//判断 err 是否为 validator 类型
		errs, ok := err.(validator.ValidationErrors)
		if ok {
			// 翻译错误
			errHandler.ResponseMsg(c, code.StatusInvalidParams, errs.Translate(trans))
			return
		}
		// 请求参数错误
		errHandler.ResponseError(c, code.StatusInvalidParams)
		return
	}

	// 用户登录
	curUser, info := logic.Login(user)
	if info != code.StatusSuccess {
		if info == code.StatusUserNotExist || info == code.StatusInvalidPwd {
			errHandler.ResponseError(c, info)
			return
		}
		errHandler.ResponseError(c, code.StatusBusy)
		return
	}

	//返回响应
	errHandler.ResponseSuccess(c, gin.H{
		"user_id":       fmt.Sprintf("%d", curUser.UserID),
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
		errHandler.ResponseMsg(c, code.StatusInvalidToken, "请求头缺少Token")
		c.Abort()
		return
	}
	// 按空格分割
	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		errHandler.ResponseMsg(c, code.StatusInvalidToken, "Token格式错误")
		c.Abort()
		return
	}
	aToken, rToken, err := jwt.RefreshToken(parts[1], rt)
	if err != nil {
		errHandler.ResponseMsg(c, code.StatusInvalidToken, "刷新Token错误")
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token":  aToken,
		"refresh_token": rToken,
	})
}
