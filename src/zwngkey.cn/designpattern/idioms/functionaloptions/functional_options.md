# Functional Options

**Go 语言中，利用 Functional Options 模式可以更为简洁优雅地完成复杂对象的构建。**

实现 Functional Options 模式有 5 个关键点：

- 定义 Functional Option 类型 ServiceProfileOption，本质上是一个入参为构建对象 ServiceProfile 的指针类型。（注意必须是指针类型，值类型无法达到修改目的）
- 定义构建 ServiceProfile 的工厂方法，以 ServiceProfileOption 的可变参数作为入参。函数的可变参数就意味着可以不传参，因此一些必须赋值的属性建议还是定义对应的函数入参。
- 可为特定的属性提供默认值，这种做法在 为配置对象赋值的场景 比较常见。
- 在工厂方法中，通过 for 循环利用 ServiceProfileOption 完成构建对象的赋值。
- 定义一系列的构建方法，以需要构建的属性作为入参，返回 ServiceProfileOption 对象，并在ServiceProfileOption 中实现属性赋值。


Functional Options 模式 的实例化逻辑是这样的：
```go
// Functional Options 模式的实例化逻辑
profile := NewServiceProfile("service1", "order",
	Status(Normal),
	Endpoint("192.168.0.1", 8080),
	SvcRegion("region1", "beijing", "China"),
	Priority(1),
	Load(100))
```

相比于传统的建造者模式，Functional Options 模式的使用方式明显更加的简洁，也更具“Go 风格”了。
