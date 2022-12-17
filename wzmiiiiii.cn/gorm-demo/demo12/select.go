package main

import (
	"log"
	"time"

	"gorm.io/gorm/clause"

	"gorm.io/gorm"
)

// SelectDemo1
// 检索单个对象
// GORM 提供了 First、Take、Last 方法，以便从数据库中检索单个对象。
// 当查询数据库时它添加了 LIMIT 1 条件，且没有找到记录时，它会返回 ErrRecordNotFound 错误
// 注意: 如果你想避免ErrRecordNotFound错误，你可以使用Find，比如db.Limit(1).Find(&user)，Find方法可以接受struct和slice的数据。
func SelectDemo1() (err error) {
	var user User

	// First()
	log.Println(db.ToSQL(func(tx *gorm.DB) *gorm.DB {
		return tx.First(&User{})
	}))
	// SELECT * FROM `users` WHERE `users`.`deleted_at` IS NULL ORDER BY `users`.`id` LIMIT 1
	err = db.First(&user).Error // 获取第一条记录（主键升序）
	if err != nil {
		return
	}
	log.Println(user)

	// Take()
	log.Println(db.ToSQL(func(tx *gorm.DB) *gorm.DB {
		return tx.Take(&User{})
	}))
	// SELECT * FROM `users` WHERE `users`.`deleted_at` IS NULL LIMIT 1
	err = db.Take(&user).Error // 获取一条记录，没有指定排序字段
	if err != nil {
		return err
	}
	log.Println(user)

	// Last()
	log.Println(db.ToSQL(func(tx *gorm.DB) *gorm.DB {
		return tx.Last(&User{})
	}))
	// SELECT * FROM `users` WHERE `users`.`deleted_at` IS NULL ORDER BY `users`.`id` DESC LIMIT 1
	err = db.Last(&user).Error // 获取最后一条记录（主键降序）
	if err != nil {
		return err
	}
	log.Println(user)

	// SELECT * FROM `users` WHERE `users`.`deleted_at` IS NULL LIMIT 1
	log.Println(db.ToSQL(func(tx *gorm.DB) *gorm.DB {
		return tx.Limit(1).Find(&User{})
	}))

	return
}

/*
	First 和 Last 会根据主键排序，分别查询第一条和最后一条记录。
	只有在目标 struct 是指针或者通过 db.Model() 指定 model 时，该方法才有效。
	此外，如果相关 model 没有定义主键，那么将按 model 的第一个字段进行排序

	// works because model is specified using `db.Model()`
	result := map[string]interface{}{}
	db.Model(&User{}).First(&result)
	// SELECT * FROM `users` ORDER BY `users`.`id` LIMIT 1

	// doesn't work
	result := map[string]interface{}{}
	db.Table("users").First(&result)

	// works with Take
	result := map[string]interface{}{}
	db.Table("users").Take(&result)

	// no primary key defined, results will be ordered by first field (i.e., `Code`)
	type Language struct {
	  Code string
	  Name string
	}
	db.First(&Language{})
	// SELECT * FROM `languages` ORDER BY `languages`.`code` LIMIT 1
*/

// SelectDemo2
// 内联条件: 查询条件可以内联到 First 和 Find 等方法中，类似于 Where.
// 如果主键是数字类型，您可以使用 内联条件 来检索对象。 传入字符串参数时，需要特别注意 SQL 注入问题
func SelectDemo2() (err error) {
	var user User
	// err = db.First(&user, 10).Error // 都行
	err = db.First(&user, "10").Error
	// SELECT * FROM `users` WHERE `users`.`id` = 10 AND `users`.`deleted_at` IS NULL ORDER BY `users`.`id` LIMIT 1
	log.Println(db.ToSQL(func(tx *gorm.DB) *gorm.DB {
		return tx.First(&User{}, 10)
	}))
	if err != nil {
		return
	}
	log.Println(user)

	// err = db.	err = db.Find(&user, "10").Error(&user, 10).Error // 都行
	err = db.Find(&user, "10").Error
	// SELECT * FROM `users` WHERE `users`.`id` = 10 AND `users`.`deleted_at` IS NULL
	log.Println(db.ToSQL(func(tx *gorm.DB) *gorm.DB {
		return tx.Find(&User{}, 10)
	}))
	if err != nil {
		return
	}
	log.Println(user)
	return
}

