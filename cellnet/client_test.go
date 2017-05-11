package cellnet

import (
	"testing"
	"github.com/davyxu/cellnet"
	"github.com/davyxu/cellnet/proto/gamedef"
	"github.com/davyxu/cellnet/socket"
	"github.com/davyxu/golog"
	"time"
)

func Test_client(t *testing.T){
	Client()
}

var log *golog.Logger = golog.New("test")

func Client(){
	queue := cellnet.NewEventQueue()

	evd := socket.NewConnector(queue).Start("127.0.0.1:1234")
	go queue.StartLoop()

	socket.RegisterMessage(evd, "gamedef.SessionConnected", func(content interface{}, ses cellnet.Session) {
		go func() {
			time.Sleep(time.Second * 5)
			ses.Close()
		}()
		for{
			ses.Send(&gamedef.TestEchoACK{
				Content: "hello",
			})
			time.Sleep(time.Second * 2)
		}
	})
	time.Sleep(time.Second *3)

	select {

	}

}

