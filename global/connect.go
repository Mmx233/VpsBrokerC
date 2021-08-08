package global

import (
	"github.com/Mmx233/VpsBrokerC/models"
	"github.com/Mmx233/VpsBrokerC/util"
	"github.com/gorilla/websocket"
	"log"
)

func ConnectWs() error {
	var e error
	if Conn, _, e = websocket.DefaultDialer.Dial(
		util.Url.Ws(), nil); e != nil {
		log.Println(e)
	}
	return e
}

func init() {
	for ConnectWs() != nil {
	}
	for Conn.WriteJSON(&models.VpsInit{
		Name: Config.Settings.Name,
		Port: Config.Settings.Port,
	}) != nil {
	}
}
