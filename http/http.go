package http

import (
	"net/http"
	"log"
)

type MyHandler struct {}

func (*MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request){
	m[r.URL.String()](w,r)
}


func Hello(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("hello"))
}


var m map[string]func(http.ResponseWriter, *http.Request)

func Http(){
	m = make(map[string]func(http.ResponseWriter, *http.Request),0)
	m["/hello"] = Hello
	server := &http.Server{
		Addr:":9999",
		Handler:&MyHandler{},
	}
	log.Fatalln(server.ListenAndServe())
}

