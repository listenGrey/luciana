package logic

import (
	"github.com/bwmarrin/snowflake"
	"luciana/model"
	"luciana/util"
)

// ChatList 获取聊天列表
func ChatList(id int64) (*[]model.FrontChatList, error) {
	list, err := util.ChatList(id)
	var res []model.FrontChatList
	for _, v := range *list {
		r := &model.FrontChatList{
			Cid:  model.Int64ToString(v.Cid),
			Name: v.Name,
		}
		res = append(res, *r)
	}
	return &res, err
}

// NewChat 创建新聊天
func NewChat(uid int64) (*model.FrontChatList, error) {
	// 生成聊天ID
	node, err := snowflake.NewNode(999)
	if err != nil {
		return nil, err
	}
	id := node.Generate()

	// 创建新聊天
	newChat := &model.Chat{
		Cid:  id.Int64(),
		Uid:  uid,
		Name: "New Chat",
	}

	// 将新聊天信息发送
	err = util.NewChat(newChat)
	if err != nil {
		return nil, err
	}

	newFrontChat := &model.FrontChatList{
		Cid:  model.Int64ToString(id.Int64()),
		Name: "New Chat",
	}

	return newFrontChat, nil
}

// GetChat 获取聊天
func GetChat(id int64) (model.FrontChat, error) {
	chat, err := util.GetChat(id)
	res := &model.FrontChat{QAs: chat.QAs}
	return *res, err
}

// RenameChat 聊天重命名
func RenameChat(chat *model.Chat) error {
	err := util.RenameChat(chat)
	return err
}

// DeleteChat 删除聊天
func DeleteChat(uid, cid int64) error {
	err := util.DeleteChat(uid, cid)
	return err
}

// Prompt 获取回答
func Prompt(request *model.FrontRequest) (string, error) {
	re := &model.Request{
		Cid:    model.StringToInt64(request.Cid),
		Prompt: request.Prompt,
	}
	res, err := util.Prompt(re)
	if err != nil {
		return "", err
	}

	return res, nil
}
