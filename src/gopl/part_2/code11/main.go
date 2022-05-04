package main

import "fmt"

func main() {
	s := make([]int, 0)

	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println(cap(s))

	for i := 0; i < 10; i++ {
		s = append(s, i)
		fmt.Println(s)
		fmt.Println(len(s))
		fmt.Println(cap(s))
	}
}
