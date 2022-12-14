package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
	err := initDB()
	if err != nil {
		log.Println(err)
		return
	}
	prepareQueryDemo()
	prepareInsertDemo()
}

func initDB() (err error) {
	dsn := `root:root@tcp(0.0.0.0:3306)/go_test`
	if db, err = sql.Open("mysql", dsn); err != nil {
		return err
	}
	if err := db.Ping(); err != nil {
		return err
	}
	log.Println("database conn establish success.")
	return nil
}

type user struct {
	id   int
	age  int
	name string
}

// database/sql中使用下面的Prepare方法来实现预处理操作。
// Prepare方法会先将sql语句发送给MySQL服务端，返回一个准备好的状态用于之后的查询和命令。返回值可以同时执行多个查询和命令。
func prepareQueryDemo() {
	sqlStr := `select id,name,age from user where id > ?`
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		log.Println("prepare failed,err:", err)
		return
	}
	defer stmt.Close()
	rows, err := stmt.Query(1)
	if err != nil {
		log.Printf("query failed, err:%v\n", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var u user
		err := rows.Scan(&u.id, &u.name, &u.age)
		if err != nil {
			log.Printf("scan failed, err:%v\n", err)
			return
		}
		log.Println(u)
	}
}

// 插入、更新和删除操作的预处理十分类似
// 预处理插入示例
func prepareInsertDemo() {
	sqlStr := "insert into user(name, age) values (?,?)"
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("prepare failed, err:%v\n", err)
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec("小王子", 18)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	_, err = stmt.Exec("沙河娜扎", 18)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	fmt.Println("insert success.")
}
