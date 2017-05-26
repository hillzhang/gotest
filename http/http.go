package http

import (
	"net/http"
	"log"
	"time"
)

type MyHandler struct {}

type Store struct {
	Host string
	Count int16
	Time time.Time
}

type Server struct {
	Writer http.ResponseWriter
	Request *http.Request
}

func(this *Store)cron(){

}

var Current = make(map[string]*Store,0)

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

