package goforrange

import (
	"fmt"
	"time"
)

type Person struct {
	Name string
	Age  int
}

func (p *Person) Print() {
	fmt.Println(p)
}

func Eg111() {
	p := []Person{
		{
			"zhangsan",
			18,
		},
		{
			"lisi",
			19,
		},
		{
			"wuwang",
			20,
		},
	}

	for _, v := range p {
		//变量v只声明一次,v的地址一直不变,变得只有v中存的值.
		//整个for循环执行完后,v中存的是p中的最后一个元素.
		//其他的goroutine中,引用了v.此时每个goroutine中的引用的v的值都为p中的最后一个元素.
		go (&v).Print() //都打印 &{wuwang 20}
	}

	for _, v := range p {
		go func(v Person) {
			v.Print() //每次打印不一样
		}(v)
	}

	p1 := []*Person{
		&Person{
			"zs",
			18,
		},
		{
			"ls",
			19,
		},
		{
			"ww",
			20,
		},
	}

	for _, v := range p1 {

		//这里每次取的是v的值.
		go v.Print() //每次打印不一样? 为什么?

		go func() {
			v.Print() //每次打印都一样? 为什么?
		}()
	}

	a := []int{
		1,
		2,
		3,
	}

	for _, v := range a {
		/*
		   for 循环中变量v只声明一次.
		*/
		go func() {
			// 因为闭包对外层词法域变量是引用的，
			// 可以想象 匿名函数 中保存着 v 的地址，它使用 v 时会直接解引用，所以 v 的值改变了会导致 匿名函数 解引用得到的值也会改变。
			fmt.Println(v) // 每次打印3
		}()
		//v每次的值,传入了Println函数中.
		go fmt.Println(v) //每次打印不一样
	}

	for _, v := range a {
		go func(v int) {
			// fmt.Println(&v, "---------")
			fmt.Println(v) // 1 2 3
		}(v)
	}
	time.Sleep(2 * time.Second)
}
