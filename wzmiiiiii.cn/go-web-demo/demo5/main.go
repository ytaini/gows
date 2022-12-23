package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	templates := loadTemplates()

	server := http.Server{Addr: ":8080"}

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fileName := request.URL.Path[1:]
		if t := templates.Lookup(fileName); t != nil {
			if err := t.Execute(writer, nil); err != nil {
				log.Println(err)
			}
		} else {
			http.NotFound(writer, request)
		}
	})

	http.Handle("/css/", http.FileServer(http.Dir("wwwroot"))) //到wwwroot中找css文件夹
	http.Handle("/img/", http.FileServer(http.Dir("wwwroot"))) //到wwwroot中找img文件夹

	log.Println(server.ListenAndServe())
}

func loadTemplates() *template.Template {
	return template.Must(template.New("").ParseGlob("templates/*.html"))
}
