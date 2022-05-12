package main

import (
	"fmt"
)

func main() {
	a := [...]int{}

	s := a[:]

	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println(cap(s))

	s = appendInt(s, 6)

	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println(cap(s))
}

func appendInt(x []int, y int) []int {
	var z []int

	zlen := len(x) + 1

	if zlen <= cap(x) { //扩展slice
		z = x[:zlen]
	} else {
		zcap := 2 * len(x)
		if zlen > zcap {
			zcap = zlen
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}

	z[len(x)] = y

	return z
}
