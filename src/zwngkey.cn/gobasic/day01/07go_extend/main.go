package main

import "fmt"

//结构体模拟继承
type Animal struct {
	name string
}

func (a Animal) move() {
	fmt.Printf("%v:在运动\n", a.name)
}

type Dog struct {
	Animal
	feet int
}

func (d Dog) jiao() {
	fmt.Printf("%v:汪汪汪\n", d.name)
}

func main() {
	// dog := Animal{
	// 	name: "小黄",
	// }
	// dog.move()
	dog := Dog{
		Animal{"xiaohei"},
		4,
	}
	dog.move()
	dog.jiao()
	fmt.Printf("%T", dog.jiao)
}
