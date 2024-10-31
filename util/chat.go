package util

import (
	"context"
	"errors"
	"github.com/listenGrey/lucianagRpcPKG/chat"
	"luciana/model"
	"luciana/pkg/grpc"
)

// ChatList 获取聊天列表
func ChatList(uid int64) (*[]model.Chat, error) {
	// 创建gRpc客户端
	client := grpc.UserClientServer(grpc.ChatList)
	if client == nil {
		return nil, errors.New("gRpc 客户端启动失败")
	}

	// 获取聊天列表
	sendId := &chat.ID{Id: uid}
	chats, err := client.(chat.GetChatListClient).GetChatList(context.Background(), sendId)
	if err != nil {
		return nil, errors.New("使用 gRpc 获取信息失败")
	}

	res := model.ChatsUnmarshal(chats)

	return res, nil
}

// GetChat 获取聊天信息
func GetChat(cid int64) (*model.Chat, error) {
	// 创建gRpc客户端
	client := grpc.UserClientServer(grpc.GetChat)
	if client == nil {
		return nil, errors.New("gRpc 客户端启动失败")
	}

	// 获取聊天列表
	sendId := &chat.ID{Id: cid}
	c, err := client.(chat.GetChatClient).GetChat(context.Background(), sendId)
	if err != nil {
		return nil, errors.New("使用 gRpc 获取信息失败")
	}

	res := model.ChatUnmarshal(c)

	return res, nil
}

// NewChat 创建新聊天
func NewChat(new *model.Chat) error {
	// 创建gRpc客户端
	client := grpc.UserClientServer(grpc.NewChat)
	if client == nil {
		return errors.New("gRpc 客户端启动失败")
	}

	// 创建新聊天
	c := &chat.Chat{
		Cid:  new.Cid,
		Uid:  new.Uid,
		Name: new.Name,
		Qas:  nil,
	}
	_, err := client.(chat.NewChatClient).NewChat(context.Background(), c)
	if err != nil {
		return errors.New("使用 gRpc 获取信息失败")
	}

	return nil
}

// RenameChat 修改聊天名
func RenameChat(ch *model.Chat) error {
	// 创建gRpc客户端
	client := grpc.UserClientServer(grpc.RenameChat)
	if client == nil {
		return errors.New("gRpc 客户端启动失败")
	}

	// 修改聊天名
	c := &chat.Chat{
		Cid:  ch.Cid,
		Uid:  ch.Uid,
		Name: ch.Name,
		Qas:  nil,
	}
	_, err := client.(chat.RenameChatClient).RenameChat(context.Background(), c)
	if err != nil {
		return errors.New("使用 gRpc 获取信息失败")
	}

	return nil
}

// DeleteChat 删除聊天
func DeleteChat(uid, cid int64) error {
	// 创建gRpc客户端
	client := grpc.UserClientServer(grpc.DeleteChat)
	if client == nil {
		return errors.New("gRpc 客户端启动失败")
	}

	// 删除聊天
	c := &chat.Chat{
		Cid:  cid,
		Uid:  uid,
		Name: "",
		Qas:  nil,
	}
	_, err := client.(chat.DeleteChatClient).DeleteChat(context.Background(), c)
	if err != nil {
		return errors.New("使用 gRpc 获取信息失败")
	}

	return nil
}
