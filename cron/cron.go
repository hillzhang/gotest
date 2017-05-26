package cron

import (
	"toolbox/cron"
	"fmt"
)

type User struct {
	_Username string "user name"
	Passord string "user password"
}

func Cron(){
	con := cron.New()
	con.AddFunc("@every 3s", func() {
		fmt.Println("hahah")
	})
	con.Start()

	select {

	}

}
