package internal

import (
	"reflect"
	"gotest/server/msg"
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
	//"github.com/golang/protobuf/proto"
	"gotest/server/core"
)

func init() {
	//handleMsg(&msg.Test{},handleLogin)
	handleMsg(&msg.User{},handleLogin)
}

func handleMsg(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func handleLogin(args []interface{}){
	//m := args[0].(*msg.Test)
	m := args[0].(*msg.User)
	// 消息的发送者
	a := args[1].(gate.Agent)

	// 输出收到的消息的内容
	log.Debug("hello %v", m.Username)

	// 给发送者回应一个 Hello 消息
	//a.WriteMsg(&msg.Test{
	//	Name: proto.String("hillzhang"),
	//})

	pid,_:= a.GetProperty("pid")
	player,_ := core.MrgObject.GetPlayer(pid.(int32))
	player.SetUser(m)
	a.WriteMsg(&msg.User{
		Username:m.Username,
		Pid: pid.(int32),
	})
}