// SelectDemo3
// 如果主键是字符串（例如像 uuid），查询将被写成这样：
func SelectDemo3() (err error) {
	var user User
	err = db.Find(&user, "id = ?", "1b74413f-f3b8-409f-ac47-e8c062e3472a").Error
	log.Println(db.ToSQL(func(tx *gorm.DB) *gorm.DB {
		return tx.Find(&User{}, "id = ?", "1b74413f-f3b8-409f-ac47-e8c062e3472a")
	}))
	if err != nil {
		return
	}
	log.Println(user)
	return
}

// 当目标对象有一个主要值时，将使用主键构建条件
/*
	var user = User{ID: 10}
	db.First(&user)
	// SELECT * FROM users WHERE id = 10;

	var result User
	db.Model(User{ID: 10}).First(&result)
	// SELECT * FROM users WHERE id = 10;
*/

// SelectDemo4
// 检索全部对象
func SelectDemo4() (err error) {
	//var users []*User // 都行
	var users []User
	ret := db.Find(&users)
	log.Println(ret.RowsAffected)
	for _, user := range users {
		log.Println(user)
	}
	return ret.Error
}

// 条件

// SelectDemo5
// string 条件
func SelectDemo5() {
	var user User
	var users []User
	// SELECT * FROM users WHERE name = 'jinzhu' ORDER BY id LIMIT 1;
	db.Where("name = ?", "jinzhu").First(&user)

	// SELECT * FROM users WHERE name <> 'jinzhu';
	db.Where("name <> ?", "jinzhu").Find(&users)

	// ELECT * FROM users WHERE name IN ('jinzhu','jinzhu 2');
	db.Where("name in ?", []string{"jinzhu", "jinzhu2"}).Find(&users)

	// SELECT * FROM users WHERE name LIKE '%jin%';
	db.Where("name LIKE ?", "%jin%").Find(&users)

	// SELECT * FROM users WHERE name = 'jinzhu' AND age >= 22;
	db.Where("name = ? AND age >= ?", "jinzhu", "22").Find(&users)

	lastWeek := time.Now()
	today := time.Now()

	// SELECT * FROM users WHERE updated_at > '2000-01-01 00:00:00';
	db.Where("updated_at < ?", lastWeek).Find(&users)

	// SELECT * FROM users WHERE created_at BETWEEN '2000-01-01 00:00:00' AND '2000-01-08 00:00:00';
	db.Where("created_at BETWEEN ? AND ?", lastWeek, today).Find(&users)
}

// 如果设置了user对象的主键，而且设置了查询条件, 条件查询不会覆盖主键的值，而是将其作为“与”条件
/*
	var user = User{ID: 10}
	db.Where("id = ?", 20}.First(&user)
	// SELECT * FROM users WHERE id = 10 and id = 20 ORDER BY id ASC LIMIT 1
	// 此查询会给出 record not found Error。
	// 因此，在要使用诸如 user 之类的变量从数据库中获取新值之前，请将主键属性（例如 id）设置为 nil。
*/

