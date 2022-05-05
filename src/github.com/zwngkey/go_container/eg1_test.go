/*
 * @Author: zwngkey
 * @Date: 2022-05-03 02:18:08
 * @LastEditTime: 2022-05-04 17:38:53
 * @Description: Go中三种一等公民容器类型：数组、切片和映射。
 */

package gocontainer

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

/*
	容器类型和容器值概述
		每个容器（值）用来表示和存储一个元素（element）序列或集合。
			一个容器中的所有元素的类型是相同的。此相同的类型称为此容器的类型的元素类型（或简称此容器的元素类型）。

		存储在一个容器中的每个元素值都关联着一个键值（key）。每个元素可以通过它的键而被访问到。
			一个映射类型的键的类型必须为一个可比较类型。 数组和切片类型的键的类型均为内置类型int。
			一个数组或切片的一个元素对应的键总是一个非负整数下标，此非负整数表示该元素在该数组或切片所有元素中的顺序位置。
			此非负整数下标亦常称为一个元素索引（index）。

		每个容器值有一个长度属性，用来表明此容器中当前存储了多少个元素。
			一个数组或切片中的每个元素所关联的非负整数索引键值的合法取值范围为左闭右开区间[0, 此数组或切片的长度)。
			 一个映射类型的容器值中的元素关联的键可以是任何此映射类型的键的类型的任何值。
			 	//map[string]int 上面这句话的意思是:key可以是string类型的任何值,value可以是int类型的任何值.

		这三种容器类型的值在使用上有很多的差别。这些差别多源于它们的内存定义的差异。
			通过值部的概念，我们得知每个数组值仅由一个直接部分组成，而一个切片或者映射值
				是由一个直接部分和一个可能的被此直接部分引用着的间接部分组成。

		一个数组或者切片的所有元素紧挨着存放在一块连续的内存中。
		 一个数组中的所有元素均存放在此数组值的直接部分，一个切片中的所有元素均存放在此切片值的间接部分。
		   在官方标准编译器和运行时中，映射是使用哈希表算法来实现的。
		   	所以一个映射中的所有元素也均存放在一块连续的内存中，但是映射中的元素并不一定紧挨着存放。
			 另外一种常用的映射实现算法是二叉树算法。
			  无论使用何种算法，一个映射中的所有元素的键也存放在此映射值（的间接部分）中。

		我们可以通过一个元素的键来访问此元素。 对于这三种容器，元素访问的时间复杂度均为O(1)。
		但是一般来说，映射元素访问消耗的时长要数倍于数组和切片元素访问消耗的时长。
		但是映射相对于数组和切片有两个优点：
			1.映射的键的类型可以是任何可比较类型。
			2.相对于使用含有大量稀疏索引的数组和切片，使用映射可以节省大量的内存。


		在任何赋值中，源值的底层间接部分不会被复制。 换句话说，当一个赋值结束后，
			一个含有间接部分的源值和目标值将共享底层间接部分。
				这就是数组和切片/映射值会有很多行为差异的原因


	非定义容器类型的字面表示形式
		数组类型：[N]T
		切片类型：[]T
		映射类型：map[K]T

		其中:
			1.T可为任意类型。它表示一个容器类型的元素类型。某个特定容器类型的值中只能存储此容器类型的元素类型的值。
			2.N必须为一个非负整数常量。它指定了一个数组类型的长度，或者说它指定了此数组类型的任何一个值中存储了多少个元素。
				 一个数组类型的长度是此数组类型的一部分。比如[5]int和[8]int是两个不同的类型。
			3.K必须为一个可比较类型。它指定了一个映射类型的键类型。
*/
/*
	一些非定义容器类型的字面表示：
		const Size = 32

		type Person struct {
			name string
			age  int
		}

		// 数组类型
		[5]string
		[Size]int
		[16][]byte  // 元素类型为一个切片类型：[]byte
		[100]Person // 元素类型为一个结构体类型：Person

		// 切片类型
		[]bool
		[]int64
		[]map[int]bool // 元素类型为一个映射类型：map[int]bool
		[]*int         // 元素类型为一个指针类型：*int

		// 映射类型
		map[string]int
		map[int]bool
		map[int16][6]string     // 元素类型为一个数组类型：[6]string
		map[bool][]string       // 元素类型为一个切片类型：[]string
		map[struct{x int}]*int8 // 元素类型为一个指针类型：*int8；
								// 键值类型为一个结构体类型。


		所有切片类型的尺寸都是一致的，所有映射类型的尺寸也都是一致的。
			一个数组类型的尺寸等于它的元素类型的尺寸和它的长度的乘积。
				长度为零的数组的尺寸为零；元素类型尺寸为零的任意长度的数组类型的尺寸也为零。
*/

/*
	容器值字面量的表示形式
		和结构体值类似，容器值的文字表示也可以用组合字面量形式（composite literal）来表示。
			比如对于一个容器类型T，它的值可以用形式T{...}来表示（除了切片和映射的零值外）。

	下面是一些容器值字面量：
		// 一个含有4个布尔元素的数组值。
		[4]bool{false, true, true, false}

		// 一个含有三个字符串值的切片值。
		[]string{"break", "continue", "fallthrough"}

		// 一个映射值。
		map[string]int{"C": 1972, "Python": 1991, "Go": 2009}

	映射组合字面量中大括号中的每一项称为一个键值对（key-value pair），或者称为一个条目（entry）。

	数组和切片组合字面量有一些微小的变种：
		// 下面这些切片字面量都是等价的。
		[]string{"break", "continue", "fallthrough"}
		[]string{0: "break", 1: "continue", 2: "fallthrough"}
		[]string{2: "fallthrough", 1: "continue", 0: "break"}
		[]string{2: "fallthrough", 0: "break", "continue"}

		// 下面这些数组字面量都是等价的。
		[4]bool{false, true, true, false}
		[4]bool{0: false, 1: true, 2: true, 3: false}
		[4]bool{1: true, true}
		[4]bool{2: true, 1: true}
		[...]bool{false, true, true, false}
		[...]bool{3: false, 1: true, true}

	上例中最后两行中的...表示让编译器推断出相应数组值的类型的长度。

	从上面的例子中，我们可以看出数组和切片组合字面量中的索引下标（即数组和切片的键值）是可选的。
		在一个数组或者切片组合字面量中：
			1.如果一个索引下标出现，它的类型不必是数组和切片类型的键值类型int，但它必须是一个可以表示为int值的非负常量；
				 如果它是一个类型确定值，则它的类型必须为一个基本整数类型。
			2.在一个数组或切片组合字面量中，如果一个元素的索引下标缺失，
				则编译器认为它的索引下标为出现在它之前的元素的索引下标加一。
			3.如果出现的第一个元素的索引下标缺失，则它的索引下标被认为是0。

	映射组合字面量中元素对应的键值不可缺失，并且它们可以为非常量。
		var a uint = 1
		var _ = map[uint]int {a : 123} // 没问题
		var _ = []int{a: 100}          // error: 下标必须为常量
		var _ = [5]int{a: 100}         // error: 下标必须为常量

	一个容器组合字面量中的常量键值（包括索引下标）不可重复。


*/
func Test1(t *testing.T) {
	var a = [4]bool{1: true, true}
	fmt.Println(a[0], a[1], a[2]) //false true true

	//索引重复
	// var b = [...]bool{2: false, 1: true, true} //error
}

