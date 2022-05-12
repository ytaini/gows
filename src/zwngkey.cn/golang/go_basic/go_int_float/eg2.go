package gointfloat

import "fmt"

func Test2() {
	a := 1.23
	b := .24
	c := 1.
	d := 01.23
	fmt.Println(a)
	fmt.Println(b)
	fmt.Printf("%v,%T\n", c, c)
	fmt.Println(d)
	e := 1.23e2
	f := 12.3e2
	g := 123.e+2
	l := 1e-1
	m := .1e0
	n := 0010e-2 //10*e-2
	h := 0e+5    //
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(l)
	fmt.Println(m)
	fmt.Println(n)
	fmt.Println(h)

	// 0x.p1    // 整数部分表示必须包含至少一个数字
	// 1p-2     // p指数形式只能出现在浮点数的十六进制字面量中
	// 0x1.5e-2 // e和E不能出现在十六进制浮点数字面量的指数部分中
	i := 0x1p-2     // 1* 2^-2
	q := 0x2.p10    //2.0 * 2^10
	r := 0x1.Fp+0   //1 + 15/16 * 2^0
	t := 0x.8p1     //8/16 * 2^1
	y := 0x1FFFp-16 //8191 * 2^-16
	fmt.Println(i)
	fmt.Println(q)
	fmt.Println(r)
	fmt.Println(t)
	fmt.Println(y)
}
