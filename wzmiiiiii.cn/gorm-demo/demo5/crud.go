package main

import (
	"fmt"
	"log"
	"time"
)

func InsertDemo1() {
	userinfo := UserInfo{
		Name: "张三",
		Age:  20,
	}
	db.Create(&userinfo) // `CreatedAt` 设为当前时间, 将 `UpdatedAt` 设为当前时间
	log.Println("新增数据的ID:", userinfo.ID)

	userinfo1 := UserInfo{
		Name:      "李四",
		Age:       31,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	// 创建记录时 userinfo1 的 `CreatedAt` 不会被修改, `UpdatedAt` 不会被修改
	db.Create(&userinfo1)
	log.Println("新增数据的ID:", userinfo1.ID)
}

func InsertDemo2() {
	userinfo := UserInfo{
		Name:      "tom",
		Age:       24,
		CreatedAt: time.Now(), // 当 autoCreateTime 标签置为 false时,CreatedAt必须指定一个时间
	}
	db.Create(&userinfo)
	log.Println("新增数据的ID:", userinfo.ID)

	userinfo1 := UserInfo{
		Name:      "jerry",
		Age:       32,
		CreatedAt: time.Now(),
	}
	db.Create(&userinfo1)
	log.Println("新增数据的ID:", userinfo1.ID)
}

func UpdateCreatedAtDemo() {
	userinfo := UserInfo{
		ID: 3,
	}
	// 修改id为3的CreatedAt字段
	db.Model(&userinfo).Update("CreatedAt", time.Now())
	if db.Error != nil {
		fmt.Println(db.Error)
		return
	}
}

func UpdateUpdatedAtDemo() {
	userinfo := UserInfo{
		ID:        1,
		Name:      "name1",
		Age:       21,
		CreatedAt: time.Now(),
	}
	// 更新id为1的UpdatedAt字段
	db.Save(&userinfo) // 同时会将 `UpdatedAt` 设为当前时间
	if db.Error != nil {
		fmt.Println(db.Error)
		return
	}
	userinfo1 := UserInfo{
		ID: 2,
	}

	// 更新id为2的name="name2"
	//db.Model(&userinfo1).Update("name", "name2") // 同时会将 `UpdatedAt` 设为当前时间

	// `UpdatedAt` 不会被修改
	db.Model(&userinfo1).UpdateColumn("name", "name3")
	if db.Error != nil {
		fmt.Println(db.Error)
		return
	}
}
