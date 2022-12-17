# SQL 构建器
- 原生SQL
  - 注意 GORM 允许缓存预编译 SQL 语句来提高性能
- 命名参数 
- DryRun模式 
  - 在不执行的情况下生成 SQL 及其参数，可以用于准备或测试生成的 SQL
```go
stmt := db.Session(&Session{DryRun: true}).First(&user, 1).Statement
stmt.SQL.String() //=> SELECT * FROM `users` WHERE `id` = $1 ORDER BY `id`
stmt.Vars         //=> []interface{}{1}
```
- ToSQL : 返回生成的 SQL 但不执行。
  - GORM使用 database/sql 的参数占位符来构建 SQL 语句，它会自动转义参数以避免 SQL 注入，但我们不保证生成 SQL 的安全，请只用于调试。
```go
sql := DB.ToSQL(func(tx *gorm.DB) *gorm.DB {
  return tx.Model(&User{}).Where("id = ?", 100).Limit(10).Order("age desc").Find(&[]User{})
})
sql //=> SELECT * FROM "users" WHERE id = 100 AND "users"."deleted_at" IS NULL ORDER BY age desc LIMIT 10
```
- Row &Rows 
```go
// 获取 *sql.Row 结果

// 使用 GORM API 构建 SQL
row := db.Table("users").Where("name = ?", "jinzhu").Select("name", "age").Row()
row.Scan(&name, &age)

// 使用原生 SQL
row := db.Raw("select name, age, email from users where name = ?", "jinzhu").Row()
row.Scan(&name, &age, &email)
```

```go
// 获取 *sql.Rows 结果

/ 使用 GORM API 构建 SQL
rows, err := db.Model(&User{}).Where("name = ?", "jinzhu").Select("name, age, email").Rows()
defer rows.Close()
for rows.Next() {
  rows.Scan(&name, &age, &email)
  
  // 业务逻辑...
}

// 原生 SQL
rows, err := db.Raw("select name, age, email from users where name = ?", "jinzhu").Rows()
defer rows.Close()

for rows.Next() {
  rows.Scan(&name, &age, &email)
  
  // 业务逻辑...
}
```
- 将 sql.Rows 扫描至 model 
```go
// 使用 ScanRows 将一行记录扫描至 struct

rows, err := db.Model(&User{}).Where("name = ?", "jinzhu").Select("name, age, email").Rows() // (*sql.Rows, error)
defer rows.Close()

var user User
for rows.Next() {
  // ScanRows 将一行扫描至 user
  db.ScanRows(rows, &user)
  
  // 业务逻辑...
}
```
- 连接 
```go
// 在一条 tcp DB 连接中运行多条 SQL (不是事务)

db.Connection(func(tx *gorm.DB) error {
  tx.Exec("SET my.role = ?", "admin")
  
  tx.First(&User{})
})
```

# 高级
[文档](https://gorm.io/zh_CN/docs/sql_builder.html#%E9%AB%98%E7%BA%A7)
## 子句（Clause）
GORM 内部使用 `SQL builder` 生成 SQL。对于每个操作，GORM 都会创建一个 `*gorm.Statement` 对象，所有的 GORM API 都是在为 statement 添加、修改 子句，最后，GORM 会根据这些子句生成 SQL

例如，当通过 `First` 进行查询时，它会在 `Statement` 中添加以下子句
```go
clause.Select{Columns: "*"}
clause.From{Tables: clause.CurrentTable}
clause.Limit{Limit: 1}
clause.OrderByColumn{
  Column: clause.Column{Table: clause.CurrentTable, Name: clause.PrimaryKey},
}
```

然后 GORM 在 `Query callback` 中构建最终的查询 SQL，像这样：

```go
Statement.Build("SELECT", "FROM", "WHERE", "GROUP BY", "ORDER BY", "LIMIT", "FOR")
```

生成 SQL:
```sql
SELECT * FROM `users` ORDER BY `users`.`id` LIMIT 1
```

您可以自定义 子句 并与 GORM 一起使用，这需要实现 `clause.Interface` 接口

## 子句构造器
不同的数据库, 子句可能会生成不同的 SQL，例如：
```go
db.Offset(10).Limit(5).Find(&users)
// SQL Server 会生成
// SELECT * FROM "users" OFFSET 10 ROW FETCH NEXT 5 ROWS ONLY
// MySQL 会生成
// SELECT * FROM `users` LIMIT 5 OFFSET 10
```

## 子句选项
GORM 定义了很多 子句，其中一些 子句提供了你可能会用到的选项

尽管很少会用到它们，但如果你发现 GORM API 与你的预期不符合。这可能可以很好地检查它们，例如：
```go
db.Clauses(clause.Insert{Modifier: "IGNORE"}).Create(&user)
// INSERT IGNORE INTO users (name,age...) VALUES ("jinzhu",18...);
```
## StatementModifier
