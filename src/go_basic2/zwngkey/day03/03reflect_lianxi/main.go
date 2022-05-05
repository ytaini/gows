package main

import (
	"fmt"
	"log"
	"reflect"
)

func main() {
	test1 := func(a, b int) {
		fmt.Println("test1 函数被调用")
		fmt.Println(a, b)
	}
	test2 := func(a, b, c, d int, e string) {
		fmt.Println("test2 函数被调用")
		fmt.Println(a, b, c, d, e)
	}
	// testReflectFunc("123", 1, 2, 3, 4, "string")
	testReflectFunc(test1, 1, 2)
	testReflectFunc(test2, 1, 2, 3, 4, "string")
}

func testReflectFunc(call any, args ...any) {
	val := reflect.ValueOf(call)

	args1 := []reflect.Value{}

	for i := 0; i < len(args); i++ {
		args1 = append(args1, reflect.ValueOf(args[i]))
	}

	if val.Kind() != reflect.Func {
		log.Fatalln("call not func")
	}
	val.Call(args1)
}
