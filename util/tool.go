package util

import (
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

func Try(e func() error, retry uint, retryMsg func(e error)) error {
	var num uint
	var err error
	for {
		num++
		err = e()
		if err != nil {
			if num < retry {
				retryMsg(err)
				time.Sleep(time.Second / 2)
				continue
			} else {
				return err
			}
		} else {
			break
		}
	}

	return nil
}

var Upper = websocket.Upgrader{
	HandshakeTimeout: time.Minute,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}
