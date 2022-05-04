package goconstvar

import "fmt"

/*
	基本类型的字面量中（除了false和true）都属于无名常量（unnamed constant），
	或者叫字面常量（literal constant）。 false和true是预声明的两个有名常量。
*/

/*
	类型不确定值（untyped value）和类型确定值（typed value）:
		在Go中，有些值的类型是不确定的。换句话说，有些值的类型有很多可能性。
		 这些值称为类型不确定值。对于大多数类型不确定值来说，它们各自都有一个默认类型， 除了预声明的nil。nil是没有默认类型的。

		与类型不确定值相对应的概念称为类型确定值

		字面常量（无名常量）都属于类型不确定值。 事实上，Go中大多数的类型不确定值都属于字面常量和本文即将介绍的有名常量。
		 少数类型不确定值包括刚提到的nil和以后会逐步接触到的 某些操作的布尔返回值 。


		一个字面（常）量的默认类型取决于它为何种字面量形式：
			一个字符串字面量的默认类型是预声明的string类型。
			一个布尔字面量的默认类型是预声明的bool类型。
			一个整数型字面量的默认类型是预声明的int类型。
			一个rune字面量的默认类型是预声明的rune（亦即int32）类型。
			一个浮点数字面量的默认类型是预声明的float64类型。
			如果一个字面量含有虚部字面量，则此字面量的默认类型是预声明的complex128类型。
*/
func Testeg1() {
	//字面量的默认类型.
	fmt.Printf("%T,%[1]v\n", 1) //int
	fmt.Printf("%T,%[1]v\n", 0b111)
	fmt.Printf("%T,%[1]v\n", 0111)
	fmt.Printf("%T,%[1]v\n", 0x111)
	fmt.Printf("%T,%[1]v\n", false)    //bool
	fmt.Printf("%T,%[1]v\n", "string") //string
	fmt.Printf("%T,%[1]v\n", 1e10)     //float64
	fmt.Printf("%T,%[1]v\n", 'a')      //int32(rune)
	fmt.Printf("%T,%[1]v\n", '\141')
	fmt.Printf("%T,%[1]v\n", '\x1f')

}

/*
	类型不确定常量的显式类型转换
		一个显式类型转换的形式为T(v)，其表示将一个值v转换为类型T。
			编译器将T(v)的转换结果视为一个类型为T的类型确定值。
			当然，对于一个特定的类型T，T(v)并非对任意的值v都合法。

		对于一个类型不确定常量值v，有两种情形显式转换T(v)是合法的：
			1.v可以表示为T类型的一个值。 转换结果为一个类型为T的类型确定常量值。
			2.v的默认类型是一个整数类型（int或者rune） 并且T是一个字符串类型。
				转换T(v)将v看作是一个Unicode码点。 转换结果为一个类型为string的字符串常量。
				 此字符串常量只包含一个Unicode码点，并且可以看作是此Unicode码点的UTF-8表示形式。
				  对于不在合法的Unicode码点取值范围内的整数v， 转换结果等同于字符串字面量"\uFFFD"（亦即"\xef\xbf\xbd"）。
				  0xFFFD是Unicode标准中的（非法码点的）替换字符值。
				  (但是请注意，今后的Go版本可能只允许rune或者byte整数被转换为字符串。
					从Go官方工具链1.15版本开始，go vet命令会对从非rune和非byte整数到字符串的转换做出警告。）

			事实上，第二种情形并不要求v必须是一个常量。 如果v是一个常量，则转换结果也是一个常量。 如果v不是一个常量，则转换结果也不是一个常量。
					const d = 'a'
					var e = '中'
					const f = string(d) //ok
					const f = e 		//error
					const f = string(e) //error
					var f = string(d) //ok
					var f = string(e) //ok


		一个类型不确定数字值所表示的值可能溢出它的默认类型的表示范围。 比如-1e1000和0x10000000000000000。
			-1e1000的默认类型为float64
			0x10000000000000000的默认类型int
		  一个溢出了它的默认类型的表示范围的类型不确定数字值是不能被转换到它的默认类型的（将编译报错）。
				即	// -1e+1000不能被表示为float64类型值。不允许溢出。
					float64(-1e1000) //error
					// 0x10000000000000000做为int值将溢出。
					int(0x10000000000000000) //error
*/
func Testeg2() {
	// 8 本来是一个类型不确定值.
	// 通过int8(8)显示类型转换后, 8 变为了一个类型为int8的类型确定值.
	const a = int8(8)
	// const b = int8(128) //error: 128超出int8类型表示范围. v 不合法.
	fmt.Printf("%T,%[1]v\n", a)

	//
	const b = 97
	// const c = 1.3
	const d byte = 'a'
	var e = '中'
	const s string = string(b)
	// const s1 string = string(c) //error
	const s2 = string(d)
	// const s5 = e //error 不能将一个变量赋值给一个常量,下面同理.
	// const s5 = string(e) //error
	var s3 = string(e)
	var s4 = string(d)
	fmt.Printf("%q\n", s)
	fmt.Printf("%q\n", s2)
	fmt.Printf("%q\n", s3)
	fmt.Printf("%q\n", s4)
}

