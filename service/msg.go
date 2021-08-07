package service

import (
	"container/list"
	"github.com/Mmx233/VpsBrokerC/global"
	"github.com/Mmx233/VpsBrokerC/models"
	"log"
	"sync"
	"time"
)

func init() {
	Msg.Init()
}

type msg struct {
	list         *list.List
	lock         *sync.RWMutex
	innerChannel chan *models.HeartBeat
	Channel      chan *models.HeartBeat
}

var Msg = msg{
	list:         &list.List{},
	lock:         &sync.RWMutex{},
	innerChannel: make(chan *models.HeartBeat),
	Channel:      make(chan *models.HeartBeat),
}

func (a *msg) push(s *models.HeartBeat) {
	a.lock.Lock()
	a.list.PushBack(s)
	a.lock.Unlock()
}

func (a *msg) front() *list.Element {
	a.lock.RLock()
	defer a.lock.RUnlock()
	return a.list.Front()
}

func (a *msg) remove(e *list.Element) {
	a.lock.Lock()
	a.list.Remove(e)
	a.lock.Unlock()
}

func (a *msg) len() int {
	a.lock.RLock()
	defer a.lock.RUnlock()
	return a.list.Len()
}

func (a *msg) Init() {
	go func() {
		for {
			e := <-a.Channel
			a.push(e)
		}
	}()

	go func() {
		for {
			if a.len() == 0 {
				a.innerChannel <- <-a.Channel
			}

			e := a.front()

			for e != nil {
				a.innerChannel <- e.Value.(*models.HeartBeat)
				t := e
				e = e.Next()
				a.remove(t)
			}
		}
	}()

	go func() {
		for {
			e := <-a.innerChannel
			for {
				if err := global.Conn.WriteJSON(e); err != nil {
					log.Println(err)
					time.Sleep(time.Second)
				} else {
					break
				}
			}
		}
	}()
}
