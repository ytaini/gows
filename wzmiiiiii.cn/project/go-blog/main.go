package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

var logger *log.Logger

func init() {
	logger = log.Default()
}

func main() {
	server := http.Server{
		Addr: ":8080",
	}
	http.HandleFunc("/index.html", indexHandleFunc)
	http.Handle("/", http.RedirectHandler("/index.html", http.StatusFound))
	if err := server.ListenAndServe(); err != nil {
		logger.Println(err)
	}
}

func indexHandleFunc(w http.ResponseWriter, r *http.Request) {
	tm := template.New("index.html")
	path, _ := os.Getwd() // 获取当前文件所在模块的路径: /Users/imzw/gows/wzmiiiiii.cn/project
	logger.Println(path)
	tm, _ = tm.ParseFiles(path + "/go-blog/template/index.html")
	_ = tm.Execute(w, "")
}
