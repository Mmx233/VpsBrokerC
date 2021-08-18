package controllers

import (
	"github.com/Mmx233/VpsBrokerC/global"
	"github.com/Mmx233/VpsBrokerC/util"
	"log"
	"os"
)

func init() {
	if util.Try(func() error {
		var e error
		global.Self, e = GetSelf()
		return e
	}, 5, func(e error) {
		log.Println("获取ip失败：\n", e)
	}) != nil {
		log.Println("超出重试次数")
		os.Exit(2)
	}

	go SReceiver()
}
