/*
 * @Author: wzmiiiiii
 * @Date: 2022-12-07 00:48:22
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-12-07 01:34:58
 */
package sfp

import (
	"fmt"
	"strings"
)

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

// 接下来，定义一个工厂类ShapeFactory，它包含一个静态方法NewShape来创建形状对象
type ShapeFactory struct{}

// 这个方法接收一个参数shapeType，表示要创建的形状类型。
// 然后根据shapeType的值，通过switch语句来决定创建哪种形状。
func (sf *ShapeFactory) NewShape(shapeType string) Shape {
	switch strings.ToLower(shapeType) {
	case "rectangle":
		return &Rectangle{}
	case "circle":
		return &Circle{}
	case "triangle":
		return &Triangle{}
	default:
		return nil
	}
}
