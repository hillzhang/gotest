package internal

import (
	"github.com/name5566/leaf/gate"
	"fmt"
	"time"
)

func init() {
	skeleton.RegisterChanRPC("NewAgent", rpcNewAgent)
	skeleton.RegisterChanRPC("CloseAgent", rpcCloseAgent)
	go func() {
		for{
			time.Sleep(time.Second * 5)
			for a := range agents {
				fmt.Println("remote_addr",a.RemoteAddr())
			}
		}
	}()
}


var agents = make(map[gate.Agent]struct{})

func rpcNewAgent(args []interface{}) {
	a := args[0].(gate.Agent)
	agents[a] = struct{}{}
}

func rpcCloseAgent(args []interface{}) {
	a := args[0].(gate.Agent)
	delete(agents, a)
}
