package main

import (
	"github.com/viphxin/xingo/fserver"
	"github.com/viphxin/xingo/iface"
	"github.com/viphxin/xingo/logger"
	"github.com/viphxin/xingo/utils"
	"gotest/xing_test/api"
	"gotest/xing_test/core"

	_ "net/http"
	_ "net/http/pprof"
	_ "runtime/pprof"
	_ "time"
	"fmt"
)

func DoConnectionMade(fconn iface.Iconnection) {
	logger.Debug("111111111111111111111111")
	p, _ := core.WorldMgrObj.AddPlayer(fconn)
	fmt.Println("player:",p.Pid,"in")
	fconn.SetProperty("pid", p.Pid)
}

func DoConnectionLost(fconn iface.Iconnection) {
	logger.Debug("222222222222222222222222")
	pid, _ := fconn.GetProperty("pid")
	p, _ := core.WorldMgrObj.GetPlayer(pid.(int32))
	//移除玩家
	core.WorldMgrObj.RemovePlayer(pid.(int32))
	fmt.Println("player:",p.Pid,"out")
}

func main() {
	server := fserver.NewServer()
	TestRouterObj := &api.TestRouter{}
	server.AddRouter(TestRouterObj)
	utils.GlobalObject.OnConnectioned = DoConnectionMade
	utils.GlobalObject.OnClosed = DoConnectionLost
	server.Serve()
}





