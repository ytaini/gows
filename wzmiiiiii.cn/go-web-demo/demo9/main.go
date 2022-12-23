// @author: wzmiiiiii
// @since: 2022/12/23 19:05:40
// @desc: TODO

package main

import (
	"log"
	"net/http"

	"wzmiiiiii.cn/gwd/demo9/controller"

	"wzmiiiiii.cn/gwd/demo9/middleware"
)

func main() {
	server := http.Server{
		Addr:    ":8080",
		Handler: new(middleware.BasicAuthMiddleware),
	}

	controller.RegisterRoutes()

	log.Println("Server starting...")
	go http.ListenAndServe(":8000", nil)
	log.Println(server.ListenAndServe())
}
