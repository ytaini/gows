// @author: wzmiiiiii
// @since: 2022/12/24 00:10:42
// @desc: TODO

package common

import (
	"log"

	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sqlx.DB

var (
	dsn = "root:root@tcp(:3306)/go_test?charset=utf8mb4&parseTime=True"
)

func init() {
	Db = sqlx.MustConnect("mysql", dsn)
	log.Println("Database Connection success...")
}
