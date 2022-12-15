package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

const DSN = "root:root@tcp(:3306)/go_test?charset=utf8mb4&parseTime=True&loc=Local"

func initDB() (err error) {
	db, err = gorm.Open(mysql.Open(DSN), &gorm.Config{})
	if err != nil {
		return
	}
	return
}
