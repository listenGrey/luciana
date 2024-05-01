package model

type Request struct {
	Prompt string `json:"prompt" binding:"required"`
}

type ResponseContent struct {
	Code    int64       `json:"code"`
	Msg     string      `json:"msg"`
	Content interface{} `json:"content,omitempty"` // content为空时不显示
}