/*
	容器类型零值的字面量表示形式
		和结构体类似，一个数组类型A的零值可以表示为A{}。 比如，数组类型[100]int的零值可以表示为[100]int{}。
		 一个数组零值中的所有元素均为对应数组元素类型的零值。

		和指针一样，所有切片和映射类型的零值均用预声明的标识符nil来表示。
			函数、通道和接口类型的零值也用预声明的标识符nil来表示。

		在运行时刻，即使一个数组变量在声明的时候未指定初始值，它的元素所占的内存空间也已经被开辟出来。
		 	但是一个nil切片或者映射值的元素的内存空间尚未被开辟出来。

		注意：[]T{}表示类型[]T的一个空切片值，它和[]T(nil)是不等价的。 同样，map[K]T{}和map[K]T(nil)也是不等价的。
*/

func Test2(t *testing.T) {

	// nil切片与空切片
	b := []int{}
	c := []int(nil)
	var d []int
	e := make([]int, 0)
	f := make([]int, 0)
	fmt.Println(b == nil) //false
	fmt.Println(c == nil) //true
	fmt.Println(d == nil) //true
	fmt.Println(e == nil) //false
	fmt.Println(f == nil) //false

	fmt.Printf("%#v\n", *(*reflect.SliceHeader)(unsafe.Pointer(&b))) //reflect.SliceHeader{Data:0x1400018be08, Len:0, Cap:0}
	fmt.Printf("%#v\n", *(*reflect.SliceHeader)(unsafe.Pointer(&c))) //reflect.SliceHeader{Data:0x0, Len:0, Cap:0}
	fmt.Printf("%#v\n", *(*reflect.SliceHeader)(unsafe.Pointer(&d))) //reflect.SliceHeader{Data:0x0, Len:0, Cap:0}
	fmt.Printf("%#v\n", *(*reflect.SliceHeader)(unsafe.Pointer(&e))) //reflect.SliceHeader{Data:0x1400018be08, Len:0, Cap:0}
	fmt.Printf("%#v\n", *(*reflect.SliceHeader)(unsafe.Pointer(&f))) //reflect.SliceHeader{Data:0x1400018be08, Len:0, Cap:0}
	fmt.Printf("%p\n", b)
	fmt.Printf("%p\n", c)
	fmt.Printf("%p\n", d)
	fmt.Printf("%p\n", e)
	fmt.Printf("%p\n", f)

	/*
		nil切片和空切片指向的地址不一样。nil切片引用数组指针地址为0（无指向任何实际地址）

		空切片的引用数组指针地址是有的，且固定为一个值，即引用地址都是一样的
	*/
}

/*
	容器字面量是不可寻址的但可以被取地址
*/
func Test3(t *testing.T) {
	pm := &map[string]int{"C": 1972, "Go": 2009}
	ps := &[]string{"break", "continue"}
	pa := &[...]bool{false, true, true, false}
	fmt.Printf("%T\n", pm) // *map[string]int
	fmt.Printf("%T\n", ps) // *[]string
	fmt.Printf("%T\n", pa) // *[4]bool
}

/*
	内嵌组合字面量可以被简化
		在某些情形下，内嵌在其它组合字面量中的组合字面量可以简化为{...}（即类型部分被省略掉了）。
			内嵌组合字面量前的取地址操作符&有时也可以被省略。
*/
func Test4(t *testing.T) {
	// heads为一个切片值。它的类型的元素类型为*[4]byte。
	// 此元素类型为一个基类型为[4]byte的指针类型。
	// 此指针基类型为一个元素类型为byte的数组类型。
	var _ = []*[4]byte{
		&[4]byte{'P', 'N', 'G', ' '},
		&[4]byte{'G', 'I', 'F', ' '},
		&[4]byte{'J', 'P', 'E', 'G'},
	}
	// 可以被简化为
	_ = []*[4]byte{
		{'P', 'N', 'G', ' '},
		{'G', 'I', 'F', ' '},
		{'J', 'P', 'E', 'G'},
	}

	// 下面这个数组组合字面量
	type language struct {
		name string
		year int
	}

	var _ = [...]language{
		language{"C", 1972},
		language{"Python", 1991},
		language{"Go", 2009},
	}
	// 可以被简化为
	var _ = [...]language{
		{"C", 1972},
		{"Python", 1991},
		{"Go", 2009},
	}

	// 下面这个映射组合字面量
	type LangCategory struct {
		dynamic bool
		strong  bool
	}

	// 此映射值的类型的键值类型为一个结构体类型，
	// 元素类型为另一个映射类型：map[string]int。
	var _ = map[LangCategory]map[string]int{
		LangCategory{true, true}: map[string]int{
			"Python": 1991,
			"Erlang": 1986,
		},
		LangCategory{true, false}: map[string]int{
			"JavaScript": 1995,
		},
		LangCategory{false, true}: map[string]int{
			"Go":   2009,
			"Rust": 2010,
		},
		LangCategory{false, false}: map[string]int{
			"C": 1972,
		},
	}
	// 可以被简化为
	var _ = map[LangCategory]map[string]int{
		{true, true}: {
			"Python": 1991,
			"Erlang": 1986,
		},
		{true, false}: {
			"JavaScript": 1995,
		},
		{false, true}: {
			"Go":   2009,
			"Rust": 2010,
		},
		{false, false}: {
			"C": 1972,
		},
	}
	// 注意，在上面的几个例子中，最后一个元素后的逗号不能被省略。
}

