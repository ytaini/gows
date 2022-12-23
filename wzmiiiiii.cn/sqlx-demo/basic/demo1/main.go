package main

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"log"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func main() {
	err := initDB()
	if err != nil {
		panic(err)
	}
	log.Println("conn establish success!!!")

	//queryRowDemo()
	//fmt.Println("---------------------")
	//queryDemo()
	//err = insertUserDemo()
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("---------------------")
	//namedQueryDemo()
	//fmt.Println("---------------------")
	//err = txDemo()
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}

	//users := []*user{
	//	{Name: "薛子涵", Age: 12},
	//	{Name: "万熠彤", Age: 16},
	//	{Name: "许嘉熙", Age: 15},
	//	{Name: "苏立诚", Age: 14},
	//	{Name: "白立果", Age: 13},
	//}
	//err = batchInsertUsers(users)

	//users := []any{
	//	&user{Name: "薛子涵", Age: 12},
	//	&user{Name: "万熠彤", Age: 16},
	//	&user{Name: "许嘉熙", Age: 15},
	//	&user{Name: "苏立诚", Age: 14},
	//	&user{Name: "白立果", Age: 13},
	//}
	//err = batchInsertUsers2(users)

	//err = batchInsertUsers3(users)

	//ids := []int{3, 5, 2, 6, 7, 8}
	//users, err := queryByIds(ids) //查询到的结果users 基于id排好序了
	//users, err := queryAndOrderByIds(ids) // 维持给定id集合的顺序。
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println("batch select success")
	//for _, user := range users {
	//	fmt.Println(user)
	//}

	//users := []*user{
	//	{Name: "name3", UserId: 28, Age: 12},
	//	{Name: "name4", UserId: 29, Age: 13},
	//	{Name: "name5", UserId: 30, Age: 14},
	//}
	//err = batchUpdateUsers(users)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println("batch update success")

	ids := []int{2, 3, 4, 5}
	err = batchDeleteUsers(ids)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("batch delete success")
}

var db *sqlx.DB

type user struct {
	UserId int    `db:"id"`
	Age    int    `db:"age"`
	Name   string `db:"name"`
}

// Value 使用sqlx.In实现批量插入,需要让user实现driver.Valuer接口
func (u *user) Value() (driver.Value, error) {
	return []any{u.Name, u.Age}, nil
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
	var users []*user
	err := db.Select(&users, sqlStr, 3)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	for _, user := range users {
		fmt.Println(user)
	}
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
	_, err = db.NamedExec(sqlStr, map[string]any{
		"name": "老刘",
		"age":  123,
	})
	return
}

// DB.NamedQuery与DB.NamedExec同理，这里是支持查询。
func namedQueryDemo() {
	sqlStr := `select * from user where name=:name`
	// 使用map做命名查询
	rows, err := db.NamedQuery(sqlStr, map[string]any{
		"name": "tom",
	})
	if err != nil {
		fmt.Printf("db.NamedQuery failed, err:%v\n", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var u user
		err := rows.StructScan(&u)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			continue
		}
		fmt.Println(u)
	}

	sqlStr = `select * from user where id > :id`

	// 使用结构体命名查询,根据结构体字段的db tag进行映射.
	u := user{UserId: 3}
	//rows, err = db.NamedQuery(sqlStr, &u) // 可以传指针,也可以传值 // 因为只是读取值.
	rows, err = db.NamedQuery(sqlStr, u)
	if err != nil {
		fmt.Printf("db.NamedQuery failed, err:%v\n", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var u user
		err := rows.StructScan(&u)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			continue
		}
		fmt.Println(u)
	}
}

// 对于事务操作，可以使用sqlx中提供的db.Beginx()和tx.Exec()方法。示例代码如下：
func txDemo() (err error) {
	tx, err := db.Beginx()
	if err != nil {
		fmt.Printf("begin trans failed, err:%v\n", err)
		return err
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			//panic(p) // re-throw panic after Rollback
			fmt.Println("panic!!,", p)
		} else if err != nil {
			tx.Rollback() // err is non-nil;dont change it.
		} else {
			err = tx.Commit() //err is nil; if Commit returns error update err
			fmt.Println("commit")
		}
	}()
	sqlStr := `update user set age = 30 where id = ?`
	ret, err := tx.Exec(sqlStr, 5)
	if err != nil {
		return err
	}
	n, err := ret.RowsAffected()
	if err != nil {
		return err
	}
	if n != 1 {
		return errors.New("exec sqlStr failed")
	}
	sqlStr = `update user set age = 40 where id = ?`
	rs := tx.MustExec(sqlStr, 8)
	n, err = rs.RowsAffected()
	if err != nil {
		return err
	}
	if n != 1 {
		return errors.New("exec sqlStr1 failed")
	}
	return err
}

// 自己拼接语句实现批量插入
func batchInsertUsers(users []*user) error {
	// 存放(?,?)的slice
	valueStrings := make([]string, 0, len(users))
	// 存放values 的slice
	valueArgs := make([]any, 0, len(users)*2)
	// 遍历users准备相关数据
	for _, u := range users {
		valueStrings = append(valueStrings, "(?,?)")
		valueArgs = append(valueArgs, u.Name)
		valueArgs = append(valueArgs, u.Age)
	}
	// 自行拼接要执行的具体语句
	stmt := fmt.Sprintf("insert into user (name, age) values %s", strings.Join(valueStrings, ","))
	fmt.Println(stmt)
	ret, err := db.Exec(stmt, valueArgs...)
	if err != nil {
		return err
	}
	n, err := ret.RowsAffected()
	if err != nil {
		return err
	}
	if n != int64(len(users)) {
		return errors.New("insert failed")
	}
	return nil
}

