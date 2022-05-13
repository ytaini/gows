package gobasictypeliteral

import (
	"fmt"
	"testing"
)

/*
	1.17种内置基本类型（type）各自属于一种Go中的类型种类（kind）。
		尽管所有的内置基本类型的名称都是非导出标识符， 但我们可以不用引入任何代码包而直接使用这些类型。

	2.除了bool和string类型，其它的15种内置基本类型都称为数值类型（整型、浮点数型和复数型）

	3.Go中有两种内置类型别名（type alias）：
		byte是uint8的内置别名。 我们可以将byte和uint8看作是同一个类型。
		rune是int32的内置别名。 我们可以将rune和int32看作是同一个类型。

	4.uintptr、int以及uint类型的值的尺寸依赖于具体编译器实现。
		通常地，在64位的架构上，int和uint类型的值是64位的；在32位的架构上，它们是32位的。
		编译器必须保证uintptr类型的值的尺寸能够存下任意一个内存地址。

	5.一个布尔值表示一个真假。在内存中，一个布尔值只有两种可能的状态。
		这两种状态使用两个预声明（或称为内置）的常量（false和true）来表示
		const (
			true  = 0 == 0 // Untyped bool.
			false = 0 != 0 // Untyped bool.
		)

	6.在内存中，一个字符串存储为一个字节（byte）序列。 此字节序列体现了此字符串所表示的文本的UTF-8编码形式.


	7.尽管布尔和字符串类型分类各自只有一种内置类型， 但我们可以声明定义更多自定义布尔和字符串类型。
		 所以，Go代码中可以出现很多布尔和字符串类型（数值类型也同样）
*/

// 一些类型定义声明
type Status bool     // status和bool是两个不同的类型
type MyString string // MyString和string是两个不同的类型
type Id uint64       // Id和uint64是两个不同的类型
type Real float32    // real和float32是两个不同的类型

// 一些类型别名声明
type boolean = bool // boolean和bool表示同一个类型
type Text = string  // Text和string表示同一个类型
type U8 = uint8     // U8、uint8和 byte表示同一个类型
type char = rune    // char、rune和int32表示同一个类型

/*
	零值(zreo value)
		每种类型都有一个零值。一个类型的零值可以看作是此类型的默认值。
			一个布尔类型的零值表示真假中的假。
			数值类型的零值都是零（但是不同类型的零在内存中占用的空间可能不同）。
			一个字符串类型的零值是一个空字符串。
			其他类型的零值为nil.

*/