/*
	容器值的比较
		任意两个映射值（或切片值）是不能相互比较的。

		尽管两个映射值和切片值是不能比较的，但是一个映射值或者切片值可以和预声明的nil标识符进行比较以检查此映射值或者切片值是否为一个零值。

		大多数数组类型都是可比较类型，除了元素类型为不可比较类型的数组类型。

		当比较两个数组值时，它们的对应元素将按照逐一被比较（可以认为按照下标顺序比较）。
			这两个数组只有在它们的对应元素都相等的情况下才相等；当一对元素被发现不相等的或者在比较中产生恐慌的时候，对数组的比较将提前结束。

	查看容器值的长度和容量
		一个数组值的容量总是和它的长度相等；一个非零映射值的容量可以被认为是无限大的。
		 	一个切片值的容量总是不小于此切片值的长度。在编程中，只有切片值的容量有实际意义。

		我们可以调用内置函数len来获取一个容器值的长度，或者调用内置函数cap来获取一个容器值的容量。
		 这两个函数都返回一个int类型确定结果值。 因为非零映射值的容量是无限大，所以cap并不适用于映射值。

		一个数组值的长度和容量永不改变。同一个数组类型的所有值的长度和容量都总是和此数组类型的长度相等。
		 切片值的长度和容量可在运行时刻改变。因为此原因，切片可以被认为是动态数组。
		  切片在使用上相比数组更为灵活，所以切片（相对数组）在编程用得更为广泛。


	读取和修改容器的元素
		一个容器值v中存储的对应着键值k的元素用语法形式v[k]来表示。 今后我们称v[k]为一个元素索引表达式。

		假设v是一个数组或者切片，在v[k]中，
			1.如果k是一个常量，则它必须满足出现在组合字面量中的索引的要求。 另外，如果v是一个数组，则k必须小于此数组的长度。
			2.如果k不是一个常量，则它必须为一个整数。 另外它必须为一个非负数并且小于len(v)，否则，在运行时刻将产生一个恐慌。
			3.如果v是一个零值切片，则在运行时刻将产生一个恐慌。

		假设v是一个映射值，在v[k]中，k的类型必须为（或者可以隐式转换为）v的类型的键值类型。另外，
			1.如果k是一个动态类型为不可比较类型的接口值，则v[k]在运行时刻将造成一个恐慌；
			2.如果v[k]被用做一个赋值语句中的目标值并且v是一个零值nil映射，则v[k]在运行时刻将造成一个恐慌；
			3.如果v[k]用来表示读取映射值v中键值k对应的元素，则它无论如何都不会产生一个恐慌，即使v是一个零值nil映射（假设k的估值没有造成恐慌）；
			4.如果v[k]用来表示读取映射值v中键值k对应的元素，并且映射值v中并不含有对应着键值k的条目，则v[k]返回一个此映射值的类型的元素类型的零值。
				一般情况下，v[k]被认为是一个单值表达式。但是在一个v[k]被用为唯一源值的赋值语句中，v[k]可以返回一个可选的第二个返回值。
					此第二个返回值是一个类型不确定布尔值，用来表示是否有对应着键值k的条目存储在映射值v中。

*/

func Test5(t *testing.T) {
	var f any = func() {

	}
	var s any = []int{}

	var m any = map[string]int{}

	var str = "string"
	var i = 1

	var m1 = map[any]int{
		i:   1,
		str: 4,
		f:   1,
		s:   2,
		m:   3,
	}
	_ = m1
	//如果k是一个动态类型为不可比较类型的接口值，则v[k]在运行时刻将造成一个恐慌；
	// var a = m1[f] //error
	// fmt.Println(a)

	// var a = m1[str] //error
	// fmt.Println(a)

	// var b = m1[1] //error
	// fmt.Println(b)
}

/*
	容器赋值
		当一个映射赋值语句执行完毕之后，目标映射值和源映射值将共享底层的元素。
			向其中一个映射中添加（或从中删除）元素将体现在另一个映射中。

		和映射一样，当一个切片赋值给另一个切片后，它们将共享底层的元素。它们的长度和容量也相等。
			但是和映射不同，如果以后其中一个切片改变了长度或者容量，此变化不会体现到另一个切片中。

		当一个数组被赋值给另一个数组，所有的元素都将被从源数组复制到目标数组。赋值完成之后，这两个数组不共享任何元素。
*/
func Test6(t *testing.T) {
	m0 := map[int]int{0: 7, 1: 8, 2: 9}
	m1 := m0
	m1[0] = 2
	fmt.Println(m0, m1) // map[0:2 1:8 2:9] map[0:2 1:8 2:9]

	s0 := []int{7, 8, 9}
	s1 := s0
	s1[0] = 2
	fmt.Println(s0, s1) // [2 8 9] [2 8 9]

	a0 := [...]int{7, 8, 9}
	a1 := a0
	a1[0] = 2
	fmt.Println(a0, a1) // [7 8 9] [2 8 9]
}

/*
	添加和删除容器元素
		向一个映射中添加一个条目的语法和修改一个映射元素的语法是一样的。
			比如，对于一个非零映射值m，如果当前m中尚未存储条目(k, e)，则下面的语法形式将把此条目存入m；
				如果存在，下面的语法形式将把键值k对应的元素值更新为e。
					m[k] = e

		内置函数delete用来从一个映射中删除一个条目。比如，下面的delete调用将把键值k对应的条目从映射m中删除。
		 如果映射m中未存储键值为k的条目，则此调用为一个空操作，它不会产生一个恐慌，即使m是一个nil零值映射。
		 	delete(m, k)

		一个数组中的元素个数总是恒定的，我们无法向其中添加元素，也无法从其中删除元素。
			但是可寻址的数组值中的元素是可以被修改的。

		我们可以通过调用内置append函数，以一个切片为基础，来添加不定数量的元素并返回一个新的切片。
			此新的结果切片包含着基础切片中所有的元素和所有被添加的元素。 注意，基础切片并未被此append函数调用所修改。
			 当然，如果我们愿意（事实上在实践中常常如此），我们可以将结果切片赋值给基础切片以修改基础切片。

		Go中并未提供一个内置方式来从一个切片中删除一个元素。 我们必须使用append函数
			和子切片语法一起来实现元素删除操作。

		注意，内置append函数是一个变长参数函数。 它有两个参数，其中第二个参数（形参）为一个变长参数。

		变长参数函数调用中的实参有两种传递方式
			对于三个点方式，append函数并不要求第二个实参的类型和第一个实参一致，但是它们的底层类型必须一致。

		请注意，当一个`append`函数调用需要为结果切片开辟内存时，结果切片的容量取决于具体编译器实现。
			在这种情况下，对于官方标准编译器，如果基础切片的容量较小，则结果切片的容量至少为基础切片的两倍。
			 这样做的目的是使结果切片有足够多的冗余元素槽位，以防止此结果切片被用做后续其它`append`函数调用的基础切片时再次开辟内存。

		`append`函数调用的第一个实参不能为类型不确定的`nil`。可以为类型确定的`nil`.如下面的代码:
				s := append([]string(nil),"string","int")

*/