// SelectDemo6
// Struct & Map 条件
func SelectDemo6() {
	var user User
	var users []User
	// SELECT * FROM users WHERE name = "jinzhu" AND age = 20 ORDER BY id LIMIT 1;
	db.Where(&User{Name: "jinzhu", Age: 20}).First(&user)

	db.Where(map[string]interface{}{"name": "jinzhu", "age": 20}).Find(&users)
	// SELECT * FROM users WHERE name = "jinzhu" AND age = 20;

	db.Where([]int64{20, 21, 22}).Find(&users)
	// SELECT * FROM users WHERE id IN (20, 21, 22);

	// 注意: 当使用 struct 查询时，GORM 只会查询非零字段，这意味着如果您的字段值为 0、''、false 或其他零值，则不会用于构建查询条件
	db.Where(&User{Name: "jinzhu", Age: 0}).Find(&users)
	// SELECT * FROM users WHERE name = "jinzhu";

	// 要在查询条件中包含零值，您可以使用映射，它将包含所有键值作为查询条件
	db.Where(map[string]interface{}{"Name": "jinzhu", "Age": 0}).Find(&users)
	// SELECT * FROM users WHERE name = "jinzhu" AND age = 0;
}

// SelectDemo7
// 指定结构体查询字段
// 使用结构体进行搜索时，可以通过将相关字段名称或数据库名称传递给 Where() 来指定要在查询条件中使用结构体中的哪些特定值
func SelectDemo7() {
	var users []User
	// SELECT * FROM users WHERE name = "jinzhu" AND age = 0;
	db.Where(&User{Name: "jinzhu"}, "name", "age").Find(&users)

	db.Where(&User{Name: "jinzhu"}, "Age").Find(&users)
	// SELECT * FROM users WHERE age = 0;
}

// SelectDemo8
// 内联条件
// 内联条件: 查询条件可以内联到 First 和 Find 等方法中，类似于 Where.
func SelectDemo8() {
	var user User
	var users []User
	// Get by primary key if it were a non-integer type
	db.Find(&user, "id = ?", "string_primary_key")
	// SELECT * FROM users WHERE id = 'string_primary_key';

	// Plain SQL
	db.Find(&user, "name = ?", "jinzhu")
	// SELECT * FROM users WHERE name = "jinzhu";

	db.Find(&users, "name <> ? AND age > ?", "jinzhu", 20)
	// SELECT * FROM users WHERE name <> "jinzhu" AND age > 20;

	// Struct
	db.Find(&users, User{Age: 20})
	// SELECT * FROM users WHERE age = 20;

	// Map
	db.Find(&users, map[string]interface{}{"age": 20})
	// SELECT * FROM users WHERE age = 20;
}

// SelectDemo9
// Not 条件
func SelectDemo9() {
	var user User
	var users []User
	db.Not("name = ?", "jinzhu").First(&user)
	// SELECT * FROM users WHERE NOT name = "jinzhu" ORDER BY id LIMIT 1;

	// Not In
	db.Not(map[string]interface{}{"name": []string{"jinzhu", "jinzhu 2"}}).Find(&users)
	// SELECT * FROM users WHERE name NOT IN ("jinzhu", "jinzhu 2");

	// Struct
	db.Not(User{Name: "jinzhu", Age: 18}).First(&user)
	// SELECT * FROM users WHERE name <> "jinzhu" AND age <> 18 ORDER BY id LIMIT 1;

	// Not In slice of primary keys
	db.Not([]int64{1, 2, 3}).First(&user)
	// SELECT * FROM users WHERE id NOT IN (1,2,3) ORDER BY id LIMIT 1;
}

// SelectDemo10
// Or 条件
func SelectDemo10() {
	var users []User
	db.Where("role = ?", "admin").Or("role = ?", "super_admin").Find(&users)
	// SELECT * FROM users WHERE role = 'admin' OR role = 'super_admin';

	// Struct
	db.Where("name = 'jinzhu'").Or(User{Name: "jinzhu 2", Age: 18}).Find(&users)
	// SELECT * FROM users WHERE name = 'jinzhu' OR (name = 'jinzhu 2' AND age = 18);

	// Map
	db.Where("name = 'jinzhu'").Or(map[string]interface{}{"name": "jinzhu 2", "age": 18}).Find(&users)
	// SELECT * FROM users WHERE name = 'jinzhu' OR (name = 'jinzhu 2' AND age = 18);
}

