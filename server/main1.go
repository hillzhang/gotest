package main

import (
	"encoding/binary"
	"net"
	"bufio"
	"os"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)
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


var(
	wchan = make(chan []byte,100)
	rchan = make(chan []byte,100)
)
func main() {
	for i := 0;i <101;i ++{
		time.Sleep(time.Millisecond * 1)
		go func() {

			conn, err := net.Dial("tcp", "127.0.0.1:3563")
			if err != nil {
				panic(err)
			}

			var user =new(User)
			var game =new(Game)


			// Hello 消息（JSON 格式）
			// 对应游戏服务器 Hello 消息结构体

			// 发送消息
			go func() {
				for {
					select {
					case data := <- wchan:
						_,err:= conn.Write(data)
						if err != nil {
							fmt.Println(err)
							os.Exit(1)
						}
					}
				}
			}()

			go func() {
				defer func() {
					if r := recover();r !=nil{
						fmt.Println(r)
					}
				}()
				for{
					var m map[string]interface{}
					buffer := make([]byte,100)
					n,err := conn.Read(buffer)
					if err != nil {
						fmt.Println(err)
					}
					err = json.Unmarshal(buffer[2:n],&m)
					if err != nil {
						fmt.Println(err)
					}
					if value,ok := m["User"];ok{
						mp := value.(map[string]interface{})
						user.Username = mp["Username"].(string)
						user.Pid = int32(mp["Pid"].(float64))
						user.RoomNum = int(mp["RoomNum"].(float64))
						fmt.Println(user)
					}

					if value,ok := m["Game"];ok{
						mp := value.(map[string]interface{})
						game.Status = int(mp["Status"].(float64))
						game.Msg = mp["Msg"].(string)
						u := mp["User"].(map[string]interface{})
						user.RoomNum = int(u["RoomNum"].(float64))
						game.User = user
						fmt.Println(game)
					}
				}
			}()

			for {
				reader := bufio.NewReader(os.Stdin)
				da,_,_ := reader.ReadLine()

				parse_line := strings.Fields(string(da))
				if parse_line[0] == "login"{
					m1 := make(map[string]User,0)
					m1["User"] = User{Username:string(parse_line[1])}

					data,err := json.Marshal(m1)
					if err != nil {
						fmt.Println(err)
					}
					// len + data
					m := make([]byte, 2+len(data))

					// 默认使用大端序
					binary.BigEndian.PutUint16(m, uint16(len(data)))

					copy(m[2:], data)
					wchan <- m
				}

				if parse_line[0] == "ready"{
					m1 := make(map[string]Game,0)
					m1["Game"] = Game{User:user,Status:1,Msg:""}
					data,err := json.Marshal(m1)
					if err != nil {
						fmt.Println(err)
					}
					// len + data
					m := make([]byte, 2+len(data))

					// 默认使用大端序
					binary.BigEndian.PutUint16(m, uint16(len(data)))

					copy(m[2:], data)
					wchan <- m
				}
				if parse_line[0] == "start"{
					m1 := make(map[string]Game,0)
					m1["Game"] = Game{User:user,Status:2,Msg:string(parse_line[1])}
					data,err := json.Marshal(m1)
					if err != nil {
						fmt.Println(err)
					}
					// len + data
					m := make([]byte, 2+len(data))

					// 默认使用大端序
					binary.BigEndian.PutUint16(m, uint16(len(data)))

					copy(m[2:], data)
					wchan <- m
				}

			}
		}()
	}

	time.Sleep(time.Minute)
}