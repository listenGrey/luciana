package util

import (
	"context"
	"errors"
	"fmt"
	"github.com/listenGrey/lucianagRpcPKG/user"
	"luciana/errHandler/code"
	"luciana/model"
	"luciana/pkg/grpc"
)

// CheckExistence 使用gRpc检查邮箱是否存在
func CheckExistence(email string) (code.Code, error) {
	// 创建gRpc客户端
	client := grpc.UserClientServer(grpc.CheckExist)
	if client == nil {
		return code.Busy, errors.New("gRpc 客户端启动失败")
	}

	// 获取用户的状态信息
	sendEmail := &user.Email{Email: email}
	res, err := client.(user.CheckExistClient).CheckExist(context.Background(), sendEmail)
	if err != nil {
		return code.Busy, errors.New("使用 gRpc 获取信息失败")
	}

	// 用户是否已经存在
	if res.Exist {
		return code.UserExist, nil
	}

	return code.Success, nil
}

// Register 发送用户注册数据
func Register(u *model.User) error {
	// 创建gRpc客户端
	client := grpc.UserClientServer(grpc.Register)
	if client == nil {
		return errors.New("gRpc 客户端启动失败")
	}
	sendUser := &user.RegisterFrom{
		Uid:        u.Uid,
		Email:      u.Email,
		Name:       u.UserName,
		Password:   u.Password,
		Invitation: u.Invitation,
	}
	_, err := client.(user.RegisterCheckClient).RegisterCheck(context.Background(), sendUser)
	if err != nil {
		return fmt.Errorf("注册失败：%s", err)
	}
	return nil
}

// LoginCheck 使用gRpc发送用户登录数据
func LoginCheck(u *model.User) (code.Code, int64, string, error) {
	// 创建gRpc客户端
	client := grpc.UserClientServer(grpc.Login)
	if client == nil {
		return code.Busy, 0, "", errors.New("gRpc 客户端启动失败")
	}
	sendUser := &user.LoginForm{
		Email:    u.Email,
		Password: u.Password,
	}
	res, err := client.(user.LoginCheckClient).LoginCheck(context.Background(), sendUser)
	if err != nil {
		return code.Busy, 0, "", errors.New("使用 gRpc 获取信息失败")
	}

	// 用户是否存在
	if !res.Exist {
		return code.UserNotExist, 0, "", nil
	}

	// 密码是否正确
	if !res.Success {
		return code.InvalidPwd, 0, "", nil
	}

	return code.Success, res.Uid, res.UserName, nil
}
