package cellnet

import (
	"testing"
	"github.com/davyxu/cellnet"
	"github.com/davyxu/cellnet/proto/gamedef"
	"github.com/davyxu/cellnet/socket"
	"github.com/davyxu/golog"
	"fmt"
)

func Test_Server(t *testing.T){
	Server1()
}

var log *golog.Logger = golog.New("test")

func Server1() {
	queue := cellnet.NewEventQueue()
	evd := socket.NewAcceptor(queue).Start("127.0.0.1:1234")
	socket.RegisterMessage(evd,"gamedef.TestEchoACK", func(content interface{}, ses cellnet.Session) {
		msg := content.(*gamedef.TestEchoACK)
		fmt.Println("recv:",msg.String())
		log.Debugln("server recv:", msg.String())

	})

	queue.StartLoop()
	select {

	}
}