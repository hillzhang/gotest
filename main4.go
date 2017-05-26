package main

import (
	"fmt"
	"gotest/cron"
)



func main() {
	user := cron.User{Passord:"hhh"}
	fmt.Println(user.Passord)
}