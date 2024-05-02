package grpc

import (
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
)

func UserClientServer(service Service) (client interface{}) {
	conn, err := grpc.Dial("localhost:8964", grpc.WithInsecure()) //server IP
	if err != nil {
		return code.StatusConnGrpcServerERR
	}
	switch service {
	case CheckExistence:
		client = user.NewCheckExistenceClient(conn)
	case Register:
		client = user.NewRegisterInfoClient(conn)
	case LoginCheck:
		client = user.NewLoginCheckClient(conn)
	default:
		client = nil
	}
	return client
}
