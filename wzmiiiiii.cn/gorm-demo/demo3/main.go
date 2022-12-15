package main

import (
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	err := initDB()
	if err != nil {
		panic(err)
	}
	sqlDBDemo()
}

func sqlDBDemo() (err error) {
	// GORM 提供了 DB 方法，可用于从当前 *gorm.DB 返回一个通用的数据库接口 *sql.DB
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	sqlDB.Ping()
	sqlDB.Close()
	sqlDB.Stats()

	// GORM 使用 database/sql 维护连接池

	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)
	return nil
}

var db *gorm.DB

const DSN = "root:root@tcp(:3306)/go_test?charset=utf8mb4&parseTime=True&loc=Local"

func initDB() (err error) {
	db, err = gorm.Open(mysql.Open(DSN), &gorm.Config{})
	if err != nil {
		return
	}
	return
}
