package controllers

import (
	"github.com/Mmx233/VpsBrokerC/controllers/c2c"
	"github.com/Mmx233/VpsBrokerC/global"
	"github.com/Mmx233/VpsBrokerC/util"
	"github.com/Mmx233/tool"
	"sync"
	"time"
)

func init() {
	go SReceiver()
}

// SReceiver c2s ws连接协程
func SReceiver() {
	util.Event.ServiceConnect()
	for {
		var t map[string]uint
		e := global.Conn.ReadJSON(&t)
		if e != nil {
			util.Event.ServiceDisconnect()
			_ = global.Conn.Close()
			for global.ConnectWs() != nil {
				time.Sleep(time.Second / 2)
			}
			util.Event.ServiceConnect()
			continue
		}

		delete(t, global.Self)
		global.Neighbors.Lock.Lock()
		for k, v := range t {
			if n, ok := global.Neighbors.Data[k]; ok {
				n.Lock.Lock()
				n.Port = v
				n.Lock.Unlock()
			} else {
				global.Neighbors.Data[k] = &global.Neighbor{
					Port:  v,
					Delay: 0,
					Lock:  &sync.Mutex{},
				}
			}
		}
		global.Neighbors.Lock.Unlock()
		controllers.Conn.Renew()
	}
}

func GetSelf() (string, error) {
	_, res, e := tool.HTTP.Get(
		util.Url.Http(global.Config.Remote.Host, global.Config.Remote.Port, global.Config.Remote.SSL)+"/c/self",
		map[string]interface{}{
			"Authorization": global.Config.Remote.AccessKey,
		},
		nil, nil, true)
	if e != nil {
		return "", e
	}
	if res["code"].(float64) != 0 {
		return "", global.ErrRemoteRefused
	}

	return res["data"].(map[string]interface{})["ip"].(string), nil
}
