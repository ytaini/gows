# Fluent API 模式

不管是传统的建造者模式，还是 Functional Options 模式，我们都没有限定属性的构建顺序，比如：
```go
// 传统建造者模式不限定属性的构建顺序
profile := NewServiceProfileBuilder().
                WithPriority(1).  // 先构建Priority也完全没问题
                WithId("service1").
                ...
// Functional Options 模式也不限定属性的构建顺序
profile := NewServiceProfile("service1", "order",
    Priority(1),  // 先构建Priority也完全没问题
	Status(Normal),
    ...
```

但是在一些特定的场景，对象的属性是要求有一定的构建顺序的，如果违反了顺序，可能会导致一些隐藏的错误。
当然，我们可以与使用者的约定好属性构建的顺序，但这种约定是不可靠的，你很难保证使用者会一直遵守该约定。所以，更好的方法应该是通过接口的设计来解决问题， Fluent API 模式 诞生了。

下面，我们使用 Fluent API 模式进行实现：
```go
type (
    // 关键点1: 为ServiceProfile定义一个Builder对象
	fluentServiceProfileBuilder struct {
        // 关键点2: 将ServiceProfile作为Builder的成员属性
		profile *ServiceProfile
	}
    // 关键点3: 定义一系列构建属性的fluent接口，通过方法的返回值控制属性的构建顺序
	idBuilder interface {
		WithId(id string) typeBuilder
	}
	typeBuilder interface {
		WithType(svcType ServiceType) statusBuilder
	}
	statusBuilder interface {
		WithStatus(status ServiceStatus) endpointBuilder
	}
	endpointBuilder interface {
		WithEndpoint(ip string, port int) regionBuilder
	}
	regionBuilder interface {
		WithRegion(regionId, regionName, regionCountry string) priorityBuilder
	}
	priorityBuilder interface {
		WithPriority(priority int) loadBuilder
	}
	loadBuilder interface {
		WithLoad(load int) endBuilder
	}
	// 关键点4: 定义一个fluent接口返回完成构建的ServiceProfile，在最后调用链的最后调用
	endBuilder interface {
		Build() *ServiceProfile
	}
)


// 关键点5: 为Builder定义一系列构建方法，也即实现关键点3中定义的Fluent接口
func (f *fluentServiceProfileBuilder) WithId(id string) typeBuilder {
	f.profile.Id = id
	return f
}

func (f *fluentServiceProfileBuilder) WithType(svcType ServiceType) statusBuilder {
	f.profile.Type = svcType
	return f
}

func (f *fluentServiceProfileBuilder) WithStatus(status ServiceStatus) endpointBuilder {
	f.profile.Status = status
	return f
}

func (f *fluentServiceProfileBuilder) WithEndpoint(ip string, port int) regionBuilder {
	f.profile.Endpoint = network.EndpointOf(ip, port)
	return f
}

func (f *fluentServiceProfileBuilder) WithRegion(regionId, regionName, regionCountry string) priorityBuilder {
	f.profile.Region = &Region{
		Id:      regionId,
		Name:    regionName,
		Country: regionCountry,
	}
	return f
}

func (f *fluentServiceProfileBuilder) WithPriority(priority int) loadBuilder {
	f.profile.Priority = priority
	return f
}

func (f *fluentServiceProfileBuilder) WithLoad(load int) endBuilder {
	f.profile.Load = load
	return f
}

func (f *fluentServiceProfileBuilder) Build() *ServiceProfile {
	return f.profile
}

// 关键点6: 定义一个实例化Builder对象的工厂方法
func NewFluentServiceProfileBuilder() idBuilder {
	return &fluentServiceProfileBuilder{profile: &ServiceProfile{}}
}
```

实现 Fluent API 模式有 6 个关键点，大部分与传统的建造者模式类似：

- 为 `ServiceProfile` 定义一个 Builder 对象 `fluentServiceProfileBuilder`。
- 把需要构建的 `ServiceProfile` 设计为 Builder 对象 `fluentServiceProfileBuilder` 的成员属性。
- 定义一系列构建属性的 Fluent 接口，通过方法的返回值控制属性的构建顺序，这是实现 Fluent API 的关键。比如 WithId 方法的返回值是 typeBuilder 类型，表示紧随其后的就是 WithType 方法。
- 定义一个 Fluent 接口（这里是 endBuilder）返回完成构建的 ServiceProfile，在最后调用链的最后调用。
- 为 Builder 定义一系列构建方法，也即实现关键点 3 中定义的 Fluent 接口，并在构建方法中返回 Builder 对象指针本身。
- 定义一个实例化 Builder 对象的工厂方法 NewFluentServiceProfileBuilder()，返回第一个 Fluent 接口，这里是 idBuilder，表示首先构建的是 Id 属性。

Fluent API 的使用与传统的建造者实现使用类似，但是它限定了方法调用的顺序。如果顺序不对，在编译期就报错了，这样就能提前把问题暴露在编译器，减少了不必要的错误使用。

```go
// Fluent API的使用方法
profile := NewFluentServiceProfileBuilder().
	WithId("service1").
	WithType("order").
	WithStatus(Normal).
	WithEndpoint("192.168.0.1", 8080).
	WithRegion("region1", "beijing", "China").
	WithPriority(1).
	WithLoad(100).
	Build()

// 如果方法调用不按照预定的顺序，编译器就会报错
profile := NewFluentServiceProfileBuilder().
	WithType("order").
	WithId("service1").
	WithStatus(Normal).
	WithEndpoint("192.168.0.1", 8080).
	WithRegion("region1", "beijing", "China").
	WithPriority(1).
	WithLoad(100).
	Build()
// 上述代码片段把WithType和WithId的调用顺序调换了，编译器会报如下错误
// NewFluentServiceProfileBuilder().WithType undefined (type idBuilder has no field or method WithType)
```

> 缺点

- 传统的建造者模式需要新增一个 Builder 对象来完成对象的构造，Fluent API 模式下甚至还要额外增加多个 Fluent 接口，一定程度上让代码更加复杂了。