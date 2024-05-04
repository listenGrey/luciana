package controller

import (
	"github.com/gin-gonic/gin"
	"luciana/errHandler"
	"luciana/errHandler/code"
	"luciana/logic"
	"luciana/model"
)

// IndexHandler 首页
func IndexHandler(c *gin.Context) {
	// 获取当前用户的ID
	id, err := errHandler.GetCurrentUserID(c)
	if err != nil {
		errHandler.ResponseError(c, code.InvalidToken)
		return
	}

	// 获取聊天列表
	res, err := logic.GetChatList(id)
	if err != nil {
		errHandler.ResponseError(c, code.Busy)
		return
	}

	errHandler.ResponseSuccess(c, *res)
}

// NewChat 创建新聊天
func NewChat(c *gin.Context) {
	// 获取新建聊天的id和name
	res, err := logic.NewChat()
	if err != nil {
		errHandler.ResponseError(c, code.Busy)
		return
	}

	errHandler.ResponseSuccess(c, *res)
}

// ChatHandler 查看聊天
func ChatHandler(c *gin.Context) {
	var chat *model.Chat
	if err := c.ShouldBindJSON(&chat); err != nil {
		errHandler.ResponseError(c, code.InvalidParams)
		return
	}
	id := chat.Id
	res, err := logic.GetChat(id)
	if err != nil {
		errHandler.ResponseError(c, code.Busy)
		return
	}

	errHandler.ResponseSuccess(c, *res)
}

// RenameHandler 聊天重命名
func RenameHandler(c *gin.Context) {
	var chat *model.Chat
	if err := c.ShouldBindJSON(&chat); err != nil {
		errHandler.ResponseError(c, code.InvalidParams)
		return
	}
	id := chat.Id
	name := chat.Name
	err := logic.RenameChat(id, name)
	if err != nil {
		errHandler.ResponseError(c, code.Busy)
		return
	}

	errHandler.ResponseSuccess(c, "修改成功")
}

// DeleteHandler 删除聊天
func DeleteHandler(c *gin.Context) {
	chat := c.Param("chat")

	err := logic.DeleteChat(chat)
	if err != nil {
		errHandler.ResponseError(c, code.Busy)
		return
	}

	errHandler.ResponseSuccess(c, "删除成功")
}

// RequestHandler 发送问题
func RequestHandler(c *gin.Context) {
	var request *model.Request
	if err := c.ShouldBindJSON(&request); err != nil {
		errHandler.ResponseError(c, code.InvalidParams)
		return
	}

	res, err := logic.Generate(request)
	if err != nil {
		errHandler.ResponseError(c, code.Busy)
		return
	}

	errHandler.ResponseSuccess(c, res)
}
