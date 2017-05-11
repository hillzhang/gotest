package core

import (
	"github.com/viphxin/xingo/iface"
)


type Player struct {
	Fconn iface.Iconnection
	Pid   int32
}

func NewPlayer(fconn iface.Iconnection, pid int32) *Player {
	p := &Player{
		Fconn: fconn,
		Pid:   pid,
	}
	return p
}


