/*
 * @Author: zwngkey
 * @Date: 2022-05-09 20:26:55
 * @LastEditors: zwngkey 18390924907@163.com
 * @LastEditTime: 2022-05-10 00:58:37
 * @Description: go 类型内嵌
 */
package gostruct

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

/*
	类型内嵌的目的和各种和类型内嵌相关的细节。

	类型内嵌:
		一个结构体类型可以拥有若干字段。 每个字段由一个字段名和一个字段类型组成。
			事实上，有时，一个字段可以仅由一个字段类型组成。 这样的字段声明方式称为类型内嵌（type embedding）。

	类型内嵌语法:
		在下面这个例子中，有六个类型被内嵌在了一个结构体类型中。每个类型内嵌形成了一个内嵌字段（embedded field）。

		因为历史原因，内嵌字段有时也称为匿名字段。但是，事实上，每个内嵌字段有一个（隐式的）名字。
		 此字段的非限定（unqualified）类型名即为此字段的名称。
		  比如，下例中的六个内嵌字段的名称分别为string、error、int、P、M和Header。

*/
func F4() {
	type P = *bool
	type M = map[int]int

	var x struct {
		string // 一个定义的非指针类型
		error  // 一个定义的接口类型
		*int   // 一个非定义指针类型
		P      // 一个非定义指针类型的别名
		M      // 一个非定义类型的别名

		http.Header // 一个定义的映射类型
	}
	x.string = "Go"
	x.error = nil
	x.int = new(int)
	x.P = new(bool)
	x.M = make(M)
	x.Header = http.Header{}
}

/*
	哪些类型可以被内嵌？
		1.一个类型名T只有在它既不表示 一个定义的指针类型 也不表示 一个基类型为指针类型 或者 接口类型的指针类型 的情况下才可以被用作内嵌字段。
		2.一个指针类型*T只有在 T为一个类型名 并且 T既不表示一个指针类型 也不表示 一个接口类型 的时候才能被用作内嵌字段。
*/
/*
下面列出了一些可以被或不可以被内嵌的类型或别名：
	type Encoder interface {Encode([]byte) []byte}
	type Person struct {name string; age int}
	type Alias = struct {name string; age int}
	type AliasPtr = *struct {name string; age int}
	type IntPtr *int
	type AliasPP = *IntPtr

	// 这些类型或别名都可以被内嵌。
	Encoder
	Person
	*Person
	Alias
	*Alias
	AliasPtr
	int
	*int

	// 这些类型或别名都不能被内嵌。
	*Encoder         // 基类型为一个接口类型
	AliasPP          // 基类型为一个指针类型IntPtr
	*AliasPtr        // 基类型为一个指针类型
	IntPtr           // 定义的指针类型
	*IntPtr          // 基类型为一个指针类型
	*chan int        // 基类型为一个非定义类型
	struct {age int} // 非定义非指针类型
	map[string]int   // 非定义非指针类型
	[]int64          // 非定义非指针类型
	func()           // 非定义非指针类型
*/
/*
	一个结构体类型中不允许有两个同名字段，此规则对匿名字段同样适用。
		 根据上述内嵌字段的隐含名称规则，一个非定义指针类型不能和它的基类型同时内嵌在同一个结构体类型中。
		 	比如，int和*int类型不能同时内嵌在同一个结构体类型中。

	一个结构体类型不能内嵌（无论间接还是直接）它自己。

	一般说来，只有内嵌含有字段或者拥有方法的类型才有意义，尽管很多既没有字段也没有方法的类型也可以被内嵌。
*/
/*
	类型内嵌的意义是什么？
		类型内嵌的主要目的是为了将被内嵌类型的功能扩展到内嵌它的结构体类型中，从而我们不必再为此结构体类型重复实现被内嵌类型的功能。

		很多其它流行面向对象的编程语言都是用继承来实现上述目的。

		 这两种方式有一个很大的不同点：
			如果类型T继承了另外一个类型，则类型T获取了另外一个类型的能力。 同时，一个T类型的值也可以被当作另外一个类型的值来使用。

			如果一个类型T内嵌了另外一个类型，则另外一个类型变成了类型T的一部分。 类型T获取了另外一个类型的能力，但是T类型的任何值都不能被当作另外一个类型的值来使用。
*/
//下面是一个展示了如何通过类型内嵌来扩展类型功能的例子：
type Person struct {
	Name string
	Age  int
}

func (p Person) PrintName() {
	fmt.Println("Name:", p.Name)
}
func (p *Person) SetAge(age int) {
	p.Age = age
}

