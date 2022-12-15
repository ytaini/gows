package main

import "time"

// 时间戳追踪
// 对于有 CreatedAt 字段的模型，创建记录时，如果该字段值为零值，则将该字段的值设为当前时间
// 对于有 UpdatedAt 字段的模型，更新记录时，将该字段的值设为当前时间。创建记录时，如果该字段值为零值，则将该字段的值设为当前时间

//type UserInfo struct {
//	ID        uint
//	Name      string
//	Age       uint
//	CreatedAt time.Time
//	UpdatedAt time.Time
//}

// 可以通过将 autoCreateTime 标签置为 false 来禁用时间戳追踪
// 可以通过将 autoUpdateTime 标签置为 false 来禁用时间戳追踪

type UserInfo struct {
	ID        uint
	Name      string
	Age       uint
	CreatedAt time.Time `gorm:"autoCreateTime:false"`
	UpdatedAt time.Time
}
