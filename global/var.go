package global

import (
	"github.com/Mmx233/VpsBrokerC/models"
	"github.com/Mmx233/VpsBrokerC/util"
	"github.com/gorilla/websocket"
	"log"
	"os"
	"sync"
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

// Conn c2s连接
var Conn *websocket.Conn

// Neighbor client信息
type Neighbor struct {
	Port  uint
	Delay int64
	Lock  *sync.Mutex
}

// Neighbors c2c列表
var Neighbors = struct {
	Data map[string]*Neighbor
	Lock *sync.RWMutex
}{
	Data: make(map[string]*Neighbor),
	Lock: &sync.RWMutex{},
}

// Self 自身ip
var Self string