/*
	打印一个映射变量时,结果中的条目顺序固定，两次打印结果相同。
		fmt.Println(m)

	循环变量一个映射时,每次打印结果中条目的输出顺序不固定.
*/
func Test7(t *testing.T) {
	var m = map[string]int{
		"张三": 1,
		"李四": 2,
		"李1": 3,
		"李2": 4,
		"李3": 5,
	}
	//结果固定
	fmt.Println(m)
	//结果不固定
	for k, v := range m {
		fmt.Println(k, "->", v)
	}

}

/*
	append函数传参方式中
		对于三个点方式，append函数并不要求第二个实参的类型和第一个实参一致，但是它们的底层类型必须一致。
*/
func Test8(t *testing.T) {
	type Myint []int
	type Myint2 Myint
	s1 := []int{1, 2, 3}
	s := Myint{1, 2, 3}
	s2 := Myint2{1, 2, 3}
	s3 := append(s1, s...)
	s4 := append(s1, s2...)
	_ = s3
	_ = s4

}

/*
	使用内置make函数来创建切片和映射
		除了使用组合字面量来创建映射和切片，我们还可以使用内置make函数来创建映射和切片。 数组不能使用内置make函数来创建。

		假设M是一个映射类型并且n是一个整数，我们可以用下面的两种函数调用来各自生成一个类型为M的映射值。
			make(M, n)
			make(M)
		 第一个函数调用形式创建了一个可以容纳至少n个条目而无需再次开辟内存的空映射值。
		 第二个函数调用形式创建了一个可以容纳一个小数目的条目而无需再次开辟内存的空映射值。此小数目的值取决于具体编译器实现。
		 	注意：第二个参数n可以为负或者零，这时对应的调用将被视为上述第二种调用形式。


		假设S是一个切片类型，length和capacity是两个非负整数，并且length小于等于capacity，
			我们可以用下面的两种函数调用来各自生成一个类型为S的切片值。length和capacity的类型必须均为整数类型（两者可以不一致）。
			make(S, length, capacity)
			make(S, length) // <=> make(S, length, length)

		使用make函数创建的切片中的所有元素值均被初始化为结果切片的元素类型的零值。


	使用内置new函数来创建容器值
		内置new函数可以用来为一个任何类型的值开辟内存并返回一个存储有此值的地址的指针。
			用new函数开辟出来的值均为零值。因为这个原因，new函数对于创建映射和切片值来说没有任何价值。


	容器元素的可寻址性
		1.如果一个数组是可寻址的，则它的元素也是可寻址的；反之亦然，即如果一个数组是不可寻址的，则它的元素也是不可寻址的。
			 原因很简单，因为一个数组只含有一个（直接）值部，并且它的所有元素和此直接值部均承载在同一个内存块上。
		2.一个切片值的任何元素都是可寻址的，即使此切片本身是不可寻址的。
			这是因为一个切片的底层元素总是存储在一个被开辟出来的内存片段（间接值部）上。
		3.任何映射元素(value部分)都是不可寻址的。

		一般来说，一个不可寻址的值的直接部分是不可修改的。但是映射元素是个例外。
			映射元素虽然不可寻址，但是每个映射元素可以被整个修改（但不能被部分修改）。注意：这个限制可能会在以后被移除。
			 对于大多数做为映射元素类型的类型，在修改它们的值的时候，一般体现不出来整个修改和部分修改的差异。
			  但是如果一个映射的元素类型为数组或者结构体类型，这个差异是很明显的。

		每个数组或者结构体值都是仅含有一个直接部分。所以
			如果一个映射类型的元素类型为一个结构体类型，
				则我们无法修改此映射类型的值中的每个结构体元素的单个字段。 我们必须整体地同时修改所有结构体字段。

			如果一个映射类型的元素类型为一个数组类型，则我们无法修改此映射类型的值中的每个数组元素的单个元素。
				 我们必须整体地同时修改所有数组元素。


*/
func Test9(t *testing.T) {
	type T struct{ age int }
	mt := map[string]T{}
	mt["John"] = T{age: 29} // 整体修改是允许的

	ma := map[int][5]int{}
	ma[1] = [5]int{1: 789} // 整体修改是允许的

	// 注意：这个限制可能会在以后被移除。
	// 这两个赋值编译不通过，因为部分修改一个映射
	// 元素是非法的。这看上去确实有些反直觉。
	/*
		ma[1][1] = 123      // error
		mt["John"].age = 30 // error
	*/

	// 读取映射元素的元素或者字段是没问题的。
	fmt.Println(ma[1][1])       // 789
	fmt.Println(mt["John"].age) // 29
}

/*
	为了让上例中的两行编译不通过的两行赋值语句编译通过，欲修改的映射元素必须先存放在一个临时变量中，
		然后修改这个临时变量，最后再用这个临时变量整体覆盖欲修改的映射元素。比如：
*/
func Test10(te *testing.T) {

	type T struct{ age int }
	mt := map[string]T{}
	mt["John"] = T{age: 29}
	ma := map[int][5]int{}
	ma[1] = [5]int{1: 789}

	t := mt["John"] // 临时变量
	t.age = 30
	mt["John"] = t // 整体修改

	a := ma[1] // 临时变量
	a[1] = 123
	ma[1] = a // 整体修改

	fmt.Println(ma[1][1], mt["John"].age) // 123 30

}

