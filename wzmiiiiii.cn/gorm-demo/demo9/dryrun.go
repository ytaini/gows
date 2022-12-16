package main

import (
	"log"

	"gorm.io/gorm"
)

func DryRunDemo1() {
	var user User
	err := db.First(&user).Error
	if err != nil {
		return
	}
	user.Age = 12
	user.Name = "张三"
	stmt := db.Session(&gorm.Session{
		DryRun: true,
	}).Save(&user).Statement
	// UPDATE `users` SET
	// `created_at`=?,`updated_at`=?,`deleted_at`=?,`name`=?,`age`=?,`birthday`=?
	// WHERE `users`.`deleted_at` IS NULL AND `id` = ?
	log.Println(stmt.SQL.String())
}

func DryRunDemo2() {
	user := User{}
	stmt := db.Session(&gorm.Session{DryRun: true}).Save(&user).Statement
	// INSERT INTO
	// `users` (`created_at`,`updated_at`,`deleted_at`,`name`,`age`,`birthday`)
	// VALUES (?,?,?,?,?,?)
	log.Println(stmt.SQL.String())
}

func DryRunDemo3() {
	stmt := db.Session(&gorm.Session{DryRun: true}).Model(&User{}).Where("id = ?", 11).
		Update("name", "updateTest").Statement
	// UPDATE `users` SET `name`=?,`updated_at`=? WHERE id = ? AND `users`.`deleted_at` IS NULL
	log.Println(stmt.SQL.String())

	var user User
	db.First(&user)
	stmt = db.Session(&gorm.Session{DryRun: true}).Model(&user).
		Update("name", "张三1").Statement
	// UPDATE `users` SET `name`=?,`updated_at`=? WHERE `users`.`deleted_at` IS NULL AND `id` = ?

	log.Println(stmt.SQL.String())

	stmt = db.Session(&gorm.Session{DryRun: true}).Model(&user).
		Where("active = ?", true).Update("age", 22).Statement
	// UPDATE `users` SET `age`=?,`updated_at`=? WHERE active = ? AND `users`.`deleted_at` IS NULL AND `id` = ?

	log.Println(stmt.SQL.String())

}
