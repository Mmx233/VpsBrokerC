package global

import (
	"fmt"
	"github.com/Mmx233/VpsBrokerC/util"
	"github.com/gorilla/websocket"
)

// ConnectWs 连接service
func ConnectWs() error {
	var e error
	Conn, _, e = websocket.DefaultDialer.Dial(
		util.Url.Ws(
			Config.Remote.Host, Config.Remote.Port, Config.Remote.SSL)+"?name="+Config.Settings.Name+"&port="+fmt.Sprint(Config.Settings.Port),
		nil)
	return e
}
