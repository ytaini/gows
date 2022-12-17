package main

// 删除一条记录.
// 删除一条记录时,删除对象需要指定主键,否则会触发批量Delete.

func DeleteDemo1() (err error) {
	var user User
	err = db.First(&user).Error
	if err != nil {
		return
	}
	err = db.Delete(&user).Error
	// delete from users where id = 1;
	if err != nil {
		return
	}
	// 带额外条件的删除
	err = db.Where("name= ?", "例子").Delete(&User{}).Error
	// delete from users where name = "例子"
	return
}

// 根据主键删除.
// GORM允许通过主键(可以是复合主键)和内联条件来删除对象,它可以使用数字.也可以使用字符串

func DeleteDemo2() (err error) {
	err = db.Delete(&User{}, 10).Error
	// delete from users where id = 10;
	if err != nil {
		return
	}
	err = db.Delete(&User{}, "5").Error
	// delete from users where id = 5;
	if err != nil {
		return
	}

	err = db.Delete(&User{}, []int{4, 5}).Error
	// delete from users where id in (4,5);
	return
}

// Delete Hook
// 对于删除操作，GORM 支持 BeforeDelete、AfterDelete Hook，在删除记录时会调用这些方法
/*
func (u *User) BeforeDelete(tx *gorm.DB) (err error) {
    if u.Role == "admin" {
        return errors.New("admin user not allowed to delete")
    }
    return
}
*/

func BeforeDeleteDemo() error {
	// 这样还是会被删除,beforeDelete中的判断条件是: u.Name = "root"
	//return db.Where("name = ?", "root").Delete(&User{}).Error

	user := User{Name: "root"}
	return db.Delete(&user).Error
}

// 批量删除
// 如果指定的值不包括主属性(主键值)，那么 GORM 会执行批量删除，它将删除所有匹配的记录

func DeleteDemo3() (err error) {
	// 一样的
	//return db.Where("name like ?","%name%").Delete(&User{}).Error
	return db.Delete(&User{}, "name like ?", "%name%").Error
}

// 阻止全局删除
// 如果在没有任何条件的情况下执行批量删除，GORM 不会执行该操作，并返回 ErrMissingWhereClause 错误
// 对此，你必须加一些条件，或者使用原生 SQL，或者启用 AllowGlobalUpdate 模式
/*
	db.Delete(&User{}).Error // gorm.ErrMissingWhereClause

	db.Where("1 = 1").Delete(&User{})
	// DELETE FROM `users` WHERE 1=1

	db.Exec("DELETE FROM users")
	// DELETE FROM users

	db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&User{})
	// DELETE FROM users
*/

// 软删除
// 如果你的模型包含一个gorm.DeleteAt字段(gorm.Model 已经包含了该字段),它将自动获得软删除的能力.
// 拥有软删除能力的模型调用 Delete 时，记录不会从数据库中被真正删除。
// - 但 GORM 会将 DeletedAt 置为当前时间， 并且你不能再通过普通的查询方法找到该记录
/*
	// user 的 ID 是 `111`
	db.Delete(&user)
	// UPDATE users SET deleted_at="2013-10-29 10:23" WHERE id = 111;

	// 批量删除
	db.Where("age = ?", 20).Delete(&User{})
	// UPDATE users SET deleted_at="2013-10-29 10:23" WHERE age = 20;

	// 在查询时会忽略被软删除的记录
	db.Where("age = 20").Find(&user)
	// SELECT * FROM users WHERE age = 20 AND deleted_at IS NULL;
*/

// 如果您不想引入 gorm.Model，您也可以这样启用软删除特性
/*
type User struct {
  ID      int
  Deleted gorm.DeletedAt
  Name    string
}
*/

// 查找被软删除的记录.
// 可以使用Unscoped()找到被软删除的记录.
/*
	db.Unscoped().Where("age = 20").Find(&users)
	// select * from user where age = 20;
*/

