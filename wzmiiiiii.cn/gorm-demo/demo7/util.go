package main

import (
	"fmt"
	"time"
)

const format = "2006-01-02"

func ParseStringData(year, month, day int) (t time.Time) {
	birthday := fmt.Sprintf("%d-%d-%d", year, month, day)
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		panic(err)
	}
	t, err = time.ParseInLocation(format, birthday, loc)
	if err != nil {
		panic(err)
	}
	return
}

func GenerateUsers() (users []*User) {
	users = make([]*User, 10000)
	for i := 0; i < 10000; i++ {
		users[i] = NewUser(fmt.Sprintf("name%d", i), 18, birthday)
	}
	return
}
