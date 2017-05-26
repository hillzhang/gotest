package internal

import (
	"github.com/name5566/leaf/log"
	"github.com/name5566/leaf/gate"
	"reflect"
	"gotest/server/msg"
	"gotest/server/core"
	"fmt"
)

func init() {
	// 向当前模块（game 模块）注册 Hello 消息的消息处理函数 handleHello
	handler(&msg.Game{}, handleHello)
}

func handler(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func handleHello(args []interface{}) {
	// 收到的 Hello 消息
	m := args[0].(*msg.Game)
	// 消息的发送者
	a := args[1].(gate.Agent)
	pid,_ := a.GetProperty("pid")
	// 输出收到的消息的内容
	if m.Status == core.Hold{
		log.Debug("user: %v in game center",pid)
	}

	if m.Status == core.Ready{
		var room *core.Room
		if len(core.GameRoom.Allocated) >0{
			room = core.GameRoom.Allocated[0]
		}else {
			room = core.GameRoom.Rooms.PopBack().(*core.Room)
			core.GameRoom.Allocated = append(core.GameRoom.Allocated,room)
		}


		player,_ :=core.MrgObject.GetPlayer(pid.(int32))
		player.User.RoomNum = room.RoomNum
		core.GameRoom.Allocated[0].Players[room.Len()+1] = player
		if room.Len() ==4 {
			core.GameRoom.Allocated = core.GameRoom.Allocated[1:]
			room.Status = core.Start
			core.GameRoom.Started[room.RoomNum] = room
			room.Broadcase()
		}
		log.Debug("allocate room num: %v,total palyers %v",room.RoomNum,room.Players)
	}


	if m.Status == core.Start{
		fmt.Println("Dddd")
		room := core.GameRoom.GetRoom(m.User.RoomNum)
		room.BroadCast(pid.(int32),m.Msg)
	}

	//log.Debug("user: %v  hello %v",pid, m.Msg)

	// 给发送者回应一个 Hello 消息
	//a.WriteMsg(&msg.Hello{
	//	Name: "hillzhang",
	//})
}