package global

import (
	"github.com/Mmx233/VpsBrokerC/models"
	"github.com/gorilla/websocket"
	"log"
)

func ConnectWs() error {
	var e error
	if Conn, _, e = websocket.DefaultDialer.Dial(
		func() string {
			return Config.RemoteAddr
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
