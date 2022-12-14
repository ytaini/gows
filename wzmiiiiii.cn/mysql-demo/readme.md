# GO操作Mysql

Go语言中的database/sql包提供了保证SQL或类SQL数据库的泛用接口，并不提供具体的数据库驱动。使用database/sql包时必须注入（至少）一个数据库驱动。

> 常用的数据库基本上都有完整的第三方实现。

> MySQL驱动

- 下载依赖
```shell
go get -u github.com/go-sql-driver/mysql
```

- 使用MySQL驱动
```go
func Open(driverName, dataSourceName string) (*DB, error)
```
Open打开一个dirverName指定的数据库，dataSourceName指定数据源，一般至少包括数据库文件名和其它连接必要的信息。
```go
import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
   // DSN:Data Source Name
	dsn := "user:password@tcp(127.0.0.1:3306)/dbname"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	defer db.Close()  // 注意这行代码要写在上面err判断的下面
}
```

> 初始化连接

Open函数可能只是验证其参数格式是否正确，实际上并不创建与数据库的连接。如果要检查数据源的名称是否真实有效，应该调用Ping方法。

返回的DB对象可以安全地被多个goroutine并发使用，并且维护其自己的空闲连接池。因此，Open函数应该仅被调用一次，很少需要关闭这个DB对象。

```go
// 定义一个全局变量db，用来保存数据库连接对象。
var db *sql.DB

func main() {
	err := initDB()
	if err != nil {
		log.Fatalln(err)
	}
    log.Println("conn establish success ...")
}

func initDB() (err error) {
	// DSN: Data Source Name
	dsn := "root:root@tcp(0.0.0.0:3306)/atguigudb?charset=utf8mb4&parseTime=True"
	// 不会校验账号密码是否正确
	// 注意！！！这里不要使用:=，我们是给全局变量赋值，然后在main函数中使用全局变量db
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	// 尝试与数据库建立连接(校验dsn是否正确)
	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}
```

其中`sql.DB`是表示连接的数据库对象（结构体实例），它保存了连接数据库相关的所有信息。**它内部维护着一个具有零到多个底层连接的连接池，它可以安全地被多个goroutine同时使用。**

> SetMaxOpenConns

```go
func (db *DB) SetMaxOpenConns(n int)
```
`SetMaxOpenConns()`设置与数据库建立连接的最大数目。 如果n大于0且小于最大闲置连接数，会将最大闲置连接数减小到匹配最大开启连接数的限制。 如果n<=0，不会限制最大开启连接数，默认为0（无限制）。


> SetMaxIdleConns

```go
func (db *DB) SetMaxIdleConns(n int)
```

`SetMaxIdleConns`设置连接池中的最大闲置连接数。 如果n大于最大开启连接数，则新的最大闲置连接数会减小到匹配最大开启连接数的限制。 如果n<=0，不会保留闲置连接。

<br>

> CRUD 代码 ./basic/main.go

<br>

## MySQL预处理

> 什么是预处理？

普通SQL语句执行过程：

- 客户端对SQL语句进行占位符替换得到完整的SQL语句。
- 客户端发送完整SQL语句到MySQL服务端
- MySQL服务端执行完整的SQL语句并将结果返回给客户端。

预处理执行过程：

- 把SQL语句分成两部分，命令部分与数据部分。
- 先把命令部分发送给MySQL服务端，MySQL服务端进行SQL预处理。
- 然后把数据部分发送给MySQL服务端，MySQL服务端对SQL语句进行占位符替换。
- MySQL服务端执行完整的SQL语句并将结果返回给客户端。

> 为什么要预处理？

1. 优化MySQL服务器重复执行SQL的方法，可以提升服务器性能，提前让服务器编译，一次编译多次执行，节省后续编译的成本。
2. 避免SQL注入问题。

<br>

> GO 实现MySQL预处理 

代码: `./prepare/main.go`

<br>

> SQL注入问题

- **我们任何时候都不应该自己拼接SQL语句！**

```go
// 这里演示一个自行拼接SQL语句的示例，编写一个根据name字段查询user表的函数如下：
// sql注入示例
func sqlInjectDemo(name string) {
	sqlStr := fmt.Sprintf("select id, name, age from user where name='%s'", name)
	fmt.Printf("SQL:%s\n", sqlStr)
	var u user
	err := db.QueryRow(sqlStr).Scan(&u.id, &u.name, &u.age)
	if err != nil {
		fmt.Printf("exec failed, err:%v\n", err)
		return
	}
	fmt.Printf("user:%#v\n", u)
}
```

此时以下输入字符串都可以引发SQL注入问题：
```go
sqlInjectDemo("xxx' or 1=1#")
sqlInjectDemo("xxx' union select * from user #")
sqlInjectDemo("xxx' and (select count(*) from user) <10 #")
```

<br>

> 补充：不同的数据库中，SQL语句使用的占位符语法不尽相同。

|    数据库     |  占位符语法  |
|:----------:|:-------:|
|   MySQL    |    ?    |
| PostgreSQL | $1, $2等 |
|   SQLite   |  ? 和$1  |
|   Oracle   |  :name  |

<br>

## GO实现MySQL事务

> 事务相关方法

Go语言中使用以下三个方法实现MySQL中的事务操作.

- 开始事务
```go
func (db *DB) Begin() (*Tx, error)
```
- 提交事务
```go
func (tx *Tx) Commit() error
```
- 回滚事务
```go
func (tx *Tx) Rollback() error
```

<br>

> 实例代码 ./tx/main.go