// SelectDemo11
// 选择特定字段
// Select 允许您指定要从数据库中检索的字段。 否则，GORM 将默认选择所有字段。
func SelectDemo11() {
	var users []User
	db.Select("name", "age").Find(&users)
	// SELECT name, age FROM users;

	db.Select([]string{"name", "age"}).Find(&users)
	// SELECT name, age FROM users;

	db.Table("users").Select("COALESCE(age,?)", 42).Rows()
	// SELECT COALESCE(age,'42') FROM users;
}

// SelectDemo12
// 排序
// 从数据库中检索记录时指定顺序
func SelectDemo12() {
	var users []User
	db.Order("age desc, name").Find(&users)
	// SELECT * FROM users ORDER BY age desc, name;

	// Multiple orders
	db.Order("age desc").Order("name").Find(&users)
	// SELECT * FROM users ORDER BY age desc, name;

	db.Clauses(clause.OrderBy{
		Expression: clause.Expr{SQL: "FIELD(id,?)", Vars: []interface{}{[]int{1, 2, 3}}, WithoutParentheses: true},
	}).Find(&User{})
	// SELECT * FROM users ORDER BY FIELD(id,1,2,3)
}

// SelectDemo13
// Limit & Offset
func SelectDemo13() {
	var users []User
	var users1 []User
	var users2 []User
	db.Limit(3).Find(&users)
	// SELECT * FROM users LIMIT 3;

	// Cancel limit condition with -1
	db.Limit(10).Find(&users1).Limit(-1).Find(&users2)
	// SELECT * FROM users LIMIT 10; (users1)
	// SELECT * FROM users; (users2)

	db.Offset(3).Find(&users)
	// SELECT * FROM users OFFSET 3;

	db.Limit(10).Offset(5).Find(&users)
	// SELECT * FROM users OFFSET 5 LIMIT 10;

	// Cancel offset condition with -1
	db.Offset(10).Find(&users1).Offset(-1).Find(&users2)
	// SELECT * FROM users OFFSET 10; (users1)
	// SELECT * FROM users; (users2)

}

// Group By & Having
/*
	type result struct {
	  Date  time.Time
	  Total int
	}

	db.Model(&User{}).Select("name, sum(age) as total").Where("name LIKE ?", "group%").Group("name").First(&result)
	// SELECT name, sum(age) as total FROM `users` WHERE name LIKE "group%" GROUP BY `name` LIMIT 1


	db.Model(&User{}).Select("name, sum(age) as total").Group("name").Having("name = ?", "group").Find(&result)
	// SELECT name, sum(age) as total FROM `users` GROUP BY `name` HAVING name = "group"

	rows, err := db.Table("orders").Select("date(created_at) as date, sum(amount) as total").
		Group("date(created_at)").Rows()
	defer rows.Close()
	for rows.Next() {
	  ...
	}

	rows, err := db.Table("orders").Select("date(created_at) as date, sum(amount) as total").
		Group("date(created_at)").Having("sum(amount) > ?", 100).Rows()
	defer rows.Close()
	for rows.Next() {
	  ...
	}

	type Result struct {
	  Date  time.Time
	  Total int64
	}
	db.Table("orders").Select("date(created_at) as date, sum(amount) as total").
		Group("date(created_at)").Having("sum(amount) > ?", 100).Scan(&results)
*/

// Distinct
// Distinct 也适用于 Pluck and Count
// db.Distinct("name", "age").Order("name, age desc").Find(&results)

// Joins

// Scan
// 将结果扫描到结构体中的工作方式与使用 Find 的方式类似
/*
	type Result struct {
	  Name string
	  Age  int
	}

	var result Result
	db.Table("users").Select("name", "age").Where("name = ?", "Antonio").Scan(&result)

	// Raw SQL
	db.Raw("SELECT name, age FROM users WHERE name = ?", "Antonio").Scan(&result)
*/
