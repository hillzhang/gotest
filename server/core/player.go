package core

import (
	"github.com/name5566/leaf/gate"
	"gotest/server/msg"
)

type Player struct {
	Agent gate.Agent
	User *msg.User
	Pid int32
}

func NewPlayer(agent gate.Agent,pid int32) *Player{
	return &Player{Agent:agent,Pid:pid}
}

func (this *Player) SendMsg(msg interface{}){
	this.Agent.WriteMsg(msg)
}

func (this *Player) SetUser(user *msg.User){
	this.User = user
}