package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"time"
)

func main() {
	for i := 0; i<50; i ++{
		go func() {
			time.Sleep(time.Second * 3)
			client := &http.Client{}


			request,err := http.NewRequest("GET","http://10.100.6.71:9999/hello",nil)
			if err != nil {
				fmt.Println(err)
			}
			response,err := client.Do(request)
			if err != nil {
				fmt.Println(err)
			}
			//defer response.Body.Close()
			result,err := ioutil.ReadAll(response.Body)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(string(result))
		}()
	}
	time.Sleep(time.Minute *2)
}