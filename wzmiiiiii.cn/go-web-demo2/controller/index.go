// @author: wzmiiiiii
// @since: 2022/12/24 01:05:39
// @desc: TODO

package controller

import (
	"net/http"
)

func indexRoute() {
	http.Handle("/", http.RedirectHandler("/index", http.StatusSeeOther))
	http.HandleFunc("/index", indexHandlerFunc)
}

func indexHandlerFunc(w http.ResponseWriter, r *http.Request) {
	var err error
	defer handleError(w, err)
	switch r.Method {
	case http.MethodGet:
		if err = indexHandlerGet(w); err != nil {
			return
		}
	}
}
func indexHandlerGet(w http.ResponseWriter) error {
	return parseTemplate(w, "", "view/index.gohtml")
}
