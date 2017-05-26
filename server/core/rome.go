package core

import (
	"toolbox/container/list"
	"gotest/server/msg"
	"fmt"
)

type Rooms struct {
	TotalRoom int
	Rooms *list.SafeList
	Allocated []*Room
	Started map[int]*Room
}
var GameRoom *Rooms

func init() {
	rooms := list.NewSafeList()
	for i := 0; i < 100; i ++{
		room := &Room{
			RoomNum:i,
			Players: make(map[int]*Player,0),
			Status: Hold,
		}
		rooms.PushFront(room)

	}
	GameRoom = &Rooms{TotalRoom:100,Rooms:rooms,Started:make(map[int]*Room,0)}
}

func (this *Rooms)GetTotalRoom()int{
	return this.TotalRoom
}
func (this *Rooms) GetRoom(key int) *Room{
	return this.Started[key]
}

const (
	Hold = iota
	Ready
	Start
	End
)

type Room struct {
	RoomNum int
	Players map[int]*Player
	Status int
}

func (this *Room) Len()int{
	return len(this.Players)
}

func (this *Room)Broadcase(){
	for _,player := range this.Players{
		player.SendMsg(&msg.Game{User:player.User,Status:Start,Msg:"started"})
	}
}
func (this *Room) BroadCast(extPid int32,sg string){
	for _,player := range this.Players{
		if player.Pid == extPid{continue}
		message := &msg.Game{User:player.User,Status:Start, Msg:fmt.Sprintf("pid %v：%v",extPid,sg)}
		player.SendMsg(message)
	}
}


