package main

import "fmt"

type person struct {
	name string
	age  int
}

type dog struct {
	name string
}

func newPerson(name string, age int) *person {
	return &person{
		name,
		age,
	}
}

func newDog(name string) dog {
	return dog{
		name,
	}
}

func (d dog) jiao() {
	fmt.Printf("%s : 再叫\n", d.name)
}

func main() {
	d1 := newDog("小黄")
	d2 := newDog("小黑")
	d1.jiao()
	d2.jiao()
	p1 := newPerson("张三", 19)
	p1.age = 20
	fmt.Printf("p1: %v\n", p1)
}