/*
	基本类型的字面量表示形式:
		布尔值的字面量形式
			Go白皮书没有定义布尔类型值字面量形式。 我们可以将false和true这两个预声明的有名常量当作布尔类型的字面量形式。
			但是，我们应该知道，从严格意义上说，它们不属于字面量。

		整数类型值的字面量形式
			整数类型值有四种字面量形式：十进制形式（decimal）、八进制形式（octal）、十六进制形式（hex）和二进制形式（binary）
				二进制:0b1111/0B1111
				十进制:15
				八进制:017/0o17/0O17
				十六进制:0xF/OXF

			整数类型的零值的字面量一般使用0表示。 当然，00和0x0等也是合法的整数类型零值的字面量形式。

		浮点数类型值的字面量形式
			一个浮点数的完整十进制字面量形式可能包含一个十进制整数部分、一个小数点、一个十进制小数部分和一个以10为底数的整数指数部分。
			整数指数部分由字母e或者E带一个十进制的整数字面量组成（xEn表示x乘以10^n的意思，而xE-n表示x除以10^n的意思）
				1.23
				01.23 // == 1.23
				.23   // == 0.23
				1.	  // == 1.0
				//一个e或者E随后的数值是指数值（底数为10）。指数值必须为一个可以带符号的十进制整数字面量。
				1.23e2  // == 123.0
				123E2   // == 12300.0
				123.E+2 // == 12300.0
				1e-1    // == 0.1
				.1e0    // == 0.1
				0010e-2 // == 0.1
				0e+5    // == 0.0

			从Go 1.13开始，Go也支持另一种浮点数字面量形式：十六进制浮点数字面量。 在一个十六进制浮点数字面量中，
				这样的一个整数,指数部分由字母p或者P带一个十进制的整数字面量组成（yPn表示y乘以2^n的意思，而yP-n表示y除以2^n的意思）。

				和整数的十六进制字面量一样，一个十六进制浮点数字面量也必须使用0x或者0X开头。
				和整数的十六进制字面量不同的是，一个十六进制浮点数字面量可以包括一个小数点和一个十六进制小数部分。

				0x1p-2     // == 1.0/4 = 0.25
				0x2.p10    // == 2.0 * 1024 == 2048.0
				0x1.Fp+0   // == 1+15.0/16 == 1.9375
				0X.8p1     // == 0 + 8.0/16 * 2 == 1.0
				0X1FFFP-16 // == 0.1249847412109375

				而下面这几个均是不合法的浮点数的十六进制字面量。
				0x.p1    // 整数部分表示必须包含至少一个数字
				1p-2     // p指数形式只能出现在浮点数的十六进制字面量中
				0x1.5e-2 // e和E不能出现在十六进制浮点数字面量的指数部分中

			浮点类型的零值的标准字面量形式为0.0。 当然其它很多形式也是合法的，比如0.、.0、0e0和0x0p0等。


		虚部字面量形式
			一个虚部值的字面量形式由一个浮点数字面量或者一个整数字面量和其后跟随的一个小写的字母i组成。
				1.23i
				1.i
				.23i
				123i
				0123i   // == 123i（兼容性使然。见下）
				1.23E2i // == 123i
				1e-1i
				011i   // == 11i（兼容性使然。见下）
				00011i // == 11i（兼容性使然。见下）
				// 下面这几行从Go 1.13开始才能编译通过。
				0o11i    // == 9i
				0x11i    // == 17i
				0b11i    // == 3i
				0X.8p-0i // == 0.5i

			虚部字面量用来表示复数的虚部。下面是一些复数值的字面量形式：
				1 + 2i       // == 1.0 + 2.0i
				1. - .1i     // == 1.0 + -0.1i
				1.23i - 7.89 // == -7.89 + 1.23i
				1.23i        // == 0.0 + 1.23i


*/

/*
	数值字面表示中使用下划线分段来增强可读性
		Go 1.13开始，下划线_可以出现在整数、浮点数和虚部数字面量中，以用做分段符以增强可读性。
		 但是要注意，在一个数值字面表示中，一个下划线_不能出现在此字面表示的首尾，并且其两侧的字符必须为（相应进制的）数字字符或者进制表示头。
		 // 合法的使用下划线的例子
			6_9          // == 69
			0_33_77_22   // == 0337722
			0x_Bad_Face  // == 0xBadFace
			0X_1F_FFP-16 // == 0X1FFFP-16
			0b1011_0111 + 0xA_B.Fp2i

			// 非法的使用下划线的例子
			_69        // 下划线不能出现在首尾
			69_        // 下划线不能出现在首尾
			6__9       // 下划线不能相连
			0_xBadFace // x不是一个合法的八进制数字
			1_.5       // .不是一个合法的十进制数字
			1._5       // .不是一个合法的十进制数字
*/

