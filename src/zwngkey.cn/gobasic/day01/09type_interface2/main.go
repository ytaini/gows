package main

import "fmt"

type speaker interface {
	speak()
}

type person struct {
	name string
}

type dog struct {
	name string
}

func (p *person) speak() {

}

func (d *dog) speak() {

}

func main() {
	var s1 speaker
	var s2 speaker
	fmt.Printf("s: %v\n", s1) //<nil>
	fmt.Printf("s: %T\n", s1) //<nil>
	// s = dog{}
	// s = person{}
	s1 = &dog{
		name: "小黑",
	}
	s2 = &person{
		name: "李四",
	}
	fmt.Printf("s1: %v\n", s1)
	fmt.Printf("s1: %T\n", s1)
	fmt.Printf("s1: %v\n", s2)
	fmt.Printf("s1: %T\n", s2)

	d, ok := s1.(*dog)
	if !ok {
		fmt.Println("err")
		return
	}
	fmt.Printf("d: %v\n", d)
	fmt.Printf("d: %v\n", d.name)
}
