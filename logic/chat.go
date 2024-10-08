package logic

import (
	"github.com/bwmarrin/snowflake"
	"luciana/model"
	"luciana/util"
)

// ChatList 获取聊天列表
func ChatList(id int64) (*[]model.Chat, error) {
	list, err := util.ChatList(id)
	return list, err
}

// NewChat 创建新聊天
func NewChat(uid int64) (*model.Chat, error) {
	// 生成聊天ID
	node, err := snowflake.NewNode(999)
	if err != nil {
		return nil, err
	}
	id := node.Generate()

	// 创建新聊天
	newChat := &model.Chat{
		Cid:  id.Int64(),
		Name: "New Chat",
	}

	// 将新聊天信息发送
	err = util.NewChat(uid, newChat)
	if err != nil {
		return nil, err
	}

	return newChat, nil
}

// GetChat 获取聊天
func GetChat(id int64) (*model.Chat, error) {
	res, err := util.GetChat(id)
	return res, err
}

// RenameChat 聊天重命名
func RenameChat(cid int64, name string) error {
	err := util.RenameChat(cid, name)
	return err
}

// DeleteChat 删除聊天
func DeleteChat(id int64) error {
	err := util.DeleteChat(id)
	return err
}

// Prompt 获取回答
func Prompt(request *model.Request) (string, error) {
	res, err := util.Prompt(request)
	if err != nil {
		return "", err
	}

	return res, nil
}
