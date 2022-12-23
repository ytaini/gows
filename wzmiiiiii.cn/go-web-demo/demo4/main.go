package main

import (
	"log"
	"net/http"
	"text/template"
)

func main() {
	server := http.Server{Addr: ":8080"}

	http.HandleFunc("/process", func(writer http.ResponseWriter, request *http.Request) {
		var err error
		defer func() {
			if err != nil {
				http.NotFound(writer, request)
			}
		}()
		t, err := template.ParseFiles("tmpl.gohtml")
		if err != nil {
			return
		}
		if err = t.Execute(writer, "<h1>hello world!!!</h1>"); err != nil {
			return
		}
	})
	log.Println(server.ListenAndServe())
}
