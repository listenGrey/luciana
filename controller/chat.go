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
	uid, err := errHandler.GetCurrentUserID(c)
	if err != nil {
		errHandler.ResponseError(c, code.Unauthorized)
		return
	}

	// 获取聊天列表
	res, err := logic.ChatList(uid)
	if err != nil {
		errHandler.ResponseError(c, code.Busy)
		return
	}

	errHandler.ResponseSuccess(c, *res)
}

// NewChat 创建新对话
func NewChat(c *gin.Context) {
	// 获取当前用户的ID
	uid, err := errHandler.GetCurrentUserID(c)
	if err != nil {
		errHandler.ResponseError(c, code.Unauthorized)
		return
	}
	// 获取新建对话的id和name
	res, err := logic.NewChat(uid)
	if err != nil {
		errHandler.ResponseError(c, code.Busy)
		return
	}

	errHandler.ResponseSuccess(c, *res)
}

// GetChatHandler 查看对话
func GetChatHandler(c *gin.Context) {
	sid := c.Param("id")
	cid, err := strconv.ParseInt(sid, 10, 64)
	if err != nil {
		errHandler.ResponseError(c, code.Busy)
		return
	}
	res, err := logic.GetChat(cid)
	if err != nil {
		errHandler.ResponseError(c, code.Busy)
		return
	}

	errHandler.ResponseSuccess(c, res)
}

// RenameChatHandler 对话重命名
func RenameChatHandler(c *gin.Context) {
	// 获取当前用户的ID
	uid, err := errHandler.GetCurrentUserID(c)
	if err != nil {
		errHandler.ResponseError(c, code.Unauthorized)
		return
	}

	var chat *model.FrontChatList
	if err = c.ShouldBind(&chat); err != nil {
		errHandler.ResponseError(c, code.InvalidParams)
		return
	}

	reChat := &model.Chat{
		Cid:  model.StringToInt64(chat.Cid),
		Uid:  uid,
		Name: chat.Name,
		QAs:  nil,
	}
	err = logic.RenameChat(reChat)
	if err != nil {
		errHandler.ResponseError(c, code.Busy)
		return
	}

	errHandler.ResponseSuccess(c, "修改成功")
}

// DeleteChatHandler 删除对话
func DeleteChatHandler(c *gin.Context) {
	sid := c.Param("id")
	cid, err := strconv.ParseInt(sid, 10, 64)
	if err != nil {
		errHandler.ResponseError(c, code.Busy)
		return
	}

	// 获取当前用户的ID
	uid, err := errHandler.GetCurrentUserID(c)
	if err != nil {
		errHandler.ResponseError(c, code.Unauthorized)
		return
	}

	err = logic.DeleteChat(uid, cid)
	if err != nil {
		errHandler.ResponseError(c, code.Busy)
		return
	}

	errHandler.ResponseSuccess(c, "删除成功")
}

// PromptHandler 发送问题
func PromptHandler(c *gin.Context) {
	var request *model.FrontRequest
	if err := c.ShouldBind(&request); err != nil {
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
