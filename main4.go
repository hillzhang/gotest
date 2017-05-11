package main

import (
	"os/exec"
	"github.com/qiniu/log"
	"fmt"
)
func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func main() {
	cmd := exec.Command("netstat","|grep snmp|wc -l")
	count ,err := cmd.CombinedOutput()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(count)
}
