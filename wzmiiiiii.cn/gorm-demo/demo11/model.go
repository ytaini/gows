package main

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"type:string;size:20"`
	Age      uint8  `gorm:"type:int;default:18"`
	Birthday time.Time
	Active   bool
}
type Result struct {
	ID   int
	Name string
	Age  int
}
