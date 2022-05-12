package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	c := sha256.Sum256([]byte("x"))
	d := sha256.Sum256([]byte("X"))

	// var s []byte = []byte("你好啊")
	// fmt.Printf("s: %v\n", s)

	fmt.Printf("c: %v\n", c)
	fmt.Printf("d: %v\n", d)

	a := [32]byte{1, 2, 3, 4}

	fmt.Printf("a: %v\n", a)
	// a: [1 2 3 4 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
	zero(&a)

	fmt.Printf("a: %v\n", a)
	// a: [0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]

}

//用于将数组清空
func zero(ptr *[32]byte) {
	*ptr = [32]byte{}
}
