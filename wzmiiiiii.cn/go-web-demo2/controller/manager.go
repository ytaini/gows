// @author: wzmiiiiii
// @since: 2022/12/24 16:17:26
// @desc: TODO

package controller

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"wzmiiiiii.cn/gwd2/model"

	"wzmiiiiii.cn/gwd2/dao"
)

func managerRoute() {
	http.HandleFunc("/manager", managerHandlerFunc)
	http.HandleFunc("/book/manager", bookManagerHandlerFunc)
	http.HandleFunc("/book/add", addBookHandlerFunc)
	http.HandleFunc("/book/delete/", deleteBookHandlerFunc)
}

func deleteBookHandlerFunc(w http.ResponseWriter, r *http.Request) {
	var err error
	defer handleError(w, err)
	switch r.Method {
	case http.MethodDelete:
		if err = deleteBookHandleDelete(r); err != nil {
			return
		}
	}
}

func deleteBookHandleDelete(r *http.Request) error {
	ids := strings.Split(r.URL.Path, "/")[3]
	id, err := strconv.Atoi(ids)
	if err != nil {
		return err
	}
	if err := dao.DeleteBookByID(id); err != nil {
		return err
	}
	return nil
}

func addBookHandlerFunc(w http.ResponseWriter, r *http.Request) {
	var err error
	defer handleError(w, err)
	switch r.Method {
	case http.MethodGet:
		if err = addBookHandleGet(w); err != nil {
			return
		}
	case http.MethodPost:
		if err = addBookHandlePost(w, r); err != nil {
			return
		}
	}
}
func addBookHandlePost(w http.ResponseWriter, r *http.Request) (err error) {
	price, _ := strconv.ParseFloat(r.PostFormValue("price"), 64)
	sales, _ := strconv.Atoi(r.PostFormValue("sales"))
	stock, _ := strconv.Atoi(r.PostFormValue("stock"))
	book := model.Book{
		Title:  r.PostFormValue("title"),
		Author: r.PostFormValue("author"),
		Price:  price,
		Sales:  sales,
		Stock:  stock,
	}
	if err = dao.SaveBook(&book); err != nil {
		return err
	}
	return bookManagerHandleGet(w)
}

func addBookHandleGet(w http.ResponseWriter) (err error) {
	return parseTemplate(w, "", "view/pages/manager/book_edit.gohtml")
}

func bookManagerHandlerFunc(w http.ResponseWriter, r *http.Request) {
	var err error
	defer handleError(w, err)
	switch r.Method {
	case http.MethodGet:
		if err = bookManagerHandleGet(w); err != nil {
			log.Println(err)
			return
		}
	}
}

func bookManagerHandleGet(w http.ResponseWriter) (err error) {
	var books []*model.Book
	books, err = dao.GetBooks()
	if err != nil {
		return
	}
	return parseTemplate(w, books, "view/pages/manager/book_manager.gohtml")
}

func managerHandlerFunc(w http.ResponseWriter, r *http.Request) {
	var err error
	defer handleError(w, err)
	switch r.Method {
	case http.MethodGet:
		if err = managerHandleGet(w); err != nil {
			return
		}
	}
}

func managerHandleGet(w http.ResponseWriter) (err error) {
	return parseTemplate(w, "", "view/pages/manager/manager.gohtml")
}
