// @author: wzmiiiiii
// @since: 2022/12/24 14:28:49
// @desc: TODO

package controller

import (
	"net/http"
	"text/template"
)

func handleError(w http.ResponseWriter, err error) {
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func parseTemplate(w http.ResponseWriter, data any, filenames ...string) error {
	return template.Must(template.ParseFiles(filenames...)).
		Execute(w, data)
}