type Singer struct {
	Person // 通过内嵌Person类型来扩展之
	works  []string
}

func TestEg1(t *testing.T) {
	var gaga = Singer{Person: Person{"Gaga", 30}}
	gaga.PrintName() // Name: Gaga
	gaga.Name = "Lady Gaga"
	(&gaga).SetAge(31)
	(&gaga).PrintName()   // Name: Lady Gaga
	fmt.Println(gaga.Age) // 31
}

/*
	从上例中，当类型Singer内嵌了类型Person之后，看上去类型Singer获取了类型Person所有的字段和方法，
		并且类型*Singer获取了类型*Person所有的方法。

	注意，类型Singer的一个值不能被当作Person类型的值用。下面的代码编译不通过：
		var gaga = Singer{}
		var _ Person = gaga
*/
/*
	当一个结构体类型内嵌了另一个类型，此结构体类型是否获取了被内嵌类型的字段和方法？

*/
// 下面这个程序使用反射列出了上一节的例子中的Singer类型的字段和方法，以及*Singer类型的方法。
func TestEg2(te *testing.T) {
	singer := Singer{}
	t := reflect.TypeOf(singer)
	fmt.Println(t, "has", t.NumField(), "fields:")
	for i := 0; i < t.NumField(); i++ {
		fmt.Print(" field#", i, ": ", t.Field(i).Name, "\n")
	}
	/*
		gostruct.Singer has 2 fields:
			field#0: Person
			field#1: works
	*/

	fmt.Println(t, "has", t.NumMethod(), "methods:")
	for i := 0; i < t.NumMethod(); i++ {
		fmt.Print(" method#", i, ": ", t.Method(i).Name, "\n")
	}
	/*
		gostruct.Singer has 1 methods:
		 	method#0: PrintName
	*/

	pt := reflect.TypeOf(&singer) // the *Singer type
	fmt.Println(pt, "has", pt.NumMethod(), "methods:")
	for i := 0; i < pt.NumMethod(); i++ {
		fmt.Print(" method#", i, ": ", pt.Method(i).Name, "\n")
	}
	/*
		*gostruct.Singer has 2 methods:
			method#0: PrintName
			method#1: SetAge
	*/

	/*
		从此输出结果中，我们可以看出类型Singer确实拥有一个PrintName方法，以及类型*Singer确实拥有两个方法：PrintName和SetAge。
		  但是类型Singer并不拥有一个Name字段。那么为什么选择器表达式gaga.Name是合法的呢？
		    毕竟gaga是Singer类型的一个值。
	*/
}

/*
	选择器的缩写形式
		对于一个值x，x.y称为一个选择器，其中y可以是一个字段名或者方法名。
		  如果y是一个字段名，那么x必须为一个结构体值或者结构体指针值。
		    一个选择器是一个表达式，它表示着一个值。 如果选择器x.y表示一个字段，
			  此字段也可能拥有自己的字段（如果此字段的类型为另一个结构体类型）和方法，
				比如x.y.z，其中z可以是一个字段名，也可是一个方法名。

		在Go中，（不考虑选择器碰撞和遮挡），如果一个选择器中的中部某项对应着一个内嵌字段，则此项可被省略掉。
		  因此内嵌字段又被称为匿名字段。
*/
type A struct {
	x int
}

func (a A) MethodA() {}

type B struct {
	*A
}

type C struct {
	B
}

func TestEg3(t *testing.T) {
	var c = &C{B: B{A: &A{x: 5}}}
	// 这几行是等价的。
	_ = c.B.A.x
	_ = c.B.x
	_ = c.A.x // A是类型C的一个提升字段
	_ = c.x   // x也是一个提升字段

	// 这几行是等价的。
	c.B.A.MethodA()
	c.B.MethodA()
	c.A.MethodA()
	c.MethodA() // MethodA是类型C的一个提升方法
}

