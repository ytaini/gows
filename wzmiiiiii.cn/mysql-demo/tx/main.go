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
	txDemo()
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

// 下面的代码演示了一个简单的事务操作，该事物操作能够确保两次更新操作要么同时成功要么同时失败，不会存在中间状态。

func txDemo() {
	tx, err := db.Begin() //开启事务
	if err != nil {
		if tx != nil {
			tx.Rollback() //回滚
		}
		fmt.Printf("begin trans failed, err:%v\n", err)
		return
	}
	sqlStr := `update user set age = 30 where id=?`
	ret1, err := tx.Exec(sqlStr, 3)
	if err != nil {
		tx.Rollback() // 回滚
		fmt.Printf("exec sql1 failed, err:%v\n", err)
		return
	}
	affRow1, err := ret1.RowsAffected()
	if err != nil {
		tx.Rollback() // 回滚
		fmt.Printf("exec ret1.RowsAffected() failed, err:%v\n", err)
		return
	}

	sqlStr2 := "Update user set age=40 where id=?"
	ret2, err := tx.Exec(sqlStr2, 7)
	if err != nil {
		tx.Rollback() // 回滚
		fmt.Printf("exec sql2 failed, err:%v\n", err)
		return
	}
	affRow2, err := ret2.RowsAffected()
	if err != nil {
		tx.Rollback() // 回滚
		fmt.Printf("exec ret1.RowsAffected() failed, err:%v\n", err)
		return
	}

	if affRow1 == 1 && affRow2 == 1 {
		fmt.Println("事务提交啦...")
		tx.Commit() // 提交事务
	} else {
		tx.Rollback()
		fmt.Println("事务回滚啦...")
	}

	fmt.Println("exec trans success!")
}
