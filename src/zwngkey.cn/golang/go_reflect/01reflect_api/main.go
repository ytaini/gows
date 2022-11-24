/*
 * @Author: wzmiiiiii
 * @Date: 2022-11-21 19:34:16
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-11-22 05:54:32
 * @Description:
 */
package main

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (p person) String() string {
	return "person"
}

func main() {
	f1()
}

// json.Unmarshal函数中,通过反射获取第二个参数的类型与值.
func f1() {
	pStr := `{"name":"张三","age":18}`

	var p person

	json.Unmarshal([]byte(pStr), &p)

	fmt.Printf("p: %#v\n", p)

	fmt.Printf("Sprint(p): %v\n", Sprint(p))
	fmt.Printf("Sprint(p): %v\n", Sprint("hello"))
	fmt.Printf("Sprint(p): %v\n", Sprint(true))
	fmt.Printf("Sprint(p): %v\n", Sprint(int8(8)))
}

// 模拟fmt.Println()
// 当其他包调用这个函数时,他能把新创建的类型传进来,此时我们该如何判断形参的类型呢?
func Sprint(x any) string {
	type stringer interface {
		String() string
	}

	switch x := x.(type) {
	case stringer:
		return x.String()

	case string:
		return x
	case int:
		return strconv.Itoa(x)
	case bool:
		if x {
			return "true"
		}
		return "false"
	default:
		return "???"
	}
}
