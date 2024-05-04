package grpc

import (
	"github.com/listenGrey/lucianagRpcPKG/chat"
	"github.com/listenGrey/lucianagRpcPKG/user"
	"google.golang.org/grpc"
	"luciana/errHandler/code"
)

// 定义gRpc客户端服务器的类型码

type Service string

const (
	CheckExistence Service = "CheckExistence"
	Register       Service = "Register"
	LoginCheck     Service = "LoginCheck"
	GetChat        Service = "GetChat"
	GetChats       Service = "GetChats"
)

func UserClientServer(service Service) (client interface{}) {
	conn, err := grpc.Dial("localhost:8964", grpc.WithInsecure()) //server IP
	if err != nil {
		return code.ConnGrpcServerERR
	}
	switch service {
	case CheckExistence:
		client = user.NewCheckExistenceClient(conn)
	case Register:
		client = user.NewRegisterInfoClient(conn)
	case LoginCheck:
		client = user.NewLoginCheckClient(conn)
	case GetChat:
		client = chat.NewGetChatServiceClient(conn)
	case GetChats:
		client = chat.NewGetChatsServiceClient(conn)
	default:
		client = nil
	}
	return client
}
