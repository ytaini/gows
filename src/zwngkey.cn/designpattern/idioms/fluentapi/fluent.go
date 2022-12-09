/*
 * @Author: wzmiiiiii
 * @Date: 2022-12-09 01:24:57
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-12-09 01:34:24
 */
package fluentapi

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
	ServiceStatus string
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

type (
	fluentServiceProfileBuilder struct {
		profile *ServiceProfile
	}
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
	endBuilder interface {
		Build() *ServiceProfile
	}
)

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
	f.profile.Endpoint = Endpoint{ip, port}
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
