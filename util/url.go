package util

import (
	"fmt"
)

type url struct{}

// Url remote链接生成
var Url url

//判断ssl
func (*url) s(ssl bool) string {
	if ssl {
		return "s"
	}

	return ""
}

//url不带协议部分
func (a *url) addr(ip string, port uint, ssl bool) string {
	return fmt.Sprintf("%s://%s:%d", a.s(ssl), ip, port)
}

// Ws websocket连接地址
func (a *url) Ws(ip string, port uint, ssl bool) string {
	return "ws" + a.addr(ip, port, ssl)
}

// Http http连接地址
func (a *url) Http(ip string, port uint, ssl bool) string {
	return "http" + a.addr(ip, port, ssl)
}
