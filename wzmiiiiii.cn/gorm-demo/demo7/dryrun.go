package main

import (
	"log"

	"gorm.io/gorm"
)

func DryRunDemo1() {
	// DryRun 模式: 在不执行的情况下生成 SQL 及其参数，可以用于准备或测试生成的 SQL
	stmt := db.Session(&gorm.Session{DryRun: true}).Create(user).Statement

	log.Println("生成的SQL:", stmt.SQL.String())
	//  INSERT INTO `users` (`created_at`,`updated_at`,`deleted_at`,`name`,`age`,`birthday`) VALUES (?,?,?,?,?,?)

	log.Println("生成的参数:")
	for i, v := range stmt.Vars {
		log.Printf("第%d个参数: %v", i, v)
	}
	//2022/12/16 15:15:22 第0个参数: 2022-12-16 15:15:22.061 +0800 CST
	//2022/12/16 15:15:22 第1个参数: 2022-12-16 15:15:22.061 +0800 CST
	//2022/12/16 15:15:22 第2个参数: {0001-01-01 00:00:00 +0000 UTC false}
	//2022/12/16 15:15:22 第3个参数: 李四
	//2022/12/16 15:15:22 第4个参数: 20
	//2022/12/16 15:15:22 第5个参数: 1999-11-13 00:00:00 +0800 CST
}

func DryRunDemo2() {
	stmt := db.Session(&gorm.Session{DryRun: true}).
		Select("Name", "Age", "createdAt").Create(user).Statement

	log.Println("生成的SQL:", stmt.SQL.String())
	// INSERT INTO `users` (`created_at`,`updated_at`,`name`,`age`) VALUES (?,?,?,?)

	log.Println("生成的参数:")
	for i, v := range stmt.Vars {
		log.Printf("第%d个参数: %v", i, v)
	}
	//2022/12/16 15:41:05 第0个参数: 2022-12-16 15:41:05.209 +0800 CST
	//2022/12/16 15:41:05 第1个参数: 2022-12-16 15:41:05.209 +0800 CST
	//2022/12/16 15:41:05 第2个参数: 张三
	//2022/12/16 15:41:05 第3个参数: 18
}

func DryRunDemo3() {
	stmt := db.Session(&gorm.Session{DryRun: true}).
		Omit("name", "birthday", "CreatedAt").Create(user).Statement

	log.Println("生成的SQL:", stmt.SQL.String())
	//  INSERT INTO `users` (`updated_at`,`deleted_at`,`age`) VALUES (?,?,?)

	log.Println("生成的参数:")
	for i, v := range stmt.Vars {
		log.Printf("第%d个参数: %v", i, v)
	}
	//2022/12/16 15:46:20 第0个参数: 2022-12-16 15:46:20.582 +0800 CST
	//2022/12/16 15:46:20 第1个参数: {0001-01-01 00:00:00 +0000 UTC false}
	//2022/12/16 15:46:20 第2个参数: 18
}

func DryRunDemo4() {
	//stmt := db.Session(&gorm.Session{DryRun: true}).Create(&users).Statement
	stmt := db.Session(&gorm.Session{DryRun: true}).Create(&users1).Statement

	log.Println("生成的SQL:", stmt.SQL.String())
	//   INSERT INTO `users` (`created_at`,`updated_at`,`deleted_at`,`name`,`age`,`birthday`)
	//   VALUES (?,?,?,?,?,?),(?,?,?,?,?,?),(?,?,?,?,?,?)

	log.Println("生成的参数:")
	for i, v := range stmt.Vars {
		log.Printf("第%d个参数: %v", i, v)
	}
}

func DryRunDemo5() {
	stmt := db.Session(&gorm.Session{DryRun: true}).
		CreateInBatches(GenerateUsers(), 100).Statement
	// 拿不到sql
	log.Println("生成的SQL:", stmt.SQL.String())
}

func DryRunDemo6() {
	stmt := db.Session(&gorm.Session{DryRun: true, CreateBatchSize: 1000}).
		Create(GenerateUsers()).Statement
	log.Println(stmt.Error)
	// 拿不到sql
	log.Println("生成的SQL:", stmt.SQL.String())
}

func DryRunDemo7() {
	// 注意: 根据 Map 创建记录,不会对创建/更新等时间进行追踪
	stmt := db.Session(&gorm.Session{DryRun: true}).Model(&User{}).Create(map[string]any{
		"Name":     "map",
		"Age":      14,
		"birthday": birthday,
	}).Statement

	log.Println("生成的SQL:", stmt.SQL.String())
	// INSERT INTO `users` (`age`,`name`,`birthday`) VALUES (?,?,?)

	log.Println("生成的参数:")
	for i, v := range stmt.Vars {
		log.Printf("第%d个参数: %v", i, v)
	}
	//2022/12/16 17:32:06 第0个参数: 14
	//2022/12/16 17:32:06 第1个参数: map
	//2022/12/16 17:32:06 第2个参数: 1999-12-13 00:00:00 +0800 CST
}
