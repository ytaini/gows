// @author: wzmiiiiii
// @since: 2022/12/23 22:58:22
// @desc: TODO

package main

import (
	"log"
	"net/http"

	"wzmiiiiii.cn/gwd2/controller"
)

func main() {

	controller.RegisterRoutes()

	log.Println("Server starting...")
	log.Println(http.ListenAndServe(":8080", nil))
}
