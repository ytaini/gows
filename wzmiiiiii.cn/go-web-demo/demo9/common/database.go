// @author: wzmiiiiii
// @since: 2022/12/23 19:16:48
// @desc: TODO

package common

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

var dsn = "root:root@tcp(:3306)/go_test?charset=utf8mb4&parseTime=True"

func init() {
	err := initDB()
	if err != nil {
		panic(err)
	}
}

func initDB() (err error) {
	Db, err = sql.Open("mysql", dsn)
	if err != nil {
		return
	}
	if err = Db.Ping(); err != nil {
		return
	}

	log.Println("Connection success!!!")

	return nil
}
