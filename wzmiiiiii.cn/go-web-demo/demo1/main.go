package main

import (
	"log"
	"net/http"
)

type helloHandler struct{}

func (*helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello go web"))
}

func welcome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("welcome"))
}

func main() {
	server := http.Server{
		Addr: ":8080",
	}

	http.Handle("/hello", &helloHandler{})

	http.HandleFunc("/welcome", welcome)
	// 上面这行代码:可以写成下面这行代码.
	// http.Handle("/welcome", http.HandlerFunc(welcome))

	log.Println(server.ListenAndServe())
}
