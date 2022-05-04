package main

import "fmt"

type animal interface {
	speaker
	mover
}

type speaker interface {
	speak()
}

type mover interface {
	move()
}

type cat struct {
}

func (c *cat) speak() {
	fmt.Println("喵喵喵!")
}

func (c cat) move() {
	fmt.Println("猫在move")
}

type dog struct {
}

func (d dog) move() {
	fmt.Println("狗在move")
}

func (d *dog) speak() {
	fmt.Println("汪汪汪!")
}

func da(s speaker) {
	s.speak()
}

func mov(m mover) {
	m.move()
}

func main() {
	var a mover
	fmt.Printf("a: %v\n", a)
	var c cat
	var d dog
	e := cat{}
	(&e).move()
	e.move() //简写,语法糖

	f := &cat{}
	(*f).move()
	f.move() //简写,语法糖

	// da(c) 不能这么传值,指针类型的接收者只能传指针类型.
	// da(d)
	da(&c)
	da(&d)
	// 值类型的接收者可以传值类型和指针类型.因为go内部会把指针对应的值传给对应的形参.
	mov(c)
	mov(d)
	mov(&c)
	mov(&d)
}
