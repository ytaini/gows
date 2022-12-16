package main

import (
	"log"

	"gorm.io/gorm/clause"

	"gorm.io/gorm"
)

func InsertDemo() (err error) {
	creditCard := CreditCard{
		Number: "123123213",
	}
	user := User{
		Name:       "zw1",
		CreditCard: creditCard,
	}
	result := db.Create(&user)
	// INSERT INTO `users` ...
	// INSERT INTO `credit_cards` ...

	// 可以通过 Select、 Omit 跳过关联保存
	//db.Omit("CreditCard").Create(&user)

	// 跳过所有关联
	//db.Omit(clause.Associations).Create(&user)

	err = result.Error
	if err != nil {
		return
	}
	log.Printf("插入数据的生成主键: %d\n", user.ID)          // 返回插入数据的主键
	log.Printf("插入数据的生成主键: %d\n", creditCard.ID)    // 返回插入数据的主键
	log.Printf("插入记录的条数:%d\n", result.RowsAffected) // 返回插入记录的条数
	return nil
}

func DryRunDemo1() {
	creditCard := CreditCard{
		Number: "123123213",
	}
	user := User{
		Name:       "zw1",
		CreditCard: creditCard,
	}
	stmt := db.Session(&gorm.Session{DryRun: true}).Create(&user).Statement
	log.Println("生成的SQL:", stmt.SQL.String())
	log.Println("生成的参数:")
	for i, v := range stmt.Vars {
		log.Printf("第%d个参数: %v", i, v)
	}
}

// upsert 是数据库插入操作的扩展，如果某个唯一字段已经存在，则将本次新增插入操作变成更新操作，否则就正常执行插入操作
// Upsert 及冲突
// GORM 为不同数据库提供了兼容的 Upsert 支持

func InsertConflictDemo1() (err error) {
	user := &UserInformation{
		Name: "zs2",
		ID:   "123",
		Age:  15,
	}
	// Do nothing on conflict
	//result := db.Clauses(clause.OnConflict{DoNothing: true}).Create(user)

	// 当产生冲突时,将除主键外的所有字段更新为新值.
	//result := db.Clauses(clause.OnConflict{UpdateAll: true}).Create(user)

	// 当产生冲突时,将指定字段更新为指定值.
	//result := db.Clauses(clause.OnConflict{
	//	Columns:   []clause.Column{{Name: "id"}},
	//	DoUpdates: clause.Assignments(map[string]interface{}{"name": "nihao"}),
	//}).Create(user)

	// 当产生冲突时,将指定字段更新为新值
	result := db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"name", "age"}),
	}).Create(user)

	err = result.Error
	if err != nil {
		return
	}
	log.Printf("影响记录的条数:%d\n", result.RowsAffected) // 返回插入记录的条数
	return nil
}
