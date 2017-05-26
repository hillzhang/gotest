package msg

import (
	"github.com/name5566/leaf/network/json"
	//"github.com/name5566/leaf/network/protobuf"
)

// 使用默认的 JSON 消息处理器（默认还提供了 protobuf 消息处理器）
//var Processor = json.NewProcessor()
//var Processor = protobuf.NewProcessor()
var Processor = json.NewProcessor()
func init() {
	// 这里我们注册了一个 JSON 消息 Hello
	//Processor.Register(&Test{})
	Processor.Register(&User{})
	Processor.Register(&Game{})
}

// 一个结构体定义了一个 JSON 消息的格式
// 消息名为 Hello
type User struct {
	Username string
	RoomNum int
	Pid int32
}

type Game struct {
	User *User
	Status int
	Msg string
}