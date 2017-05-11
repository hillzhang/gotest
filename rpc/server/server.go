package main


import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"fmt"
)
type MetricValue struct {
	Endpoint  string      `json:"endpoint"`
	Metric    string      `json:"metric"`
	Value     interface{} `json:"value"`
	Step      int64       `json:"step"`
	Type      string      `json:"counterType"`
	Tags      string      `json:"tags"`
	Timestamp int64       `json:"timestamp"`
}
func main() {

	addr := "127.0.0.1:1112"
	tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		log.Fatalf("net.ResolveTCPAddr fail: %s", err)
	}

	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		log.Fatalf("listen %s fail: %s", addr, err)
	} else {
		log.Println("rpc listening", addr)
	}

	server := rpc.NewServer()
	server.Register(new(Transfer))

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("listener.Accept occur error:", err)
			continue
		}
		// go rpc.ServeConn(conn)
		go server.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}

type Transfer int


func (this *Transfer) Update(name string, reply *string)error{
	fmt.Println(name)
	return nil
}
