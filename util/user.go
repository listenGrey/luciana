package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/listenGrey/lucianagRpcPKG/user"
	"github.com/segmentio/kafka-go"
	"luciana/errHandler/code"
	"luciana/model"
	"luciana/pkg/grpc"
	"time"

	"context"
)

// CheckExistence 使用gRpc检查邮箱是否存在
func CheckExistence(email string) (code.Code, error) {
	// 创建gRpc客户端
	client := grpc.UserClientServer(grpc.CheckExistence)
	if client == nil {
		return code.Busy, errors.New("gRpc 客户端启动失败")
	}

	// 获取用户的状态信息
	sendEmail := &user.Email{Email: email}
	res, err := client.(user.CheckExistClient).RegisterCheck(context.Background(), sendEmail)
	if err != nil {
		return code.Busy, errors.New("使用 gRpc 获取信息失败")
	}

	// 用户是否已经存在
	if res.Exist {
		return code.UserExist, nil
	}

	return code.Success, nil
}

// Register 使用kafka发送用户注册数据
func Register(u *model.User) error {
	ctx := context.Background()
	// 创建 Kafka 生产者
	writer := &kafka.Writer{
		Addr:                   kafka.TCP("localhost:9092"),
		Topic:                  "register",
		Balancer:               &kafka.Hash{},
		WriteTimeout:           1 * time.Second,
		RequiredAcks:           kafka.RequireNone,
		AllowAutoTopicCreation: false,
	}

	defer writer.Close()

	// 构造消息
	key := []byte(fmt.Sprintf("%d", u.UserID)) // key = id
	value, err := json.Marshal(*u)             // value = data
	if err != nil {
		return err
	}

	// 发送消息
	err = writer.WriteMessages(
		ctx,
		kafka.Message{
			Key:   key,
			Value: value,
		},
	)
	if err != nil {
		return err
	}

	return nil
}

// LoginCheck 使用gRpc发送用户登录数据
func LoginCheck(u *model.User) (code.Code, int64, string, error) {
	// 创建gRpc客户端
	client := grpc.UserClientServer(grpc.LoginCheck)
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

	return code.Success, res.UserId, res.UserName, nil
}
