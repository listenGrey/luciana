package model

import (
	"github.com/listenGrey/lucianagRpcPKG/chat"
	"strconv"
)

type ChatList struct {
	Uid   int64      `json:"uid"`
	Chats []ChatInfo `json:"chats"`
}

type ChatInfo struct {
	Cid  int64  `bson:"cid"`
	Name string `bson:"name"`
}

type Chat struct {
	Cid  int64  `json:"cid" binding:"required"`
	Uid  int64  `json:"uid"`
	Name string `json:"name"`
	QAs  []QA   `json:"qas"`
}

type QA struct {
	Request  string `json:"request"`
	Response string `json:"response"`
}

type Request struct {
	Cid    int64  `json:"cid" binding:"required"`
	Prompt string `json:"prompt" binding:"required"`
}

type ResponseContent struct {
	Code int64       `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"` // content为空时不显示
}

func ChatsUnmarshal(c *chat.ChatList) *[]Chat {
	var chats []Chat

	for _, v := range c.GetChats() {
		var ch Chat

		ch.Cid = v.Cid
		ch.Name = v.Name

		chats = append(chats, ch)
	}

	return &chats
}

func ChatUnmarshal(c *chat.Chat) *Chat {
	var ch Chat

	ch.Cid = c.Cid
	ch.Uid = c.Uid
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

type FrontChatList struct {
	Cid  string `json:"cid"`
	Name string `json:"name"`
}

type FrontChat struct {
	QAs []QA `json:"qas"`
}

type FrontRequest struct {
	Cid    string `json:"cid" binding:"required"`
	Prompt string `json:"prompt" binding:"required"`
}

func Int64ToString(id int64) string {
	return strconv.FormatInt(id, 10)
}

func StringToInt64(id string) int64 {
	i, _ := strconv.ParseInt(id, 10, 64)
	return i
}
