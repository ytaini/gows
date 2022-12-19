package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

const DSN = "root:root@tcp(:3306)/go_test?charset=utf8mb4&parseTime=True&loc=Local"

func initDB() (err error) {
	db, err = gorm.Open(mysql.Open(DSN), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		return
	}
	return
}
