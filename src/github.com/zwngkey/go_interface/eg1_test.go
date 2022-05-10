/*
 * @Author: zwngkey
 * @Date: 2022-05-07 03:32:33
 * @LastEditors: zwngkey 18390924907@163.com
 * @LastEditTime: 2022-05-09 20:19:17
 * @Description: go interface
 */
package gointerface

import (
	"fmt"
	"reflect"
	"testing"
)

/*
	接口
		接口类型是Go中的一种很特别的类型。接口类型在Go中扮演着重要的角色。
			首先，在Go中，接口值可以用来包裹 非接口值；
			  然后，通过值包裹，反射和多态得以实现。

	什么是接口类型？
		一个接口类型指定了一个方法原型的集合。 换句话说，一个接口类型定义了一个方法集。
			事实上，我们可以把一个接口类型看作是一个方法集。 接口类型中指定的任何方法原型中的方法名称都不能为空标识符_。

		我们也常说一个接口类型规定了一个（用此接口类型指定的方法集表示的）行为规范。
*/

// ReadWriter是一个定义的接口类型。
type ReadWriter interface {
	Read(buf []byte) (n int, err error)
	Write(buf []byte) (n int, err error)
}

// Runnable是一个非定义接口类型的别名。
type Runnable = interface {
	Run()
}

/*
	特别地，一个没有指定任何方法原型的接口类型称为一个空接口类型。 下面是两个空接口类型：
		// 一个非定义空接口类型。
		interface{}

		// 类型I是一个定义的空接口类型。
		type I interface{}


	类型的方法集
		每个类型都有一个方法集。
			1.对于一个非接口类型，它的方法集由所有为它声明的（包括显式和隐式的）方法的原型组成。
			2.对于一个接口类型，它的方法集就是它所指定的方法集。

		为了解释方便，一个类型的方法集常常也可称为它的任何值的方法集。

		如果两个非定义接口类型指定的方法集是等价的，则这两个接口类型为同一个类型。
			但是请注意：不同代码包中的同名非导出方法名将总被认为是不同名的。


	什么是实现（implementation）？
		如果一个任意类型T的方法集为一个接口类型的方法集的超集，则我们说类型T实现了此接口类型。
			T可以是一个非接口类型，也可以是一个接口类型。

		实现关系在Go中是隐式的。两个类型之间的实现关系不需要在代码中显式地表示出来。Go中没有类似于implements的关键字。
			 Go编译器将自动在需要的时候检查两个类型之间的实现关系。

		一个接口类型总是实现了它自己。两个指定了一个相同的方法集的接口类型相互实现了对方。
*/
/*
	在下面的例子中，类型*Book、MyInt和*MyInt都拥有一个原型为About() string的方法，
		所以它们都实现了接口类型interface {About() string}。
*/
type Book struct {
	name string
	// 更多其它字段……
}

func (book *Book) About() string {
	return "Book: " + book.name
}

type MyInt int

func (MyInt) About() string {
	return "我是一个自定义整数类型"
}

func TestEg1(t *testing.T) {
	var a interface{ About() string } = MyInt(1)
	a.About()
}

