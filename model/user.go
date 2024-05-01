package model

import (
	"encoding/json"
	"errors"
)

type RegisterFrom struct {
	UserName   string `json:"user_name" binding:"required,min=1,max=10"`
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password" binding:"required,min=8,max=16"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
	Invitation string `json:"invitation" binding:"required,min=6,max=6"`
}

func (r *RegisterFrom) UnmarshalJSON(data []byte) (err error) {
	required := struct {
		UserName   string `json:"user_name"`
		Email      string `json:"email"`
		Password   string `json:"password"`
		RePassword string `json:"re_password"`
		Invitation string `json:"invitation"`
	}{}
	err = json.Unmarshal(data, &required)

	if err != nil {
		return
	} else if len(required.UserName) < 1 {
		err = errors.New("用户名不能为空")
	} else if len(required.Email) == 0 {
		err = errors.New("邮箱不能为空")
	} else if len(required.Password) < 8 {
		err = errors.New("密码至少为8位")
	} else if len(required.Password) > 16 {
		err = errors.New("密码最多为16位")
	} else if required.Password != required.RePassword {
		err = errors.New("两次密码不一致")
	} else if len(required.Invitation) != 6 {
		err = errors.New("邀请码为6位")
	} else {
		r.UserName = required.UserName
		r.Email = required.Email
		r.Password = required.Password
		r.RePassword = required.RePassword
		r.Invitation = required.Invitation
	}
	return
}

type LoginForm struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (l *LoginForm) UnmarshalJSON(data []byte) (err error) {
	required := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}
	err = json.Unmarshal(data, &required)

	if err != nil {
		return
	} else if len(required.Email) == 0 {
		err = errors.New("邮箱不能为空")
	} else if len(required.Password) < 8 {
		err = errors.New("密码至少为8位")
	} else if len(required.Password) > 16 {
		err = errors.New("密码最多为16位")
	} else {
		l.Email = required.Email
		l.Password = required.Password
	}
	return
}

type User struct {
	UserID       int64  `json:"user_id" bson:"user_id"`
	UserName     string `json:"user_name" bson:"user_name"`
	Password     string `json:"password" bson:"password"`
	Email        string `json:"email" bson:"email"`
	AccessToken  string
	RefreshToken string
}