/*
	这就是为什么在上一节的例子中选择器表达式gaga.Name是合法的， 因为它只不过是gaga.Person.Name的一个缩写形式。

	类似的，选择器gaga.PrintName可以被看作是gaga.Person.PrintName的缩写形式。
		但是，我们也可以不把它看作是一个缩写。毕竟，类型Singer确实拥有一个PrintName方法， 尽管此方法是被隐式声明的

	同样的原因，选择器(&gaga).PrintName和(&gaga).SetAge可以看作（也可以不看作）是(&gaga.Person).PrintName和(&gaga.Person).SetAge的缩写。

	Name被称为类型Singer的一个提升字段（promoted field）。 PrintName被称为类型Singer的一个提升方法（promoted method）。

	注意：我们也可以使用选择器gaga.SetAge，但是只有在gaga是一个可寻址的类型为Singer的值的情况下。
	  它只不过是(&gaga).SetAge的一个语法糖。

	在上面的例子中，c.B.A.FieldX称为选择器表达式c.FieldX、c.B.FieldX和c.A.FieldX的完整形式。
		类似的，c.B.A.MethodA可以称为c.MethodA、c.B.MethodA和c.A.MethodA的完整形式。

	如果一个选择器的完整形式中的所有中部项均对应着一个内嵌字段，则中部项的数量称为此选择器的深度。
		比如，上面的例子中的选择器c.MethodA的深度为2，因为此选择器的完整形式为c.B.A.MethodA，并且B和A都对应着一个内嵌字段。

*/
/*
	选择器遮挡和碰撞
		一个值x（这里我们总认为它是可寻址的）可能同时拥有多个最后一项相同的选择器，并且这些选择器的中间项均对应着一个内嵌字段。
		  对于这种情形（假设最后一项为y）：
			1.只有深度最浅的一个完整形式的选择器（并且最浅者只有一个）可以被缩写为x.y。 换句话说，x.y表示深度最浅的一个选择器。
				其它完整形式的选择器被此最浅者所遮挡（压制）。
			2.如果有多个完整形式的选择器同时拥有最浅深度，则任何完整形式的选择器都不能被缩写为x.y。
				我们称这些同时拥有最浅深度的完整形式的选择器发生了碰撞。

	如果一个方法选择器被另一个方法选择器所遮挡，并且它们对应的方法原型是一致的，那么我们可以说第一个方法被第二个覆盖（overridden）了。

*/
type D struct {
	x string
}

func (D) y(int) bool {
	return false
}

type E struct {
	y bool
}

func (E) x(string) {}

type F struct {
	E
}

func TestEg4(t *testing.T) {
	/*
		下面这段代码编译不通过，原因是选择器v1.A.x和v1.B.x的深度一样，所以它们发生了碰撞，结果导致它们都不能被缩写为v1.x。
		 同样的情况发生在选择器v1.A.y和v1.B.y身上。
	*/
	var v1 struct {
		D
		E
	}
	_ = v1
	// _ = v1.x // error: 模棱两可的v1.x
	// _ = v1.y // error: 模棱两可的v1.y

	/*
		下面的代码编译没问题。选择器v2.C.B.x被另一个选择器v2.A.x遮挡了，所以v2.x实际上是选择器v2.A.x的缩写形式。
			因为同样的原因，v2.y是选择器v2.A.y（而不是选择器v2.C.B.y）的缩写形式。
	*/
	var v2 struct {
		E
		F
	}

	_ = v2.x
	_ = v2.y
}

/*
	一个被遮挡或者碰撞的选择器并不妨碍更深层的选择器被提升

	一个不寻常的但需要注意的细节是：
		来自不同库包的两个非导出方法（或者字段）将总是被认为是两个不同的标识符，即使它们的名字完全一致。
		 因此，当它们的属主类型被同时内嵌在同一个结构体类型中的时候，它们绝对不会相互碰撞或者遮挡。
*/

