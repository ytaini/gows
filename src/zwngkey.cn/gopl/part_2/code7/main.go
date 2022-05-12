package main

import "fmt"

func main() {

	//定义切片
	// a := []int{122, 231, 551}

	// fmt.Printf("a: %v\n", a)
	// fmt.Printf("a: %v\n", len(a))
	// fmt.Printf("a: %v\n", cap(a))

	arr := [10]int{1, 2, 3, 4, 5, 6}
	s := arr[:5]

	fmt.Printf("arr: %v\n", arr)
	// arr: [1 2 3 4 5 6 0 0 0 0]
	fmt.Printf("s: %v\n", s)
	// s: [1 2 3 4 5]
	fmt.Printf("the cap of the s: %v\n", cap(s))
	// the cap of the s: 10

	reverse(s)

	fmt.Printf("reverse: arr: %v\n", arr)
	// reverse: arr: [5 4 3 2 1 6 0 0 0 0]
	fmt.Printf("reverse: s: %v\n", s)
	// reverse: s: [5 4 3 2 1]

}

func reverse(s []int) {

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
