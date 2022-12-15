package main

import (
	"log"
)

func main() {
	err := initDB()
	if err != nil {
		panic(err)
	}
	//err = CreateTableDemo1()
	err = CreateTableDemo2()
	//err = CreateTableDemo3()
	if err != nil {
		panic(err)
	}
	log.Println("create table user_infos success")
}
