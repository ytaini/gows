package main

import (
	"database/sql"
	"fmt"
	"strings"

	"gorm.io/gorm/schema"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	initDB3()
	//initDB2()
	//initDB1()
}

const DSN = "root:root@tcp(:3306)/go_test?charset=utf8mb4&parseTime=True&loc=Local"

func initDB3() {
	// GORM 允许通过一个现有的数据库连接来初始化 *gorm.DB
	sqlDB, err := sql.Open("mysql", DSN)
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println(db)
}

func initDB2() {

	// MySQL 驱动程序提供了 一些高级配置 可以在初始化过程中使用，例如：
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       DSN,
		DefaultStringSize:         256,   // string类型字段的默认长度.
		DisableDatetimePrecision:  true,  // 禁用datetime 精度,Mysql5.6之前不支持.
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前MySQL版本自动配置
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "t_",                              // table name prefix, table for `User` would be `t_users`
			SingularTable: true,                              // use singular table name, table for `User` would be `user` with this option enabled
			NoLowerCase:   true,                              // skip the snake_casing of names
			NameReplacer:  strings.NewReplacer("CID", "Cid"), // use name replacer to change struct/field name before convert it to db name

		},
	})
	// 命名策略
	// GORM 允许用户通过覆盖默认的NamingStrategy来更改命名约定，这需要实现接口 Namer

	// 默认 NamingStrategy 也提供了几个选项，如上.

	if err != nil {
		panic(err)
	}
	fmt.Println(db)
}

func initDB1() {
	// 注意：想要正确的处理 time.Time ，需要带上 parseTime 参数，
	// 要支持完整的 UTF-8 编码，需要将 charset=utf8 更改为 charset=utf8mb4
	db, err := gorm.Open(mysql.Open(DSN), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println(db)
}
