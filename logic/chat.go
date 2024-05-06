package logic

import (
	"github.com/bwmarrin/snowflake"
	"github.com/tmc/langchaingo/llms/ollama"
	"luciana/model"
	"luciana/util"

	"context"
)

// GetChatList 获取聊天列表
func GetChatList(id int64) (*[]model.Chat, error) {
	list, err := util.GetChatList(id)
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
		ChatID: id.Int64(),
		Name:   "New Chat",
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
func RenameChat(id int64, name string) error {
	err := util.RenameChat(id, name)
	return err
}

// DeleteChat 删除聊天
func DeleteChat(id string) error {
	err := util.DeleteChat(id)
	return err
}

// Generate 生成回答
func Generate(request *model.Request) (string, error) {
	llm, err := ollama.New(ollama.WithModel("llama3"))
	if err != nil {
		return "", err
	}

	res, err := llm.Call(context.Background(), request.Prompt)
	if err != nil {
		return "", err
	}

	qa := &model.QA{
		Request:  request.Prompt,
		Response: res,
	}
	err = util.SendQA(qa, request.Id)
	if err != nil {
		return "", err
	}

	return res, nil
}
