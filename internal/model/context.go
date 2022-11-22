package model

import "github.com/gogf/gf/v2/net/ghttp"

type Context struct {
	Session *ghttp.Session
	User    *ContextUser
}

type ContextUser struct {
	Id       uint
	Passport string
	Nickname string
}
