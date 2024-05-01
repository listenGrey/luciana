package grpc

import (
	"github.com/listenGrey/TmagegRpcPKG/userInfo"
	"google.golang.org/grpc"
	"luciana/errHandler/code"
)

// 定义gRpc客户端服务器的类型码

type GrpcService string

const (
	CheckExistence GrpcService = "CheckExistence"
	Register       GrpcService = "Register"
	LoginCheck     GrpcService = "LoginCheck"
)

func UserClientServer(funcCode GrpcService) (client interface{}) {
	conn, err := grpc.Dial("localhost:8964", grpc.WithInsecure()) //server IP
	if err != nil {
		return code.StatusConnGrpcServerERR
	}
	switch funcCode {
	case CheckExistence:
		client = userInfo.NewCheckExistenceClient(conn)
	case Register:
		client = userInfo.NewRegisterInfoClient(conn)
	case LoginCheck:
		client = userInfo.NewLoginCheckClient(conn)
	default:
		client = nil
	}
	return client
}
