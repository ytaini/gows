package main

import (
	"fmt"
)

const (
	a, b = iota + 1, iota + 2
	c, d
	e, f
)

func main() {
	fmt.Println(a, b)
	fmt.Println(c, d)
	fmt.Println(e, f)
}
