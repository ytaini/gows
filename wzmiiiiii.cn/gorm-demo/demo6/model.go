package main

import (
	"database/sql"
	"time"
)

// 模型定义
// 模型是标准的 struct，由 Go 的基本数据类型、实现了 Scanner 和 Valuer 接口的自定义类型及其指针或别名组成
// 例如:

//type UserInfo struct {
//	ID           uint
//	Name         string
//	Email        *string
//	Age          uint8
//	Birthday     time.Time
//	MemberNumber sql.NullString
//	ActivatedAt  sql.NullTime
//	CreatedAt    time.Time
//	UpdatedAt    time.Time
//}

// UserInfo
// 注意 对于声明了默认值的字段，像 0、”、false 等零值是不会保存到数据库。
// 您需要使用指针类型或 Scanner/Valuer 来避免这个问题. 例如:
type UserInfo struct {
	ID           uint
	Name         string
	Email        *string // 这样
	Age          uint8
	Addr         *Addr `gorm:"embedded;embeddedPrefix:addr_"`
	Birthday     time.Time
	MemberNumber sql.NullString // 或这样
	ActivatedAt  sql.NullTime   // 或这样
	CreatedAt    int64          `gorm:"autoCreateTime:milli"`
	UpdatedAt    int
	flag         bool `gorm:"-:all"`
}

type Addr struct {
	City    string
	Country string
}

// GORM 倾向于约定优于配置
// 默认情况下，GORM 使用 ID 作为主键，使用结构体名的 蛇形复数 作为表名，字段名的 蛇形 作为列名，
// 并使用 CreatedAt、UpdatedAt 字段追踪创建、更新时间
// 如果遵循 GORM 的约定，就可以少写的配置、代码。 如果约定不符合实际要求，GORM 允许你配置它们

// GORM 定义一个 gorm.Model 结构体，其包括字段 ID、CreatedAt、UpdatedAt、DeletedAt
// gorm.Model 的定义
//type Model struct {
//	ID        uint `gorm:"primaryKey"`
//	CreatedAt time.Time
//	UpdatedAt time.Time
//	DeletedAt gorm.DeletedAt `gorm:"index"`
//}
// 可以将它嵌入到你的结构体中，以包含这几个字段

//
// 字段级权限控制
// 可导出的字段在使用 GORM 进行 CRUD 时拥有全部的权限，此外，GORM 允许您用标签控制字段级别的权限。
// - 这样您就可以让一个字段的权限是只读、只写、只创建、只更新或者被忽略
// 注意： 使用 GORM Migrator 创建表时，不会创建被忽略的字段

//type User struct {
//	Name string `gorm:"<-:create"`          //允许读和创建
//	Name string `gorm:"<-:update"`          //允许读和更新
//	Name string `gorm:"<-"`                 //允许读和写(创建和更新)
//	Name string `gorm:"<-:false"`           //允许读,禁止写
//	Name string `gorm:"->"`                 //只读（除非有自定义配置，否则禁止写）
//	Name string `gorm:"->;<-:create"`       //允许读和写
//	Name string `gorm:"->:false;<-:create"` // 仅创建（禁止从 db 读）
//	Name string `gorm:"-"`                  // 通过 struct 读写会忽略该字段
//	Name string `gorm:"-:migration"`        // 通过 struct 迁移会忽略该字段
//	Name string `gorm:"-:all"`              // 通过 struct 读写、迁移会忽略该字段
//}

// 创建/更新时间追踪（纳秒、毫秒、秒、Time）
// GORM 约定使用 CreatedAt、UpdatedAt 追踪创建/更新时间。
// - 如果您定义了这种字段，GORM 在创建、更新时会自动填充为 当前时间(time.Now())

// 要使用不同名称的字段，您可以配置 autoCreateTime、autoUpdateTime 标签
// 如果您想要保存 UNIX（毫/纳）秒时间戳，而不是 time.Time，您只需简单地将 time.Time 修改为 int 即可

//type User struct {
//	CreatedAt time.Time // 在创建时，如果该字段值为零值，则使用当前时间填充
//	UpdatedAt int       // 在创建时该字段值为零值或者在更新时，使用当前时间戳秒数填充
//	Updated   int64 `gorm:"autoUpdateTime:nano"` // 使用时间戳填纳秒数充更新时间
//	Updated   int64 `gorm:"autoUpdateTime:milli"` // 使用时间戳毫秒数填充更新时间
//	Created   int64 `gorm:"autoCreateTime"`      // 使用时间戳秒数填充创建时间
//}

// 嵌入结构体
// 对于匿名字段，GORM 会将其字段包含在父结构体中，例如：
//type User struct {
//	gorm.Model
//	Name string
//}
//// 等效于
//type User struct {
//	ID        uint           `gorm:"primaryKey"`
//	CreatedAt time.Time
//	UpdatedAt time.Time
//	DeletedAt gorm.DeletedAt `gorm:"index"`
//	Name string
//}

// 对于正常的结构体字段，你也可以通过标签 embedded 将其嵌入，例如：

//type Author struct {
//	Name  string
//	Email string
//}
//
//type Blog struct {
//	ID      int
//	Author  Author `gorm:"embedded"`
//	Upvotes int32
//}
//// 等效于
//type Blog struct {
//	ID    int64
//	Name  string
//	Email string
//	Upvotes  int32
//}

// 并且，您可以使用标签 embeddedPrefix 来为 db 中的字段名添加前缀，例如：

//type Blog struct {
//	ID      int
//	Author  Author `gorm:"embedded;embeddedPrefix:author_"`
//	Upvotes int32
//}
//// 等效于
//type Blog struct {
//	ID          int64
//	AuthorName  string
//	AuthorEmail string
//	Upvotes     int32
//}

// 字段标签
// 声明 model 时，tag 是可选的，GORM 支持以下 tag： tag 名大小写不敏感，但建议使用 camelCase(驼峰) 风格