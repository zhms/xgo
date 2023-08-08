package main

import (
	"github.com/zhms/xgo/xgo"
)

var db *xgo.XDb = &xgo.XDb{}
var redis *xgo.XRedis = &xgo.XRedis{}
var http *xgo.XHttp = &xgo.XHttp{}
var fullatuh = `
{
	"系统首页": { "查" : 1},
	"系统管理": {
		"运营商管理": { "查": 1,"增": 1,"删": 1,"改": 1},
		"渠道管理":   { "查": 1,"增": 1,"删": 1,"改": 1},
		"系统设置":   { "查": 1,"增": 1,"删": 1,"改": 1},
		"账号管理":   { "查": 1,"增": 1,"删": 1,"改": 1},
		"角色管理":   { "查": 1,"增": 1,"删": 1,"改": 1},
		"登录日志":   { "查": 1},
		"操作日志":   { "查": 1}
	}
}
`

func main() {
	xgo.Init()

	db.Init("server.db")
	redis.Init("server.redis")
	http.Init("server.http", redis)
	http.InitWs("/api/ws")
	xgo.AdminInit(http, db, redis, fullatuh)

	http.OnPostNoAuth("/test", func(ctx *xgo.XHttpContent) {
		ctx.RespOK("123")
	})

	xgo.Run()
}
