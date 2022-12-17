package main

import (
	"errors"
	"strings"
	"time"

	"gorm.io/gorm"
)

var birthday = ParseStringData(1999, 12, 13)

// User
// 以通过标签 default 为字段定义默认值
// 插入记录到数据库时，默认值 会被用于 填充值为 零值 的字段
// 注意: 对于声明了默认值的字段，像 0、”、false 等零值是不会保存到数据库。
// - 您需要使用指针类型或 Scanner/Valuer 来避免这个问题
type User struct {
	gorm.Model
	Name     string `gorm:"type:string;size:20"`
	Age      uint8  `gorm:"type:int;default:18"`
	Birthday time.Time
	Active   bool
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

func (u *User) BeforeDelete(tx *gorm.DB) error {
	if strings.ToLower(u.Name) == "root" {
		return errors.New("root user not allowed to delete")
	}
	return nil
}
