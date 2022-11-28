/*
 * @Author: wzmiiiiii
 * @Date: 2022-11-27 20:04:46
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-11-27 20:14:00
 * @Description:
 */
package main

import (
	"fmt"
	"sync"
)

var spool = &sync.Pool{
	New: func() any {
		fmt.Println("creating a new Person")
		return new(Person)
	},
}

type Person struct {
	Name string
}

func main() {
	p := spool.Get().(*Person)
	fmt.Println("首次从 pool 里获取：", p)

	p.Name = "first"
	fmt.Printf("设置 p.Name = %s\n", p.Name)

	spool.Put(p)

	fmt.Println("Pool 里已有一个对象：&{first}，调用 Get: ", spool.Get().(*Person))
	fmt.Println("Pool 没有对象了，调用 Get: ", spool.Get().(*Person))
}
