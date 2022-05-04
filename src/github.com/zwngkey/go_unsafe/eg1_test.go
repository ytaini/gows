/*
 * @Author: zwngkey
 * @Date: 2022-04-27 13:21:02
 * @LastEditTime: 2022-05-02 15:18:07
 * @Description:
 */
package gounsafe

import (
	"fmt"
	"testing"
	"unsafe"
)

func Test(t *testing.T) {
	var i int8 = 10

	//输出i变量本身的地址的十进制与十六进制
	fmt.Printf("%d,%x\n", &i, &i) //1374390714370,0x14000120002

	//p变量中存的是i变量的地址的十六进制数
	p := unsafe.Pointer(&i)
	fmt.Println(p) //0x14000120002

	//u变量中存的是 i变量中存的值.
	u := uintptr(i)
	fmt.Println(u) //10

	//k变量中存的是i的地址值的十六进制数
	k := &i

	//u变量中存的是 k变量中存的值的十进制数.
	u = uintptr(unsafe.Pointer(k))
	fmt.Println(u) //1374390714370

	//获取指针大小
	u = unsafe.Sizeof(p) //8
	fmt.Println(u)
	//获取变量的大小
	u = unsafe.Sizeof(i) //1
	fmt.Println(u)

	//接下来演示一下内存对齐,猜一猜下面l两个打印值是多少呢?
	person1 := Person1{a: true, b: 1, c: 1, d: "spw"}
	fmt.Println(unsafe.Sizeof(person1))  //40
	fmt.Println(unsafe.Alignof(person1)) //8
	person2 := Person2{b: 1, c: 1, a: true, d: "spw"}
	fmt.Println(unsafe.Sizeof(person2))  //32
	fmt.Println(unsafe.Alignof(person2)) //8
	//第一个结果是40,第二个结果是32，为什么会有这些差距呢？其实就是内存对齐做的鬼，我来详细解释一下

	p2 := Person3{a: true, c: 1}
	fmt.Println(unsafe.Sizeof(p2))  //2
	fmt.Println(unsafe.Alignof(p2)) //1

	//在结构体中，内存对齐是按照结构体中最大字节数对齐的(但不会超过8)

	//指针运算.
	var w = new(W)
	fmt.Println(w.b, w.c)

	//现在我们通过指针运算给b变量赋值为10
	q := uintptr(unsafe.Pointer(w)) //获取了w的指针起始值，
	e := unsafe.Offsetof(w.b)       // 获取b变量的偏移量
	fmt.Println(q)
	fmt.Println(e)
	b := unsafe.Pointer(q + e) //两个相加就得到了b的地址值，
	//通过 * 符号取值，然后赋值，*((*int)(b)) 相当于把（*int） 转换成 int了，最后对变量重新赋值成10，这样指针运算就完成了。
	*((*int)(b)) = 10 //将通用指针Pointer转换成具体指针((*int)(b))

	//此时结果就变成了10，0
	fmt.Println(w.b, w.c)

	//*((*int)(b)) = 10 这行代码的解释
	var d int
	var f *int = &d
	g := unsafe.Pointer(f) //将具体指针*int 转为通用指针Pointer.
	h := (*int)(g)         //将通用指针Pointer转换成具体指针((*int)(b))
	*h = 10
	fmt.Println(d) //10
}

type W struct {
	c int64
	b int32
}

//创建两个结构体
type Person1 struct {
	a bool
	b int64
	c int8
	d string
}
type Person2 struct {
	b int64
	c int8
	a bool
	d string
}
type Person3 struct {
	a bool
	c int8
}
