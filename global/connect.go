package global

import (
	"github.com/Mmx233/VpsBrokerC/util"
	"github.com/gorilla/websocket"
	"log"
)

// ConnectWs 连接service
func ConnectWs() error {
	var e error
	if Conn, _, e = websocket.DefaultDialer.Dial(
		util.Url.Ws(Config.Remote.Ip, Config.Remote.Port, Config.Remote.SSL), nil); e != nil {
		log.Println(e)
	}
	return e
}
