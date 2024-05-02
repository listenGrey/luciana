package util

import (
	"github.com/listenGrey/lucianagRpcPKG/user"
	"luciana/errHandler/code"
	"luciana/model"
	"luciana/pkg/grpc"

	"context"
)

// CheckExistence 检查邮箱是否存在
func CheckExistence(email string) code.Code {
	// 创建gRpc客户端
	client := grpc.UserClientServer(grpc.CheckExistence)
	if client == code.StatusConnGrpcServerERR {
		return code.StatusConnGrpcServerERR
	}

	// 获取用户的状态信息
	sendEmail := &user.RegisterEmail{Email: email}
	res, err := client.(user.CheckExistenceClient).RegisterCheck(context.Background(), sendEmail)
	if err != nil {
		return code.StatusRecvGrpcSerInfoERR
	}

	// 获取用户信息节点的状态
	if res.ServerErr {
		return code.StatusRecvGrpcSerInfoERR
	}

	// 用户是否已经存在
	if res.Exist {
		return code.StatusUserExist
	}

	return code.StatusSuccess
}

// Register 用户注册
func Register(u *model.User) code.Code {
	// 创建gRpc客户端
	client := grpc.UserClientServer(grpc.Register)
	if client == code.StatusConnGrpcServerERR {
		return code.StatusConnGrpcServerERR
	}
	sendUser := &user.RegisterForm{
		UserID:   u.UserID,
		Email:    u.Email,
		UserName: u.UserName,
		Password: u.Password,
	}
	res, err := client.(user.RegisterInfoClient).Register(context.Background(), sendUser)
	if err != nil {
		return code.StatusRecvGrpcSerInfoERR
	}

	// 获取用户信息节点的状态
	if res.ServerErr {
		return code.StatusRecvGrpcSerInfoERR
	}

	// 用户注册情况
	if !res.Success {
		return code.StatusRegisterERR
	}

	return code.StatusSuccess
}

// LoginCheck 用户登录
func LoginCheck(u *model.User) (info code.Code, userID int64) {
	// 创建gRpc客户端
	client := grpc.UserClientServer(grpc.LoginCheck)
	if client == code.StatusConnGrpcServerERR {
		return code.StatusConnGrpcServerERR, 0
	}
	sendUser := &user.LoginForm{
		Email:    u.Email,
		Password: u.Password,
	}
	res, err := client.(user.LoginCheckClient).LoginCheck(context.Background(), sendUser)
	if err != nil {
		return code.StatusRecvGrpcSerInfoERR, 0
	}

	// 获取用户信息节点的状态
	if res.ServerErr {
		return code.StatusRecvGrpcSerInfoERR, 0
	}

	// 用户是否存在
	if !res.Exist {
		return code.StatusUserNotExist, 0
	}

	// 密码是否正确
	if !res.Success {
		return code.StatusInvalidPwd, 0
	}

	return code.StatusSuccess, userID
}
