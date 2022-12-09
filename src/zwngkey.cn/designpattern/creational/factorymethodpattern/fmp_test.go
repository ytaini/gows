/*
 * @Author: wzmiiiiii
 * @Date: 2022-12-07 01:40:36
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-12-07 01:44:18
 */
package factorymethodpattern

import "testing"

func Test(t *testing.T) {
	// 用户可以根据需要来选择使用哪一个具体工厂来创建产品实例，并调用这些实例的方法
	var factory Factory

	factory = &RectangleFactory{}
	rectangle := factory.CreateShape()
	rectangle.Draw()

	rectangle1 := factory.CreateShape()
	rectangle1.Draw()

	factory = &CircleFactory{}
	circle := factory.CreateShape()
	circle.Draw()

	circle1 := factory.CreateShape()
	circle1.Draw()
}
