package util

import (
	"fmt"
	"github.com/Mmx233/VpsBrokerC/global"
)

type url struct{}

// Url remote链接生成
var Url url

//判断ssl
func (*url) s() string {
	if global.Config.Remote.SSL {
		return "s"
	}

	return ""
}

//url不带协议部分
func (a *url) addr() string {
	return fmt.Sprintf("%s://%s:%d", a.s(), global.Config.Remote.Ip, global.Config.Remote.Port)
}

// Ws websocket连接地址
func (a *url) Ws() string {
	return "ws" + a.addr() + "/c/"
}

// Http http连接地址
func (a *url) Http() string {
	return "http" + a.addr()
}

// AuthHeader 鉴权头
func (a *url) AuthHeader() map[string]interface{} {
	return map[string]interface{}{
		"Authorization": global.Config.Remote.AccessKey,
	}
}
