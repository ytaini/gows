package gooperator

import (
	"fmt"
	"testing"
)

//基本类型值的各种运算操作符。

/*
	一个二元运算符运算需要其涉及的两个操作数的类型必须一样时，是指：
		1.如果这两个操作数都是类型确定值，则它们的类型必须相同，或者其中一个操作数可以被隐式转换到另一个操作数的类型
		2.如果其中只有一个操作数是类型确定的，则要么另外一个类型不确定操作数可以表示为此类型确定操作数的类型的值，
			要么此类型不确定操作数的默认类型的任何值可以被隐式转换到此类型确定操作数的类型。
		3.如果这两个操作数都是类型不确定的，则它们必须同时都为两个布尔值，或同时都为两个字符串值，或者同时都为两个基本数字值。

	一个运算符（一元或者二元）运算要求其涉及的某个操作数的类型必须为某个特定类型时，是指：
		1.如果这个操作数是类型确定的，则它的类型必须为所要求的特定类型，或者此操作数可以被隐式转换为所要求的特定类型。
		2.如果这个操作数是类型不确定的，则此操作数要么可以表示为所要求的特定类型值，要么此操作数的默认类型的任何值可以被隐式转换为所要求的特定类型。
*/

/*
	常量表达式:当一个表达式中涉及到的所有操作数都是常量时，此表达式称为一个常量表达式。
			一个常量表达式的估值是在编译阶段进行的。一个常量表达式的估值结果依然是一个常量
	非常量表达式: 如果一个表达式中涉及到的操作数中至少有一个不为常量，则此表达式称为非常量表达式。
*/

func Test1(t *testing.T) {
	//类型确定的操作值
	var (
		a, b float32 = 12.0, 13.1
		h, f float64 = 14.2, 94.1
		c, d int16   = 42, -2
		e    uint8   = 7
		g    byte    = 'a'
	)
	//两个类型确定的操作数
	_ = a * b //同类型
	_ = e * g //同类型
	_ = c % d
	// _ = a * h //error 不同类型不能操作 需要显示转换
	_ = a * float32(f)

	//一个类型确定,一个不确定
	_ = 12 - a  // 12将被当做a的类型（float32）的值使用。
	_ = 'a' - h // 'a'将被当做h的类型（float64）的值使用。

	// 两个类型不确定的操作数.
	_ = 12.3 + 5e-10       // (都为数值类型)
	_ = 12 + 'a'           // (都为数值类型)
	_ = true || false      // (都为布尔类型)
	_ = "string" + "false" // (都为字符串类型)

}

/*

	+,-,*,/ : 这四个运算符操作数的要求:
		- 两个运算数的类型必须相同并且为基本数值类型

	% :
	&,|,^,&^ :
		- 两个运算数的类型必须相同并且为基本 整数 数值类型。

	<<, >> :
		- 左操作数必须为一个整数，右操作数也必须为一个整数（如果它是一个常数，则它必须非负），但它们的类型可以不同。
		- 一个负右操作数（非常数）将在运行时刻造成panic。
		右移运算符规则:
			在右移运算中，左边空出来的位（即高位）全部用左操作数的最高位（即正负号位）填充。
			比如如果左操作数-128的类型为int8（-128二进制补码表示为10000000），
			则10000000 >> 2的二进制补码结果为11100000（即-32）

	^:

*/
func Test2(t *testing.T) {
	var (
		a, b float32 = 12.0, 13.1
		c, d int16   = 42, -2
		e    uint8   = 7
	)
	_, _ = c+int16(e), uint8(c)+e
	_, _, _, _ = a/b, c/d, -100/-9, 1.23/1.2
	_, _, _, _ = c|d, c&d, c^d, c&^d
	_, _, _, _ = d<<e, 123>>e, e>>3, 0xF<<0
	_, _, _, _ = -b, +c, ^e, ^-1

	// 这些行编译将失败。
	/*
		_ = a % b   // error: a和b都不是整数
		_ = a | b   // error: a和b都不是整数
		_ = c + e   // error: c和e的类型不匹配
		_ = b >> 5  // error: b不是一个整数
		_ = c >> -5 // error: -5不是一个无符号整数

		_ = e << uint(c) // 编译没问题
		_ = e << c       // 从Go 1.13开始，此行才能编译过
		_ = e << -c      // 从Go 1.13开始，此行才能编译过。
						 // 将在运行时刻造成恐慌。
		_ = e << -1 // error: 右操作数不能为负（常数）
	*/
}

