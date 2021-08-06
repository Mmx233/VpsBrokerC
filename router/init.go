package router

import (
	"github.com/Mmx233/VpsBrokerC/global"
	"github.com/Mmx233/VpsBrokerC/models"
)

func init() {
	for connect() != nil {
	}
	for global.Conn.WriteJSON(&models.VpsInit{
		Name: global.Config.Name,
		Port: global.Config.Port,
	}) != nil {
	}
}
