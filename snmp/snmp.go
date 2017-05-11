package main

import (
	snmp "sys_monitor/gosnmp"
	"time"
	"fmt"
)

func Snmp(){
	smp := &snmp.GoSNMP{
		Target:    "10.100.6.78",
		Port:      161,
		Community: "public",
		Version:   snmp.Version2c,
		Timeout:   time.Duration(5) * time.Second,
		Retries:3,
	}
	err := smp.Connect()

	if err != nil {
		fmt.Println("connect",err)
	}
	defer func() {
		smp.Conn.Close()
		if r := recover();r != nil{
			return
		}
	}()

	result,err := smp.Get([]string{".1.3.6.1.2.1.1.5.0"})
	if err != nil {
		fmt.Println("get",err)
	}
	fmt.Println(string(result.Variables[0].Value.([]byte)))
}

func main() {
	for i := 0; i < 10000; i ++{
		go Snmp()
	}
	time.Sleep(time.Minute)
}