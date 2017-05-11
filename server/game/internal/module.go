package internal

import (
	"github.com/name5566/leaf/module"
	"github.com/name5566/leaf/gate"
	"gotest/server/base"
)

var (
	skeleton = base.NewSkeleton()
	ChanRPC  = skeleton.ChanRPCServer
	Agents = make(map[gate.Agent]struct{})
)

type Module struct {
	*module.Skeleton
}

func (m *Module) OnInit() {
	m.Skeleton = skeleton
}

func (m *Module) OnDestroy() {

}