/*
	溢出:
		一个类型确定数字型常量所表示的值是不能溢出它的类型的表示范围的。
		一个类型不确定数字型常量所表示的值是可以溢出它的默认类型的表示范围的。
			当一个类型不确定数字常量值溢出它的默认类型的表示范围时，此数值不会被截断。
		将一个非常量数字值转换为其它数字类型时，此非常量数字值可以溢出转化结果的类型。
			在此转换中，当溢出发生时，转化结果为此非常量数字值的截断表示

		对于一个算数运算的结果，上述规则同样适用。
*/

func Test3(t *testing.T) {
	// 结果为非常量
	var a, b uint8 = 255, 1
	var c = a + b  // c==0。a+b是一个非常量表达式，结果中溢出的高位比特将被截断舍弃。
	var d = a << b // d == 254。同样，结果中溢出的高位比特将被截断舍弃。
	_, _ = c, d

	// 结果为类型不确定常量，允许溢出其默认类型。
	const X = 0x1FFFFFFFF * 0x1FFFFFFFF // 没问题，尽管X溢出
	const R = 'a' + 0x7FFFFFFF          // 没问题，尽管R溢出

	// 运算结果或者转换结果为类型确定常量
	// var e = X                // error: X溢出int。
	// var h = R                // error: R溢出rune。
	// const Y = 128 - int8(1)  // error: 128溢出int8。
	// const Z = uint8(255) + 1 // error: 256溢出uint8。
}

/*
	关于算术运算的结果:除了移位运算，对于一个二元算术运算:
		- 如果两个操作数都为类型确定值，则此运算的结果也是一个和这两个操作数类型相同的类型确定值。
		- 如果只有一个操作数是类型确定的，则此运算的结果也是一个和此类型确定操作数类型相同的类型确定值。
			另一个类型不确定操作数的类型将被推断为（或隐式转换为）此类型确定操作数的类型。
		- 如果它的两个操作数均为类型不确定值，则此运算的结果也是一个类型不确定值。
			在运算中，两个操作数的类型将被设想为它们的默认类型中一个（按照此优先级来选择：complex128高于float64高于rune高于int）。
			运算结果的默认类型同样为此设想类型。
		  比如，如果一个类型不确定操作数的默认类型为int，另一个类型不确定操作数的默认类型为rune，
			则前者的类型在运算中也被视为rune，运算结果为一个默认类型为rune的类型不确定值。
*/
func Test4(t *testing.T) {
	// 三个类型不确定常量。它们的默认类型
	// 分别为：int、rune和complex64.
	const X, Y, Z = 2, 'A', 3i

	var a, b int = X, Y // 两个类型确定值

	// 变量d的类型被推断为Y的默认类型：rune（亦即int32）。
	d := X + Y
	// 变量e的类型被推断为a的类型：int。
	e := Y - a
	// 变量f的类型和a及b的类型一样：int。
	f := a * b
	// 变量g的类型被推断为Z的默认类型：complex64。
	g := Z * Y

	// 2 65 (+0.000000e+000+3.000000e+000i)
	println(X, Y, Z)
	// 67 63 130 (+0.000000e+000+1.950000e+002i)
	println(d, e, f, g)
}

/*
	对于位移运算: 首先移位运算的结果肯定都是整数。
		- 如果左操作数是一个类型确定值（则它的类型必定为整数），则此移位运算的结果也是一个和左操作数类型相同的类型确定值。
		- 如果左操作数是一个类型不确定值并且右操作数是一个常量，则左操作数将总是被视为一个整数。
			如果它的默认类型不是一个整数（rune或int），则它的默认类型将被视为int。
			此移位运算的结果也是一个类型不确定值并且它的默认类型和左操作数的默认类型一致。
		- 如果左操作数是一个类型不确定值并且右操作数是一个非常量，则左操作数将被首先转化为运算结果的期待设想类型。
			如果期待设想类型并没有被指定，则左操作数的默认类型被视为运算结果的类型。
			如果此期待设想类型不是一个基本整数类型，则编译报错。
			当然最终运算结果是一个左操作数的默认类型的类型确定值。

*/
func Test5(t *testing.T) {
	const N = 2
	// A == 12，它是一个默认类型为int的类型不确定值。
	const Y = 3.0
	const A = Y << N
	// B == 12，它是一个类型为int8的类型确定值。
	const B = int8(3.0) << N

	var m = uint(32)
	// 如果期待设想类型并没有被指定，则左操作数的默认类型将被视为运算结果的类型。
	// 如果此期待设想类型不是一个基本整数类型，则编译报错。
	// var x1 = 1.0 << m //error
	// 下面的三行是相互等价的。
	var x int64 = 1 << m  // 1的类型将被设想为int64，而非int
	var y = int64(1 << m) // 同上
	var z = int64(1) << m // 同上

	// 下面这行编译不通过。
	/*
	   var _ = 1.23 << m // error: 浮点数不能被移位
	*/
	_, _, _, _, _ = x, y, z, A, B
}