/*
	从数组或者切片派生切片（取子切片）
		我们可以从一个基础切片或者一个可寻址的基础数组派生出另一个切片。此派生操作也常称为一个取子切片操作
		 派生出来的切片的元素和基础切片（或者数组）的元素位于同一个内存片段上。
		 或者说，派生出来的切片和基础切片（或者数组）将共享一些元素。

		Go中有两种取子切片的语法形式（假设baseContainer是一个切片或者数组）：
			baseContainer[low : high]       // 双下标形式
			baseContainer[low : high : max] // 三下标形式

		上面所示的双下标形式等价于下面的三下标形式：
			baseContainer[low : high : cap(baseContainer)]

		上面所示的取子切片表达式的语法形式中的下标必须满足下列关系，否则代码要么编译不通过，要么在运行时刻将造成恐慌
			// 双下标形式
			0 <= low <= high <= cap(baseContainer)

			// 三下标形式
			0 <= low <= high <= max <= cap(baseContainer)

		注意：
			只要上述关系均满足，下标low和high都可以大于len(baseContainer)。但是它们一定不能大于cap(baseContainer)。

			如果baseContainer是一个零值nil切片，只要上面所示的子切片表达式中下标的值均为0，
				则这两个子切片表达式不会造成恐慌。 在这种情况下，结果切片也是一个nil切片。

		子切片表达式的结果切片的长度为high - low、容量为max - low。 派生出来的结果切片的长度可能大于基础切片的长度，
			但结果切片的容量绝不可能大于基础切片的容量。

		在实践中，我们常常在子切片表达式中省略若干下标，以使代码看上去更加简洁。省略规则如下：
			1.如果下标low为零，则它可被省略。此条规则同时适用于双下标形式和三下标形式。
			2.如果下标high等于len(baseContainer)，则它可被省略。此条规则同时只适用于双下标形式。
			3.三下标形式中的下标max在任何情况下都不可被省略。
			比如，下面的子切片表达式都是相互等价的：
				baseContainer[0 : len(baseContainer)]
				baseContainer[: len(baseContainer)]
				baseContainer[0 :]
				baseContainer[:]
				baseContainer[0 : len(baseContainer) : cap(baseContainer)]
				baseContainer[: len(baseContainer) : cap(baseContainer)]


		请注意，子切片操作有可能会造成暂时性的内存泄露。 比如，下面在这个函数中开辟的内存块中的前50个元素槽位在它的调用返回之后将不再可见。 这50个元素槽位所占内存浪费了，这属于暂时性的内存泄露。 当这个函数中开辟的内存块今后不再被任何切片所引用，此内存块将被回收，这时内存才不再继续泄漏。
			func f() []int {
				s := make([]int, 10, 100)
				return s[50:60]
			}
		请注意，在上面这个函数中，子切片表达式中的起始下标（50）比s的长度（10）要大，这是允许的。

*/
func f() []int {
	return make([]int, 10, 100)[50:60]
}
func Test12(t *testing.T) {
	a := []int(nil)
	b := a[:]
	fmt.Println(b == nil) //true
	c := f()
	fmt.Println(c)
	fmt.Println(len(c))
	fmt.Println(cap(c))
}

/*
	切片转化为数组指针
		从Go 1.17开始，一个切片可以被转化为一个相同元素类型的数组的指针类型。
			但是如果数组的长度大于被转化切片的长度，则将导致恐慌产生。
*/
type S []int
type A [2]int
type P *A

func Test11(t *testing.T) {
	var x []int
	var y = make([]int, 0)
	var x0 = (*[0]int)(x) // okay, x0 == nil
	var y0 = (*[0]int)(y) // okay, y0 != nil
	_, _ = x0, y0

	//z,a,b 引用同一内存段.
	var z = make([]int, 3, 5)
	var a = (*[3]int)(z) // okay
	var b = (*[2]int)(z) // okay
	var _ = (*A)(z)      // okay
	var _ = P(z)         // okay

	fmt.Printf("%p\n", z)     //0x1400010a120
	fmt.Printf("%p\n", a)     //0x1400010a120
	fmt.Printf("%p\n", b)     //0x1400010a120
	fmt.Printf("%p\n", &z[0]) //0x1400010a120
	fmt.Printf("%p\n", &a[0]) //0x1400010a120
	fmt.Printf("%p\n", &b[0]) //0x1400010a120

	fmt.Println(z[1], a[1], b[1])
	z[1] = 1
	fmt.Println(z[1], a[1], b[1])

	var w S = z
	var _ = (*[3]int)(w) // okay
	var _ = (*[2]int)(w) // okay
	var _ = (*A)(w)      // okay
	var _ = P(w)         // okay

	// var _ = (*[4]int)(z) // 会产生恐慌
}

/*
	使用内置copy函数来复制切片元素
		我们可以使用内置copy函数来将一个切片中的元素复制到另一个切片。 这两个切片的类型可以不同，但是它们的元素类型必须相同。
			换句话说，这两个切片的类型的底层类型必须相同。 copy函数的第一个参数为目标切片，第二个参数为源切片。
			 传递给一个copy函数调用的两个实参可以共享一些底层元素。
			 	copy函数返回复制了多少个元素，此值（int类型）为这两个切片的长度的较小值。

		结合子切片语法，我们可以使用copy函数来在两个数组之间或者一个数组与一个切片之间复制元素。

		注意，做为一个特例，copy函数可以用来将一个字符串中的字节复制到一个字节切片。

		截至目前（Go 1.17），copy函数调用的两个实参均不能为类型不确定的nil。
*/
func Test13(t *testing.T) {
	arr := [...]int{1, 2, 3, 4, 5}
	s1 := arr[:3]
	s2 := arr[1:5]
	fmt.Println(arr)    //[1,2,3,4,5]
	fmt.Println(s1, s2) //[1,2,3],[2,3,4,5]
	copy(s1, s2)
	fmt.Println(arr) //[2,3,4,4,5]
	fmt.Println(s1)  //[2,3,4]
	fmt.Println(s2)  //[3,4,4,5]

	fmt.Printf("%p\n", &arr)    // 0x1400010c120
	fmt.Printf("%p\n", &arr[0]) // 0x1400010c120
	fmt.Printf("%p\n", &arr[1]) // 0x1400010c128
	fmt.Printf("%p\n", s1)      // 0x1400010c120
	fmt.Printf("%p\n", s2)      // 0x1400010c128

	s := "string"
	b := make([]byte, 6)
	copy(b, s)
	fmt.Println(string(b))
}

