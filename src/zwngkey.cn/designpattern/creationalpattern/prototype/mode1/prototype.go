/*
 * @Author: wzmiiiiii
 * @Date: 2022-12-07 22:22:34
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-12-07 22:49:23
 */
package mode1

// 原型模式使对象能复制自身，并且暴露到接口中，使客户端面向接口编程时，不知道接口实际对象的情况下生成新的对象。
// 原型模式配合原型管理器使用，使得客户端在不知道具体类的情况下，通过接口管理器得到新的实例，并且包含部分预设定配置。

// Clonable 是原型对象需要实现的接口
type Clonable interface {
	Clone() Clonable
}

type PrototypeManager struct {
	prototypes map[string]Clonable
}

func NewPrototypeManager() *PrototypeManager {
	return &PrototypeManager{
		prototypes: make(map[string]Clonable),
	}
}

func (pm *PrototypeManager) Get(name string) Clonable {
	return pm.prototypes[name].Clone()
}

func (pm *PrototypeManager) Set(name string, prototype Clonable) {
	pm.prototypes[name] = prototype
}
