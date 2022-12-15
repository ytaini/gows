package main

import (
	"gorm.io/gorm"
)

// 约定:
// 默认情况下，GORM 会使用 ID 作为表的主键。

//type UserInfo struct {
//	ID   uint
//	Name string
//	Age  uint
//}

//type UserInfo struct {
//	// 可以通过标签 `gorm:"primaryKey"` 将其它字段设为主键
//	UserId uint `gorm:"primaryKey"`
//	Name   string
//	Age    uint
//}

// 复合主键
// 通过将多个字段设为主键，以创建复合主键
// 注意：默认情况下，整型 PrioritizedPrimaryField 启用了 AutoIncrement，
// - 如果需要禁用它，需要为整型字段关闭 autoIncrement

//type UserInfo struct {
//	UserId uint `gorm:"primaryKey;autoIncrement:false"`
//	XueHao uint `gorm:"primaryKey;autoIncrement:false"`
//	Name   string
//	Age    uint
//}

// 复数表名
// GORM 使用结构体名的 蛇形命名 作为表名。对于结构体 UserInfo，根据约定，其表名为 user_infos
// 可以实现 Tabler 接口来更改默认表名

//type UserInfo struct {
//	Id   uint `gorm:"primaryKey"`
//	Name string
//	Age  uint
//}
//
//func (UserInfo) TableName() string {
//	return "table_name"
//}

// 注意： TableName 不支持动态变化，它会被缓存下来以便后续使用。想要使用动态表名，你可以使用 Scopes

type UserInfo struct {
	Id   uint `gorm:"primaryKey"`
	Name string
	Age  uint
}

func UserInfoTableName(userInfo UserInfo) func(tx *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		if userInfo.Name == "root" {
			return tx.Table("admin_user_infos")
		} else {
			return tx.Table("user_infos")
		}
	}
}

// 列名
// 根据约定，数据表的列名使用的是 struct 字段名的 蛇形命名(a_b)
//type UserInfo struct {
//	Id        uint      // 列名:`id`
//	Name      string    // 列名:`name`
//	Age       uint      // 列名:`age`
//	CreatedAt time.Time // 列名:`created_at`
//}

// 可以使用 column 标签或 命名策略 来覆盖列名
//type Animal struct {
//	AnimalId int64     `gorm:"column:beast_id"`  // 将列名设为 `beast_id`
//	Birthday time.Time `gorm:"day_of_the_beast"` // 将列名设为 `day_of_the_beast`
//	Age      int64     `gorm:"age_of_the_beast"` // 将列名设为 `age_of_the_beast`
//}
