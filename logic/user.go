package logic

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/bwmarrin/snowflake"
	"luciana/errHandler/code"
	"luciana/model"
	"luciana/pkg/jwt"
	"luciana/util"
)

// encryptPwd 将密码加密
func encryptPwd(pwdByte []byte) (res string) {
	hashedPassword := md5.Sum(pwdByte)
	return hex.EncodeToString(hashedPassword[:])
}

// Register 用户注册
func Register(client *model.RegisterForm) (code.Code, error) {
	// 判断邀请码
	if client.Invitation != "ae86se" {
		return code.InvalidInvitation, nil
	}
	// 将注册邮箱通过gRpc发送到用户信息节点去判断用户是否存在
	existence, err := util.CheckExistence(client.Email)
	if err != nil {
		return code.Busy, err
	}

	if existence != code.Success {
		return existence, nil
	}

	// 生成用户ID，并对密码加密
	node, err := snowflake.NewNode(1)
	if err != nil {
		return code.Busy, err
	}
	userId := node.Generate()
	pwdByte := []byte(client.Password)
	userPwd := encryptPwd(pwdByte)

	// 创建一个用户
	user := &model.User{
		Uid:        userId.Int64(),
		Email:      client.Email,
		UserName:   client.UserName,
		Password:   userPwd,
		Invitation: client.Invitation,
	}

	// 发送用户信息
	err = util.Register(user)
	if err != nil {
		return code.Busy, err
	}

	return code.Success, nil
}

// Login 用户登录
func Login(form *model.LoginForm) (*model.User, code.Code, error) {
	// 对密码加密
	pwdByte := []byte(form.Password)
	userPwd := encryptPwd(pwdByte)

	user := &model.User{
		Email:    form.Email,
		Password: userPwd,
	}

	// 将登录信息通过gRpc发送到用户信息节点去判断用户和密码是否正确
	info, userID, userName, err := util.LoginCheck(user)
	if err != nil {
		return nil, code.Busy, err
	}
	if info != code.Success {
		return nil, info, nil
	}
	user.Uid = userID
	user.UserName = userName

	// 生成JWT
	aToken, rToken, err := jwt.GenToken(user.Uid, user.UserName)
	if err != nil {
		return nil, code.Busy, err
	}
	user.AccessToken = aToken
	user.RefreshToken = rToken

	return user, code.Success, nil
}
