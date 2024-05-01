package util

import (
	"github.com/listenGrey/TmagegRpcPKG/userInfo"
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

	// 发送邮箱
	sendEmail := &userInfo.RegisterEmail{Email: email}
	res, err := client.(userInfo.CheckExistenceClient).RegisterCheck(context.Background(), sendEmail)
	if err != nil {
		return code.StatusRecvGrpcSerInfoERR
	}

	// 获取用户的状态信息
	exist := res.Exist
	info := res.Info

	if info == code.StatusConnDBERR.Code() {
		return code.StatusConnDBERR
	} else if info == code.StatusBusy.Code() {
		return code.StatusBusy
	}

	if exist {
		return code.StatusUserExist
	}

	return code.StatusSuccess
}

// Register 用户注册
func Register(user *model.User) code.Code {
	client := grpc.UserClientServer(grpc.Register)
	if client == code.StatusConnGrpcServerERR {
		return code.StatusConnGrpcServerERR
	}
	sendUser := &userInfo.RegisterForm{
		UserID:   user.UserID,
		Email:    user.Email,
		UserName: user.UserName,
		Password: user.Password,
	}
	res, err := client.(userInfo.RegisterInfoClient).Register(context.Background(), sendUser)
	if err != nil {
		return code.StatusRecvGrpcSerInfoERR
	}

	sta := res.Success
	info := res.Info

	if info == code.StatusConnDBERR.Code() {
		return code.StatusConnDBERR
	} else if info == code.StatusBusy.Code() {
		return code.StatusBusy
	}

	if !sta {
		return code.StatusRegisterERR
	}

	return code.StatusSuccess
}

// LoginCheck 用户登录
func LoginCheck(user *model.User) (info code.Code, userID int64) {
	client := grpc.UserClientServer(grpc.LoginCheck)
	if client == code.StatusConnGrpcServerERR {
		return code.StatusConnGrpcServerERR, 0
	}
	sendUser := &userInfo.LoginForm{
		Email:    user.Email,
		Password: user.Password,
	}
	res, err := client.(userInfo.LoginCheckClient).LoginCheck(context.Background(), sendUser)
	if err != nil {
		return code.StatusRecvGrpcSerInfoERR, 0
	}
	sta := res.Info
	userID = res.UserID
	if sta == code.StatusConnDBERR.Code() {
		return code.StatusConnDBERR, 0
	} else if sta == code.StatusUserNotExist.Code() {
		return code.StatusUserNotExist, 0
	} else if sta == code.StatusInvalidPwd.Code() {
		return code.StatusInvalidPwd, 0
	} else if sta == code.StatusBusy.Code() {
		return code.StatusBusy, 0
	}

	return code.StatusSuccess, userID
}
