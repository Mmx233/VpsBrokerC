package util

import (
	"fmt"
	"log"
)

type event struct{}

var Event event

func (*event) e(name string, content string) {
	log.Println(fmt.Sprintf("[%s] %s", name, content))
}

func (a *event) NewPeer(name string) {
	a.e("peer", "new "+name)
}

func (a *event) LostPeer(name string) {
	a.e("peer", "lost "+name)
}