/*
	注意：因为任何方法集都是一个空方法集的超集，所以任何类型都实现了任何空接口类型。 这是Go中的一个重要事实。

	隐式实现关系的设计使得一个声明在另一个代码包（包括标准库包）中的类型可以被动地实现一个用户代码包中的接口类型。
		比如，如果我们声明一个像下面这样的接口类型，则database/sql标准库包中声明的DB和Tx类型都实现了这个接口类型，
			因为它们都拥有此接口类型指定的三个方法。
		type DatabaseStorer interface {
			Exec(query string, args ...interface{}) (sql.Result, error)
			Prepare(query string) (*sql.Stmt, error)
			Query(query string, args ...interface{}) (*sql.Rows, error)
		}
*/
/*
	值包裹
		每个接口值都可以看作是一个用来包裹一个非接口值的盒子。
			欲将一个非接口值包裹在一个接口值中，此非接口值的类型必须实现了此接口值的类型。

		在Go中，如果类型T实现了一个接口类型I，则类型T的值都可以隐式转换到类型I。
			换句话说，类型T的值可以赋给类型I的可修改值。 当一个T值被转换到类型I（或者赋给一个I值）的时候，
				1.如果类型T是一个非接口类型，则此T值的一个复制将被包裹在结果（或者目标）I值中。 此操作的时间复杂度为O(n)，其中n为T值的尺寸。
				2.如果类型T也为一个接口类型，则此T值中当前包裹的（非接口）值将被复制一份到结果（或者目标）I值中。
				   官方标准编译器为此操作做了优化，使得此操作的时间复杂度为O(1)，而不是O(n)。

		包裹在一个接口值中的非接口值的类型信息也和此非接口值一起被包裹在此接口值中

		当一个非接口值被包裹在一个接口值中，此非接口值称为此接口值的动态值，此非接口值的类型称为此接口值的动态类型。

		接口值的动态值的直接部分是不可修改的，除非它的动态值被整体替换为另一个动态值。

		接口类型的零值也用预声明的nil标识符来表示。 一个nil接口值中什么也没包裹。将一个接口值修改为nil将清空包裹在此接口值中的非接口值。

		（注意，在Go中，很多其它非接口类型的零值也使用nil标识符来表示。 非接口类型的nil零值也可以被包裹在接口值中。
			一个包裹了一个nil非接口值的接口值不是一个nil接口值，因为它并非什么都没包裹。）

		因为任何类型都实现了空接口类型，所以任何非接口值都可以被包裹在任何一个空接口类型的接口值中。
		 （以后，一个空接口类型的接口值将被称为一个空接口值。注意空接口值和nil接口值是两个不同的概念。）
		   因为这个原因，空接口值可以被认为是很多其它语言中的any类型。

		当一个类型不确定值（除了类型不确定的nil）被转换为一个空接口类型（或者赋给一个空接口值），此类型不确定值将首先转换为它的默认类型。
			（或者说，此类型不确定值将被推断为一个它的默认类型的类型确定值。）
*/
type Aboutable interface {
	About() string
}

// 一些目标值为接口类型的赋值
func TestEg2(t *testing.T) {
	// 一个*Book类型的值被包裹在了一个Aboutable类型的值中。
	var a Aboutable = &Book{"Go语言101"}
	fmt.Println(a) // &{Go语言101}

	// i是一个空接口值。类型*Book实现了任何空接口类型。
	var i interface{} = &Book{"Rust 101"}
	fmt.Println(i) // &{Rust 101}

	// Aboutable实现了空接口类型interface{}。
	i = a
	fmt.Println(i) // &{Go语言101}
}

// 一个空接口类型的值包裹着各种非接口值的例子
func TestEg3(t *testing.T) {
	var i interface{}
	i = []int{1, 2, 3}
	vKind := reflect.TypeOf(i).Kind()
	if vKind == reflect.Slice {
		fmt.Println(" vkind is a slice type")
	}
	fmt.Println(i) // [1 2 3]

	i = map[string]int{"Go": 2012}
	vKind = reflect.TypeOf(i).Kind()
	if vKind == reflect.Map {
		fmt.Println("vkind is a map type")
	}
	fmt.Println(i) // map[Go:2012]

	i = true
	fmt.Println(i) // true
	i = 1
	fmt.Println(i) // 1
	i = "abc"
	fmt.Println(i) // abc

	// 将接口值i中包裹的值清除掉。
	i = nil
	fmt.Println(i) //

}

