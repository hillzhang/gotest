package main

import (
"io"
"log"
"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello world!")
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	err := http.ListenAndServeTLS(":8080", "cert.pem", "key.pem", nil)
	//err := http.ListenAndServe(":8080",nil)
	if err != nil {
		log.Fatal("ListenAndServeTLS:", err.Error())
	}
	mux := http.NewServeMux()
	mux.Handle()
	server := http.Server{}
	server.l

}