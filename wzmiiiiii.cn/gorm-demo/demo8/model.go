package main

import "gorm.io/gorm"

// 创建关联数据时，如果关联值是非零值，这些关联会被 upsert，且它们的 Hook 方法也会被调用

type CreditCard struct {
	gorm.Model
	Number string
	UserID uint // 会作为外键
}

type User struct {
	gorm.Model
	Name       string
	CreditCard CreditCard //关联表
}

// UserInfo
// 注意: 若要数据库有默认、虚拟/生成的值，你必须为字段设置 default 标签。
// - 若要在迁移时跳过默认值定义，你可以使用 default:(-)，例如：
type UserInfo struct {
	ID        string `gorm:"default:uuid_generate_v3()"` // db func
	FirstName string
	LastName  string
	Age       uint8
	// 注意: 使用虚拟/生成的值时，你可能需要禁用它的创建、更新权限
	FullName string `gorm:"->;type:GENERATED ALWAYS AS (concat(firstname,' ',lastname));default:(-);"`
}

type UserInformation struct {
	ID   string
	Name string
	Age  int
}
