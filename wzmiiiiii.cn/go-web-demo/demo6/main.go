package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		// 因为使用define action,所以现在它们对应的模版名分别为:layout,content
		t, _ := template.ParseFiles("static/layout.gohtml", "static/home.gohtml")
		_ = t.ExecuteTemplate(w, "layout", "hello world")
	})

	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		// 因为使用define action,所以现在它们对应的模版名分别为:layout,content
		t, _ := template.ParseFiles("static/layout.gohtml", "static/about.gohtml")
		_ = t.ExecuteTemplate(w, "layout", "hello world")
	})

	http.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request) {
		// 因为使用define action,所以现在它们对应的模版名分别为:layout,content
		t, _ := template.ParseFiles("static/layout.gohtml")
		_ = t.ExecuteTemplate(w, "layout", nil)
	})

	log.Println(http.ListenAndServe(":8080", nil))
}
