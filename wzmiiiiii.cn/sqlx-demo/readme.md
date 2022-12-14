# sqlx库使用指南

在项目中我们通常可能会使用database/sql连接MySQL数据库。
本文借助使用sqlx实现批量插入数据的例子，介绍了sqlx中可能被你忽视了的`sqlx.In`和`DB.NamedExec`方法。

> sqlx介绍

在项目中我们通常可能会使用database/sql连接MySQL数据库。sqlx可以认为是Go语言内置database/sql的超集，它在优秀的内置database/sql基础上提供了一组扩展。这些扩展中除了大家常用来查询的Get(dest interface{}, ...) error和Select(dest interface{}, ...) error外还有很多其他强大的功能。

> 安装

```shell
go get -u github.com/jmoiron/sqlx
```

> 基本使用

示例代码: `./basic/main.go`

