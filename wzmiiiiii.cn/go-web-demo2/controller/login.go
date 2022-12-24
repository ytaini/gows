// @author: wzmiiiiii
// @since: 2022/12/24 01:07:13
// @desc: TODO

package controller

import (
	"net/http"
	"text/template"

	"wzmiiiiii.cn/gwd2/model"

	"wzmiiiiii.cn/gwd2/dao"
)

func loginRoute() {
	http.HandleFunc("/login", loginHandlerFunc)
}

func loginHandlerFunc(w http.ResponseWriter, r *http.Request) {
	var err error
	defer handleError(w, err)

	switch r.Method {
	case http.MethodGet:
		if err = loginHandleGet(w); err != nil {
			return
		}
	case http.MethodPost:
		if err = loginHandlePost(w, r); err != nil {
			return
		}
	}
}

func loginHandleGet(w http.ResponseWriter) (err error) {
	return template.Must(template.ParseFiles("view/pages/user/login.gohtml")).Execute(w, "")
}

func loginHandlePost(w http.ResponseWriter, r *http.Request) (err error) {
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	var user *model.User
	if user, err = dao.UserAuthentication(username, password); err != nil {
		return
	}
	if user != nil {
		if err = template.Must(template.ParseFiles("view/pages/user/login_success.gohtml")).Execute(w, username); err != nil {
			return
		}
	} else {
		if err = template.Must(template.ParseFiles("view/pages/user/login.gohtml")).
			Execute(w, "用户名或密码错误!!!"); err != nil {
			return
		}
	}
	return
}
