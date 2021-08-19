package service

import (
	"github.com/Mmx233/VpsBrokerC/global"
	"github.com/Mmx233/VpsBrokerC/models"
	"sync"
	"time"
)

type stat struct {
	downList map[string]bool
	lock     *sync.RWMutex
}

var Stat = stat{
	downList: make(map[string]bool),
	lock:     &sync.RWMutex{},
}

func (a *stat) Down(ip string) {
	a.lock.Lock()
	defer a.lock.Unlock()

	if v, ok := a.downList[ip]; ok && v {
		return
	}

	a.downList[ip] = true
	Msg.Channel <- &models.HeartBeat{
		Type:     "down",
		TargetIp: ip,
		Time:     time.Now().Add(-time.Second * 5).UnixNano(),
	}
}

func (a *stat) Up(ip string, t int64) {
	a.lock.Lock()
	defer a.lock.Unlock()

	global.Neighbors.Lock.RLock()
	n := global.Neighbors.Data[ip]
	global.Neighbors.Lock.RUnlock()
	n.Lock.Lock()
	n.Delay = time.Now().UnixNano() - t
	n.Lock.Unlock()

	if v, ok := a.downList[ip]; ok && !v {
		return
	} else if !ok {
		a.downList[ip] = false
		return
	}

	a.downList[ip] = false
	Msg.Channel <- &models.HeartBeat{
		Type:     "up",
		TargetIp: ip,
		Time:     t,
	}
}