// sqlx.In是sqlx提供的一个非常方便的函数。
// sqlx.In的批量插入示例
// 前提是需要让结构体实现driver.Valuer接口
// 使用sqlx.In帮我们拼接语句和参数, 注意传入的参数是[]any
func batchInsertUsers2(users []any) error {
	sqlStr := `insert into user (name, age) values `
	for range users {
		sqlStr += "(?),"
	}
	sqlStr = sqlStr[:len(sqlStr)-1]
	// insert into user (name, age) values (?),(?),(?),(?),(?) 这里问号的个数取决于users的长度.
	query, args, err := sqlx.In(sqlStr, users...) // 如果arg实现了 driver.Valuer, sqlx.In 会通过调用 Value()来展开它
	if err != nil {
		return err
	}
	// insert into user (name, age) values (?, ?),(?, ?),(?, ?),(?, ?),(?, ?)
	fmt.Println(query) // 查看生成的querystring
	// [薛子涵 12 万熠彤 16 许嘉熙 15 苏立诚 14 白立果 13]
	fmt.Println(args) // 查看生成的args
	_, err = db.Exec(query, args...)
	return err
}

// 使用NamedExec实现批量插入
// 注意 ：该功能需1.3.1版本以上
func batchInsertUsers3(users []*user) error {
	sqlStr := `insert into user (name, age) values (:name,:age)`
	ret, err := db.NamedExec(sqlStr, users)
	if err != nil {
		return err
	}
	n, err := ret.RowsAffected()
	if err != nil {
		return err
	}
	if n != int64(len(users)) {
		return errors.New("insert failed")
	}
	return nil
}

// 批量更新
// replace into test_tbl (id,dr) values (1,'2'),(2,'3'),...(x,'y');
// insert into test_tbl (id,dr) values (1,'2'),(2,'3'),...(x,'y') on duplicate key update dr=values(dr);
func batchUpdateUsers(users []*user) error {
	//sqlStr := `replace into user (id,name,age) values (:id,:name,:age)`
	sqlStr := `insert into user (id, name) values (:id,:name) on duplicate key update name=values(name)`
	ret, err := db.NamedExec(sqlStr, users)
	if err != nil {
		return err
	}
	n, err := ret.RowsAffected()
	if err != nil {
		return err
	}
	fmt.Println(n) // 6
	return nil
}

// 批量删除
func batchDeleteUsers(ids []int) error {
	sqlStr := `delete from user where id in (?)`
	delStmt, args, err := sqlx.In(sqlStr, ids)
	if err != nil {
		return err
	}
	fmt.Println(delStmt) //delete from user where id in (?, ?, ?, ?)
	fmt.Println(args)    //[2 3 4 5]
	delStmt = db.Rebind(delStmt)
	ret, err := db.Exec(delStmt, args...)
	if err != nil {
		return err
	}
	n, err := ret.RowsAffected()
	if err != nil {
		return err
	}
	fmt.Println(n)
	return nil
}

// 关于sqlx.In这里再补充一个用法，在sqlx查询语句中实现In查询和FIND_IN_SET函数。
// 即实现 SELECT * FROM user WHERE id in (3, 2, 1);
// 和 SELECT * FROM user WHERE id in (3, 2, 1) ORDER BY FIND_IN_SET(id, '3,2,1');。

// in 查询
// 查询id在给定id集合中的数据。
func queryByIds(ids []int) (users []*user, err error) {
	//动态填充id
	query, args, err := sqlx.In("select id,name,age from user where id in (?)", ids)
	if err != nil {
		return
	}
	fmt.Println(query) // select id,name,age from user where id in (?, ?, ?)
	// sqlx.In 返回带 `?` bindvar的查询语句, 可以使用Rebind()生成对应数据库的bindvar类型.
	query = db.Rebind(query)
	// 因为mysql对应的bindvars就是使用'?',所以没有变化.
	// 如果现在我们使用的是PostgreSQL数据库,调用Rebind方法后,query 将变成 select id,name,age from user where id in ($1, $2, $3).
	fmt.Println(query) // select id,name,age from user where id in (?, ?, ?)
	err = db.Select(&users, query, args...)
	return
}

// in查询和FIND_IN_SET函数
// 查询id在给定id集合的数据并维持给定id集合的顺序。
func queryAndOrderByIds(ids []int) (users []*user, err error) {
	// 动态填充id
	strIds := make([]string, 0, len(ids))
	for _, id := range ids {
		strIds = append(strIds, fmt.Sprintf("%d", id))
	}
	sqlStr := `select * from user where id in (?) order by find_in_set(id,?)`
	query, args, err := sqlx.In(sqlStr, ids, strings.Join(strIds, ","))
	if err != nil {
		return
	}
	query = db.Rebind(query)
	err = db.Select(&users, query, args...)
	return
}