/*
	在编译时刻，Go编译器将构建一个全局表用来存储代码中要用到的各个类型的信息。
		对于一个类型来说，这些信息包括：此类型的种类（kind）、此类型的所有方法和字段信息、此类型的尺寸，等等。
			这个全局表将在程序启动的时候被加载到内存中。

	在运行时刻，当一个非接口值被包裹到一个接口值，Go运行时（至少对于官方标准运行时来说）将分析和构建这两个值的类型的实现关系信息，
	  并将此实现关系信息存入到此接口值内。 对每一对这样的类型，它们的实现关系信息将仅被最多构建一次。
	  	并且为了程序效率考虑，此实现关系信息将被缓存在内存中的一个全局映射中，以备后用。
		  	所以此全局映射中的条目数永不减少。
			 事实上，一个非零接口值在内部只是使用一个指针字段来引用着此全局映射中的一个实现关系信息条目。

	对于一个非接口类型和接口类型对，它们的实现关系信息包括两部分的内容：
		1.动态类型（即此非接口类型）的信息。
		2.一个方法表（切片类型），其中存储了所有此接口类型指定的并且为此非接口类型（动态类型）声明的方法。

	这两部分的内容对于实现Go中的两个特性起着至关重要的作用。
		动态类型信息是实现反射的关键。
		方法表是实现多态的关键。

*/
/*
	多态（polymorphism）
		多态是接口的一个关键功能和Go语言的一个重要特性。

		当非接口类型T的一个值t被包裹在接口类型I的一个接口值i中，通过i调用接口类型I指定的一个方法时，
			事实上为非接口类型T声明的对应方法将通过非接口值t被调用。
			 换句话说，调用一个接口值的方法实际上将调用此接口值的动态值的对应方法。
			 	 比如，当方法i.m被调用时，其实被调用的是方法t.m。
				  一个接口值可以通过包裹不同动态类型的动态值来表现出各种不同的行为，这称为多态。

		当方法i.m被调用时，i存储的实现关系信息的方法表中的方法t.m将被找到并被调用。
			此方法表是一个切片，所以此寻找过程只不过是一个切片元素访问操作，不会消耗很多时间。

		注意，在nil接口值上调用方法将产生一个恐慌，因为没有具体的方法可被调用。

*/

type Filter interface {
	About() string
	Process([]int) []int
}

// UniqueFilter用来删除重复的数字。
type UniqueFilter struct{}

func (UniqueFilter) About() string {
	return "删除重复的数字"
}
func (UniqueFilter) Process(inputs []int) []int {
	outs := make([]int, 0, len(inputs))
	pusheds := make(map[int]bool)
	for _, n := range inputs {
		if !pusheds[n] {
			pusheds[n] = true
			outs = append(outs, n)
		}
	}
	return outs
}

// MultipleFilter用来只保留某个整数的倍数数字。
type MultipleFilter int

func (mf MultipleFilter) About() string {
	return fmt.Sprintf("保留%v的倍数", mf)
}
func (mf MultipleFilter) Process(inputs []int) []int {
	var outs = make([]int, 0, len(inputs))
	for _, n := range inputs {
		if n%int(mf) == 0 {
			outs = append(outs, n)
		}
	}
	return outs
}

// 在多态特性的帮助下，只需要一个filteAndPrint函数。
func filteAndPrint(fltr Filter, unfiltered []int) []int {
	// 在fltr参数上调用方法其实是调用fltr的动态值的方法。
	filtered := fltr.Process(unfiltered)
	fmt.Println(fltr.About()+":\n\t", filtered)
	return filtered
}

func TestEg4(t *testing.T) {

	numbers := []int{12, 7, 21, 12, 12, 26, 25, 21, 30}
	fmt.Println("过滤之前：\n\t", numbers)

	// 三个非接口值被包裹在一个Filter切片的三个接口元素中。
	filters := []Filter{
		UniqueFilter{},
		MultipleFilter(2),
		MultipleFilter(3),
	}

	// 每个切片元素将被赋值给类型为Filter的循环变量fltr。
	// 每个元素中的动态值也将被同时复制并被包裹在循环变量fltr中。
	for _, fltr := range filters {
		numbers = filteAndPrint(fltr, numbers)
	}
}

/*
	上例中,多态使得我们不必为每个过滤器类型写一个单独的filteAndPrint函数。

	除了上述这个好处，多态也使得一个代码包的开发者可以在此代码包中声明一个接口类型并声明一个拥有此接口类型参数的函数（或者方法），
	   从而此代码包的一个用户可以在用户包中声明一个实现了此接口类型的用户类型，
	    并且将此用户类型的值做为实参传递给此代码包中声明的函数（或者方法）的调用。
		 此代码包的开发者并不用关心一个用户类型具体是如何声明的，只要此用户类型满足此代码包中声明的接口类型规定的行为即可。

	事实上，多态对于一个语言来说并非一个不可或缺的特性。我们可以通过其它途径来实现多态的作用。
	  但是，多态可以使得我们的代码更加简洁和优雅。
*/

