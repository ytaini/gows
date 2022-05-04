package gotypeos

import (
	"fmt"
	"time"
)

/*
	类型别名: https://studygolang.com/articles/10282
*/
//如果定义的类型别名是exported (首字母大写)的，那么别的包中就可以使用，它和原始类型是否可exported没关系。
//也就是说，你可以为unexported类型定义一个exported的类型别名
type ty1 struct {
	S string
}
type Ty2 = ty1

// 类型定义的类型的方法集和原始类型的方法集没有任何关系，而类型别名和原始类型的方法集是一样的
// 如果类型别名和原始类型定义了相同的方法，代码编译的时候会报错，因为有重复的方法定义。

type T1 struct{}
type T3 = T1

func (t1 T1) say()       {}
func (t3 *T3) greeting() {}

func TestEg26() {
	var t1 T1
	var t3 T3
	t1.say()
	t1.greeting()
	t3.say()
	t3.greeting()
}

// 另一个有趣的现象是 embedded type, 比如下面的例子， T3是T1的别名。
// 在定义结构体S的时候，我们使用了匿名嵌入类型，那么这个时候调用s.say会怎么样呢？
// 实际是,你会编译出错，因为s.say｀不知道该调用s.T1.say还是s.T3.say`，所以这个时候你需要明确的调用
type S struct {
	T1
	T3
}

func TestEg27() {
	var s *S = &S{
		T1{},
		T3{},
	}
	// s.say() 编译出错
	s.T1.say() //ok
}

//既然类型别名和原类型是相同的，那么在switch - type中，你不能将原类型和类型别名作为两个分支，因为这是重复的case
type D = int

func TestEg25() {
	var x D = 10
	var y interface{} = x
	switch y.(type) {
	case int:
		// case D: 重复了.
	}
}

// 类型别名在定义的时候不允许出现循环定义别名的情况:下面的情况不允许出现.
// type T2 = T1
// type T1 = struct {
// 	next *T2
// }

type MyTime = time.Time
type ITime = MyTime //为别名定义别名

func TestEg24() {
	var t MyTime = time.Now()
	var i ITime = time.Now()
	fmt.Println(t)
	fmt.Println(i)
}

type G = interface{}

func TestEg23() {
	var g G = "hello, go"
	fmt.Println(g)

}

type F = func()

func TestEg22() {
	var f F = func() {
		fmt.Println("func()")
	}
	f()
}

func TestEg21() {
	var x int32 = 32.0
	// var y int = x

	// 为什么?
	var z rune = x //rune是int32的类型别名.

	_ = z
}
