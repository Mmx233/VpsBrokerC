package router

import (
	"github.com/Mmx233/VpsBrokerC/global"
	"github.com/gorilla/websocket"
	"log"
)

func connect() error {
	var e error
	if global.Conn, _, e = websocket.DefaultDialer.Dial(
		func() string {
			return global.Config.RemoteAddr
		}(), nil); e != nil {
		log.Println(e)
	}
	return e
}
