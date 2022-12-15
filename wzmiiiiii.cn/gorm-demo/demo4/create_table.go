package main

// CreateTableDemo3 使用动态表名.
func CreateTableDemo3() (err error) {
	rootUserInfo := UserInfo{
		Name: "root",
	}
	// 注意： TableName 不支持动态变化，它会被缓存下来以便后续使用。
	// 想要使用动态表名，可以使用 Scopes
	// 基于 UserInfo 中 Name 的值,使用不同的表名.
	// name=root时: 表名=admin_user_infos
	// 其他的,表名=user_infos
	err = db.Scopes(UserInfoTableName(rootUserInfo)).AutoMigrate(&UserInfo{})
	if err != nil {
		return
	}
	otherUserInfo := UserInfo{
		Name: "tom",
	}
	err = db.Scopes(UserInfoTableName(otherUserInfo)).AutoMigrate(&UserInfo{})
	if err != nil {
		return
	}
	return
}

func CreateTableDemo2() (err error) {
	// 临时指定表名
	// 可以使用 Table 方法临时指定表名，例如：

	// 根据 UserInfo 的字段创建 `my_user_infos` 表
	err = db.Table("my_user_infos").AutoMigrate(&UserInfo{})
	if err != nil {
		return
	}
	return
}

func CreateTableDemo1() (err error) {
	// db.AutoMigrate()相当于把结构体变成数据库中的表，如果数据库中不存在表就创建；
	// 如果存在就更新；同时还会自动调整约束索引之类
	err = db.AutoMigrate(&UserInfo{})
	if err != nil {
		return
	}
	return
}
