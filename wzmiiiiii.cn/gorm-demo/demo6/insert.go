package main

import (
	"database/sql"
	"log"
	"time"
)

func InsertDemo() (err error) {
	userinfo := UserInfo{
		Name:     "tom",
		Age:      15,
		Email:    new(string),
		Birthday: time.Now(),
		Addr: &Addr{
			City:    "湖南",
			Country: "China",
		},
		MemberNumber: sql.NullString{
			String: "",
			Valid:  true,
		},
		ActivatedAt: sql.NullTime{
			Time:  time.UnixMilli(1124314134),
			Valid: true,
		},
	}
	err = db.Create(&userinfo).Error
	if err != nil {
		return
	}
	log.Println(db.RowsAffected)
	log.Println(userinfo.ID)
	return nil
}
