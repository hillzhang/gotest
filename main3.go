package main

import (
	"encoding/binary"
	"net"
	"gotest/server/msg"
	"github.com/golang/protobuf/proto"
	"bytes"
	"fmt"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:3563")
	if err != nil {
		panic(err)
	}

	// Hello 消息（JSON 格式）
	// 对应游戏服务器 Hello 消息结构体
	m := new(msg.Test)
	m.Name= proto.String("zhanglinshan")

	data,err := proto.Marshal(m)
	// len + data

	// 默认使用大端序
	buf := new(bytes.Buffer)
	binary.Write(buf,binary.LittleEndian,uint16(len(data)+2))
	binary.Write(buf,binary.LittleEndian,uint16(0))
	binary.Write(buf,binary.LittleEndian,data)

	// 发送消息
	conn.Write(buf.Bytes())

	b := make([]byte,100)

	for {
		fmt.Println("start")
		n,err := conn.Read(b)
		if err != nil {
			fmt.Println(err)
		}
		m := new(msg.Test)
		err = proto.Unmarshal(b[4:n],m)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(m.String())
	}
}