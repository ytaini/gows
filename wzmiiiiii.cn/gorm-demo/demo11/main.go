package main

import (
	"database/sql"
	"fmt"

	"gorm.io/gorm"
)

func main() {
	err := initDB()
	if err != nil {
		panic(err)
	}

	// RawSQLDemo()

}

// 命名参数
// GORM 支持 sql.NamedArg、map[string]interface{}{} 或 struct 形式的命名参数

func NamedArgDemo() (err error) {
	var user User
	err = db.Where("name = @name", sql.Named("name", "张三")).Find(&user).Error
	// select * from users where name = '张三'
	err = db.Where("name = @name", map[string]any{"name": "张三"}).Find(&user).Error
	// select * from users where name = '张三'

	// 原生SQL 及命名参数
	db.Raw("SELECT * FROM users WHERE name1 = @name OR name2 = @name2 OR name3 = @name",
		sql.Named("name", "jinzhu1"), sql.Named("name2", "jinzhu2")).Find(&user)
	// SELECT * FROM users WHERE name1 = "jinzhu1" OR name2 = "jinzhu2" OR name3 = "jinzhu1"

	db.Exec("UPDATE users SET name1 = @name, name2 = @name2, name3 = @name",
		sql.Named("name", "jinzhunew"), sql.Named("name2", "jinzhunew2"))
	// UPDATE users SET name1 = "jinzhunew", name2 = "jinzhunew2", name3 = "jinzhunew"

	db.Raw("SELECT * FROM users WHERE (name1 = @name AND name3 = @name) AND name2 = @name2",
		map[string]interface{}{"name": "jinzhu", "name2": "jinzhu2"}).Find(&user)
	// SELECT * FROM users WHERE (name1 = "jinzhu" AND name3 = "jinzhu") AND name2 = "jinzhu2"

	type NamedArgument struct {
		Name  string
		Name2 string
	}

	db.Raw("SELECT * FROM users WHERE (name1 = @Name AND name3 = @Name) AND name2 = @Name2",
		NamedArgument{Name: "jinzhu", Name2: "jinzhu2"}).Find(&user)
	// SELECT * FROM users WHERE (name1 = "jinzhu" AND name3 = "jinzhu") AND name2 = "jinzhu2"
	return err
}

// 原生查询 SQL 和 Scan
// Exec 原生 SQL
// 注意 GORM 允许缓存预编译 SQL 语句来提高性能

func RawSQLDemo() {
	//var ress []*Result //也可以
	var ress []Result
	err := db.Raw("select id,name,age from users where name = ?", "name1").Scan(&ress).Error
	if err != nil {
		panic(err)
	}
	for _, result := range ress {
		fmt.Println(result)
	}

	var ageSum int
	err = db.Raw("select sum(age) from users").Scan(&ageSum).Error
	if err != nil {
		panic(err)
	}
	fmt.Println(ageSum)

	db.Exec("update users set name = ? , age = ? where id = ?", "你好", 13, 2)
	db.Exec("update users set name = ? , age = ? where id = ?", "你好", gorm.Expr("age / ?", 2), 3)
	//db.Exec("Drop Table users")
}
