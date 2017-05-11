package snmp

import (
	"testing"
	"time"
)

func TestSnmp(t *testing.T) {
	for i := 0; i < 10; i ++{
		go Snmp()
	}
	time.Sleep(time.Minute)

}

