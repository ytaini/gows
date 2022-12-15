package main

func CreateTableDemo() (err error) {
	// db.AutoMigrate()相当于把结构体变成数据库中的表，如果数据库中不存在表就创建；
	// 如果存在就更新；同时还会自动调整约束索引之类
	err = db.AutoMigrate(&UserInfo{})
	if err != nil {
		return
	}
	return
}
