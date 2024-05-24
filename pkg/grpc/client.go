package grpc

import (
	"github.com/listenGrey/lucianagRpcPKG/ask"
	"github.com/listenGrey/lucianagRpcPKG/chat"
	"github.com/listenGrey/lucianagRpcPKG/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// 定义gRpc客户端服务器的类型码

type Service string

const (
	CheckExistence Service = "CheckExistence"
	LoginCheck     Service = "LoginCheck"
	GetChat        Service = "GetChat"
	GetChats       Service = "GetChats"
	SendPrompt     Service = "SendPrompt"
)

func UserClientServer(service Service) (client interface{}) {
	creds1, err := credentials.NewClientTLSFromFile("./pkg/ca/server1", "")
	if err != nil {
		return nil
	}
	creds2, err := credentials.NewClientTLSFromFile("./pkg/ca/server2", "")
	if err != nil {
		return nil
	}
	creds3, err := credentials.NewClientTLSFromFile("./pkg/ca/server3", "")
	if err != nil {
		return nil
	}

	userConn, err := grpc.Dial("localhost:8964", grpc.WithTransportCredentials(creds1)) //server IP
	chatConn, err := grpc.Dial("localhost:8964", grpc.WithTransportCredentials(creds2)) //server IP
	askConn, err := grpc.Dial("localhost:8964", grpc.WithTransportCredentials(creds3))  //server IP
	if err != nil {
		return nil
	}
	switch service {
	case CheckExistence:
		client = user.NewCheckExistClient(userConn)
	case LoginCheck:
		client = user.NewLoginCheckClient(userConn)
	case GetChat:
		client = chat.NewGetChatServiceClient(chatConn)
	case GetChats:
		client = chat.NewGetChatsServiceClient(chatConn)
	case SendPrompt:
		client = ask.NewChatServiceClient(askConn)
	default:
		client = nil
	}
	return client
}