/*
	遍历容器元素:for-range循环
		遍历一个nil映射或者nil切片是允许的。这样的遍历可以看作是一个空操作。

		一些关于遍历映射条目的细节：
			1.映射中的条目的遍历顺序是不确定的（可以认为是随机的）。或者说，同一个映射中的条目的两次遍历中，
				条目的顺序很可能是不一致的，即使在这两次遍历之间，此映射并未发生任何改变。
			2.如果在一个映射中的条目的遍历过程中，一个还没有被遍历到的条目被删除了，则此条目保证不会被遍历出来。
			3.如果在一个映射中的条目的遍历过程中，一个新的条目被添加入此映射，则此条目并不保证将在此遍历过程中被遍历出来。

		对一个for-range循环代码块:
			for key, element := range aContainer {...}
			有三个重要的事实存在：
				1.被遍历的容器值是aContainer的一个副本。 注意，只有aContainer的直接部分被复制了。
					此副本是一个匿名的值，所以它是不可被修改的。
					 1.如果aContainer是一个数组，那么在遍历过程中对此数组元素的修改不会体现到循环变量中。
					    原因是此数组的副本（被真正遍历的容器）和此数组不共享任何元素。
					 2.如果aContainer是一个切片（或者映射），那么在遍历过程中对此切片（或者映射）元素的修改将体现到循环变量中。
					 	原因是此切片（或者映射）的副本和此切片（或者映射）共享元素（或条目）。
				2.在遍历中的每个循环步，aContainer副本中的一个键值元素对将被赋值（复制）给循环变量。
					所以对循环变量的直接部分的修改将不会体现在aContainer中的对应元素中。
					（因为这个原因，并且for-range循环是遍历映射条目的唯一途径，所以最好不要使用大尺寸的映射键值和元素类型，以避免较大的复制负担。）
				3.所有被遍历的键值对将被赋值给同一对循环变量实例。


		复制一个切片或者映射的代价很小，但是复制一个大尺寸的数组的代价比较大。
			所以，一般来说，range关键字后跟随一个大尺寸数组不是一个好主意。
			 如果我们要遍历一个大尺寸数组中的元素，我们以遍历从此数组派生出来的一个切片，或者遍历一个指向此数组的指针

*/
// 验证上述第一个和第二个事实。
func Test14(t *testing.T) {
	type Person struct {
		name string
		age  int
	}
	persons := [2]Person{{"Alice", 28}, {"Bob", 25}}
	for i, p := range persons {
		fmt.Println(i, p)
		// 此修改将不会体现在这个遍历过程中，
		// 因为被遍历的数组是persons的一个副本。
		persons[1].name = "Jack"

		// 此修改不会反映到persons数组中，因为p
		// 是persons数组的副本中的一个元素的副本。
		p.age = 31
	}
	fmt.Println("persons:", persons)
	/*
		输出结果:
		0 {Alice 28}
		1 {Bob 25}
		persons: &[{Alice 28} {Jack 25}]
	*/
}

// 如果我们将上例中的数组改为一个切片，则在循环中对此切片的修改将在循环过程中体现出来。 但是对循环变量的修改仍然不会体现在此切片中。
func Test15(t *testing.T) {
	type Person struct {
		name string
		age  int
	}
	// 改为一个切片。
	persons := []Person{{"Alice", 28}, {"Bob", 25}}
	for i, p := range persons {
		fmt.Println(i, p)
		// 这次，此修改将反映在此次遍历过程中。
		persons[1].name = "Jack"
		// 这个修改仍然不会体现在persons切片容器中。
		p.age = 31
	}
	fmt.Println("persons:", &persons)
	/*
		输出结果:
			0 {Alice 28}
			1 {Jack 25}
			persons: &[{Alice 28} {Jack 25}]
	*/
}

// 验证上述的第二个和第三个事实：
func Test16(t *testing.T) {
	langs := map[struct{ dynamic, strong bool }]map[string]int{
		{true, false}:  {"JavaScript": 1995},
		{false, true}:  {"Go": 2009},
		{false, false}: {"C": 1972},
	}
	// 此映射的键值和元素类型均为指针类型。
	m0 := map[*struct{ dynamic, strong bool }]*map[string]int{}

	for category, langInfo := range langs {
		m0[&category] = &langInfo
		// 下面这行修改对映射langs没有任何影响。
		category.dynamic, category.strong = true, true
	}

	for category, langInfo := range langs {
		fmt.Println(category, langInfo)
	}

	m1 := map[struct{ dynamic, strong bool }]map[string]int{}
	for category, langInfo := range m0 {
		m1[*category] = *langInfo
	}
	// 映射m0和m1中均只有一个条目。
	fmt.Println(len(m0), len(m1)) // 1 1
	fmt.Println(m1)               // m1的值是不确定的.
}

/*
	对于一个数组或者切片，如果它的元素类型的尺寸较大，则一般来说，用第二个循环变量来存储每个循环步中被遍历的元素不是一个好主意。
		对于这样的数组或者切片，我们最好忽略或者舍弃for-range代码块中的第二个循环变量，或者使用传统的for循环来遍历元素。
		 比如，在下面这个例子中，函数fa中的循环效率比函数fb中的循环低得多。
*/

type Buffer struct {
	start, end int
	data       [1024]byte
}

func Fa(buffers []Buffer) int {
	numUnreads := 0
	for _, buf := range buffers {
		numUnreads += buf.end - buf.start
	}
	return numUnreads
}

func Fb(buffers []Buffer) int {
	numUnreads := 0
	for i := range buffers {
		numUnreads += buffers[i].end - buffers[i].start
	}
	return numUnreads
}

/*
	把数组指针当做数组来使用
		对于某些情形，我们可以把数组指针当做数组来使用。
		我们可以通过在range关键字后跟随一个数组的指针来遍历此数组中的元素。
			 对于大尺寸的数组，这种方法比较高效，因为复制一个指针比复制一个大尺寸数组的代价低得多。
			  下面的例子中的两个循环是等价的，它们的效率也基本相同
*/
func Test17(t *testing.T) {
	var a [100]int

	for i, n := range &a { // 复制一个指针的开销很小
		fmt.Println(i, n)
	}

	for i, n := range a[:] { // 复制一个切片的开销很小
		fmt.Println(i, n)
	}
}

/*
	如果一个for-range循环中的第二个循环变量既没有被忽略，也没有被舍弃，并且range关键字后跟随一个nil数组指针
		，则此循环将造成一个恐慌。 在下面这个例子中，前两个循环都将打印出5个下标，但最后一个循环将导致一个恐慌。
*/

func Test18(t *testing.T) {
	var p *[5]int // nil

	for i, _ := range p { // okay
		fmt.Println(i)
	}

	for i := range p { // okay
		fmt.Println(i)
	}

	for i, n := range p { // panic
		fmt.Println(i, n)
	}
}

/*
	我们可以通过数组的指针来访问和修改此数组中的元素。如果此指针是一个nil指针，将导致一个恐慌。
*/
func Test19(t *testing.T) {
	a := [5]int{2, 3, 5, 7, 11}
	p := &a
	p[0], p[1] = 17, 19
	fmt.Println(a) // [17 19 5 7 11]
	p = nil
	_ = p[0] // panic
}

