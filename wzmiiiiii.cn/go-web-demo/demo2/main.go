package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "wwwroot"+r.URL.Path)
	})
	//http.ListenAndServe(":8080", nil)

	// 这行代码跟上面的代码效果一样.
	http.ListenAndServe(":8080", http.FileServer(http.Dir("wwwroot")))
}
