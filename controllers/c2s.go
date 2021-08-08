package controllers

import (
	"github.com/Mmx233/VpsBrokerC/global"
	"github.com/Mmx233/VpsBrokerC/util"
	"github.com/Mmx233/tool"
)

func init() {
	go SReceiver()
}

func SReceiver() {
	for {
		e := global.Conn.ReadJSON(&global.Neighbors)
		if e != nil {
			_ = global.Conn.Close()
			for global.ConnectWs() != nil {
			}
		}
	}
}

func GetSelf() error {
	tool.HTTP.Get(
		util.Url.Http(),
	)
}