/*
	我们可以从一个数组的指针派生出一个切片。从一个nil数组指针派生切片将导致一个恐慌。
*/
func Test20(t *testing.T) {
	pa := &[5]int{2, 3, 5, 7, 11}
	s := (*pa)[1:3]
	fmt.Println(s) // [3 5]
	pa = nil
	s = pa[0:0] // panic
	// 如果下一行能被执行到，则它也会产生恐慌。
	_ = (*[0]byte)(nil)[:]
}

/*
	内置len和cap函数调用接受数组指针做为实参。 nil数组指针实参不会导致恐慌。
		var pa *[5]int // == nil
		fmt.Println(len(pa), cap(pa)) // 5 5
*/
func Test21(t *testing.T) {
	var pa *[5]int                // == nil
	fmt.Println(len(pa), cap(pa)) // 5 5
}

/*
	memclr优化
	 	数组和切片元素的重置操作将被优化为一个内部的memclr函数调用.

		假设t0是一个类型T的零值字面量，并且a是一个元素类型为T的数组或者切片，
			则官方标准编译器将把下面的单循环变量for-range代码块优化为一个内部的memclr调用。
				大多数情况下，此memclr调用比一个一个地重置元素要快。
					for i := range a {
						a[i] = t0
					}

		此优化在官方标准编译器1.5版本中被引入。

		此优化不适用于a为一个数组指针的情形，至少到目前（Go 1.17）为止是这样。 所以，如果你打算重置一个数组，最好不要在range关键字后跟随此数组的指针。

*/

/*
	内置函数len和cap的调用可能会在编译时刻被估值
		如果传递给内置函数len或者cap的一个调用的实参是一个数组或者数组指针，则此调用将在编译时刻被估值。
			此估值结果是一个类型为内置类型int的类型确定常量值。
*/
var a1 [5]int
var p1 *[7]string

// N和M都是类型为int的类型确定值。
const N = len(a1)
const M = cap(p1)

func Test22(t *testing.T) {
	fmt.Println(N) // 5
	fmt.Println(M) // 7
}

/*
	单独修改一个切片的长度或者容量
		上面已经提到了，一般来说，一个切片的长度和容量不能被单独修改。一个切片只有通过赋值的方式被整体修改。
			但是，事实上，我们可以通过反射的途径来单独修改一个切片的长度或者容量。

		传递给函数SetLen的实参值必须不大于原切片值的容量。
		传递给函数SetCap的实参值必须不小于原切片值的长度并且须不大于原切片值的容量。 否则，在运行时刻将产生一个恐慌。

		此反射方法的效率很低，远低于一个切片的赋值。
*/
func Test23(t *testing.T) {
	s := make([]int, 2, 6)
	fmt.Println(len(s), cap(s)) // 2 6
	fmt.Printf("%p\n", s)

	reflect.ValueOf(&s).Elem().SetLen(3)
	fmt.Println(len(s), cap(s)) // 3 6

	reflect.ValueOf(&s).Elem().SetCap(5)
	fmt.Println(len(s), cap(s)) // 3 5

}

/*
	更多切片操作
		Go不支持更多的内置切片操作，比如切片克隆、元素删除和插入。 我们必须用上面提到的各种内置操作来实现这些操作。

		在下面当前大节中的例子中，假设s是被谈到的切片、T是它的元素类型、t0是类型T的零值字面量。
*/
// 切片克隆
func Test24(t *testing.T) {
	s := []int{1, 2, 3, 4, 5, 6}
	s1 := []int(nil)
	s2 := []int{}
	fmt.Printf("%p\n", s)
	fmt.Printf("%p\n", s1)
	fmt.Printf("%p\n", s2)

	// 对于目前的标准编译器（1.17版本），最简单的克隆一个切片的方法为：
	sClone := append(s[:0:0], s...)
	_ = sClone

	sClone1 := append(s1[:0:0], s1...)
	fmt.Println(sClone1 == nil) //true

	sClone2 := append(s2[:0:0], s2...)
	fmt.Println(sClone2 == nil) //false
	fmt.Printf("%p\n", sClone2)

	// 我们也可以使用下面这种实现。但是和上面这个实现相比，它有一个不完美之处：
	//	如果源切片s是一个空切片（但是非nil），则结果切片是一个nil切片。
	_ = append([]int(nil), s...)

	sClone3 := append([]int(nil), s2...)
	fmt.Println(sClone3 == nil) //true

	// 上面这两种append实现都有一个缺点：它们开辟的内存块常常会比需要的略大一些从而可能造成一点小小的不必要的性能损失。
	// 我们可以使用这两种方法来避免这个缺点：
	// make+copy实现：
	sClone4 := make([]int, len(s))
	copy(sClone4, s)
	fmt.Println(sClone4)

	sClone5 := make([]int, len(s1)) //空切片
	copy(sClone5, s1)               //源:nil切片,目标:空切片
	fmt.Println(sClone5 == nil)     //false

	sClone6 := make([]int, len(s2)) //空切片
	copy(sClone6, s2)               //源:空切片,目标:空切片
	fmt.Println(sClone6 == nil)     //false

	// 或者下面的make+append实现。
	// 对于目前的官方Go工具链v1.17来说，这种实现比上面的make+copy实现略慢一点。
	_ = append(make([]int, 0, len(s)), s...)
	b := append(make([]int, 0, len(s1)), s1...) //空切片
	c := append(make([]int, 0, len(s2)), s2...) //空切片
	fmt.Println(b == nil)                       //false
	fmt.Println(c == nil)                       //false

	// 上面这两种make方法都有一个缺点：如果s是一个nil切片，则使用此方法将得到一个非nil切片。
	// 不过，在编程实践中，我们常常并不需要追求克隆的完美性。如果我们确实需要，则需要多写几行：
	if s1 != nil {
		sClone = make([]int, len(s1))
		copy(sClone, s1)
	}
}

/*
	限定标识符:
		限定标识符为使用包名前缀限定的标识符。包名与标识符均不能为空白的。
			一个限定标识符代表了对另一个代码包中的某个标识符的访问.
			  包必须被导入。 标识符必须是已导出的.
				math.Sin	// 表示math包中的Sin函数
*/

