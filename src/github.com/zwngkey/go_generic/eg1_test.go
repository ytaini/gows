/*
 * @Author: zwngkey
 * @Date: 2022-05-05 20:50:28
 * @LastEditors: zwngkey 18390924907@163.com
 * @LastEditTime: 2022-05-06 01:19:21
 * @Description: go泛型
 */
package gogeneric

import (
	"fmt"
	"testing"
)

/*
	泛型类型(generic type)
		类型形参、类型实参、类型约束和泛型类型
*/
type IntSlice []int
type StringSlice []string
type Float32Slie []float32
type Float64Slice []float64

type Slice[T int | string | float32 | float64] []T

/*
	不同于一般的类型定义，这里类型名称 Slice 后带了中括号，各个部分的含义：
		1.T 是类型形参(Type parameter)，在定义Slice类型的时候 T 代表的具体类型并不确定，类似一个占位符
		2. int|float32|float64 这部分被称为类型约束(Type constraint)，中间的 | 的意思是告诉编译器，
			类型形参 T 只可以接收 int 或 float32 或 float64 这三种类型的实参
		3.中括号里的 T int|float32|float64 这一整串因为定义了所有的类型形参(在这个例子里只有一个类型形参T），
			所以我们称其为 类型形参列表(type parameter list)
		这里新定义的类型名称叫 Slice[T]

	类型定义中带 类型形参 的类型，称之为 泛型类型(Generic type)

	泛型类型不能直接拿来使用，必须传入类型实参(Type argument) 将其确定为具体的类型之后才可使用。
		而传入类型实参确定具体类型的操作被称为 类型实例化(Instantiations)
*/

func Test1(t *testing.T) {
	// var a []int = []int{1, 2, 3}

	// 这里传入了类型实参int，泛型类型Slice[T]被实例化为具体的类型 Slice[int]
	var a Slice[int] = []int{1, 2, 3}
	fmt.Printf("%T\n", a)

	var b Slice[string] = []string{"你", "好", "哇"}
	fmt.Printf("%#v\n", b)

	var c Slice[float32] = []float32{1, 2, 3}
	_ = c

	var d Slice[float64] = []float64{1, 2, 3}
	_ = d
}

/*
	所有类型定义都可使用类型形参，所以下面这种结构体以及接口的定义也可以使用类型形参：
		方法不支持泛型
		匿名结构体不支持泛型
		匿名函数不支持泛型
*/
type MyInterface[K any, T any] interface {
	Print(K, T) K
}
type MyStruct[K any, T any] struct {
	Anyth T
	name  string
	age   int
}

func (m MyStruct[K, T]) Print(p K, data T) K {
	return p
}

type MyChan[T any] chan T

func TestEg2(t *testing.T) {
	var ms = MyStruct[int, int]{
		Anyth: 15,
		name:  "张三",
		age:   20,
	}
	fmt.Printf("%#v\n", ms)

	var mi MyInterface[int, int] = ms
	r := mi.Print(123, 123)
	fmt.Println(r)

	// mc := make(MyChan[int])

	// mc <- 1
	// <-mc

}

// 一个可以容纳所有int,uint以及浮点类型的泛型切片
// type MySlice[T int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64] []T

/*
	Go支持将类型约束单独拿出来定义到接口中，从而让代码更容易维护：

		下面的代码把类型约束给单独拿出来，写入了接口类型 NumberType 当中。需要指定类型约束的时候直接使用接口 NumberType 即可。
*/

type NumberType interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint |
		~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64
}
type MySlice[T NumberType] []T

/*
	不过这样的代码依旧不好维护，而接口和接口、接口和普通类型之间也是可以通过 | 进行组合：
	同时，在接口里也能直接组合其他接口.
*/
type Int interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}
type Uint interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}
type Float interface {
	~float32 | ~float64
}
type SliceEleType interface {
	Int | Uint | Float | ~string
}

func MyFunc[T SliceEleType](a T, b T) T {
	return a + b
}

