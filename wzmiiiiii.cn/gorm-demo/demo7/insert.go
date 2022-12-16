package main

import (
	"log"

	"gorm.io/gorm"
)

// 单条插入

// InsertDemo1 创建记录
func InsertDemo1() (err error) {
	user := User{Name: "张三", Age: 13, Birthday: birthday}
	result := db.Create(&user) // 通过数据的指针来创建
	err = result.Error         // 返回 error
	if err != nil {
		return
	}
	log.Printf("插入数据的生成主键: %d\n", user.ID)          // 返回插入数据的主键
	log.Printf("插入记录的条数:%d\n", result.RowsAffected) // 返回插入记录的条数
	return nil
}

// InsertDemo2 用指定的字段创建记录
func InsertDemo2() (err error) {
	// 创建记录并更新给出的字段。
	result := db.Select("Name", "Age", "createdAt").Create(user)

	err = result.Error // 返回 error
	if err != nil {
		return
	}
	log.Printf("插入数据的生成主键: %d\n", user.ID)          // 返回插入数据的主键
	log.Printf("插入记录的条数:%d\n", result.RowsAffected) // 返回插入记录的条数
	return nil
}

func InsertDemo3() (err error) {
	// 创建一个记录且一同忽略传递给略去的字段值。
	result := db.Omit("name", "birthday", "CreatedAt").Create(user)

	err = result.Error // 返回 error
	if err != nil {
		return
	}
	log.Printf("插入数据的生成主键: %d\n", user.ID)          // 返回插入数据的主键
	log.Printf("插入记录的条数:%d\n", result.RowsAffected) // 返回插入记录的条数
	return nil
}

// 批量插入

// BatchInsertDemo1
// 要有效地插入大量记录，请将一个 slice 传递给 Create 方法。
// GORM 将生成单独一条SQL语句来插入所有数据，并回填主键的值，钩子方法也会被调用。
func BatchInsertDemo1() (err error) {
	//result := db.Create(&users) // 结构体指针slice
	//result := db.Create(&users1) // 结构体值slice 都可以
	//result := db.Create(users)  // 可以传slice的指针
	result := db.Create(users1) //	也可以直接传slice

	err = result.Error
	if err != nil {
		return
	}
	log.Printf("插入记录的条数:%d\n", result.RowsAffected) // 返回插入记录的条数
	for _, user := range users1 /*user*/ {
		log.Printf("插入数据的生成主键: %d\n", user.ID) // 返回插入数据的主键
	}
	return nil
}

// BatchInsertDemo2
// 使用 CreateInBatches 分批创建时，你可以指定每批的数量，例如：
func BatchInsertDemo2() (err error) {
	//result := db.CreateInBatches(users, 1) // batchSize 太小不行...
	generateUser := GenerateUsers()
	result := db.CreateInBatches(generateUser, 1000)

	err = result.Error
	if err != nil {
		return
	}
	log.Printf("插入记录的条数:%d\n", result.RowsAffected) // 返回插入记录的条数
	//for _, user := range generateUser {
	//	log.Printf("插入数据的生成主键: %d\n", user.ID) // 返回插入数据的主键
	//}
	return nil
}

// Upsert 和 Create With Associations 也支持批量插入

// 注意: 使用CreateBatchSize 选项初始化 GORM 时，所有的创建& 关联 INSERT 都将遵循该选项
/*
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{
	  CreateBatchSize: 1000,
	})

	db := db.Session(&gorm.Session{CreateBatchSize: 1000})

	users = [5000]User{{Name: "jinzhu", Pets: []Pet{pet1, pet2, pet3}}...}

	db.Create(&users)
	// INSERT INTO users xxx (5 batches)
	// INSERT INTO pets xxx (15 batches)
*/

func BatchInsertDemo3() (err error) {
	result := db.Session(&gorm.Session{CreateBatchSize: 1000}).
		Create(GenerateUsers())

	err = result.Error
	if err != nil {
		return
	}
	log.Println(result.CreateBatchSize)
	log.Printf("插入记录的条数:%d\n", result.RowsAffected) // 返回插入记录的条数
	return nil
}

// GORM 支持根据 map[string]interface{} 和 []map[string]interface{}{} 创建记录，例如：
// 注意： 根据 map 创建记录时，association 不会被调用，且主键也不会自动填充
// 注意: 根据 Map 创建记录,不会对创建/更新等时间进行追踪

func InsertByMapDemo1() (err error) {
	mp := map[string]any{
		"Name":     "map12",
		"Age":      14,
		"birthday": birthday,
	}
	result := db.Model(&User{}).Create(mp)
	err = result.Error
	if err != nil {
		return
	}
	log.Printf("插入记录的条数:%d\n", result.RowsAffected) // 返回插入记录的条数
	return nil
}

func InsertByMapDemo2() (err error) {
	result := db.Model(&User{}).Create([]map[string]any{
		{
			"Name":     "map2",
			"Age":      14,
			"birthday": birthday,
		},
		{
			"Name":     "map3",
			"Age":      14,
			"birthday": birthday,
		},
	})
	err = result.Error
	if err != nil {
		return
	}
	log.Printf("插入记录的条数:%d\n", result.RowsAffected) // 返回插入记录的条数
	return nil
}

// 使用 SQL 表达式、Context Valuer 创建记录
// GORM 允许使用 SQL 表达式插入数据，有两种方法实现这个目标。
// - 根据 map[string]interface{} 或 自定义数据类型 创建，例如:
/*
	// 通过 map 创建记录
	db.Model(User{}).Create(map[string]interface{}{
	  "Name": "jinzhu",
	  "Location": clause.Expr{SQL: "ST_PointFromText(?)", Vars: []interface{}{"POINT(100 100)"}},
	})
	// INSERT INTO `users` (`name`,`location`) VALUES ("jinzhu",ST_PointFromText("POINT(100 100)"));

	// 通过自定义类型创建记录
	type Location struct {
		X, Y int
	}

	// Scan 方法实现了 sql.Scanner 接口
	func (loc *Location) Scan(v interface{}) error {
	  // Scan a value into struct from database driver
	}

	func (loc Location) GormDataType() string {
	  return "geometry"
	}

	func (loc Location) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
	  return clause.Expr{
		SQL:  "ST_PointFromText(?)",
		Vars: []interface{}{fmt.Sprintf("POINT(%d %d)", loc.X, loc.Y)},
	  }
	}

	type User struct {
	  Name     string
	  Location Location
	}

	db.Create(&User{
	  Name:     "jinzhu",
	  Location: Location{X: 100, Y: 100},
	})
	// INSERT INTO `users` (`name`,`location`) VALUES ("jinzhu",ST_PointFromText("POINT(100 100)"))
*/