/*
	下面这段代码展示了对于左操作数为类型不确定值的移位运算，编译结果因右操作数是否为常量而带来的不同结果：
*/
func TestEg1(t *testing.T) {
	const n = uint(2)
	var m = uint(2)
	_ = m
	// 这两行编译没问题。
	var _ float64 = 1 << n
	var _ = float64(1 << n)

	// 这两行编译失败。
	// var _ float64 = 1 << m  // error
	// var _ = float64(1 << m) // error

	// 上面这段代码最后两行编译失败是因为它们都等价于下面这两行：
	// var _ = float64(1) << m
	// var _ = 1.0 << m // error: shift of type float64
}

//另一个例子：
func TestEg2(t *testing.T) {
	const n = uint(8)
	var m = uint(8)

	var a byte = 1 << n / 128
	var b byte = 1 << m / 128
	fmt.Println(a, b) // 2 0
	// 上面这个程序打印出2 0，因为最后两行等价于：
	// var a = byte(int(1) << n / 128)
	// var b = byte(1) << m / 128
}

/*
	除法和余数运算
		如果除数y是非整数型的0，则运算结果为一个无穷大（+Inf，当被除数不为0时）或者NaN（not a number，当被除数为0时）
*/
func TestEg3(t *testing.T) {
	//余数的符号与除数有关.
	println(5/3, 5%3)     // 1 2
	println(5/-3, 5%-3)   // -1 2
	println(-5/3, -5%3)   // -1 -2
	println(-5/-3, -5%-3) // 1 -2

	println(5.0 / 3.0)               // 1.666667
	fmt.Println((1 - 1i) / (1 + 1i)) // -1i

	var a, b = 1.0, 0.0
	const c = 0.0
	println(a/b, b/b) // +Inf NaN
	println(b / c)    //NaN
	println(a / c)    //+Inf

	_ = int(a) / int(b) // 编译没问题，但在运行时刻将造成恐慌。

	// 这两行编译不通过。
	// println(1.0 / 0.0) // error: 除数为0
	// println(0.0 / 0.0) // error: 除数为0
}

/*
	自增和自减操作符
		和其它语言不一样的是，自增（aNumber++）和自减（aNumber--）操作没有返回值， 所以它们不能当做表达式来使用
*/

/*
	字符串衔接运算符
		两个操作数必须为同一类型的字符串值。

		如果一个字符串衔接运算中的一个操作值为类型确定的，则结果字符串是一个类型和此操作数类型相同的类型确定值。
			否则，结果字符串是一个类型不确定值（肯定是一个常量）。
*/
type Mystring string

func TestEg4(t *testing.T) {
	const a string = "string"
	const a1 = "string"
	b := Mystring("mystring")
	// c := a + b //error 不同类型
	c := a1 + b //c 的类型为Mystring
	const d = a1 + "hello"
	_, _ = a, d
	fmt.Println(c)

}

/*
	比较运算符
		> , < , <= , >=
			两个操作值的类型必须相同并且它们的类型必须为整数类型、浮点数类型或者字符串类型。
		== , !=
			如果两个操作数都为类型确定的，则它们的类型必须一样，或者其中一个操作数可以隐式转换为另一个操作数的类型。
				两者的类型必须都为可比较类型
			如果只有一个操作数是类型确定的，则另一个类型不确定操作数必须可以隐式转换到类型确定操作数的类型。
			如果两个操作数都是类型不确定的，则它们必须同时为两个类型不确定布尔值、两个类型不确定字符串值或者另个类型不确定数字值。

		注意，并非所有的实数在内存中都可以被精确地表示，所以比较两个浮点数或者复数的结果并不是很可靠。
		 在编程中，我们常常比较两个浮点数的差值是否小于一个阙值来检查两个浮点数是否相等。
*/
