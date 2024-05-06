package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/listenGrey/lucianagRpcPKG/chat"
	"github.com/segmentio/kafka-go"
	"luciana/model"
	"luciana/pkg/grpc"
	"time"

	"context"
)

// GetChatList 使用gRpc获取聊天列表
func GetChatList(uid int64) (*[]model.Chat, error) {
	// 创建gRpc客户端
	client := grpc.UserClientServer(grpc.GetChats)
	if client == nil {
		return nil, errors.New("gRpc 客户端启动失败")
	}

	// 获取聊天列表
	sendId := &chat.ID{Id: uid}
	chats, err := client.(chat.GetChatsServiceClient).GetChats(context.Background(), sendId)
	if err != nil {
		return nil, errors.New("使用 gRpc 获取信息失败")
	}

	res := model.ChatsUnmarshal(chats)

	return res, nil
}

// NewChat 使用kafka发送新聊天信息
func NewChat(uid int64, new *model.Chat) error {
	ctx := context.Background()
	// 创建 Kafka 生产者
	writer := &kafka.Writer{
		Addr:                   kafka.TCP("localhost:9092"),
		Topic:                  "new_chat",
		Balancer:               &kafka.Hash{},
		WriteTimeout:           1 * time.Second,
		RequiredAcks:           kafka.RequireNone,
		AllowAutoTopicCreation: false,
	}

	defer writer.Close()

	// 构造消息
	key := []byte(fmt.Sprintf("%d", uid)) // key = uid
	value, err := json.Marshal(new)       // value = data
	if err != nil {
		return err
	}

	// 发送消息
	err = writer.WriteMessages(
		ctx,
		kafka.Message{
			Key:   key,
			Value: value,
		},
	)
	if err != nil {
		return err
	}

	return nil
}

// GetChat 使用gRpc来获取聊天信息
func GetChat(cid int64) (*model.Chat, error) {
	// 创建gRpc客户端
	client := grpc.UserClientServer(grpc.GetChat)
	if client == nil {
		return nil, errors.New("gRpc 客户端启动失败")
	}

	// 获取聊天列表
	sendId := &chat.ID{Id: cid}
	c, err := client.(chat.GetChatServiceClient).GetChat(context.Background(), sendId)
	if err != nil {
		return nil, errors.New("使用 gRpc 获取信息失败")
	}

	res := model.ChatUnmarshal(c)

	return res, nil
}

// RenameChat 使用kafka发送修改聊天名
func RenameChat(cid int64, name string) error {
	ctx := context.Background()
	// 创建 Kafka 生产者
	writer := &kafka.Writer{
		Addr:                   kafka.TCP("localhost:9092"),
		Topic:                  "rename",
		Balancer:               &kafka.Hash{},
		WriteTimeout:           1 * time.Second,
		RequiredAcks:           kafka.RequireNone,
		AllowAutoTopicCreation: false,
	}

	defer writer.Close()

	// 构造消息
	key := []byte(fmt.Sprintf("%d", cid)) // key = cid
	value := []byte(name)                 // value = name

	// 发送消息
	err := writer.WriteMessages(
		ctx,
		kafka.Message{
			Key:   key,
			Value: value,
		},
	)
	if err != nil {
		return err
	}

	return nil
}

// DeleteChat 使用kafka发送删除聊天
func DeleteChat(cid string) error {
	ctx := context.Background()
	// 创建 Kafka 生产者
	writer := &kafka.Writer{
		Addr:                   kafka.TCP("localhost:9092"),
		Topic:                  "delete",
		Balancer:               &kafka.Hash{},
		WriteTimeout:           1 * time.Second,
		RequiredAcks:           kafka.RequireNone,
		AllowAutoTopicCreation: false,
	}

	defer writer.Close()

	// 构造消息
	key := []byte(cid) // key = cid
	var value []byte   // value = nil

	// 发送消息
	err := writer.WriteMessages(
		ctx,
		kafka.Message{
			Key:   key,
			Value: value,
		},
	)
	if err != nil {
		return err
	}

	return nil
}

// SendQA 使用kafka发送QA
func SendQA(qa *model.QA, cid int64) error {
	ctx := context.Background()
	// 创建 Kafka 生产者
	writer := &kafka.Writer{
		Addr:                   kafka.TCP("localhost:9092"),
		Topic:                  "send_qa",
		Balancer:               &kafka.Hash{},
		WriteTimeout:           1 * time.Second,
		RequiredAcks:           kafka.RequireNone,
		AllowAutoTopicCreation: false,
	}

	defer writer.Close()

	// 构造消息
	key := []byte(fmt.Sprintf("%d", cid)) // key = cid
	value, err := json.Marshal(*qa)       // value = qa
	if err != nil {
		return err
	}

	// 发送消息
	err = writer.WriteMessages(
		ctx,
		kafka.Message{
			Key:   key,
			Value: value,
		},
	)
	if err != nil {
		return err
	}

	return nil
}