/*
	rune类型是int32类型的别名。 因此，rune类型（泛指）是特殊的整数类型。
		一个rune值可以用上面已经介绍的整数类型的字面量形式表示。
			var r rune = 10
			r = 0b1011_0111
			r = 01276
			r = 0x12af
			r = 1e6
		另一方面，各种整数类型的值也可以用rune字面量形式来表示。
		var r1 rune = 'a'
		fmt.Println(r1 == 97) //true

	在Go中，一个rune值表示一个Unicode码点。 一般说来，我们可以将一个Unicode码点看作是一个Unicode字符。
	但是,有些Unicode字符由多个Unicode码点组成。 每个英文或中文Unicode字符值含有一个Unicode码点。

	一个rune字面量由若干包在一对单引号中的字符组成。
	包在单引号中的字符序列表示一个Unicode码点值。
	rune字面量形式有几个变种，其中最常用的一种变种是将一个rune值对应的Unicode字符直接包在一对单引号中
*/
func Test1(te *testing.T) {
	//将Unicode字符直接包在一对单引号中.
	a := 'o'
	b := '中'
	c := 'π'
	fmt.Printf("%c,%q\n", a, a)
	fmt.Printf("%c,%q\n", b, b)
	fmt.Printf("%c,%q\n", c, c)

	//  注意：
	// 		\之后必须跟随三个八进制数字字符（0-7）表示一个byte值，
	//      \x之后必须跟随两个十六进制数字字符（0-9，a-f和A-F）表示一个byte值，
	//		\u之后必须跟随四个十六进制数字字符表示一个rune值（此rune值的高四位都为0），
	//		\U之后必须跟随八个十六进制数字字符表示一个rune值。
	//		这些八进制和十六进制的数字字符序列表示的整数必须是一个合法的Unicode码点值，否则编译将失败。
	d := '\141'
	e := '\x61'
	f := '\u0061'
	g := '\U00000061'
	fmt.Printf("%c,%q\n", d, d)
	fmt.Printf("%c,%q\n", e, e)
	fmt.Printf("%c,%q\n", f, f)
	fmt.Printf("%c,%q\n", g, g)

	//如果一个rune字面量中被单引号包起来的部分含有两个字符， 并且第一个字符是\，第二个字符不是x、 u和U,
	//那么这两个字符将被转义为一个特殊字符
	q := '\a'
	w := '\b'
	r2 := '\f'
	t := '\n'
	y := '\t'
	u := '\\'
	i := '\''
	o := '\r'
	fmt.Printf("%c,%q\n", q, q)
	fmt.Printf("%c,%q\n", w, w)
	fmt.Printf("%c,%q\n", r2, r2)
	fmt.Printf("%c,%q\n", t, t)
	fmt.Printf("%c,%q\n", y, y)
	fmt.Printf("%c,%q\n", u, u)
	fmt.Printf("%c,%q\n", i, i)
	fmt.Printf("%c,%q\n", o, o)

	//rune类型的零值常用 '\000'、'\x00'或'\u0000'等来表示
	var k rune
	fmt.Printf("%c,%q\n", k, k) //'\x00'
}

/*
	字符串值的字面量形式
		在Go中，字符串值是UTF-8编码的， 甚至所有的Go源代码都必须是UTF-8编码的。

		Go字符串的字面量形式有两种。
			一种是解释型字面表示（interpreted string literal，双引号风格）。
			另一种是直白字面表示（raw string literal，反引号风格）

		双引号风格的字符串字面量中支持\"转义，但不支持\'转义；而rune字面量则刚好相反。

		以\、\x、\u和\U开头的rune字面量（不包括两个单引号）也可以出现在双引号风格的字符串字面量中
			// 这几个字符串字面量是等价的。
			"\141\142\143"
			"\x61\x62\x63"
			"\x61b\x63"
			"abc"

		在UTF-8编码中，一个Unicode码点（rune）可能由1到4个字节组成。
		每个英文字母的UTF-8编码只需要一个字节；每个中文字符的UTF-8编码需要三个字节

		直白反引号风格的字面表示中是不支持转义字符的。 除了首尾两个反引号，直白反引号风格的字面表示中不能包含反引号

		字符串类型的零值在代码里用 ""或``表示。
*/

/*
	基本数值类型字面量的适用范围
		每种数值类型有一个能够表示的数值范围。 如果一个字面量超出了一个类型能够表示的数值范围（溢出），
		则在编译时刻，此字面量不能用来表示此类型的值。

	注意几个溢出的例子：
		字面量0x10000000000000000需要65个比特才能表示，所以在运行时刻，任何 基本整数类型 都不能精确表示此字面量。

		在IEEE-754标准中，最大的可以精确表示的float32类型数值为3.40282346638528859811704183484516925440e+38，
		所以 3.5e38 不能表示任何float32和complex64类型的值。

		在IEEE-754标准中，最大的可以精确表示的float64类型数值为1.797693134862315708145274237317043567981e+308，
		因此 2e+308 不能表示任何基本数值类型的值。

		尽管0x10000000000000000可以用来表示float32类型的值，但是它不能被任何float32类型的值所精确表示。
		上面已经提到了，当使用字面量来表示非整数基本数值类型的时候，精度丢失是允许的（但溢出是不允许的）。

*/
func Test2(t *testing.T) {
	var v float32 = 0x10000000000000000
	fmt.Println(v)
}