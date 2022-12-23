// @author: wzmiiiiii
// @since: 2022/12/24 01:05:39
// @desc: TODO

package controller

import (
	"net/http"
	"text/template"
)

func indexRoute() {
	http.Handle("/", http.RedirectHandler("/index", http.StatusSeeOther))
	http.HandleFunc("/index", indexHandlerFunc)
}
func indexHandlerFunc(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("view/index.gohtml"))
	if err := t.Execute(w, nil); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