// 永久删除.
// 可以使用Unscoped()永久删除匹配的记录
/*
	db.Unscoped().Delete(&order)
	// delete from orders where id = 10;
*/

// Delete Flag
// 默认情况下，gorm.Model 使用 *time.Time 作为 DeletedAt 字段的值。
// 此外，通过 gorm.io/plugin/soft_delete 插件还支持其它数据格式
/*
	提示 当使用 DeletedAt 字段创建唯一复合索引时，你必须通过 gorm.io/plugin/soft_delete 等插件将字段定义为时间戳之类的数据格式，例如：

	import "gorm.io/plugin/soft_delete"

	type User struct {
	  ID        uint
	  Name      string                `gorm:"uniqueIndex:udx_name"`
	  DeletedAt soft_delete.DeletedAt `gorm:"uniqueIndex:udx_name"`
	}
*/

// 将 unix 秒级时间戳作为 delete flag
/*
	import "gorm.io/plugin/soft_delete"

	type User struct {
	  ID        uint
	  Name      string
	  DeletedAt soft_delete.DeletedAt
	}

	// 查询
	SELECT * FROM users WHERE deleted_at = 0;

	// 删除
	UPDATE users SET deleted_at = 当前时间戳  WHERE ID = 1;
*/

// 还可以指定 milli、nano 使用毫秒、纳秒作为值，例如：
/*
	type User struct {
	  ID    uint
	  Name  string
	  DeletedAt soft_delete.DeletedAt `gorm:"softDelete:milli"`
	  // DeletedAt soft_delete.DeletedAt `gorm:"softDelete:nano"`
	}

	// 查询
	SELECT * FROM users WHERE deleted_at = 0;

	// 删除
	UPDATE users SET deleted_at =  当前毫秒、纳秒时间戳  WHERE ID = 1;
*/

// 使用 1 / 0 作为 Delete Flag
/*
	import "gorm.io/plugin/soft_delete"

	type User struct {
	  ID    uint
	  Name  string
	  IsDel soft_delete.DeletedAt `gorm:"softDelete:flag"`
	}

	// 查询
	SELECT * FROM users WHERE is_del = 0;

	// 删除
	UPDATE users SET is_del = 1 WHERE ID = 1;
*/

// 混合模式
/*
	type User struct {
	  ID        uint
	  Name      string
	  DeletedAt time.Time
	  IsDel     soft_delete.DeletedAt `gorm:"softDelete:flag,DeletedAtField:DeletedAt"` // 使用 `1` `0` 标识
	  // IsDel     soft_delete.DeletedAt `gorm:"softDelete:,DeletedAtField:DeletedAt"` // 使用 `unix second` 标识
	  // IsDel     soft_delete.DeletedAt `gorm:"softDelete:nano,DeletedAtField:DeletedAt"` // 使用 `unix nano second` 标识
	}

	// 查询
	SELECT * FROM users WHERE is_del = 0;

	// 删除
	UPDATE users SET is_del = 1, deleted_at = 'current unix second' WHERE ID = 1;
*/

// 返回删除行的数据
// 返回被删除的数据，仅适用于支持 Returning 的数据库
/*
	// 返回所有列
	var users []User
	DB.Clauses(clause.Returning{}).Where("role = ?", "admin").Delete(&users)
	// DELETE FROM `users` WHERE role = "admin" RETURNING *
	// users => []User{{ID: 1, Name: "jinzhu", Role: "admin", Salary: 100}, {ID: 2, Name: "jinzhu.2", Role: "admin", Salary: 1000}}

	// 返回指定的列
	DB.Clauses(clause.Returning{Columns: []clause.Column{{Name: "name"}, {Name: "salary"}}}).Where("role = ?", "admin").Delete(&users)
	// DELETE FROM `users` WHERE role = "admin" RETURNING `name`, `salary`
	// users => []User{{ID: 0, Name: "jinzhu", Role: "", Salary: 100}, {ID: 0, Name: "jinzhu.2", Role: "", Salary: 1000}}
*/