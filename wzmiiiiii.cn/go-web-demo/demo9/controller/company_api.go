// @author: wzmiiiiii
// @since: 2022/12/23 19:54:09
// @desc: TODO

package controller

import (
	"encoding/json"
	"net/http"

	"wzmiiiiii.cn/gwd/demo9/model"

	"wzmiiiiii.cn/gwd/demo9/dao"
)

func registerAPIRoutes() {
	http.HandleFunc("/api/companies", getCompanies)
}

func getCompanies(w http.ResponseWriter, r *http.Request) {
	var err error
	defer func() {
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
	}()
	switch r.Method {
	case http.MethodGet:
		companies, err := dao.GetAllCompanies()
		if err != nil {
			return
		}
		if err := json.NewEncoder(w).Encode(&companies); err != nil {
			return
		}
	case http.MethodPost:
		var c model.Company
		if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
			return
		}
		if err := dao.InsertCompany(&c); err != nil {
			return
		}
	}
}
