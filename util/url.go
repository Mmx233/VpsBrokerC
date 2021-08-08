package util

import (
	"fmt"
	"github.com/Mmx233/VpsBrokerC/global"
)

type url struct{}

var Url url

func (*url) s() string {
	if global.Config.Remote.SSL {
		return "s"
	}

	return ""
}

func (a *url) addr() string {
	return fmt.Sprintf("%s://%s:%d", a.s(), global.Config.Remote.Ip, global.Config.Remote.Port)
}

func (a *url) Ws() string {
	return "ws" + a.addr()
}

func (a *url) Http() string {
	return "http" + a.addr()
}
