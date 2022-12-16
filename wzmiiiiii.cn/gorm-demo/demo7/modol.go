package main

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

var birthday = ParseStringData(1999, 12, 13)

var user = NewUser("张三", 18, birthday)

var users = []*User{
	NewUser("tom", 13, birthday),
	NewUser("tom1", 13, birthday),
	NewUser("tom2", 13, birthday),
}

var users1 = []User{
	{Name: "tom", Age: 13, Birthday: birthday},
	{Name: "tom1", Age: 13, Birthday: birthday},
	{Name: "tom2", Age: 13, Birthday: birthday},
}

type User struct {
	gorm.Model
	Name     string `gorm:"type:string;size:20"`
	Age      uint8  `gorm:"type:int;default:18"`
	Birthday time.Time
}

// BeforeCreate
// GORM 允许用户定义的钩子有 BeforeSave, BeforeCreate, AfterSave, AfterCreate 等创建记录时将调用这些钩子方法
// 如果您想跳过 钩子 方法，您可以使用 SkipHooks 会话模式，例如：
// - DB.Session(&gorm.Session{SkipHooks: true}).Create(&user)
// - DB.Session(&gorm.Session{SkipHooks: true}).Create(&users)
// - DB.Session(&gorm.Session{SkipHooks: true}).CreateInBatches(users, 100)
func (u *User) BeforeCreate(tx *gorm.DB) error {
	if u.Name == "admin" {
		return errors.New("invalid name")
	}
	return nil
}

func NewUser(name string, age uint8, birthday time.Time) *User {
	return &User{Name: name, Age: age, Birthday: birthday}
}
