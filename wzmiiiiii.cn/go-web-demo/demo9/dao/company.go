// @author: wzmiiiiii
// @since: 2022/12/23 19:15:58
// @desc: TODO

package dao

import (
	"log"

	"wzmiiiiii.cn/gwd/demo9/common"
	"wzmiiiiii.cn/gwd/demo9/model"
)

func GetAllCompanies() (companies []*model.Company, err error) {
	sql := `select id,name,nick_name from companies`
	rows, err := common.Db.Query(sql)
	if err != nil {
		return
	}
	for rows.Next() {
		var c model.Company
		err = rows.Scan(&c.ID, &c.Name, &c.NickName)
		if err != nil {
			return
		}
		companies = append(companies, &c)
	}
	return
}

func GetCompanyById(id string) (company model.Company, err error) {
	sql := `select id,name,nick_name from companies where id = ?`
	err = common.Db.QueryRow(sql, id).Scan(&company.ID, &company.Name, &company.NickName)
	log.Println(err)
	return
}

func InsertCompany(company *model.Company) (err error) {
	sql := `insert into companies (id, name, nick_name) VALUES (?,?,?)`
	stmt, err := common.Db.Prepare(sql)
	if err != nil {
		return
	}
	if _, err = stmt.Exec(company.ID, company.Name, company.NickName); err != nil {
		return
	}
	return
}

func UpdateCompany(company *model.Company) (err error) {
	sql := `update companies set name = ? , nick_name = ? where id = ?`
	_, err = common.Db.Exec(sql, company.Name, company.NickName, company.ID)
	return
}

func DeleteCompanyById(id string) (err error) {
	sql := `delete from companies where id = ?`
	_, err = common.Db.Exec(sql, id)
	return
}
