package main

func CreateTableDemo() (err error) {
	err = db.AutoMigrate(&User{}, &CreditCard{}, &UserInformation{})
	if err != nil {
		return
	}
	return
}
