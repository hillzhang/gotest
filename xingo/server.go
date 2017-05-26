package main


import (
	"github.com/viphxin/xingo/iface"
	"github.com/viphxin/xingo/fserver"
	"github.com/viphxin/xingo/logger"
	"github.com/viphxin/xingo/utils"
	"gotest/xingo/core"
	"gotest/xingo/api"
	"fmt"
)
func DoConnectionMade(fconn iface.Iconnection) {
	logger.Debug("111111111111111111111111")
	p := core.MrgObj.AddPlayer(fconn)
	fconn.SetProperty("pid", p.Pid)
	fmt.Println(p.Pid)
}

func DoConnectionLost(fconn iface.Iconnection) {
	logger.Debug("222222222222222222222222")
	pid, _ := fconn.GetProperty("pid")
	p,err := core.MrgObj.GetPlayer(pid.(int32))
	fmt.Println(p.Pid)
	if err != nil {
		fmt.Println(err)
	}
	//移除玩家
	core.MrgObj.RemovePlayer(pid.(int32))
}

func main() {
	server := fserver.NewServer()
	server.AddRouter(&api.Api{})
	utils.GlobalObject.OnConnectioned = DoConnectionMade
	utils.GlobalObject.OnClosed = DoConnectionLost

	server.Serve()
}