/*
	go中的类型推断
		类型推断是指在某些场合下，程序员可以在代码中使用一些类型不确定值，编译器会自动推断出这些类型不确定值在特定情景下应该被视为某种特定类型的值。

		在Go代码中，如果某处需要一个特定类型的值并且一个类型不确定值可以表示为此特定类型的值，
			则此类型不确定值可以使用在此处。Go编译器将此类型不确定值视为此特定类型的类型确定值。
			这种情形常常出现在运算符运算、函数调用和赋值语句中。

		有些场景对某些类型不确定值并没有特定的类型要求。在这种情况下，Go编译器将这些类型不确定值视为它们各自的默认类型的类型确定值。
*/

/*

	（有名）常量声明（constant declaration）
		和无名字面常量一样，有名常量也必须都是布尔、数字或者字符串值。


		Go白皮书把每行含有一个等号=的语句称为一个常量描述（constant specification）。
		每个const关键字对应一个常量声明。一个常量声明中可以有若干个常量描述。

		常量声明中的等号=表示“绑定”而非“赋值”。 每个常量描述将一个或多个字面量绑定到各自对应的有名常量上。
			 或者说，每个有名常量其实代表着一个字面常量。

		在下面的例子中，有名常量π和Pi都绑定到（或者说代表着）字面常量3.1416,
		这两个有名常量可以在程序代码中被多次使用，
		从而有效避免了字面常量3.1416在代码中出现在多处。
		如果字面常量3.1416在代码中出现在多处，当我们以后欲将3.1416改为3.14的时候，
		所有出现在代码中的3.1416都得逐个修改。
		有了有名常量的帮助，我们只需修改对应常量描述中的3.1416即可。
		这是常量声明的主要作用。当然常量声明也可常常增加代码的可读性


		下面例子中声明的所有常量都是类型不确定的。
		它们各自的默认类型和它们各自代表的字面量的默认类型是一样的。
*/

//	常量是不可改变的（不可寻址的），所以常量不能做为目标值出现在纯赋值语句的左边，而只能出现在右边用做源值。

const π = 3.1416
const Pi = π

const (
	// 包级常量声明中的常量描述的顺序并不重要
	Yes = true
	No  = !Yes
	Num = ^1
)

/*
	类型确定的有名常量
		我们可以在声明一些常量的时候指定这些常量的确切类型。 这样声明的常量称为类型确定有名常量。

*/
const X float32 = 3.14

const (
	A, B int64   = -3, 5
	Y    float32 = 2.718
)

//也可以使用显式类型转换来声明类型确定常量
const Z = float32(3.14)

