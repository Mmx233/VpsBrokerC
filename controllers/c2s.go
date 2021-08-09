package controllers

import (
	"github.com/Mmx233/VpsBrokerC/global"
	"github.com/Mmx233/VpsBrokerC/util"
	"github.com/Mmx233/tool"
	"time"
)

func init() {
	go SReceiver()
}

func SReceiver() {
	for {
		var t map[string]uint
		e := global.Conn.ReadJSON(&t)
		global.Neighbors.Lock.Lock()
		global.Neighbors.Data=t
		global.Neighbors.Lock.Unlock()
		if e != nil {
			_ = global.Conn.Close()
			for global.ConnectWs() != nil {
				time.Sleep(time.Second/2)
			}
		}


	}
}

func GetSelf() (string, error) {
	_, res, e := tool.HTTP.Get(
		util.Url.Http()+"/c/self",
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