/*
	反射（reflection）
		一个接口值中存储的动态类型信息可以被用来检视此接口值的动态值和操纵此动态值所引用的值。 这称为反射。

		这里只说明Go中的内置反射机制。在Go中，内置反射机制包括类型断言（type assertion）和type-switch流程控制代码块。

	类型断言
		Go中有四种接口相关的类型转换情形：
			1.将一个非接口值转换为一个接口类型。在这样的转换中，此非接口值的类型必须实现了此接口类型。
			2.将一个接口值转换为另一个接口类型（前者接口值的类型实现了后者目标接口类型）。
			3.将一个接口值转换为一个非接口类型（此非接口类型必须实现了此接口值的接口类型）。
			4.将一个接口值转换为另一个接口类型（前者接口值的类型未实现后者目标接口类型，但是前者的动态类型有可能实现了目标接口类型）。

		前两种情形都要求源值的类型必须实现了目标接口类型。 它们都是通过普通类型转换（无论是隐式的还是显式的）来达成的。
			这两种情形的合法性是在编译时刻验证的。

		后两种情形的合法性是在运行时刻通过类型断言来验证的。 事实上，类型断言也适用于上面列出的第二种情形。

		一个类型断言表达式的语法为i.(T)，其中i为一个接口值，T为一个类型名或者类型字面表示。 类型T可以为:
			1.任意一个非接口类型。
			2.或者一个任意接口类型。

		在一个类型断言表达式i.(T)中，i称为断言值，T称为断言类型。 一个断言可能成功或者失败。
			1.对于T是一个非接口类型的情况，如果断言值i的动态类型存在并且此动态类型和T为同一类型，则此断言将成功；否则，此断言失败。
				当此断言成功时，此类型断言表达式的估值结果为断言值i的动态值的一个复制。 我们可以把此种情况看作是一次拆封动态值的尝试。
			2.对于T是一个接口类型的情况，当断言值i的动态类型存在并且此动态类型实现了接口类型T，则此断言将成功；否则，此断言失败。
			 	当此断言成功时，此类型断言表达式的估值结果为一个包裹了断言值i的动态值的一个复制的T值。

		一个失败的类型断言的估值结果为断言类型的零值。

		按照上述规则，如果一个类型断言中的断言值是一个零值nil接口值，则此断言必定失败。

		对于大多数场合，一个类型断言被用做一个单值表达式。 但是，当一个类型断言被用做一个赋值语句中的唯一源值时，
			此断言可以返回一个可选的第二个结果并被视作为一个多值表达式。 此可选的第二个结果为一个类型不确定的布尔值，用来表示此断言是否成功了

		注意：如果一个断言失败并且它的可选的第二个结果未呈现，则此断言将造成一个恐慌。

		事实上，对于一个动态类型为T的接口值i，方法调用i.m(...)等价于i.(T).m(...)。

*/
// 如何使用类型断言的例子（断言类型为非接口类型）：
func TestEg5(t *testing.T) {
	// 编译器将把123的类型推断为内置类型int。
	var x any = 123

	// 情形一：
	n, ok := x.(int)
	fmt.Println(n, ok) // 123 true

	n = x.(int)
	fmt.Println(n) // 123

	// 情形二：
	a, ok := x.(float64)
	fmt.Println(a, ok) // 0 false

	// 情形三：
	a = x.(float64) // 将产生一个恐慌
	_ = a
}

type Writer interface {
	Write(buf []byte) (int, error)
}

type DummyWriter struct{}

func (DummyWriter) Write(buf []byte) (int, error) {
	return len(buf), nil
}

// 如何使用类型断言的例子（断言类型为接口类型）：
func TestEg6(t *testing.T) {
	var x interface{} = DummyWriter{}

	// y的动态类型为内置类型string。
	var y interface{} = "abc"
	var w Writer
	var ok bool

	// DummyWriter既实现了Writer，也实现了interface{}。
	w, ok = x.(Writer)
	fmt.Println(w, ok) // {} true
	x, ok = w.(interface{})
	fmt.Println(x, ok) // {} true

	// y的动态类型为string。string类型并没有实现Writer。
	w, ok = y.(Writer)
	fmt.Println(w, ok) //  false
	// w = y.(Writer)     // 将产生一个恐慌
}

