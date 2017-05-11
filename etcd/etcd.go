package etcd

import (
	"log"
	"time"
	"context"
	"github.com/coreos/etcd/client"
	"fmt"
)

func Etcd() {
	cfg := client.Config{
		Endpoints:               []string{"http://127.0.0.1:2379"},
		Transport:               client.DefaultTransport,
		// set timeout per request to fail fast when the target endpoint is unavailable
		HeaderTimeoutPerRequest: time.Second,
	}
	c, err := client.New(cfg)
	if err != nil {
		log.Fatal(err)
	}
	kapi := client.NewKeysAPI(c)
	// set "/foo" key with "bar" value
	//_,err = kapi.Set(context.Background(),"monitor","",&client.SetOptions{Dir:true})
	//if err != nil {
	//	fmt.Println(err)
	//}
	//_,err = kapi.Set(context.Background(),"/monitor/addr3","beijing",nil)
	//if err != nil {
	//	fmt.Println(err)
	//}

	//resp,err := kapi.Get(context.Background(),"/monitor",nil)
	//sort.Sort(resp.Node.Nodes)
	//for _,node := range resp.Node.Nodes{
	//	fmt.Println(node.Key,node.Value)
	//}
	//
	//resp,err := kapi.Get(context.Background(),"/monitor",nil)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//for _,node := range resp.Node.Nodes{
	//	fmt.Println(node.Key)
	//}

	//resp,err:= kapi.Set(context.Background(),"test","value",&client.SetOptions{TTL:time.Second * 60})
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(resp.Action,resp.PrevNode)


	watcher := kapi.Watcher("/monitor",&client.WatcherOptions{Recursive:true})

	for{
		resp,err := watcher.Next(context.Background())
		if err != nil {
			fmt.Println(err)
		}
		switch resp.Action {
		case "set":
			fmt.Println("set",resp.Node)
		case "get":
			fmt.Println("get",resp.Node)
		case "delete":
			fmt.Println("delete",resp.Node)
		case "expire":
			fmt.Println("expire",resp.Node)
		}
	}
}