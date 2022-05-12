package gointfloat

import "fmt"

func Test1() {
	a := 0xF
	b := 15
	c := 0b_1_1_1_1
	d := 017
	fmt.Println(a == b, a == c, a == d)
}
