package main

import (
	"fmt"
	_ "github.com/Mmx233/VpsBrokerC/controllers/c2s"
	"github.com/Mmx233/VpsBrokerC/global"
	"github.com/Mmx233/VpsBrokerC/router"
	"log"
)

func main() {
	if e := router.G.Run(":" + fmt.Sprint(global.Config.Settings.Port)); e != nil {
		log.Println(e)
	}
}
