package main

import "fmt"

type person struct {
	name string
	age  int
}

func f(x *person) {
	// (*x).name = "李四"
	x.name = "李四"
}

func main() {
	p1 := person{
		name: "张三",
		age:  19,
	}

	fmt.Printf("p1: %v\n", p1)
	f(&p1)
	fmt.Printf("p1: %v\n", p1)
}
