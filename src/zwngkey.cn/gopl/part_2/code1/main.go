package main

import "fmt"

func main() {
	var x, y = 1<<1 | 1<<5, 1<<1 | 1<<2
	// 00000001 << 1
	// 00000010
	// 00000001 << 5
	// 00100000 |
	// 00000001
	// => 00100001
	fmt.Printf("%08b\n", x)
	fmt.Printf("%08b\n", y)
}
