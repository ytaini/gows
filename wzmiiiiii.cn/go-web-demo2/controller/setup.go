// @author: wzmiiiiii
// @since: 2022/12/24 01:04:03
// @desc: TODO

package controller

import "net/http"

func RegisterRoutes() {
	indexRoute()
	loginRoute()
	registerRoute()
	managerRoute()
	staticRoute()
}

func staticRoute() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("view/static/"))))
}
