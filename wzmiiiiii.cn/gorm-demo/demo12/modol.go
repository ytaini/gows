package main

import (
	"errors"
	"strings"
	"time"

	"gorm.io/gorm"
)

var birthday = ParseStringData(1999, 12, 13)

type User struct {
	gorm.Model
	Name     string `gorm:"type:string;size:20"`
	Age      uint8  `gorm:"type:int;default:18"`
	Birthday time.Time
	Active   bool
}

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
