package internal

import (
	"github.com/name5566/leaf/gate"
	"gotest/server/core"
	"log"
	"fmt"
)

func init() {
	skeleton.RegisterChanRPC("NewAgent", rpcNewAgent)
	skeleton.RegisterChanRPC("CloseAgent", rpcCloseAgent)
	//go func() {
	//	time.Sleep(time.Minute)
	//	fmt.Println("dddd")
	//	core.MrgObject.Broadcast()
	//}()
}

//var agents = make(map[gate.Agent]struct{})

func rpcNewAgent(args []interface{}) {
	a := args[0].(gate.Agent)
	p,err := core.MrgObject.AddPlayer(a)
	if err != nil {
		log.Println(err)
	}
	a.SetProperty("pid",p.Pid)
	core.MrgObject.Display()
	//msg := &msg.Game{Msg:"hillzhang"}
	//p.SendMsg(msg)
	//agents[a] = struct{}{}
}

func rpcCloseAgent(args []interface{}) {
	a := args[0].(gate.Agent)
	pid,err := a.GetProperty("pid")
	if err != nil {
		fmt.Println(err)
	}
	p,err := core.MrgObject.GetPlayer(pid.(int32))
	if err != nil {
		fmt.Println(err)
	}
	core.MrgObject.RemovePlayer(p.Pid)

	a.RemoveProperty("pid")
	core.MrgObject.Display()
	//delete(agents, a)
}
