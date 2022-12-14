package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// 定义一个全局变量db，用来保存数据库连接对象。
var db *sql.DB

func main() {
	err := initDB()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("conn establish success ...")
	//queryRowDemo()
	//queryDemo()
	//insertDemo()
	//updateDemo()
	deleteDemo()
}

func initDB() (err error) {
	// DSN: Data Source Name
	dsn := "root:root@tcp(0.0.0.0:3306)/go_test?charset=utf8mb4&parseTime=True"
	// 不会校验账号密码是否正确
	// 注意！！！这里不要使用:=，我们是给全局变量赋值，然后在main函数中使用全局变量db
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	// 尝试与数据库建立连接(校验dsn是否正确)
	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}

type user struct {
	id   int
	age  int
	name string
}

// CRUD

// 单行查询db.QueryRow()执行一次查询，并期望返回最多一行结果（即Row）.
// QueryRow总是返回非nil的值，直到返回值的Scan方法被调用时，才会返回被延迟的错误。（如：未找到结果）
func queryRowDemo() {
	//sqlStr := `select id,name,age from user where id=?`
	sqlStr := `select * from user where id=?`
	var u user
	// 非常重要: 确保QueryRow之后调用Scan方法,否则持有的数据库连接不会被释放.
	err := db.QueryRow(sqlStr, 1).Scan(&u.id, &u.name, &u.age)
	if err != nil {
		log.Println("scan failed, err:", err)
		return
	}
	log.Println(u)
}

// 多行查询db.Query()执行一次查询，返回多行结果（即Rows），一般用于执行select命令。参数args表示query中的占位参数。
func queryDemo() {
	sqlStr := `select id,name,age from user where id > ?`
	rows, err := db.Query(sqlStr, 1)
	if err != nil {
		log.Println("query failed,err:", err)
		return
	}
	// 非常重要: 关闭rows释放持有的数据库连接
	defer rows.Close()

	// 循环读取结果集中的数据
	for rows.Next() {
		var u user
		err := rows.Scan(&u.id, &u.name, &u.age)
		if err != nil {
			log.Println("scan failed,err:", err)
			return
		}
		log.Println(u)
	}
}

// 插入、更新和删除操作都使用Exec方法。
// Exec执行一次命令（包括查询、删除、更新、插入等），返回的Result是对已执行的SQL命令的总结。参数args表示query中的占位参数。
func insertDemo() {
	sqlStr := `insert into user(name,age) values(?,?)`
	ret, err := db.Exec(sqlStr, "tom", 22)
	if err != nil {
		log.Println("insert data failed, err:", err)
		return
	}
	id, err := ret.LastInsertId() // 新插入数据的id
	if err != nil {
		log.Println("get lastInsertId failed, err:", err)
		return
	}
	log.Printf("insert success, the id is %d", id)
}

func updateDemo() {
	sqlStr := `update user set name = ? where id = ?`
	ret, err := db.Exec(sqlStr, "jerry", 1)
	if err != nil {
		log.Println("update data failed, err:", err)
		return
	}
	n, err := ret.RowsAffected()
	if err != nil {
		log.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	log.Printf("update success, affected rows:%d", n)
}

func deleteDemo() {
	sqlStr := `delete from user where id = ?`
	ret, err := db.Exec(sqlStr, 1)
	if err != nil {
		log.Println("delete data failed, err:", err)
		return
	}
	n, err := ret.RowsAffected()
	if err != nil {
		log.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	log.Printf("delete success, affected rows:%d", n)
}

// MySQL预处理
