/*
 * @Author: wzmiiiiii
 * @Date: 2022-12-06 22:40:51
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-12-06 23:42:30
 */
package type1

// 优点: 不会并发安全问题.
// 缺点: 没有Lazy Loading. 可能造成内存浪费

// 饿汉式(通过包级全局变量+私有类型+一个公有函数)
type singleton map[string]string

var instance = make(singleton)

func New() singleton {
	return instance
}

// 或者
type singleton1 struct{}

var instance1 = &singleton1{}

func New1() *singleton1 {
	return instance1
}
