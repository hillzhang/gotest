package main

import (
)
import "fmt"

type User struct {
	Username string
	RoomNum int
	Pid int32
}
func main() {
	a := make([]*User,0)
	a = append(a,nil)
	fmt.Println(a)

}