func TestEg3(t *testing.T) {
	var a MySlice[int] = []int{}
	_ = a

	c := MyFunc(1, 2)
	b := MyFunc("hello", "world")

	fmt.Println(c)
	fmt.Println(b)

}

/*
	~ : 指定底层类型

	现在来看一个情况:
		var s1 Slice[int] // 正确
		type MyInt int
		var s2 Slice[MyInt] // ✗ 错误。MyInt类型底层类型是int但并不是int类型，不符合 Slice[T] 的类型约束

	这里发生错误的原因是，泛型类型 Slice[T] 允许的是 int 作为类型实参，
		而不是 MyInt （虽然 MyInt 类型底层类型是 int ，但它依旧不是 int 类型）

	为了从根本上解决这个问题，Go新增了一个符号 ~ ，在类型约束中使用类似 ~int 这种写法的话，就代表着不光是 int ，
		所有以 int 为底层类型的类型也都可用于实例化。

	限制：使用 ~ 时有一定的限制：
		~后面的类型不能为接口
		~后面的类型必须为基本类型
*/

/*
	上面的例子中，我们学习到了一种接口的全新写法，而这种写法在Go1.18之前是不存在的。如果你比较敏锐的话，
		一定会隐约认识到这种写法的改变这也一定意味着Go语言中 接口(interface) 这个概念发生了非常大的变化。

	在Go1.18之前，Go官方对 接口(interface) 的定义是：接口是一个方法集(method set)

	就如下面这个代码一样， ReadWriter 接口定义了一个接口(方法集)，这个集合中包含了 Read() 和 Write() 这两个方法。
		所有同时定义了这两种方法的类型被视为实现了这一接口。

	type ReadWriter interface {
		Read(p []byte) (n int, err error)
		Write(p []byte) (n int, err error)
	}

	我们如果换一个角度来重新思考上面这个接口的话，会发现接口的定义实际上还能这样理解：
		可以我们把 ReaderWriter 接口看成代表了一个 类型的集合，
			所有实现了 Read() Writer() 这两个方法的类型都在接口代表的类型集合当中

	通过换个角度看待接口，在我们眼中接口的定义就从 方法集(method set) 变为了 类型集(type set)。
		而Go1.18开始就是依据这一点将接口的定义正式更改为了 类型集(Type set)


	type Float interface {
		~float32 | ~float64
	}

	type Slice[T Float] []T

	用 类型集 的概念重新理解上面的代码的话就是：
		接口类型 Float 代表了一个 类型集合， 所有以 float32 或 float64 为底层类型的类型，都在这一类型集之中

	而 type Slice[T Float] []T 中， 类型约束 的真正意思是：
		类型约束 指定了类型形参可接受的类型集合，只有属于这个集合中的类型才能替换形参用于实例化


	既然接口定义发生了变化，那么从Go1.18开始 接口实现(implement) 的定义自然也发生了变化：
		当满足以下条件时，我们可以说 类型 T 实现了接口 I ( type T implements interface I)：
			T 不是接口时：类型 T 是接口 I 代表的类型集中的一个成员 (T is an element of the type set of I)
			T 是接口时： T 接口代表的类型集是 I 代表的类型集的子集(Type set of T is a subset of the type set of I)


	类型的并集
		之前一直使用的 | 符号就是求类型的并集( union )


	类型的交集
		type A interface {
			NumberType
			Uint
		}
		接口 A 代表的是 NumberType 与 Uint 的 交集，即 ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64

	除了上面的交集，下面也是一种交集：
		type C interface {
			~int
			int
		}
	很显然，~int 和 int 的交集只有int一种类型，所以接口C代表的类型集中只有int一种类型


	空集:
		 Bad 这个接口代表的类型集为一个空集：
			type Bad interface {
				int
				float32
			} // 类型 int 和 float32 没有相交的类型，所以接口 Bad 代表的类型集为空
		没有任何一种类型属于空集。虽然 Bad 这样的写法是可以编译的，但实际上并没有什么意义


	 空接口和 any
		因为，Go1.18开始接口的定义发生了改变，所以 interface{} 的定义也发生了一些变更：
			空接口代表了所有类型的集合

		所以，对于Go1.18之后的空接口应该这样理解：
			1.虽然空接口内没有写入任何的类型，但它代表的是所有类型的集合，而非一个 空集
			2.类型约束中指定 空接口 的意思是指定了一个包含所有类型的类型集，并不是类型约束限定了只能使用 空接口 来做类型形参

	使用下面这行命令直接把整个项目中的空接口全都替换成 any。
		gofmt -w -r 'interface{} -> any' ./...

*/
/*
	comparable(可比较) 和 可排序(ordered)
        对于一些数据类型，我们需要在类型约束中限制只接受能 != 和 == 对比的类型，如map：
            // 错误。因为 map 中键的类型必须是可进行 != 和 == 比较的类型
            type MyMap[KEY any, VALUE any] map[KEY]VALUE

        所以Go直接内置了一个叫 comparable 的接口，它代表了所有可用 != 以及 == 对比的类型：
            type MyMap[KEY comparable, VALUE any] map[KEY]VALUE // 正确

        comparable 比较容易引起误解的一点是很多人容易把他与可排序搞混淆。
            可比较指的是 可以执行 != == 操作的类型，并没确保这个类型可以执行大小比较（ >,<,<=,>= ）

        而可进行大小比较的类型被称为 Orderd 。目前Go语言并没有像 comparable 这样直接内置对应的关键词，
            所以想要的话需要自己来定义相关接口，比如我们可以参考Go官方包golang.org/x/exp/constraints 如何定义.
            但是这个包属于实验性质的 x 包，今后可能会发生非常大变动，所以并不推荐直接使用
*/

