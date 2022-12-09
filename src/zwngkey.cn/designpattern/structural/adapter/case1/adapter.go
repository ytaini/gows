/*
 * @Author: wzmiiiiii
 * @Date: 2022-12-09 02:03:33
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-12-09 03:22:26
 */
package adapter

// Target 是适配的目标接口
type Target interface {
	Request() string
}

// Adaptee 被适配的目标接口
type Adaptee interface {
	SpecificRequest() string
}

// NewAdaptee 是被适配接口的工厂函数
func NewAdaptee() Adaptee {
	return &adapteeImpl{}
}

// AdapteeImpl 是被适配的目标类
type adapteeImpl struct{}

// SpecificRequest 是目标类的一个方法
func (*adapteeImpl) SpecificRequest() string {
	return "adaptee method"
}

// NewAdapter 是Adapter的工厂函数
func NewAdapter(adaptee Adaptee) Target {
	return &adapter{
		Adaptee: adaptee,
	}
}

type adapter struct {
	Adaptee
}

func (a *adapter) Request() string {
	return a.SpecificRequest()
}
