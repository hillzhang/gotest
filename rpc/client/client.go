package main

import (
	"sys_monitor/toolkits/net"
	"time"
	"fmt"
	"gotest/rpc/common"
)

func main() {
	rpcClient, err := net.JsonRpcClient("tcp","devops.people.lo:8432",time.Second*5)
	if err != nil {
		fmt.Println(err)
	}
	var reply *common.Reply
	err = rpcClient.Call("Transfer.Update","zhanglinshan", &reply)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(reply)
}


