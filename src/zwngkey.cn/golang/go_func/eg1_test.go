/*
 * @Author: zwngkey
 * @Date: 2022-05-05 20:40:09
 * @LastEditors: zwngkey 18390924907@163.com
 * @LastEditTime: 2022-05-06 03:12:00
 * @Description: go 函数
 */
package gofunc

import (
	"fmt"
	"testing"
)

/*
	事实上，在Go中，函数是一种一等公民类型。换句话说，我们可以把函数当作值来使用。
		尽管Go是一门静态语言，但是Go函数的灵活性宛如甚至超越了很多动态语言。

	Go中有一些内置函数，这些函数展示在builtin和unsafe标准包中。 内置函数和自定义函数有很多差别。
*/
/*
	函数签名（function signature）和函数类型
		一个函数类型的字面表示形式由一个func关键字和一个函数签名字面表示形式组成。
		 一个函数签名由一个输入参数类型列表和一个输出结果类型列表组成。
		 	参数名称和结果名称可以出现函数签名的字面表示形式中，但是它们并不重要。

	func关键字可以出现在函数签名的字面形式中，也可以不出现。 鉴于此，我们常常混淆使用函数类型（见下）和函数签名这两个概念。

	下面是一个函数类型的字面形式:
		func (a int, b, c string) (x, y int, z bool)

	 参数名和结果名可以是空标识符_。上面的字面形式等价于：
		func (_ int, _, _ string) (_, _ int, _ bool)

	 函数参数列表中的参数名或者结果列表中的结果名可以同时省略（即匿名）。上面的字面形式等价于：
		func (int, string, string) (int, int, bool) // 标准函数字面形式

	所有上面列出的函数类型字面形式表示同一个（非定义）函数类型。

*/
/*
	变长参数和变长参数函数类型
		一个函数的最后一个参数可以是一个变长参数。一个函数可以最多有一个变长参数。一个变长参数的类型总为一个切片类型。
		 变长参数在声明的时候必须在它的（切片）类型的元素类型前面前置三个点...，以示这是一个变长参数。 两个变长函数类型的例子：
			func (values ...int64) (sum int64)
			func (sep string, tokens ...string) string

		一个变长函数类型和一个非变长函数类型绝对不可能是同一个类型。


	所有的函数类型都属于不可比较类型
		一个函数值可以和类型不确定的nil比较。
		因为函数类型属于不可比较类型，所以函数类型不可用做映射类型的键值类型。
*/
/*
	函数原型（function prototype）
		一个函数原型由一个函数名称和一个函数类型（或者说一个函数签名）组成。
		 它的字面形式由一个func关键字、一个函数名和一个函数签名字面形式组成。

		一个函数原型的例子：
			func Double(n int) (result int)
		换句话说，一个函数原型可以看作是一个不带函数体的函数声明； 或者说一个函数声明由一个函数原型和一个函数体组成。


	变长函数声明和变长函数调用
		变长函数声明
			变长函数声明和普通函数声明类似，只不过最后一个参数必须为变长参数。 一个变长参数在函数体内将被视为一个切片。

			如果一个变长参数的类型部分为...T，则此变长参数的类型实际为[]T。
*/

// Sum返回所有输入实参的和。
func Sum(values ...int64) (sum int64) {
	// values的类型为[]int64。
	sum = 0
	for _, v := range values {
		sum += v
	}
	return
}
func Concat(sep string, tokens ...string) (r string) {
	for i, t := range tokens {
		if i != 0 {
			r += sep
		}
		r += t
	}
	return
}

/*
	变长参数函数调用
		在变长参数函数调用中，可以使用两种风格的方式将实参传递给类型为[]T的变长形参：
			1.传递一个切片做为实参。此切片必须可以被赋值给类型为[]T的值（或者说此切片可以被隐式转换为类型[]T）。
				此实参切片后必须跟随三个点...。
			2.传递零个或者多个可以被隐式转换为T的实参（或者说这些实参可以赋值给类型为T的值）。
				这些实参将被添加入一个匿名的在运行时刻创建的类型为[]T的切片中，然后此切片将被传递给此函数调用。

			注意: 这两种方式不能混用.
*/
func TestEg11(t *testing.T) {
	a0 := Sum()
	a1 := Sum(2)
	a3 := Sum(2, 3, 5)
	// 上面三行和下面三行是等价的。
	s1 := []int64{}
	s2 := []int64{2}
	s3 := []int64{2, 3, 5}

	b0 := Sum(s1...) // <=> Sum(nil...)
	b1 := Sum(s2...)
	b3 := Sum(s3...)
	fmt.Println(a0, a1, a3) // 0 2 10
	fmt.Println(b0, b1, b3) // 0 2 10
}

