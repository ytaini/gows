package main

func CreateTableDemo() (err error) {
	err = db.AutoMigrate(&User{})
	if err != nil {
		return
	}
	return
}
