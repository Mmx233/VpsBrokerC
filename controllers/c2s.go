package controllers

import "github.com/Mmx233/VpsBrokerC/global"

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

func GetSelf() {

}
