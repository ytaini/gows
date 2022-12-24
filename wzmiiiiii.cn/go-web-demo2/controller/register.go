// @author: wzmiiiiii
// @since: 2022/12/24 01:39:00
// @desc: TODO

package controller

import (
	"net/http"

	"github.com/google/uuid"
	"wzmiiiiii.cn/gwd2/dao"
	"wzmiiiiii.cn/gwd2/model"
)

func registerRoute() {
	http.HandleFunc("/register", registerHandlerFunc)
	http.HandleFunc("/checkusername", checkUserNameHandlerFunc)
}

func checkUserNameHandlerFunc(w http.ResponseWriter, r *http.Request) {
	var err error
	defer handleError(w, err)

	switch r.Method {
	case http.MethodPost:
		username := r.PostFormValue("username")
		var exist bool
		if exist, err = dao.CheckUserName(username); err != nil {
			return
		}
		if !exist {
			w.Write([]byte("用户名可用!"))
		} else {
			w.Write([]byte("用户名已存在!"))
		}
	}
}

func registerHandlerFunc(w http.ResponseWriter, r *http.Request) {
	var err error
	defer handleError(w, err)

	switch r.Method {
	case http.MethodGet:
		if err = registerHandleGet(w); err != nil {
			return
		}
	case http.MethodPost:
		if err = registerHandlePost(w, r); err != nil {
			return
		}
	}
}

func registerHandleGet(w http.ResponseWriter) (err error) {
	return parseTemplate(w, "", "view/pages/user/regist.gohtml")
}

func registerHandlePost(w http.ResponseWriter, r *http.Request) (err error) {
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	rePassword := r.PostFormValue("repwd")
	email := r.PostFormValue("email")
	if password != rePassword {
		if err = parseTemplate(w, "两次密码不匹配", "view/pages/user/regist.gohtml"); err != nil {
			return
		}
	} else {
		var exist bool
		if exist, err = dao.CheckUserName(username); err != nil {
			return
		}
		if !exist {
			user := model.User{
				ID:       uuid.NewString(),
				Passwd:   password,
				UserName: username,
				Email:    email,
			}
			if err = dao.SavaUser(&user); err != nil {
				return
			}
			if err = parseTemplate(w, username, "view/pages/user/regist_success.gohtml"); err != nil {
				return
			}
		} else {
			if err = parseTemplate(w, "用户名已存在", "view/pages/user/regist.gohtml"); err != nil {
				return
			}
		}
	}
	return
}
