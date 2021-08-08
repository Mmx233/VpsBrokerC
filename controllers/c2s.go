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

func GetSelf() (string, error) {
	_, res, e := tool.HTTP.Get(
		util.Url.Http(),
		util.Url.AuthHeader(),
		nil, nil, true)
	if e != nil {
		return "", e
	}
	if res["code"].(float64) != 0 {
		return "", global.ErrRemoteRefused
	}

	return res["data"].(map[string]interface{})["ip"].(string), nil
}
