// @author: wzmiiiiii
// @since: 2022/12/23 20:04:04
// @desc: TODO

package controller

import (
	"fmt"
	"html/template"
	"net/http"
	"regexp"
	"strings"

	"github.com/google/uuid"

	"wzmiiiiii.cn/gwd/demo9/funcs"

	"wzmiiiiii.cn/gwd/demo9/dao"
	"wzmiiiiii.cn/gwd/demo9/model"
)

func registerRoutes() {
	http.HandleFunc("/", listCompanies)
	http.HandleFunc("/companies", listCompanies)
	http.HandleFunc("/companies/add", addCompany)
	http.HandleFunc("/companies/delete/", deleteCompany)
	http.HandleFunc("/companies/seed", seed)
	http.HandleFunc("/companies/edit/", editCompany)
}

func editCompany(w http.ResponseWriter, r *http.Request) {
	re := regexp.MustCompile(`/companies/edit/([a-zA-Z0-9-]*$)`)
	matched := re.FindStringSubmatch(r.URL.Path)

	if len(matched) > 0 {
		id := matched[1]
		switch r.Method {
		case http.MethodGet:
			company, err := dao.GetCompanyById(id)
			if err == nil {
				t, err := template.New("company-edit").
					ParseFiles("./template/_layout.gohtml", "./template/company/edit.gohtml")
				if err == nil {
					t.ExecuteTemplate(w, "_layout", company)
					return
				}
			}
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		case http.MethodPost:
			company := &model.Company{}
			company.ID = r.PostFormValue("id")
			company.Name = r.PostFormValue("name")
			company.NickName = r.PostFormValue("nickName")
			err := dao.UpdateCompany(company)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
			} else {
				http.Redirect(w, r, "/", http.StatusSeeOther)
			}
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

func seed(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "ok")
}

func deleteCompany(w http.ResponseWriter, r *http.Request) {
	id := strings.Split(r.URL.Path, "/")[3]
	//log.Println(id)
	if r.Method == http.MethodDelete {
		err := dao.DeleteCompanyById(id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func addCompany(w http.ResponseWriter, r *http.Request) {
	var err error
	defer func() {
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
	}()
	switch r.Method {
	case http.MethodGet:
		t := template.New("company-add")
		t, err = t.ParseFiles("./template/_layout.gohtml", "./template/company/add.gohtml")
		if err != nil {
			return
		}
		t.ExecuteTemplate(w, "_layout", nil)
	case http.MethodPost:
		c := model.Company{}
		c.ID = uuid.NewString()
		c.Name = r.PostFormValue("name")
		c.NickName = r.PostFormValue("nickName")
		if err = dao.InsertCompany(&c); err != nil {
			return
		}
		http.Redirect(w, r, "/companies", http.StatusSeeOther)
	}
}

func listCompanies(w http.ResponseWriter, r *http.Request) {
	var err error
	var companies []*model.Company
	defer func() {
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
	}()

	companies, err = dao.GetAllCompanies()

	if err != nil {
		return
	}
	funcMap := template.FuncMap{"add": funcs.Add}
	t := template.New("").Funcs(funcMap)
	t, err = t.ParseFiles("./template/_layout.gohtml", "./template/company/list.gohtml")
	if err != nil {
		return
	}
	if err = t.ExecuteTemplate(w, "_layout", companies); err != nil {
		return
	}
}
