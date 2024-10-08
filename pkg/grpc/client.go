package grpc

import (
	"github.com/listenGrey/lucianagRpcPKG/ask"
	"github.com/listenGrey/lucianagRpcPKG/chat"
	"github.com/listenGrey/lucianagRpcPKG/user"
	"google.golang.org/grpc"
)

// 定义gRpc客户端服务器的类型码

type Service string

const (
	CheckExist Service = "CheckExist"
	Login      Service = "Login"
	Register   Service = "Register"

	ChatList   Service = "ChatList"
	GetChat    Service = "GetChat"
	NewChat    Service = "NewChat"
	RenameChat Service = "RenameChat"
	DeleteChat Service = "DeleteChat"

	Prompt Service = "Prompt"
)

func UserClientServer(service Service) (client interface{}) {
	/*
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
	*/
	userConn, err := grpc.Dial("localhost:8964", grpc.WithInsecure()) //server IP
	chatConn, err := grpc.Dial("localhost:8964", grpc.WithInsecure()) //server IP
	askConn, err := grpc.Dial("localhost:8964", grpc.WithInsecure())  //server IP
	if err != nil {
		return nil
	}
	switch service {
	case CheckExist:
		client = user.NewCheckExistClient(userConn)
	case Login:
		client = user.NewLoginCheckClient(userConn)
	case Register:
		client = user.NewRegisterCheckClient(userConn)
	case ChatList:
		client = chat.NewGetChatListClient(chatConn)
	case GetChat:
		client = chat.NewGetChatClient(chatConn)
	case NewChat:
		client = chat.NewNewChatClient(chatConn)
	case RenameChat:
		client = chat.NewRenameChatClient(chatConn)
	case DeleteChat:
		client = chat.NewDeleteChatClient(chatConn)
	case Prompt:
		client = ask.NewRequestClient(askConn)
	default:
		client = nil
	}
	return client
}
