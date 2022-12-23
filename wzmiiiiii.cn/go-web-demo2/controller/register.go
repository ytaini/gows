// @author: wzmiiiiii
// @since: 2022/12/24 01:39:00
// @desc: TODO

package controller

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/google/uuid"
	"wzmiiiiii.cn/gwd2/model"

	"wzmiiiiii.cn/gwd2/dao"
)

func registerRoute() {
	http.HandleFunc("/register", registerHandlerFunc)
}

func registerHandlerFunc(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	rePassword := r.PostFormValue("repwd")
	email := r.PostFormValue("email")

	if password != rePassword {
		fmt.Println("asd")
	}

	exist, err := dao.CheckUserName(username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if !exist {
		user := model.User{
			ID:       uuid.NewString(),
			Passwd:   password,
			UserName: username,
			Email:    email,
		}

		if err := dao.SavaUser(&user); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		err := template.Must(template.ParseFiles("view/pages/user/regist_success.gohtml")).
			Execute(w, user)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	} else {
		err := template.Must(template.ParseFiles("view/pages/user/regist.gohtml")).
			Execute(w, "")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}
