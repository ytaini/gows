// @author: wzmiiiiii
// @since: 2022/12/24 01:07:13
// @desc: TODO

package controller

import (
	"net/http"
	"text/template"

	"wzmiiiiii.cn/gwd2/dao"
)

func loginRoute() {
	http.HandleFunc("/login", loginHandlerFunc)
}

func loginHandlerFunc(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	user, err := dao.UserAuthentication(username, password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if user != nil {
		err := template.Must(template.ParseFiles("view/pages/user/login_success.gohtml")).
			Execute(w, "")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	} else {
		err := template.Must(template.ParseFiles("view/pages/user/login.gohtml")).
			Execute(w, "用户名或密码错误!!!")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}
