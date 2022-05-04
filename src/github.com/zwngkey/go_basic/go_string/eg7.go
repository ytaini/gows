package main

import (
	"fmt"
	"unicode/utf8"
    "reflect"
    "unsafe"
)

/*
   go中string,本质是一个只读的[]byte,仅比切片少了一个Cap属性,因此它们可以相互转换,[]byte的长度,就是字符串的长度.
   字符串的值一旦确定,就不能进行修改了.

   rune表示一个uft8字符,一个rune字符可以由一个或多个byte组成.

   可以通过[]rune(str) 将一个字符串转换为一个[]rune.

   对于同一字面量,不同的字符串变量指向相同的底层数组,这是因为字符串是只读的,为了节省内容,相同的字符串
   通常对应于同一字符串.
*/

func test(){
    s := "hello,world"
    s1 := "hello,world"
    s2 := "hello,world"[7:]
    //4332692213
    //4332692213 
    //4332692220 
    fmt.Printf("%d \n", (*reflect.StringHeader)(unsafe.Pointer(&s)).Data)
    fmt.Printf("%d \n", (*reflect.StringHeader)(unsafe.Pointer(&s1)).Data)
    fmt.Printf("%d \n", (*reflect.StringHeader)(unsafe.Pointer(&s2)).Data)
    fmt.Printf("%T\n",s2) //string
    
}

func main(){
    test()
}

func test1() {
	x := "test你好"
	//x[0]= 'T' //error
	a := x[0]
	fmt.Printf("%v,%T\n", a, a) //返回116 而不是一个字符  uint8也就是byte
	//要想返回一个字符
	fmt.Printf("%c\n", a) //t
	fmt.Printf("%q\n", a) //'t'
	fmt.Println("-------------------------------------")
	fmt.Printf("%c\n", x[4])
	fmt.Printf("%q\n", x[4])

	fmt.Println("-------------------------------------")

	fmt.Println(len(x))                    //10
	fmt.Println(utf8.RuneCountInString(x)) //6

	fmt.Println("-------------------------------------")
    
    b := []byte(x)
    c := []rune(x)
    x = string(b)
    x = string(c)
	
    fmt.Println("-------------------------------------")

	for i, v := range x {
		//对字符串进行遍历时,这里的v为当前"字符"
		//而i是这个字符的第一个byte的索引
		fmt.Println(i)
		fmt.Printf("%q\n", v) //
	}

	fmt.Println("-------------------------")

	/*
       每次迭代go都会用utf8解码出一个rune类型的字符,
	   对于它无法理解的任何byte序列,他将返回0xfffd runes(即unicode替换字符),
	   而不是真实的数据.

	   string变量的for range 是unicode码类型遍历
	   string遍历的普通for是字节码(ASCII码)类型变量
	*/
	data := "A\xfe\x02\xff\x04"
	for _, v := range data {
		fmt.Printf("%#x,%q\n", v, v)
	}
	for _, v := range []byte(data) {
		fmt.Printf("%#x,%q\n", v, v)
	}
}
