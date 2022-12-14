package main

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func main() {
	err := initDB()
	if err != nil {
		panic(err)
	}
	log.Println("conn establish success!!!")
	queryRowDemo()
	queryDemo()
	err = insertUserDemo()
	if err != nil {
		panic(err)
	}
}

var db *sqlx.DB

type user struct {
	Id   int
	Age  int
	Name string
}

// 连接数据库
func initDB() (err error) {
	dsn := `root:root@tcp(0.0.0.0:3306)/go_test`
	//  也可以使用MustConnect连接不成功就panic
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		return err
	}
	return
}

// 查询单行数据
func queryRowDemo() {
	sqlStr := `select id,name,age from user where id = ?`
	var u user
	err := db.Get(&u, sqlStr, 2)
	if err != nil {
		fmt.Printf("get failed, err:%v\n", err)
		return
	}
	fmt.Println(u)
}

// 查询多行数据
func queryDemo() {
	sqlStr := `select id,name,age from user where id > ?`
	var users []user
	err := db.Select(&users, sqlStr, 3)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	fmt.Println(users)
}

// 插入数据
func insertRowDemo() {
	sqlStr := "insert into user(name, age) values (?,?)"
	ret, err := db.Exec(sqlStr, "ls", 19)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	theID, err := ret.LastInsertId() // 新插入数据的id
	if err != nil {
		fmt.Printf("get lastinsert ID failed, err:%v\n", err)
		return
	}
	fmt.Printf("insert success, the id is %d.\n", theID)
}

// 更新数据
func updateRowDemo() {
	sqlStr := "update user set age=? where id = ?"
	ret, err := db.Exec(sqlStr, 39, 6)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("update success, affected rows:%d\n", n)
}

// 删除数据
func deleteRowDemo() {
	sqlStr := "delete from user where id = ?"
	ret, err := db.Exec(sqlStr, 6)
	if err != nil {
		fmt.Printf("delete failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("delete success, affected rows:%d\n", n)
}

// DB.NamedExec方法用来绑定SQL语句与结构体或map中的同名字段。
func insertUserDemo() (err error) {
	sqlStr := `insert into user (name, age) values (:name,:age)`
	_, err = db.NamedQuery(sqlStr, map[string]any{
		"name": "老刘",
		"age":  123,
	})
	return
}

// DB.NamedQuery与DB.NamedExec同理，这里是支持查询。