func TestEg12(t *testing.T) {
	tokens := []string{"Go", "C", "Rust"}
	langsA := Concat(",", tokens...)         // 风格1
	langsB := Concat(",", "Go", "C", "Rust") // 风格2
	fmt.Println(langsA == langsB)            // true
}

/*
	更多关于函数声明和函数调用的事实
		同一个包中可以同名的函数
			一般来说，同一个包中声明的函数的名称不能重复，但有两个例外：
				1.同一个包内可以声明若干个原型为func ()的名称为init的函数。
				2.多个函数的名称可以被声明为空标识符_。这样声明的函数不可被调用。

		某些函数调用是在编译时刻被估值的
			大多数函数调用都是在运行时刻被估值的。 但unsafe标准库包中的函数的调用都是在编译时刻估值的。
				另外，某些其它内置函数（比如len和cap等）的调用在所传实参满足一定的条件的时候也将在编译时刻估值。


		所有的函数调用的传参均属于值复制
			和赋值一样，传参也属于值（浅）复制。当一个值被复制时，只有它的直接部分被复制了。

		不含函数体的函数声明
			我们可以使用Go汇编（Go assembly）来实现一个Go函数。 Go汇编代码放在后缀为.a的文件中。
				一个使用Go汇编实现的函数依旧必须在一个*.go文件中声明，但是它的声明必须不能含有函数体。
				 换句话说，一个使用Go汇编实现的函数的声明中只含有它的原型。

			或者使用//go:linkname funcName package.funcName

		某些有返回值的函数可以不必返回
			如果一个函数有返回值，则它的函数体内的最后一条语句必须为一条终止语句。 Go中有多种终止语句，return语句只是其中一种。
			  所以一个有返回值的函数的体内不一定需要一个return语句。 比如下面两个函数（它们均可编译通过）：
				func fa() int {
					a:
					goto a
				}

				func fb() bool {
					for{}
				}


		自定义函数的调用返回结果可以被舍弃，但是某些内置函数的调用返回结果不可被舍弃
			自定义函数的调用结果都是可以被舍弃掉的。 但是大多数内置函数（除了recover和copy）的调用结果都是不可被舍弃的。
				 调用结果不可被舍弃的函数是不可以被用做延迟调用函数和协程起始函数的，比如append函数。
*/
func TestEg13(t *testing.T) {
	var s = []int{1, 2, 3}
	var comp = 1 + 1i
	s = append(s, 1, 2, 3)
	_ = cap(s)
	_ = len(s)
	_ = make([]int, 1)
	_ = imag(comp)
	_ = real(comp)
	_ = complex(1, 2)
}

/*
	有返回值的函数的调用是一种表达式
		一个有且只有一个返回值的函数的每个调用总可以被当成一个单值表达式使用。
			比如，它可以被内嵌在其它函数调用中当作实参使用，或者可以被当作其它表达式中的操作数使用。

		如果一个有多个返回结果的函数的一个调用的返回结果没有被舍弃，则此调用可以当作一个多值表达式使用在两种场合：
			1.此调用可以在一个赋值语句中当作源值来使用，但是它不能和其它源值掺和到一块。
			2.此调用可以内嵌在另一个函数调用中当作实参来使用，但是它不能和其它实参掺和到一块。

		注意，在目前的标准编译器的实现中，有几个内置函数破坏了上述规则的普遍性。
*/
func HalfAndNegative(n int) (int, int) {
	return n / 2, -n
}

func AddSub(a, b int) (int, int) {
	return a + b, a - b
}

func Dummy(values ...int) {}

func TestEg14(t *testing.T) {
	_, _ = AddSub(HalfAndNegative(6))
	_, _ = AddSub(AddSub(AddSub(7, 5)))
	_, _ = AddSub(AddSub(HalfAndNegative(6)))
	Dummy(HalfAndNegative(6))
	Dummy(AddSub(1, 2))
	_, _ = AddSub(7, 5)

	// 下面这几行编译不通过。
	/*
		_, _, _ = 6, AddSub(7, 5)
		Dummy(AddSub(7, 5), 9)
		Dummy(AddSub(7, 5), HalfAndNegative(6))
	*/
}