/*
	为内嵌了其它类型的结构体类型声明的隐式方法
		上面已经提到过，类型Singer和*Singer都有一个PrintName方法，并且类型*Singer还有一个SetAge方法。
			但是，我们从没有为这两个类型声明过这几个方法。这几个方法从哪来的呢？

		事实上，假设结构体类型S内嵌了一个类型（或者类型别名）T，并且此内嵌是合法的，
			1.对内嵌类型T的每一个方法，如果此方法对应的选择器既不和其它选择器碰撞也未被其它选择器遮挡，
				则编译器将会隐式地为结构体类型S声明一个同样原型的方法。 继而，编译器也将为指针类型*S隐式声明一个相应的方法。
			2.对类型*T的每一个方法，如果此方法对应的选择器既不和其它选择器碰撞也未被其它选择器遮挡，
				则编译器将会隐式地为类型*S声明一个同样原型的方法。

		简单说来，
			类型struct{T}和*struct{T}均将获取类型T的所有方法。
			类型*struct{T}、struct{*T}和*struct{*T}都将获取类型*T的所有方法。

		下面展示了编译器为类型Singer和*Singer隐式声明的三个（提升）方法：
			// 注意：这些声明不是合法的Go语法。这里这样表示只是为了解释目的。
			// 它们有助于解释提升方法值是如何被估值的。
			func (s Singer) PrintName = s.Person.PrintName
			func (s *Singer) PrintName = s.Person.PrintName
			func (s *Singer) SetAge = s.Person.SetAge
		右边的部分为各个提升方法相应的完整形式选择器形式。

		从方法中，我们得知我们不能为非定义的结构体类型（和基类型为非定义结构体类型的指针类型）声明方法。
			但是，通过类型内嵌，这样的类型也可以拥有方法。

		如果一个结构体类型内嵌了一个实现了一个接口类型的类型（此内嵌类型可以是此接口类型自己），
			则一般说来，此结构体类型也实现了此接口类型，除非发生了选择器碰撞和遮挡。

		请注意：一个类型将只会获取它（直接或者间接）内嵌了的类型的方法。 换句话说，
			一个类型的方法集由为类型直接（显式或者隐式）声明的方法和此类型的内嵌类型的方法集组成。 比如，在下面的例子中，
			1.类型Age没有方法，因为代码中既没有为它声明任何方法，它也没有内嵌任何类型，。
			2.类型X有两个方法：IsOdd和Double。 其中IsOdd方法是通过内嵌类型MyInt而得来的。
			3.类型Y没有方法，因为它所内嵌的类型Age没有方法，另外代码中也没有为它声明任何方法。
			4.类型Z只有一个方法：IsOdd。 此方法是通过内嵌类型MyInt而得来的。 它没有获取到类型X的Double方法，因为它并没有内嵌类型X。

*/
type MyInt int

func (mi MyInt) IsOdd() bool {
	return mi%2 == 1
}

type Age MyInt

type X struct {
	MyInt
}

func (x X) Double() MyInt {
	return x.MyInt + x.MyInt
}

type Y struct {
	Age
}

type Z X

/*
	提升方法值的正规化和估值
		假设v.m是一个合法的提升方法表达式，在编译时刻，编译器将把此提升方法表达式正规化。
		 正规化过程分为两步：首先找出此提升方法表达式的完整形式；然后将此完整形式中的隐式取地址和解引用操作均转换为显式操作。

		和其它方法值估值的规则一样，对于一个已经正规化的方法值表达式v.m，在运行时刻，当v.m被估值的时候，
			属主实参v的估值结果的一个副本将被存储下来以供后面调用此方法值的时候使用。

		以下面的代码为例：
			1.提升方法表达式s.M1的完整形式为s.T.X.M1。 将此完整形式中的隐式取地址和解引用操作转换为显式操作之后的结果为(*s.T).X.M1。
				 在运行时刻，属主实参(*s.T).X被估值并且估值结果的一个副本被存储下来以供后用。
				  此估值结果为1，这就是为什么调用f()总是打印出1。

			2.提升方法表达式s.M2的完整形式为s.T.X.M2。 将此完整形式中的隐式取地址和解引用操作转换为显式操作之后的结果为(&(*s.T).X).M2。
			  在运行时刻，属主实参&(*s.T).X被估值并且估值结果的一个副本被存储下来以供后用。
			   此估值结果为提升字段s.X（也就是(*s.T).X）的地址。 任何对s.X的修改都可以通过解引用此地址而反映出来，
			   	但是对s.T的修改是不会通过此地址反映出来的。 这就是为什么两个g()调用都打印出了2。
*/
type H int

func (h H) M1() {
	fmt.Println(h)
}

func (h *H) M2() {
	fmt.Println(*h)
}

type J struct{ *H }

type K struct{ *J }

func TestEg6(te *testing.T) {
	var h = H(1)
	var j = J{H: &h}
	// var k = K{J: j}
	// var f = k.M1 // <=> (*k.T).X.M1
	// var g = k.M2 // <=> (&(*k.T).X).M2
	// k.H = 2
	// f() // 1
	// g() // 2
	// k.J = &J{H: 3}
	// f() // 1
	// g() // 2

	// fmt.Println(reflect.TypeOf(*j).NumMethod()) //1
	// fmt.Println(reflect.TypeOf(j).NumMethod())  //2
	// fmt.Println(reflect.TypeOf(k).NumMethod())  //2
	// fmt.Println(reflect.TypeOf(&k).NumMethod()) //2
	fmt.Println(reflect.TypeOf(j).NumMethod())  //2
	fmt.Println(reflect.TypeOf(&j).NumMethod()) //2
}

// 一个有趣的类型内嵌的例子. 这个程序有什么问题?
type I interface {
	m()
}

type T struct {
	I
}

func TestEg5(te *testing.T) {
	var t T
	var i = &t
	t.I = i
	i.m()
}
