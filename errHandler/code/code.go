package code

type Code int64

// 可以暴露给外部的错误码
const (
	Success           Code = 200
	InvalidParams     Code = 400
	Unauthorized      Code = 401
	InvalidInvitation Code = 403
	NotFound          Code = 404
	InvalidAuthFormat Code = 422
	Busy              Code = 503
	UserExist         Code = 1001
	UserNotExist      Code = 1002
	InvalidPwd        Code = 1003
)

var msgFlags = map[Code]string{
	Success:           "成功",
	InvalidParams:     "请求参数错误",
	Unauthorized:      "需要用户认证",
	InvalidInvitation: "邀请码错误",
	NotFound:          "资源不存在",
	InvalidAuthFormat: "认证格式有误",
	Busy:              "业务繁忙，请稍后重试",
	UserExist:         "用户已存在",
	UserNotExist:      "用户不存在",
	InvalidPwd:        "用户名或密码错误",
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