/*
	type-switch流程控制代码块
		type-switch流程控制的语法或许是Go中最古怪的语法。 它可以被看作是类型断言的增强版。它和switch-case流程控制代码块有些相似。
		  一个type-switch流程控制代码块的语法如下所示：
			switch aSimpleStatement; v := x.(type) {
			case TypeA:
				...
			case TypeB, TypeC:
				...
			case nil:
				...
			default:
				...
			}
		其中aSimpleStatement;部分是可选的。 aSimpleStatement必须是一个简单语句。
			x必须为一个估值结果为接口值的表达式，它称为此代码块中的断言值。 v称为此代码块中的断言结果，它必须出现在一个短变量声明形式中。

		在一个type-switch代码块中，每个case关键字后可以跟随一个nil标识符和若干个类型名或类型字面表示形式。
			在同一个type-switch代码块中，这些跟随在所有case关键字后的条目不可重复。

		如果跟随在某个case关键字后的条目为一个非接口类型（用一个类型名或类型字面表示），则此非接口类型必须实现了断言值x的（接口）类型。


		如果我们不关心一个type-switch代码块中的断言结果，则此type-switch代码块可简写为switch x.(type) {...}。

		type-switch代码块和switch-case代码块有很多共同点：
			1.在一个type-switch代码块中，最多只能有一个default分支存在。
			2.在一个type-switch代码块中，如果default分支存在，它可以为最后一个分支、第一个分支或者中间的某个分支。
			3.一个type-switch代码块可以不包含任何分支，它可以被视为一个空操作。

		但是，和switch-case代码块不一样，fallthrough语句不能使用在type-switch代码块中。
*/
// 下面是一个使用了type-switch代码块的例子：

func TestEg7(t *testing.T) {
	values := []any{
		456, "abc", true, 0.33, int32(789),
		[]int{1, 2, 3}, map[int]bool{}, nil,
	}
	for _, x := range values {
		// 这里，虽然变量v只被声明了一次，但是它在
		// 不同分支中可以表现为多个类型的变量值。
		switch v := x.(type) {
		case []int: // 一个类型字面表示
			// 在此分支中，v的类型为[]int。
			fmt.Println("int slice:", v)
		case string: // 一个类型名
			// 在此分支中，v的类型为string。
			fmt.Println("string:", v)
		case int, float64, int32: // 多个类型名
			// 在此分支中，v的类型为x的类型interface{}。
			fmt.Println("number:", v)
		case nil:
			// 在此分支中，v的类型为x的类型interface{}。
			fmt.Println(v)
		default:
			// 在此分支中，v的类型为x的类型interface{}。
			fmt.Println("others:", v)
		}
		// 注意：在最后的三个分支中，v均为接口值x的一个复制。
	}

}

// 上面这个例子程序在逻辑上等价于下面这个：
func TestEg8(t *testing.T) {
	values := []interface{}{
		456, "abc", true, 0.33, int32(789),
		[]int{1, 2, 3}, map[int]bool{}, nil,
	}
	for _, x := range values {
		if v, ok := x.([]int); ok {
			fmt.Println("int slice:", v)
		} else if v, ok := x.(string); ok {
			fmt.Println("string:", v)
		} else if x == nil {
			v := x
			fmt.Println(v)
		} else {
			_, isInt := x.(int)
			_, isFloat64 := x.(float64)
			_, isInt32 := x.(int32)
			if isInt || isFloat64 || isInt32 {
				v := x
				fmt.Println("number:", v)
			} else {
				v := x
				fmt.Println("others:", v)
			}
		}
	}

}

/*
	更多接口相关的内容
		接口类型内嵌
			一个接口类型可以内嵌另一个接口类型的名称。 此内嵌的效果相当于将此被内嵌的接口类型所指定的所有方法原型展开到内嵌它的接口类型的定义体内.

			一个接口类型不能内嵌（无论是直接还是间接）它自己。
*/
// 比如，在下面的例子中，接口类型Ic、Id和Ie的所指定的方法集完全一样。
type Ia interface {
	fa()
}

type Ib = interface {
	fb()
}

type Ic interface {
	fa()
	fb()
}

type Id = interface {
	Ia // 内嵌Ia
	Ib // 内嵌Ib
}

type Ie interface {
	Ia // 内嵌Ia
	fb()
}

