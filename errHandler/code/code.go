package code

type Code int64

// 可以暴露给外部的错误码
const (
	Success           Code = 1000
	InvalidParams     Code = 1001
	UserExist         Code = 1002
	UserNotExist      Code = 1003
	InvalidPwd        Code = 1004
	Busy              Code = 1005
	InvalidToken      Code = 1006
	InvalidAuthFormat Code = 1007
	InvalidInvitation Code = 1008
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
