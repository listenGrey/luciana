package grpc

import (
	"crypto/tls"
	"crypto/x509"
	"github.com/listenGrey/lucianagRpcPKG/chat"
	"github.com/listenGrey/lucianagRpcPKG/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
)

// 定义gRpc客户端服务器的类型码

type Service string

const (
	CheckExistence Service = "CheckExistence"
	LoginCheck     Service = "LoginCheck"
	GetChat        Service = "GetChat"
	GetChats       Service = "GetChats"
)

func UserClientServer(service Service) (client interface{}) {
	clientCert, err := tls.LoadX509KeyPair("client.crt", "client.key") //签名和证书的位置
	if err != nil {
		return nil
	}

	certPool := x509.NewCertPool()
	caCert, err := ioutil.ReadFile("ca.crt") // 读取根证书
	if err != nil {
		return nil
	}

	if ok := certPool.AppendCertsFromPEM(caCert); !ok {
		return nil
	}

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{clientCert},
		RootCAs:      certPool,
		ServerName:   "your_server_hostname", // 配置
	}

	userConn, err := grpc.Dial("localhost:8964", grpc.WithTransportCredentials(credentials.NewTLS(tlsConfig))) //server IP
	chatConn, err := grpc.Dial("localhost:8964", grpc.WithTransportCredentials(credentials.NewTLS(tlsConfig))) //server IP
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
	default:
		client = nil
	}
	return client
}
