package main

import (
	"log"
	"time"

	"gorm.io/gorm"
)

// SaveDemo
// Save 会保存所有的字段，即使字段是零值.
// 如果主键不存在,Save会插入.
func SaveDemo() (err error) {
	var user User
	err = db.First(&user).Error
	if err != nil {
		return
	}
	user.Age = 12
	user.Name = "张三"
	ret := db.Save(&user)
	log.Println(ret.RowsAffected)
	return ret.Error
}

// 更新单个列
// 使用 Update 更新单个列时，它需要有Where子句，否则会引发错误 ErrMissingWhereClause.
// 当使用 Model 方法并且其值具有主键值时，这时将使用主键来构建Where条件.例如:

func UpdateDemo1() (err error) {
	err = db.Model(&User{}).Where("id = ?", 11).
		Update("name", "updateTest").Error
	// UPDATE users SET name='updateTest', updated_at='2022-12-16 22:38:55.253' WHERE id = 11;
	if err != nil {
		return
	}
	var user User
	db.First(&user)
	// 当使用 Model 方法并且其值具有主键值时，这时将使用主键来构建Where条件
	err = db.Model(&user).Update("name", "张三1").Error
	// UPDATE users SET name='张三1', updated_at='2022-12-16 22:38:55.253' WHERE id = 1;
	if err != nil {
		return err
	}

	// 根据条件和 model 的值进行更新
	// 如果user中有主键值,这时会使用主键 加 Where 中的条件进行更新.
	// 如果没有,则只会根据Where中的条件进行更新.
	err = db.Model(&user).Where("active = ?", true).Update("age", 22).Error
	//  UPDATE users SET age=22, updated_at='2022-12-16 22:38:55.253' WHERE id = 13 AND active = true;

	return
}

// 更新多列
// Updates 方法支持 struct 和 map[string]interface{} 参数。
// 注意 当使用 struct 进行更新时，GORM 只会更新非零值的字段。
// - 你可以使用 map 更新字段，或者使用 Select 指定要更新的字段

func UpdateDemo2() (err error) {
	var user User
	err = db.First(&user).Error
	if err != nil {
		return err
	}
	err = db.Where("name=? and age = ?", "例子", 13).Updates(User{Birthday: time.Now()}).Error
	// 根据 `struct` 更新属性，只会更新非零值的字段
	err = db.Model(&user).Updates(User{Name: "lisi", Age: 31, Birthday: time.Now()}).Error
	err = db.Model(&User{}).Where("id=?", 2).Updates(User{Name: "lisi", Age: 31, Birthday: time.Now()}).Error
	err = db.Model(&User{}).Where("name=? and age = ?", "例子", 22).Updates(User{Birthday: time.Now()}).Error
	// UPDATE users SET name='hello', age=18, updated_at = '2013-11-17 21:34:10' WHERE id = 111;

	//  根据 `map` 更新属性
	err = db.Model(&user).Updates(map[string]interface{}{"name": "hello", "age": 18, "active": false}).Error
	if err != nil {
		return err
	}
	return nil
}

// 更新选定字段
// 如果您想要在更新时选定、忽略某些字段，您可以使用 Select、Omit
/*
	// 使用 Map 进行 Select
	// User's ID is `111`:
	db.Model(&user).Select("name").Updates(map[string]interface{}{"name": "hello", "age": 18, "active": false})
	// UPDATE users SET name='hello' WHERE id=111;

	db.Model(&user).Omit("name").Updates(map[string]interface{}{"name": "hello", "age": 18, "active": false})
	// UPDATE users SET age=18, active=false, updated_at='2013-11-17 21:34:10' WHERE id=111;

	// 使用 Struct 进行 Select（会 select 零值的字段）
	db.Model(&user).Select("Name", "Age").Updates(User{Name: "new_name", Age: 0})
	// UPDATE users SET name='new_name', age=0 WHERE id=111;

	// Select 所有字段（查询包括零值字段的所有字段）
	db.Model(&user).Select("*").Update(User{Name: "jinzhu", Role: "admin", Age: 0})

	// Select 除 Role 外的所有字段（包括零值字段的所有字段）
	db.Model(&user).Select("*").Omit("Role").Update(User{Name: "jinzhu", Role: "admin", Age: 0})
*/

