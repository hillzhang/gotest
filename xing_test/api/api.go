package api

import (
	"github.com/viphxin/xingo/fnet"
	"github.com/viphxin/xingo/logger"
	"github.com/viphxin/xingo/utils"
	"gotest/xingo_demo/pb"
	"fmt"
	"github.com/golang/protobuf/proto"
)

type TestRouter struct {
}

func (this *TestRouter) Api_0(request *fnet.PkgAll) {
	logger.Debug("call Api_0")
	// request.Fconn.SendBuff(0, nil)
	data := pb.SyncPid{}
	err := proto.Unmarshal(request.Pdata.Data,&data)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data.String())
	data1 := &pb.SyncPid{}
	data1.Pid = int32(32)
	packdata, err := utils.GlobalObject.Protoc.GetDataPack().Pack(1, data1)
	if err == nil{
		request.Fconn.Send(packdata)
	}else{
		logger.Error("pack data error")
	}
}

