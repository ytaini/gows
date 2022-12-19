package main

import (
	_ "embed"
)

// 这是一个指令,而不是普通的注释
//
//go:embed test/1.txt
var s string

//go:embed test/2.txt
var arr []byte

func main() {
	print(s)
	print(string(arr))
}