// 更新 Hook
// GORM 支持的 hook 点包括：BeforeSave, BeforeUpdate, AfterSave, AfterUpdate. 更新记录时将调用这些方法
/*
	func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
		if u.Role == "admin" {
			return errors.New("admin user not allowed to update")
		}
		return
	}
*/

// 批量更新
// 如果您尚未通过 Model 指定记录的主键，则 GORM 会执行批量更新
/*
	// 根据 struct 更新
	db.Model(User{}).Where("role = ?", "admin").Updates(User{Name: "hello", Age: 18})
	// UPDATE users SET name='hello', age=18 WHERE role = 'admin';

	// 根据 map 更新
	db.Table("users").Where("id IN ?", []int{10, 11}).Updates(map[string]interface{}{"name": "hello", "age": 18})
	// UPDATE users SET name='hello', age=18 WHERE id IN (10, 11);
*/

// 阻止全局更新
// 如果在没有任何条件的情况下执行批量更新，默认情况下，GORM 不会执行该操作，并返回 ErrMissingWhereClause 错误
// 对此，你必须加一些条件，或者使用原生 SQL，或者启用 AllowGlobalUpdate 模式，例如：
/*
	db.Model(&User{}).Update("name", "jinzhu").Error // gorm.ErrMissingWhereClause

	db.Model(&User{}).Where("1 = 1").Update("name", "jinzhu")
	// UPDATE users SET `name` = "jinzhu" WHERE 1=1

	db.Exec("UPDATE users SET name = ?", "jinzhu")
	// UPDATE users SET name = "jinzhu"

	db.Session(&gorm.Session{AllowGlobalUpdate: true}).Model(&User{}).Update("name", "jinzhu")
	// UPDATE users SET `name` = "jinzhu"
*/

// 更新的记录数
// 获取受更新影响的行数
/*
	// 通过 `RowsAffected` 得到更新的记录数
	result := db.Model(User{}).Where("role = ?", "admin").Updates(User{Name: "hello", Age: 18})
	// UPDATE users SET name='hello', age=18 WHERE role = 'admin';

	result.RowsAffected // 更新的记录数
	result.Error        // 更新的错误
*/

// GORM 允许使用 SQL 表达式更新列

func UpdateDemo3() (err error) {
	err = db.Table("users").Where("id = ?", 1).
		Update("age", gorm.Expr("age / ? * ?", 2, 3)).Error
	// UPDATE "users" SET "age" = age * 2 / 3, "updated_at" = '2013-11-17 21:34:10' WHERE "id" = 1

	err = db.Model(&User{}).Where("id=?", 2).
		Updates(map[string]any{"age": gorm.Expr("age / ? * ?", 2, 3)}).Error
	// UPDATE "users" SET "age" = age * 2 / 3, "updated_at" = '2013-11-17 21:34:10' WHERE "id" = 2
	return
}

// 不使用 Hook 和时间追踪
// 上面的更新操作会自动运行 model 的 BeforeUpdate, AfterUpdate 方法，更新 UpdatedAt 时间戳,
// - 在更新时保存其 Associations, 如果你不想调用这些方法，你可以使用 UpdateColumn， UpdateColumns
// 其用法类似于 Update、Updates

func UpdateDemo4() (err error) {
	err = db.Model(&User{}).Where("age > 40").
		UpdateColumn("age", gorm.Expr("age / ?", 2)).Error
	// Update user set age = age / 2 where age > 40;
	return
}

// GORM 也允许使用自定义数据类型的 Context Valuer 来更新
/*
	// 根据自定义数据类型创建
	type Location struct {
		X, Y int
	}

	func (loc Location) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
	  return clause.Expr{
		SQL:  "ST_PointFromText(?)",
		Vars: []interface{}{fmt.Sprintf("POINT(%d %d)", loc.X, loc.Y)},
	  }
	}

	db.Model(&User{ID: 1}).Updates(User{
	  Name:  "jinzhu",
	  Location: Location{X: 100, Y: 100},
	})
	// UPDATE `user_with_points` SET `name`="jinzhu",`location`=ST_PointFromText("POINT(100 100)") WHERE `id` = 1
*/