/*

	嵌套函数调用
		基本规则：
			如果一个函数（包括方法）调用的返回值个数不为零，并且它的返回值列表可以用做另一个函数调用的实参列表，
				则此前者调用可以被内嵌在后者调用之中，但是此前者调用不可和其它的实参混合出现在后者调用的实参列表中。

		语法糖：
			如果一个函数调用刚好返回一个结果，则此函数调用总是可以被当作一个单值实参用在其它函数调用中，
				并且此函数调用可以和其它实参混合出现在其它函数调用的实参列表中。
*/
func f0() float64                  { return 1 }
func f1() (float64, float64)       { return 1, 2 }
func f2(float64, float64)          {}
func f3(float64, float64, float64) {}
func f4() (x, y []int)             { return }
func f5() (x map[int]int, y int)   { return }

type I interface{ m() (float64, float64) }
type T struct{}

func (T) m() (float64, float64) { return 1, 2 }

func TestEg15(t *testing.T) {
	f2(f0(), 12)
	f2(f1())
	fmt.Println(f1())
	_ = complex(f1())
	_ = complex(T{}.m())
	f2(I(T{}).m())

	f3(1, 2, 3)
	//error
	// f3(123,f1())
	// f3(f1(), 123)
	// f3(f0(), f1())

	// 此行从Go官方工具链1.15开始才能够编译通过。
	println(f1())

	// 下面这三行从Go官方工具链1.13开始才能够编译通过。
	copy(f4())
	delete(f5())
	_ = complex(I(T{}).m())

}

/*
	函数值
		函数类型的值称为函数值。 在字面上，函数类型的零值也使用预定义的nil来表示。

		当我们声明了一个函数的时候，我们实际上同时声明了一个不可修改的函数值。
			此函数值用此函数的名称来标识。此函数值的类型的字面表示形式为此函数的原型刨去函数名部分。

		注意：内置函数和init函数不可被用做函数值。
			var f = len //error
			var f = init //error

		任何函数值都可以被当作普通声明函数来调用。
			调用一个nil函数来开启一个协程将产生一个致命的不可恢复的错误，此错误将使整个程序崩溃。
			 在其它情况下调用一个nil函数将产生一个可恢复的恐慌。

		从值部一文，我们得知，当一个函数值被赋给另一个函数值后，这两个函数值将共享底层部分（内部的函数结构）。
			换句话说，这两个函数值表示的函数可以看作是同一个函数。调用它们的效果是相同的。
*/
func TestEg16(t *testing.T) {
	defer func() {
		recover()
		fmt.Println("over")
	}()
	var f = Sum
	f = nil
	fmt.Printf("f(1, 2): %v\n", f(1, 2))
}

func Double(n int) int {
	return n + n
}

func Apply(n int, f func(int) int) int {
	return f(n)
}
func TestEg17(t *testing.T) {
	fmt.Printf("%T\n", Double) // func(int) int
	// Double = nil // error: Double是不可修改的

	var f func(int) int = Double // 默认值为nil
	g := Apply
	fmt.Printf("%T\n", g) // func(int, func(int) int) int

	fmt.Println(f(9))         // 18
	fmt.Println(g(6, Double)) // 12
	fmt.Println(Apply(6, f))  // 12

}

/*
	在实践中，我们常常将一个匿名函数赋值给一个函数类型的变量，从而可以在以后多次调用此匿名函数。

	Go中所有的函数都可以看作是闭包，这是Go函数如此灵活及使用体验如此统一的原因。
*/
func TestEg18(t *testing.T) {
	//判断n 是否为 x 的倍数.
	isMultipleOfX := func(x int) func(int) bool {
		return func(n int) bool {
			return n%x == 0
		}
	}
	//isMultipleOf2函数判断n是否为2的倍数
	isMultipleOf2 := isMultipleOfX(2)
	isMultipleOf2(2)

	//isMultipleOf3函数判断n是否为3的倍数
	isMultipleOf3 := isMultipleOfX(3)
	isMultipleOf3(3)

	isMultipleOf5 := isMultipleOfX(5)

	//isMultipleOf15函数判断n是否为15的倍数
	isMultipleOf15 := func(x int) bool {
		return isMultipleOf3(x) && isMultipleOf5(x)
	}

	isMultipleOf15(75)
}
