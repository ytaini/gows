package main

import "log"

func CreateTableDemo() (err error) {
	//err = db.AutoMigrate(&User{})
	migrator := db.Migrator()
	err = migrator.AutoMigrate(&User{})
	if err != nil {
		return
	}
	log.Println(migrator.CurrentDatabase()) // 获取当前数据库名.
	tableList, err := migrator.GetTables()  // 获取当前数据库中所有表名.
	if err != nil {
		return
	}
	log.Println(tableList)
	return
}
