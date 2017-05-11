package main

import (
	"github.com/gorilla/rpc/json"
)
type TestEchoACK struct {
	Content string `protobuf:"bytes,1,opt,name=Content,json=content" json:"Content,omitempty"`
}
func main() {
	json.EncodeClientRequest()

}