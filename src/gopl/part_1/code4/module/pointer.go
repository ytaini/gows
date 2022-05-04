package test_pointer

import "fmt"

func TestPointer() {
	x := 1
	p := &x
	fmt.Println(*p)

	*p = 2

	fmt.Println(x)

}
