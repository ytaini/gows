/*
 * @Author: wzmiiiiii
 * @Date: 2022-12-07 01:33:43
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-12-07 01:53:37
 */
package factorymethodpattern

import "fmt"

// 例如，假设你需要创建一个画图软件，该软件可以画出各种形状，包括矩形、圆形和三角形。
// 那么你可以使用简单工厂模式来实现这个功能。首先，定义一个接口Shape来表示一个形状，该接口包含一个Draw方法来绘制形状
type Shape interface {
	Draw()
}

// 然后，可以定义三个具体的形状类，分别表示矩形、圆形和三角形
type Rectangle struct{}

func (r *Rectangle) Draw() {
	// 绘制矩形
	fmt.Println("绘制矩形")
}

type Circle struct{}

func (c *Circle) Draw() {
	// 绘制圆形
	fmt.Println("绘制圆形")
}

type Triangle struct{}

func (t *Triangle) Draw() {
	// 绘制三角形
	fmt.Println("绘制三角形")
}

// 工厂父类负责定义创建产品对象的公共接口，而工厂子类则负责生成具体的产品对象
// 目的是将产品类的实例化操作延迟到工厂子类中完成，即通过工厂子类来确定究竟应该实例化哪一个具体产品类

// 我们可以定义一个抽象工厂类来声明一个创建产品的接口：
type Factory interface {
	CreateShape() Shape
}

// 接着，我们可以定义两个具体工厂类来实现这个接口，分别用来创建Rectangle和Circle等的实例
type RectangleFactory struct{}

func (rf *RectangleFactory) CreateShape() Shape {
	return &Rectangle{}
}

type CircleFactory struct{}

func (cf *CircleFactory) CreateShape() Shape {
	return &Circle{}
}

type TriangleFactory struct{}

func (tf *TriangleFactory) CreateShape() Shape {
	return &Triangle{}
}
