/*
 * @Author: wzmiiiiii
 * @Date: 2022-12-09 00:22:52
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-12-09 01:09:19
 */
package builder

// 在程序设计中，我们会经常遇到一些复杂的对象，其中有很多成员属性，甚至嵌套着多个复杂的对象。
// 这种情况下，创建这个复杂对象就会变得很繁琐。对于 C++/Java 而言，最常见的表现就是构造函数有着长长的参数列表：
// MyObject obj = new MyObject(param1, param2, param3, param4, param5, param6, ...)

// 对于 Go 语言来说，最常见的表现就是多层的嵌套实例化：
// obj := &MyObject{
// 	Field1: &Field1 {
// 	  Param1: &Param1 {
// 		Val: 0,
// 	  },
// 	  Param2: &Param2 {
// 		Val: 1,
// 	  },
// 	  ...
// 	},
// 	Field2: &Field2 {
// 	  Param3: &Param3 {
// 		Val: 2,
// 	  },
// 	  ...
// 	},
// 	...
// }

// 上述的对象创建方法有两个明显的缺点：（1）对使用者不友好，使用者在创建对象时需要知道的细节太多；（2）代码可读性很差。

// 针对这种对象成员较多，创建对象逻辑较为繁琐的场景，非常适合使用建造者模式来进行优化。

// 建造者模式的作用有如下几个：1、封装复杂对象的创建过程，使对象使用者不感知复杂的创建逻辑。
// 2、可以一步步按照顺序对成员进行赋值，或者创建嵌套对象，并最终完成目标对象的创建。
// 3、对多个对象复用同样的对象创建逻辑。
// 其中，第1和第2点比较常用，下面对建造者模式的实现也主要是针对这两点进行示例。

// 示例

// 在简单的分布式应用系统（示例代码工程）中，我们定义了服务注册中心，提供服务注册、去注册、更新、 发现等功能。
// 要实现这些功能，服务注册中心就必须保存服务的信息，我们把这些信息放在了 ServiceProfile 这个数据结构上，定义如下：

// ServiceProfile 服务档案，其中服务ID唯一标识一个服务实例，一种服务类型可以有多个服务实例
type ServiceProfile struct {
	Id       string        // 服务ID
	Type     ServiceType   // 服务类型
	Status   ServiceStatus // 服务状态
	Endpoint Endpoint      // 服务Endpoint
	Region   *Region       // 服务所属region
	Priority int           // 服务优先级，范围0～100，值越低，优先级越高
	Load     int           // 服务负载，负载越高表示服务处理的业务压力越大
}

type (
	ServiceType   string
	ServiceStatus int
)

// Region 值对象，每个服务都唯一属于一个Region
type Region struct {
	Id      string
	Name    string
	Country string
}

// Endpoint 值对象，其中ip和port属性为不可变，如果需要变更，需要整对象替换
type Endpoint struct {
	ip   string
	port int
}

// 如果按照直接实例化方式应该是这样的：
// 多层的嵌套实例化
// profile := &ServiceProfile{
// 	Id:       "service1",
// 	Type:     "order",
// 	Status:   Normal,
// 	Endpoint: EndpointOf("192.168.0.1", 8080),
// 	Region: &Region{ // 需要知道对象的实现细节
// 		Id:      "region1",
// 		Name:    "beijing",
// 		Country: "China",
// 	},
// 	Priority: 1,
// 	Load:     100,
// }

// 虽然 ServiceProfile 结构体嵌套的层次不多，但是从上述直接实例化的代码来看，确实存在对使用者不友好和代码可读性较差的缺点。
// 比如，使用者必须先对 Endpoint 和 Region 进行实例化，这实际上是将 ServiceProfile 的实现细节暴露给使用者了。

// 下面我们引入建造者模式对代码进行优化重构：
// 关键点1: 为ServiceProfile定义一个Builder对象
type serviceProfileBuild struct {
	// 关键点2: 将ServiceProfile作为Builder的成员属性
	profile *ServiceProfile
}

// 关键点3: 定义构建ServiceProfile的方法
func (s *serviceProfileBuild) WithId(id string) *serviceProfileBuild {
	s.profile.Id = id
	// 关键点4: 返回Builder接收者指针，支持链式调用
	return s
}

func (s *serviceProfileBuild) WithType(serviceType ServiceType) *serviceProfileBuild {
	s.profile.Type = serviceType
	return s
}

func (s *serviceProfileBuild) WithStatus(status ServiceStatus) *serviceProfileBuild {
	s.profile.Status = status
	return s
}

func (s *serviceProfileBuild) WithEndpoint(ip string, port int) *serviceProfileBuild {
	s.profile.Endpoint = Endpoint{ip, port}
	return s
}

func (s *serviceProfileBuild) WithRegion(regionId, regionName, regionCountry string) *serviceProfileBuild {
	s.profile.Region = &Region{Id: regionId, Name: regionName, Country: regionCountry}
	return s
}

func (s *serviceProfileBuild) WithPriority(priority int) *serviceProfileBuild {
	s.profile.Priority = priority
	return s
}

func (s *serviceProfileBuild) WithLoad(load int) *serviceProfileBuild {
	s.profile.Load = load
	return s
}

// 关键点5: 定义Build方法，在链式调用的最后调用，返回构建好的ServiceProfile
func (s *serviceProfileBuild) Build() *ServiceProfile {
	return s.profile
}

// 关键点6: 定义一个实例化Builder对象的工厂方法
func NewServiceProfileBuilder() *serviceProfileBuild {
	return &serviceProfileBuild{profile: &ServiceProfile{}}
}