/*
	接口两种类型
		这个例子是阐述接口是类型集最好的例子：
			type ReadWriter interface {
				~string | ~[]rune

				Read([]byte) (int, error)
				Write([]byte) (int, error)
			}

		接口类型 ReadWriter 代表了一个类型集合，所有以 string 或 []rune 为底层类型，
			并且 实现了 Read() Write() 这两个方法的类型都在 ReadWriter 代表的类型集当中
*/

type ReadWriter interface {
	~string | ~[]rune

	Read([]byte) (int, error)
	Write([]byte) (int, error)
}

type MyString string

func (m MyString) Read(p []byte) (n int, err error) {
	return 1, nil
}

func (m MyString) Write(p []byte) (n int, err error) {
	return 1, nil
}

func F1[T ReadWriter](a T) T {
	return a
}

func TestEg4(t *testing.T) {
	// s := "string"
	s1 := MyString("string")
	// F1(s) //error
	F1(s1) //ok
}

/*
	为了解决这个问题也为了保持Go语言的兼容性，Go1.18开始将接口分为了两种类型
		基本接口(Basic interface)
		一般接口(General interface)

	基本接口(Basic interface)
		接口定义中如果只有方法的话，那么这种接口被称为基本接口(Basic interface)。

		这种接口就是Go1.18之前的接口，用法也基本和Go1.18之前保持一致。基本接口大致可以用于如下几个地方：
			1.最常用的，定义接口变量并赋值
			2.基本接口因为也代表了一个类型集，所以也可用在类型约束中
				// io.Reader 和 io.Writer 都是基本接口，也可以用在类型约束中
				type MySlice[T io.Reader | io.Writer] []Slice


	一般接口(General interface)
		如果接口内不光只有方法，还有类型的话，这种接口被称为 一般接口(General interface)

		一般接口类型不能用来定义变量，只能用于泛型的类型约束中。
*/
/*
	泛型接口

*/
type DataProcessor[T any] interface {
	Process(oriData T) (newData T)
	Save(data T) error
}
type CSVProcessor struct {
}

func (c CSVProcessor) Process(oriData string) (newData string) {
	return "string"
}

func (c CSVProcessor) Save(oriData string) error {
	return nil
}

