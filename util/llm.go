package util

import (
	"context"
	"errors"
	"github.com/listenGrey/lucianagRpcPKG/ask"
	"luciana/model"
	"luciana/pkg/grpc"
)

// Prompt 使用gRpc发送问题，接收回答
func Prompt(prompt *model.Request) (string, error) {
	// 创建gRpc客户端
	client := grpc.UserClientServer(grpc.Prompt)
	if client == nil {
		return "", errors.New("gRpc 客户端启动失败")
	}

	// 获取回答
	request := &ask.Prompt{
		Cid:    prompt.Cid,
		Prompt: prompt.Prompt,
	}

	res, err := client.(ask.RequestClient).Request(context.Background(), request)
	if err != nil {
		return "", errors.New("使用 gRpc 获取信息失败")
	}

	return res.Response, nil
}
