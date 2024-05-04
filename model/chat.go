package model

import "github.com/listenGrey/lucianagRpcPKG/chat"

type Request struct {
	Id     int64  `json:"id" binding:"required"`
	Prompt string `json:"prompt" binding:"required"`
}

type QA struct {
	Request  string `json:"request"`
	Response string `json:"response"`
}

type Chat struct {
	Id   int64  `json:"id" binding:"required"`
	Name string `json:"name"`
	QAs  []QA   `json:"qa_s"`
}

type ResponseContent struct {
	Code    int64       `json:"code"`
	Msg     string      `json:"msg"`
	Content interface{} `json:"content,omitempty"` // content为空时不显示
}

func ChatsUnmarshal(c *chat.Chats) *[]Chat {
	var chats []Chat

	for _, v := range c.Chats {
		var ch Chat

		ch.Id = v.Id
		ch.Name = v.Name

		chats = append(chats, ch)
	}

	return &chats
}

func ChatUnmarshal(c *chat.Chat) *Chat {
	var ch Chat

	ch.Id = c.Id
	ch.Name = c.Name

	var qas []QA
	for _, v := range c.Qas {
		var qa QA

		qa.Request = v.Request
		qa.Response = v.Response

		qas = append(qas, qa)
	}
	ch.QAs = qas

	return &ch
}
