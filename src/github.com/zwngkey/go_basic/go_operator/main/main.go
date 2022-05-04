package main

import (
	"fmt"

	gooperator "github.com/zwngkey/go_basic/go_operator"
)

func main() {
	// gooperator.Test1()
	var a int8 = ^0
	fmt.Printf("%b\n", a) // -1
	gooperator.Testeg2()
	gooperator.Testeg3()
}
