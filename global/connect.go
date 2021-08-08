package global

import (
	"fmt"
	"github.com/Mmx233/VpsBrokerC/models"
	"github.com/gorilla/websocket"
	"log"
)

func ConnectWs() error {
	var e error
	if Conn, _, e = websocket.DefaultDialer.Dial(
		func() string {
			t := "ws"
			if Config.RemoteSSL {
				t += "s"
			}
			return t + "://" + Config.RemoteIp + ":" + fmt.Sprint(Config.RemotePort)
		}(), nil); e != nil {
		log.Println(e)
	}
	return e
}

func init() {
	for ConnectWs() != nil {
	}
	for Conn.WriteJSON(&models.VpsInit{
		Name: Config.Name,
		Port: Config.Port,
	}) != nil {
	}
}
