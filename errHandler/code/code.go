package code

type Code int64

const (
	// 可以暴露给外部的错误码
	Success           Code = 1000
	InvalidParams     Code = 1001
	UserExist         Code = 1002
	UserNotExist      Code = 1003
	InvalidPwd        Code = 1004
	Busy              Code = 1005
	InvalidToken      Code = 1006
	InvalidAuthFormat Code = 1007
	InvalidInvitation Code = 1008

	// 内部错误码
	InvalidGenID       Code = 1100
	RegisterERR        Code = 1101
	ConnGrpcServerERR  Code = 1102
	RecvGrpcSerInfoERR Code = 1103
	KafkaSendERR       Code = 1104
	KafkaReceiveERR    Code = 1105
)

var msgFlags = map[Code]string{
	Success:           "成功",
	InvalidParams:     "请求参数错误",
	UserExist:         "用户已存在",
	UserNotExist:      "用户不存在",
	InvalidPwd:        "用户名或密码错误",
	Busy:              "业务繁忙，请稍后重试",
	InvalidToken:      "无效的Token",
	InvalidAuthFormat: "认证格式有误",
	InvalidInvitation: "邀请码错误",

	InvalidGenID:       "生成ID失败",
	RegisterERR:        "用户注册失败",
	ConnGrpcServerERR:  "无法连接到gRpc服务器",
	RecvGrpcSerInfoERR: "从gRpc服务器获取信息失败",
	KafkaSendERR:       "向kafka中发送数据失败",
	KafkaReceiveERR:    "从kafka中获取数据失败",
}

func (c Code) Msg() string {
	msg, ok := msgFlags[c]
	if ok {
		return msg
	}
	return msgFlags[Busy]
}

func (c Code) Code() int64 {
	return int64(c)
}