const (
	Q, W = int64(-3), int64(5)
	E    = float32(2.718)
)

/*
	欲将一个字面常量绑定到一个类型确定有名常量上，此字面常量必须能够表示为此常量的确定类型的值。 否则，编译将报错。
*/

func Testeg3() {
	// const a uint8 = 256             // error: 256溢出uint8
	// const b = uint8(255) + uint8(1) // error: 256溢出uint8
	// const c = int8(-128) / int8(-1) // error: 128溢出int8
	// const MaxUint_a = uint(^0)      // error: -1溢出uint
	// const MaxUint_b uint = ^0       // error: -1溢出uint
}

/*
	下面这个类型确定常量声明在64位的操作系统上是合法的，但在32位的操作系统上是非法的。
	因为一个uint值在32位操作系统上的尺寸是32位， (1 << 64) - 1将溢出uint
*/
const MaxUint uint = (1 << 64) - 1 //32位系统报错.

//如何声明一个代表着最大uint值的常量呢？
const MaxUint2 uint = ^uint(0)

//声明一个有名常量来表示最大的int值
const MaxInt = int(^uint(0) >> 1)

//使用类似的方法，我们可以声明一个常量来表示当前操作系统的位数，或者检查当前操作系统是32位的还是64位的
const NativeWordBits int = 32 << (^uint(0) >> 63) // 64 or 32
const Is64bitOS = ^uint(0)>>63 != 0
const Is32bitOS = ^uint(0)>>32 == 0

/*
	常量声明中的自动补全
	与
	在常量声明中使用iota
*/

const (
	X1 float32 = 3.14
	Y1         // 这里必须只有一个标识符
	Z1         // 这里必须只有一个标识符

	A1, B1 = "Go", "language"
	C1, _
	// 上一行中的空标识符是必需的（如果
	// 上一行是一个不完整的常量描述）。
)

func Test5() {
	const (
		k = 3 // 在此处，iota == 0

		m float32 = iota + .5 // m float32 = 1 + .5
		n                     // n float32 = 2 + .5

		p = 9        // 在此处，iota == 3
		q = iota * 2 // q = 4 * 2
		_            // _ = 5 * 2
		r            // r = 6 * 2

		s, t = iota, iota // s, t = 7, 7
		u, v              // u, v = 8, 8
		_, w              // _, w = 9, 9

	)
	const x = iota // x = 0 （iota == 0）
	const (
		y = iota // y = 0 （iota == 0）
		z        // z = 1
	)

}

/*
	一个类型不确定常量所表示的值可以溢出其默认类型
		下例中的三个类型不确定常量均溢出了它们各自的默认类型，但是此程序编译和运行都没问题。
*/
// 三个类型不确定常量。
const n = 1 << 64          // 默认类型为int
const r = 'a' + 0x7FFFFFFF // 默认类型为rune
const x12 = 2e+308         // 默认类型为float64
func Testeg11() {
	_, _, _ = n>>2, r>>2, x12/2
}

// 但是下面这些编译不通过，因为三个声明的常量为类型确定常量。
// 三个类型确定常量。
// const n int = 1 << 64           // error: 溢出int
// const r rune = 'a' + 0x7FFFFFFF // error: 溢出rune
// const x float64 = 2e+308        // error: 溢出float64

/*
	每个常量标识符将在编译的时候被其绑定的字面量所替代.
	常量声明可以看作是增强型的C语言中的#define宏。 在编译阶段，所有的标识符将被它们各自绑定的字面量所替代。
	如果一个运算中的所有运算数都为常量，则此运算的结果也为常量。或者说，此运算将在编译阶段就被估值。

const X = 3
const Y = X + X
var a = X

func main() {
	b := Y
	println(a, b, X, Y)
}

上面这段程序代码将在编译阶段被重写为下面这样：
var a = 3

func main() {
	b := 6
	println(a, b, 3, 6)
}

*/
