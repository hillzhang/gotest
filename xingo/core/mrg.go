package core

import (
	"github.com/viphxin/xingo/iface"
	"sync"
	"github.com/qiniu/errors"
)
var MrgObj *Mrg

func init() {
	MrgObj = &Mrg{
		PlayerNumGen:    0,
		Players:         make(map[int32]*Player),
	}
}
type Mrg struct {
	PlayerNumGen int32
	Players map[int32]*Player
	sync.RWMutex
}


func (this *Mrg) GetPlayerNumGen()int32{
	return this.PlayerNumGen
}

func (this *Mrg) AddPlayer(fconn iface.Iconnection) *Player{
	this.Lock()
	defer this.Unlock()
	this.PlayerNumGen += 1
	Player := NewPlayer(fconn, this.PlayerNumGen)
	this.Players[Player.Pid] = Player

	return Player
}

func (this *Mrg) GetPlayer(pid int32) (*Player,error){
	this.Lock()
	defer this.Unlock()
	p,ok := this.Players[pid]
	if ok{
		return p, nil
	}else {
		return nil, errors.New("no player in the game")
	}
}

func (this *Mrg) RemovePlayer(pid int32 ){
	this.Lock()
	defer this.Unlock()
	delete(this.Players,pid)
}

