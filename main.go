package main

import (
	"github.com/Mmx233/VpsBrokerC/router"
	"log"
)

func main() {
	if e := router.G.Run(); e != nil {
		log.Println(e)
	}
}
