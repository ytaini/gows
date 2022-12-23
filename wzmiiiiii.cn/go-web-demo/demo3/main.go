package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(10240); err != nil {
		log.Println(err)
		return
	}
	for key, strings := range r.Header {
		log.Printf("%v:%v\v", key, strings)
	}
	log.Println(r.URL.String())  // /process?hello=world&id=123
	log.Println(r.URL.Query())   // map[hello:[world] id:[123]]
	log.Println(r.Form)          // map[hello:[world asdda] id:[123] post:[456]]
	log.Println(r.PostForm)      // map[hello:[asdda] post:[456]]
	log.Println(r.MultipartForm) // &{map[hello:[asdda] post:[456]] map[uploaded:[0x1400010e280]]}
	log.Println(r.Method)        // POST
	log.Println(r.Host)          // localhost:8080
	log.Println(r.RemoteAddr)    // [::1]:57306

	//fileHeader := r.MultipartForm.File["uploaded"][0]
	//file, err := fileHeader.Open()
	//上面这段代码和下面的功能一样.
	file, _, err := r.FormFile("uploaded")

	if err == nil {
		data, err := io.ReadAll(file)
		if err == nil {
			if _, err := fmt.Fprintln(w, string(data)); err != nil {
				log.Println(err)
				return
			}
		}
	}
}

func main() {
	server := http.Server{
		Addr: ":8080",
	}
	http.HandleFunc("/process", process)
	log.Println(server.ListenAndServe())
}
