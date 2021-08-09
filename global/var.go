package global

import (
	"github.com/Mmx233/VpsBrokerC/models"
	"github.com/Mmx233/VpsBrokerC/util"
	"github.com/gorilla/websocket"
	"log"
	"os"
)

func init() {
	if util.Try(ConnectWs, 5, func(e error) {
		log.Println("连接service失败：/n", e)
	}) != nil {
		log.Println("超出重试次数")
		os.Exit(2)
	}

	if util.Try(func() error {
		return Conn.WriteJSON(&models.VpsInit{
			Name: Config.Settings.Name,
			Port: Config.Settings.Port,
		})
	}, 5, func(e error) {
		log.Println("发送初始化信息失败：/n", e)
	}) != nil {
		log.Println("超出重试次数")
		os.Exit(2)
	}

}

var Conn *websocket.Conn

var Neighbors map[string]uint

var Self string
