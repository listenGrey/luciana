package controller

import (
	"github.com/gin-gonic/gin"
	"luciana/errHandler"
	"luciana/errHandler/code"
	"luciana/logic"
	"luciana/model"
	"strconv"
)

// ChatListHandler 首页，查看聊天列表
func ChatListHandler(c *gin.Context) {
	// 获取当前用户的ID
	id, err := errHandler.GetCurrentUserID(c)
	if err != nil {
		errHandler.ResponseError(c, code.Unauthorized)
		return
	}

	// 获取聊天列表
	res, err := logic.ChatList(id)
	if err != nil {
		errHandler.ResponseError(c, code.Busy)
		return
	}

	errHandler.ResponseSuccess(c, *res)
}

// NewChat 创建新对话
func NewChat(c *gin.Context) {
	// 获取当前用户的ID
	id, err := errHandler.GetCurrentUserID(c)
	if err != nil {
		errHandler.ResponseError(c, code.Unauthorized)
		return
	}
	// 获取新建对话的id和name
	res, err := logic.NewChat(id)
	if err != nil {
		errHandler.ResponseError(c, code.Busy)
		return
	}

	errHandler.ResponseSuccess(c, *res)
}

// GetChatHandler 查看对话
func GetChatHandler(c *gin.Context) {
	sid := c.Param("id")
	id, err := strconv.ParseInt(sid, 10, 64)
	if err != nil {
		errHandler.ResponseError(c, code.Busy)
		return
	}
	res, err := logic.GetChat(id)
	if err != nil {
		errHandler.ResponseError(c, code.Busy)
		return
	}

	errHandler.ResponseSuccess(c, *res)
}

// RenameChatHandler 对话重命名
func RenameChatHandler(c *gin.Context) {
	var chat *model.Chat
	if err := c.ShouldBindJSON(&chat); err != nil {
		errHandler.ResponseError(c, code.InvalidParams)
		return
	}
	cid := chat.Cid
	name := chat.Name
	err := logic.RenameChat(cid, name)
	if err != nil {
		errHandler.ResponseError(c, code.Busy)
		return
	}

	errHandler.ResponseSuccess(c, "修改成功")
}

// DeleteChatHandler 删除对话
func DeleteChatHandler(c *gin.Context) {
	sid := c.Param("id")
	id, err := strconv.ParseInt(sid, 10, 64)
	if err != nil {
		errHandler.ResponseError(c, code.Busy)
		return
	}

	err = logic.DeleteChat(id)
	if err != nil {
		errHandler.ResponseError(c, code.Busy)
		return
	}

	errHandler.ResponseSuccess(c, "删除成功")
}

// PromptHandler 发送问题
func PromptHandler(c *gin.Context) {
	var request *model.Request
	if err := c.ShouldBindJSON(&request); err != nil {
		errHandler.ResponseError(c, code.InvalidParams)
		return
	}

	res, err := logic.Prompt(request)
	if err != nil {
		errHandler.ResponseError(c, code.Busy)
		return
	}

	errHandler.ResponseSuccess(c, res)
}
