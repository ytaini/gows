// @author: wzmiiiiii
// @since: 2022/12/24 01:04:03
// @desc: TODO

package controller

import "net/http"

func RegisterRoutes() {
	indexRoute()
	loginRoute()
	registerRoute()
	staticRoute()
}

func staticRoute() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("view/static/"))))
	http.Handle("/pages/", http.StripPrefix("/pages/", http.FileServer(http.Dir("view/pages/"))))
}