// 从Go 1.14版本开始，下面这段代码中的几个接口类型声明均被认为是合法的，它们指定的方法集其实和Ic是一样的。
type Ix interface {
	Ia
	Ic
}

type Iy = interface {
	Ib
	Ic
}

type Iz interface {
	Ic
	fa()
}

/*
	接口值相关的比较
		接口值相关的比较有两种情形：
			1.比较一个非接口值和接口值；
			2.比较两个接口值。

		对于第一种情形，非接口值的类型必须实现了接口值的类型（假设此接口类型为I），
			所以此非接口值可以被隐式转化为（包裹到）一个I值中。 这意味着非接口值和接口值的比较可以转化为两个接口值的比较。

		比较两个接口值其实是比较这两个接口值的动态类型和和动态值。

		下面是（使用==比较运算符）比较两个接口值的步骤：
			1.如果其中一个接口值是一个nil接口值，则比较结果为另一个接口值是否也为一个nil接口值。
			2.如果这两个接口值的动态类型不一样，则比较结果为false。
			3.对于这两个接口值的动态类型一样的情形，
				1.如果它们的动态类型为一个不可比较类型，则将产生一个恐慌。
				2.否则，比较结果为它们的动态值的比较结果。

		简而言之，两个接口值的比较结果只有在下面两种任一情况下才为true：
			1.这两个接口值都为nil接口值。
			2.这两个接口值的动态类型相同、动态类型为可比较类型、并且动态值相等。
*/
func TestEg9(t *testing.T) {
	var a, b, c interface{} = "abc", 123, "a" + "b" + "c"
	fmt.Println(a == b) // 第二步的情形。输出"false"。
	fmt.Println(a == c) // 第三步的情形。输出"true"。

	var x *int = nil
	var y *bool = nil
	var ix, iy interface{} = x, y
	var i interface{} = nil
	fmt.Println(ix == iy) // 第二步的情形。输出"false"。
	fmt.Println(ix == i)  // 第一步的情形。输出"false"。
	fmt.Println(iy == i)  // 第一步的情形。输出"false"。

	var s []int = nil // []int为一个不可比较类型。
	i = s
	fmt.Println(i == nil) // 第一步的情形。输出"false"。
	fmt.Println(i == i)   // 第三步的情形。将产生一个恐慌。
}

/*
	指针动态值和非指针动态值
		标准编译器/运行时对接口值的动态值为指针类型的情况做了特别的优化。
			此优化使得接口值包裹指针动态值比包裹非指针动态值的效率更高。 对于小尺寸值，此优化的作用不大；
			 但是对于大尺寸值，包裹它的指针比包裹此值本身的效率高得多。 对于类型断言，此结论亦成立。

		所以尽量避免在接口值中包裹大尺寸值。对于大尺寸值，应该尽量包裹它的指针。

		一个[]T类型的值不能直接被转换为类型[]I，即使类型T实现了接口类型I.
*/
//比如，我们不能直接将一个[]string值转换为类型[]interface{}。
func TestEg10(t *testing.T) {
	words := []string{"hi", "you", "are", "a", "good", "boy"}

	// fmt.Println函数的原型为：
	// func Println(a ...any) (n int, err error)
	// 所以words...不能传递给此函数的调用。

	// fmt.Println(words...) // 编译不通过
	// 将[]string值转换为类型[]any。
	iw := make([]any, 0, len(words))
	for _, w := range words {
		iw = append(iw, w)
	}
	fmt.Println(iw...) // 编译没问题
}

/*
	一个接口类型每个指定的方法都对应着一个隐式声明的函数
		如果接口类型I指定了一个名为m的方法原型，则编译器将隐式声明一个与之对应的函数名为I.m的函数。
		 此函数比m的方法原型中的参数多一个。此多出来的参数为函数I.m的第一个参数，它的类型为I。
		  对于一个类型为I的值i，方法调用i.m(...)和函数调用I.m(i, ...)是等价的。
*/
type I interface {
	m(int) bool
}

type T string

func (t T) m(n int) bool {
	return len(t) > n
}
func TestEg11(t *testing.T) {
	var i I = T("string")
	fmt.Println(i.m(6))
	fmt.Println(I.m(i, 1))
	fmt.Println(interface{ m(int) bool }.m(i, 5))
}
