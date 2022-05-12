package main

import "fmt"

func main() {
	v := 1
	a := incr(&v) // 相当于a=++v
	fmt.Println(v)
	fmt.Println(a)

}

func incr(p *int) int {
	*p++
	return *p
}