// 根据子查询进行更新
/*
	db.Model(&user).Update("company_name", db.Model(&Company{}).Select("name").Where("companies.id = users.company_id"))
	// UPDATE "users" SET "company_name" = (SELECT name FROM companies WHERE companies.id = users.company_id);

	db.Table("users as u").Where("name = ?", "jinzhu").
		Update("company_name", db.Table("companies as c").Select("name").Where("c.id = u.company_id"))

	db.Table("users as u").Where("name = ?", "jinzhu").
		Updates(map[string]interface{}{"company_name": db.Table("companies as c").Select("name").Where("c.id = u.company_id")})
*/

// 在 Update 后修改值
// 若要在 Before 钩子中改变要更新的值，如果它是一个完整的更新，可以使用 Save；否则，应该使用 SetColumn
/*
	func (user *User) BeforeSave(tx *gorm.DB) (err error) {
	  if pw, err := bcrypt.GenerateFromPassword(user.Password, 0); err == nil {
		tx.Statement.SetColumn("EncryptedPassword", pw)
	  }

	  if tx.Statement.Changed("Code") {
		user.Age += 20
		tx.Statement.SetColumn("Age", user.Age)
	  }
	}
*/

// 返回修改行的数据
// 返回被修改的数据，仅适用于支持 Returning 的数据库，例如：
/*
	// 返回所有列
	var users []User
	DB.Model(&users).Clauses(clause.Returning{}).Where("role = ?", "admin").Update("salary", gorm.Expr("salary * ?", 2))
	// UPDATE `users` SET `salary`=salary * 2,`updated_at`="2021-10-28 17:37:23.19" WHERE role = "admin" RETURNING *
	// users => []User{{ID: 1, Name: "jinzhu", Role: "admin", Salary: 100}, {ID: 2, Name: "jinzhu.2", Role: "admin", Salary: 1000}}

	// 返回指定的列
	DB.Model(&users).Clauses(clause.Returning{Columns: []clause.Column{{Name: "name"}, {Name: "salary"}}}).Where("role = ?", "admin").Update("salary", gorm.Expr("salary * ?", 2))
	// UPDATE `users` SET `salary`=salary * 2,`updated_at`="2021-10-28 17:37:23.19" WHERE role = "admin" RETURNING `name`, `salary`
	// users => []User{{ID: 0, Name: "jinzhu", Role: "", Salary: 100}, {ID: 0, Name: "jinzhu.2", Role: "", Salary: 1000}}
*/

// 检查字段是否有变更？
// GORM 提供了 Changed 方法，它可以被用在 Before Update Hook 里，它会返回字段是否有变更的布尔值.
// Changed 方法只能与 Update、Updates 方法一起使用，并且它只是检查 Model() 中对象字段的值与 Update()、Updates() 中的值是否相等。
// 如果值有变更，且字段没有被忽略，则返回 true.
/*
	func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	  // 如果 Role 字段有变更
		if tx.Statement.Changed("Role") {
		return errors.New("role not allowed to change")
		}

	  if tx.Statement.Changed("Name", "Admin") { // 如果 Name 或 Role 字段有变更
		tx.Statement.SetColumn("Age", 18)
	  }

	  // 如果任意字段有变更
		if tx.Statement.Changed() {
			tx.Statement.SetColumn("RefreshedAt", time.Now())
		}
		return nil
	}

	db.Model(&User{ID: 1, Name: "jinzhu"}).Updates(map[string]interface{"name": "jinzhu2"})
	// Changed("Name") => true
	db.Model(&User{ID: 1, Name: "jinzhu"}).Updates(map[string]interface{"name": "jinzhu"})
	// Changed("Name") => false, 因为 `Name` 没有变更
	db.Model(&User{ID: 1, Name: "jinzhu"}).Select("Admin").Updates(map[string]interface{
	  "name": "jinzhu2", "admin": false,
	})
	// Changed("Name") => false, 因为 `Name` 没有被 Select 选中并更新

	db.Model(&User{ID: 1, Name: "jinzhu"}).Updates(User{Name: "jinzhu2"})
	// Changed("Name") => true
	db.Model(&User{ID: 1, Name: "jinzhu"}).Updates(User{Name: "jinzhu"})
	// Changed("Name") => false, 因为 `Name` 没有变更
	db.Model(&User{ID: 1, Name: "jinzhu"}).Select("Admin").Updates(User{Name: "jinzhu2"})
	// Changed("Name") => false, 因为 `Name` 没有被 Select 选中并更新
*/
