package core

import (
	"github.com/name5566/leaf/gate"
	"sync"
	"github.com/qiniu/errors"
	"fmt"
	"gotest/server/msg"
)

type iface interface {
	AddPlayer(agent gate.Agent)(*Player,error)
	RemovePlayer(pid int32) (*Player,error)
	GetPlayer(pid int32) (*Player,error)
	Online_player_len() int
	Offline_player_len() int
	InGame_player_len() int
}

type Mrg struct {
	PlayerGenNum int32
	Players map[int32] *Player
	Online_players map[int32] *Player
	Offline_players map[int32] *Player
	InGame_players map[int32] *Player
	sync.RWMutex
}

var MrgObject *Mrg

func init() {
	MrgObject = &Mrg{PlayerGenNum:0, Players:make(map[int32]*Player,0)}
}

func (this *Mrg) AddPlayer(agent gate.Agent)(*Player,error) {
	this.Lock()

	this.PlayerGenNum += 1
	p := NewPlayer(agent,this.PlayerGenNum)

	if _,ok := this.Players[p.Pid];ok{
		return nil,errors.New("current user already in player map")
	}
	this.Players[p.Pid] = p
	defer this.Unlock()
	return p,nil
}

func (this *Mrg) GetPlayer(pid int32) (*Player,error){
	this.Lock()
	defer this.Unlock()
	if _,ok := this.Players[pid];ok{
		return this.Players[pid],nil
	}else {
		return nil,errors.New("not found")
	}
}

func (this *Mrg) RemovePlayer(pid int32) {
	this.Lock()
	defer this.Unlock()
	delete(this.Players,pid)
}

func (this *Mrg) Len()int{
	return len(this.Players)
}

func (this *Mrg) Online_player_len()int{
	return len(this.Online_players)
}

func (this *Mrg) Offline_player_len()int{
	return len(this.Offline_players)
}
func (this *Mrg) InGame_player_len()int{
	return len(this.InGame_players)
}

func (this *Mrg) Display(){
	fmt.Println("current users:",this.Len())
	fmt.Print("current users:")
	for key,_ := range this.Players{
		fmt.Print(key," ")
	}
}

func (this *Mrg) Broadcast(){
	for _,val := range this.Players{
		val.Agent.WriteMsg(&msg.Game{Msg:"you too"})
	}
}