/*
	在Go官方工具链1.15版本之前，对于一些常见的使用场景，使用append来克隆切片比使用make加copy要高效得多。
	 但是从1.15版本开始，官方标准编译器对make+copy这种方法做了特殊的优化，从而使得此方法总是比使用append来克隆切片高效。
		但是请注意：此优化只在被克隆的切片呈现为一个标识符（包括限定标识符）并且make调用只有两个实参的时候才有效。
		 比如，在下面的代码中，它只在第一种情况中才有效：
		 // 情况一： make只有2个实参,copy调用中s为标识符
		var s = make([]byte, 10000)
		y = make([]T, len(s)) // works
		copy(y, s)

		// 情况二： make调用3个实参
		var s = make([]byte, 10000)
		y = make([]T, len(s), len(s)) // not work
		copy(y, s)

		// 情况三： copy调用中a[0]不是一个标识符（包括限定标识符）
		var a = [1][]byte{s}
		y = make([]T, len(a[0])) // not work
		copy(y, a[0])

		// 情况四：
		type T struct {x []byte}
		var t = T{x: s}
		y = make([]T, len(t.x)) // not work
		copy(y, t.x)
*/

/*
	删除一段切片元素
		切片的元素在内存中是连续存储的，相邻元素之间是没有间隙的。所以，当切片的一个元素段被删除时，
			如果剩余元素的次序必须保持原样，则被删除的元素段后面的每个元素都得前移。
			如果剩余元素的次序不需要保持原样，则我们可以将尾部的一些元素移到被删除的元素的位置上。
*/

//删除切片[from,to)的元素
func DeleteSliceFromTo(s []int, from, to int) []int {
	s = append(s[:from], s[to:]...)

	// _ = s[:2+copy(s[2:], s[6:])]

	// t0是类型T的零值字面量
	var t0 int
	// 如果切片的元素可能引用着其它值，则我们应该重置因为删除元素而多出来的元素槽位上的元素值，以避免暂时性的内存泄露
	// "len(s)+to-from"是删除操作之前切片s的长度。
	temp := s[len(s) : len(s)+to-from]
	for i := range temp {
		temp[i] = t0
	}
	return s
}
func Test25(t *testing.T) {
	target := []int{1, 2, 3, 4, 5, 6, 7}
	target = DeleteSliceFromTo(target, 2, 6)
	fmt.Println(target)
	// fmt.Println(target[:cap(target)])
}

/*
	删除一个元素
*/
func Test26(t *testing.T) {
	target := []int{1, 2, 3, 4, 5, 6, 7}
	//删除一个指定index的元素
	i := 2
	_ = append(target[:i], target[i+1:]...)
	_ = target[:i+copy(target[i:], target[i+1:])]

	// 第三种方法（不保持剩余元素的次序）：
	// s[i] = s[len(s)-1]
	// s = s[:len(s)-1]

	// 如果切片的元素可能引用着其它值，则我们应该重置刚多出来的元素槽位上的元素值，以避免暂时性的内存泄露：
	// s[len(s):len(s)+1][0] = t0
	// // 或者
	// s[:len(s)+1][len(s)] = t0
}

/*
	条件性地删除切片元素

	// 假设T是一个小尺寸类型。
	func DeleteElements(s []T, keep func(T) bool, clear bool) []T {
		// result := make([]T, 0, len(s))
		result := s[:0] // 无须开辟内存
		for _, v := range s {
			if keep(v) {
				result = append(result, v)
			}
		}
		if clear { // 避免暂时性的内存泄露。
			temp := s[len(result):]
			for i := range temp {
				temp[i] = t0 // t0是类型T的零值
			}
		}
		return result
	}

	// 注意：如果T是一个大尺寸类型，请慎用T做为参数类型和使用双循环变量for-range代码块遍历元素类型为T的切片。
*/

/*
	将一个切片中的所有元素插入到另一个切片中
*/
func Test27(t *testing.T) {
	s := make([]int, 7, 20)
	for i := range s {
		s[i] = i
	}
	s1 := []int{10, 11, 12, 13}
	i := 2

	{
		//下面两种方式可能最多导致两次内存开辟（最少一次）。

		//不会修改s
		s2 := append(append(s[:i:i], s1...), s[i:]...)
		fmt.Println(s2)

		//修改了s
		// s3 := append(s[:i], append(s1, s[i:]...)...)
		// fmt.Println(s3)

		// Push（插入到结尾）。
		s4 := append(s, s1...)
		// Unshift（插入到开头）。
		s5 := append(s1, s...)

		_, _ = s4, s5
	}

	{
		//这种方式最多只会导致一次内存开辟（最少零次）。
		//此繁琐实现中的make调用将会把一些刚开辟出来的元素清零。这其实是没有必要的。
		//所以此繁琐实现并非总是比上面的单行实现效率更高。事实上，它仅在处理小切片时更高效。
		totalLen := len(s) + len(s1)
		if cap(s) < totalLen {
			x := make([]int, 0, totalLen)
			x = append(x, s[:i]...)
			x = append(x, s1...)
			x = append(x, s[i:]...)
			s = x
		} else {
			// s = append(s[:i+len(s1)], s[i:]...)
			// s = append(s[:i], s1...)
			// s = s[:totalLen]
			s = s[:totalLen]
			copy(s[i+len(s1):], s[i:])
			copy(s[i:], s1)
		}
		fmt.Println(s)
	}
}

/*
	插入若干独立的元素
		插入若干独立的元素和插入一个切片中的所有元素类似。
*/

/*
	特殊的插入和删除：前推/后推，前弹出/后弹出
*/
//
func Push(target []int, eles ...int) []int {
	return append(target, eles...)
}
func Pop(target []int) ([]int, int) {
	temp := len(target) - 1
	return target[:temp], target[temp]
}
func Shift(target []int) ([]int, int) {
	return target[1:], target[0]
}

func Unshift(target []int, eles ...int) []int {
	if cap(target) < len(target)+len(eles) {
		return append(eles, target...)
	}
	target = target[:len(target)+len(eles)]
	copy(target[1+len(eles):], target[1:])
	copy(target[1:], target)
	return target
}
func Test28(t *testing.T) {
	a := []int{1, 2, 3}
	b := []int{3, 4, 5}
	// a = Push(a, b...)
	// a, ele := Pop(a)
	// a, ele := Shift(a)
	// fmt.Println(a, ele)

	a = Unshift(a, b...)
	fmt.Println(a)
}

/*
	上述各种容器操作内部都未同步
		请注意，上述所有各种容器操作的内部实现都未进行同步。
			如果不使用各种并发同步技术，在没有协程修改一个容器值和它的元素的时候，
				多个协程并发读取此容器值和它的元素是安全的。
					但是并发修改同一个容器值则是不安全的。 不使用并发同步技术而并发修改同一个容器值将会造成数据竞争。
*/