type DataProcessor2[T any] interface {
	int | ~struct{ Data any }

	Process(data T) (newData T)
	Save(data T) error
}

// JsonProcessor 实现了接口 DataProcessor2[string] 的两个方法，
//	同时底层类型是 struct{ Data interface{} }。所以实现了接口 DataProcessor2[string]
type JsonProcessor struct {
	Data any
}

func (c JsonProcessor) Process(oriData string) (newData string) {
	return "STRING"
}

func (c JsonProcessor) Save(oriData string) error {
	return nil
}

/*
	DataProcessor2[string] 因为带有类型并集所以它是 一般接口(General interface)，所以实例化之后的这个接口代表的意思是：
		1.只有实现了 Process(string) string 和 Save(string) error 这两个方法，并且以 int 或 struct{ Data interface{} } 为底层类型的类型才算实现了这个接口
		2.一般接口(General interface) 不能用于变量定义只能用于类型约束，所以接口 DataProcessor2[string] 只是定义了一个用于类型约束的类型集

*/
func F2[T DataProcessor2[string]](a T) {
	fmt.Println("F2")
}

func TestEg5(t *testing.T) {
	var a DataProcessor[string] = CSVProcessor{}
	_ = a

	b := JsonProcessor{Data: "string"}
	F2(b)
}

/*
	接口定义的种种限制规则
		Go1.18从开始，在定义类型集(接口)的时候增加了非常多十分琐碎的限制规则，其中很多规则都在之前的内容中介绍过了，
			但剩下还有一些规则因为找不到好的地方介绍，在这里统一介绍下：
*/
/*
	用 | 连接多个类型的时候，类型之间不能有相交的部分(即必须是不交集):

	type MyInt int

	// 错误，MyInt的底层类型是int,和 ~int 有相交的部分
	type _ interface {
		~int | MyInt
	}

	但是相交的类型中是接口的话，则不受这一限制：
	type _ interface {
		~int | interface{ MyInt }  // 正确
	}

	type _ interface {
		interface{ ~int } | MyInt // 也正确
	}

	type _ interface {
		interface{ ~int } | interface{ MyInt }  // 也正确
	}

*/
/*
	类型的并集中不能有类型形参

	type MyInf[T ~int | ~string] interface {
		~float32 | T  // 错误。T是类型形参
	}

	type MyInf2[T ~int | ~string] interface {
		T  // 错误
	}

*/
/*
	接口不能直接或间接地并入自己
		type Bad interface {
			Bad // 错误，接口不能直接并入自己
		}

		type Bad2 interface {
			Bad1
		}
		type Bad1 interface {
			Bad2 // 错误，接口Bad1通过Bad2间接并入了自己
		}

		type Bad3 interface {
			~int | ~string | Bad3 // 错误，通过类型的并集并入了自己
		}
*/
/*
	接口的并集成员个数大于一的时候不能直接或间接并入 comparable 接口
		type OK interface {
    		comparable // 正确。只有一个类型的时候可以使用 comparable
		}

		type Bad1 interface {
			[]int | comparable // 错误，类型并集不能直接并入 comparable 接口
		}

		type CmpInf interface {
			comparable
		}
		type Bad2 interface {
			chan int | CmpInf  // 错误，类型并集通过 CmpInf 间接并入了comparable
		}
*/
/*
	带方法的接口(无论是基本接口还是一般接口)，都不能写入接口的并集中：
		type _ interface {
			~int | ~string | error // 错误，error是带方法的接口(一般接口) 不能写入并集中
		}

		type DataProcessor[T any] interface {
			~string | ~[]byte

			Process(data T) (newData T)
			Save(data T) error
		}

		// 错误，实例化之后的 DataProcessor[string] 是带方法的一般接口，不能写入类型并集
		type _ interface {
			~int | ~string | DataProcessor[string]
		}

		type Bad[T any] interface {
			~int | ~string | DataProcessor[T]  // 也不行
		}
*/
