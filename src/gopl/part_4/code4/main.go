package main

import "fmt"

// 版本一:
// type Circle struct {
// 	X, Y, Radius int
// }

// type Wheel struct {
// 	X, Y, Radius, Spokes int
// }

// 版本二
// type Point struct {
// 	X, Y int
// }

// type Circle struct {
// 	Point  Point
// 	Radius int
// }

// type Wheel struct {
// 	Circle Circle
// 	Spokes int
// }

type Point struct {
	X, Y int
}

type circle struct {
	Point
	Radius int
}

type Wheel struct {
	circle
	Spokes int
}

func main() {
	var w Wheel
	w.X = 1
	w.Y = 2
	w.Radius = 3
	w.Spokes = 4

	//匿名成员的结构体字面值
	a := Wheel{circle{Point{1, 2}, 3}, 4}
	fmt.Printf("a: %#v\n", a)

	b := Wheel{
		circle: circle{
			Point:  Point{X: 8, Y: 9},
			Radius: 9,
		},
		Spokes: 10,
	}
	fmt.Printf("b: %#v\n", b)
	// 访问每个成员变得很繁琐
	// w.Circle.Point.X = 6
	// w.Circle.Point.Y = 8
	// w.Circle.Radius = 7
	// w.Spokes = 9
}
