/*
 * @Author: zwngkey
 * @Date: 2022-05-04 18:06:06
 * @LastEditors: zwngkey 18390924907@163.com
 * @LastEditTime: 2022-05-05 18:49:39
 * @Description: go 字符串
 */
package gostring

import (
	"bytes"
	"fmt"
	"testing"
	"time"
	"unicode/utf8"
)

/*
	字符串类型的内部结构定义
		对于标准编译器，字符串类型的内部结构声明如下：
			type _string struct {
				elements *byte // 引用着底层的字节
				len      int   // 字符串中的字节数
			}
		从这个声明来看，我们可以将一个字符串的内部定义看作为一个字节序列。
			事实上，我们确实可以把一个字符串看作是一个元素类型为byte的（且元素不可修改的）切片。


	关于字符串的一些简单事实
		1.字符串值（和布尔以及各种数值类型的值）可以被用做常量。
		2.Go支持两种风格的字符串字面量表示形式：双引号风格（解释型字面表示）和反引号风格（直白字面表示）
		3.字符串类型的零值为空字符串。一个空字符串在字面上可以用""或者``来表示。
		4.我们可以用运算符+和+=来衔接字符串。
		5.字符串类型都是可比较类型。同一个字符串类型的值可以用==和!=比较运算符来比较。
			 并且和整数/浮点数一样，同一个字符串类型的值也可以用>、<、>=和<=比较运算符来比较。
			 	当比较两个字符串值的时候，它们的底层字节将逐一进行比较。
*/
func TestEg1(t *testing.T) {
	const s = "string"
	var s1 = "string中"
	var s2 = `123`
	var s3 string
	var s4 = s1 + s2
	fmt.Println(s == s1)
	fmt.Println(s1 == s2)
	fmt.Println(s3 > s4)
	fmt.Println([]byte(s1))
	fmt.Println([]byte(s2))
	fmt.Println([]rune(s1))
	fmt.Println([]rune(s2))
}

/*
	更多关于字符串类型和值的事实：
		1.和Java语言一样，字符串值的内容（即底层字节）是不可更改的。 字符串值的长度也是不可独立被更改的。
			一个可寻址的字符串只能通过将另一个字符串赋值给它来整体修改它。
		2.对于同一字符串字面量,不同的字符串变量指向相同的底层数组,这是因为字符串是只读的,为了节省内容.
		3.字符串类型没有内置的方法。我们可以
			使用strings标准库提供的函数来进行各种字符串操作。
			调用内置函数len来获取一个字符串值的长度（此字符串中存储的字节数）。
			使用容器元素索引语法aString[i]来获取aString中的第i个字节。
				 表达式aString[i]是不可寻址的。换句话说，aString[i]不可被修改。
			使用子切片语法aString[start:end]来获取aString的一个子字符串。 这里，start和end均为aString中存储的字节的下标。
		4.对于标准编译器来说，一个字符串的赋值完成之后，此赋值中的目标值和源值将共享底层字节。
			 一个子切片表达式aString[start:end]的估值结果也将和基础字符串aString共享一部分底层字节。
*/
func TestEg2(t *testing.T) {
	var helloWorld = "hello world!"

	var hello = helloWorld[:5] // 取子字符串
	// 104是英文字符h的ASCII（和Unicode）码。
	fmt.Println(hello[0])         // 104
	fmt.Printf("%T \n", hello[0]) // uint8

	// hello[0]是不可寻址和不可修改的，所以下面
	// 两行编译不通过。
	/*
		hello[0] = 'H'         // error
		fmt.Println(&hello[0]) // error
	*/
}

/*
	注意：如果在aString[i]和aString[start:end]中，aString和各个下标均为常量，
			则编译器将在编译时刻验证这些下标的合法性，但是这样的元素访问和子切片表达式的估值结果总是非常量
*/
func TestEg3(t *testing.T) {
	const s = "string"
	const i = 2
	// const j = 7
	// const a = s[i] //error
	// const a = s[i:j]//error
	// var a = s[j] //error
	// var a = s[i:j] error
	const k = 5
	_ = s[k]   //ok
	_ = s[i:k] //ok
}

/*
	len()函数注意点:如果表达式s表示一个字符串常量，则表达式len(s)将在编译时刻估值；
*/
const s = "Go101.org" // len(s) == 9

// len(s)是一个常量表达式，但len(s[:])却不是。
var a byte = 1 << len(s) / 128
var b byte = 1 << len(s[:]) / 128

func TestEg4(t *testing.T) {
	fmt.Println(a, b) //4 0
}

/*
	字符串编码和Unicode码点
		Unicode标准为全球各种人类语言中的每个字符制定了一个独一无二的值。 但
			Unicode标准中的基本单位不是字符，而是码点（code point）。
				大多数的码点实际上就对应着一个字符。 但也有少数一些字符是由多个码点组成的。

		码点值在Go中用rune值来表示。 内置rune类型为内置int32类型的一个别名。

		在具体应用中，码点值的编码方式有很多，比如UTF-8编码和UTF-16编码等。
			目前最流行编码方式为UTF-8编码。在Go中，所有的字符串常量都被视为是UTF-8编码的。
			 在编译时刻，非法UTF-8编码的字符串常量将导致编译失败。
			  在运行时刻，Go运行时无法阻止一个字符串是非法UTF-8编码的。

		在UTF-8编码中，一个码点值可能由1到4个字节组成。
		 比如，每个英语码点值（均对应一个英语字符）均由一个字节组成，而每个中文码点值（均对应一个中文字符）均由三个字节组成。
*/

/*
	字符串相关的类型转换
		1.整数可以被显式转换为字符串类型（但是反之不行）。
		2.一个字符串值可以被显式转换为一个字节切片（byte slice），反之亦然。
		   一个字节切片类型是一个元素类型为内置类型byte的切片类型。 或者说，一个字节切片类型的底层类型为[]byte（亦即[]uint8）。
		3.一个字符串值可以被显式转换为一个码点切片（rune slice），反之亦然。
		   一个码点切片类型是一个元素类型为内置类型rune的切片类型。 或者说，一个码点切片类型的底层类型为[]rune（亦即[]int32）。

		在一个从码点切片到字符串的转换中，码点切片中的每个码点值将被UTF-8编码为一到四个字节至结果字符串中。
			 如果一个码点值是一个不合法的Unicode码点值，则它将被视为Unicode替换字符（码点）值0xFFFD（Unicode replacement character）。
			  替换字符值0xFFFD将被UTF-8编码为三个字节0xef 0xbf 0xbd

		当一个字符串被转换为一个码点切片时，此字符串中存储的字节序列将被解读为一个一个码点的UTF-8编码序列。
		 非法的UTF-8编码字节序列将被转化为Unicode替换字符值0xFFFD。

		当一个字符串被转换为一个字节切片时，结果切片中的底层字节序列是此字符串中存储的字节序列的一份深复制。
		  即Go运行时将为结果切片开辟一块足够大的内存来容纳被复制过来的所有字节。当此字符串的长度较长时，此转换开销是比较大的。
		    同样，当一个字节切片被转换为一个字符串时，此字节切片中的字节序列也将被深复制到结果字符串中。当此字节切片的长度较长时，此转换开销同样是比较大的。
			 在这两种转换中，必须使用深复制的原因是字节切片中的字节元素是可修改的，但是字符串中的字节是不可修改的，
			  所以一个字节切片和一个字符串是不能共享底层字节序列的。

		请注意，在字符串和字节切片之间的转换中，
			1.非法的UTF-8编码字节序列将被保持原样不变。
			2.标准编译器做了一些优化，从而使得这些转换在某些情形下将不用深复制.

		Go并不支持字节切片和码点切片之间的直接转换。我们可以用下面列出的方法来实现这样的转换：
			利用字符串做为中间过渡。这种方法相对方便但效率较低，因为需要做两次深复制。
			使用unicode/utf8标准库包中的函数来实现这些转换。 这种方法效率较高，但使用起来不太方便。
			使用bytes标准库包中的Runes函数来将一个字节切片转换为码点切片。 但此包中没有将码点切片转换为字节切片的函数。
*/
func Runes2Bytes(r []rune) []byte {
	sizes := 0
	for _, v := range r {
		sizes += utf8.RuneLen(v)
	}
	b := make([]byte, sizes)
	size := 0
	for _, v := range r {
		size += utf8.EncodeRune(b[size:], v)
	}
	return b
}
func TestEg5(t *testing.T) {
	s := "你好哇,李芳芳"
	// string <==>[]byte
	bs := []byte(s)
	s = string(bs)

	// string <==> []rune
	rs := []rune(s)
	s = string(rs)

	//[]byte <==> []rune
	rs = bytes.Runes(bs)
	bs = Runes2Bytes(rs)

	_, _ = s, bs

}

/*
	字符串和字节切片之间的转换的编译器优化
		字符串和字节切片之间的转换将深复制它们的底层字节序列。 标准编译器做了一些优化，从而在某些情形下避免了深复制。
		 至少这些优化在当前（Go官方工具链1.17）是存在的。 这样的情形包括：
			1.一个for-range循环中跟随range关键字后面的是从字符串到字节切片的转换；
			2.一个在映射元素读取索引语法中被用做键值的从字节切片到字符串的转换（注意：对修改写入索引语法无效）；
			3.一个字符串比较表达式中被用做比较值的从字节切片到字符串的转换；
			4.一个（至少有一个被衔接的字符串值为非空字符串常量的）字符串衔接表达式中的从字节切片到字符串的转换。
*/
func TestEg6(t *testing.T) {
	var str = "world"
	// 这里，转换[]byte(str)将不需要一个深复制。
	for i, b := range []byte(str) {
		fmt.Println(i, ":", b)
	}

	key := []byte{'k', 'e', 'y'}
	m := map[string]string{}
	// 这个string(key)转换仍然需要深复制。
	m[string(key)] = "value"
	// 这里的转换string(key)将不需要一个深复制。
	// 即使key是一个包级变量，此优化仍然有效。
	fmt.Println(m[string(key)]) // value

	// 下面的四个转换都不需要深复制。
	{
		var s string
		var x = []byte{1023: 'x'}
		var y = []byte{1023: 'y'}
		// 下面的四个转换都不需要深复制。
		if string(x) != string(y) {
			s = (" " + string(x) + string(y))[1:]
		}

		// 两个在比较表达式中的转换不需要深复制，
		// 但两个字符串衔接中的转换仍需要深复制。
		// 请注意此字符串衔接和fc中的衔接的差别。
		if string(x) != string(y) {
			s = string(x) + string(y)
		}
		_ = s
	}
}

/*
	使用for-range循环遍历字符串中的码点
		for-range循环控制中的range关键字后可以跟随一个字符串，用来遍历此字符串中的码点（而非字节元素）。
			 字符串中非法的UTF-8编码字节序列将被解读为Unicode替换码点值0xFFFD。

		从下面程序输出结果可以看出：
			1.下标循环变量的值并非连续。原因是下标循环变量为字符串中字节的下标，而一个码点可能需要多个字节进行UTF-8编码。
			2.第一个字符é由两个码点（共三字节）组成，其中一个码点需要两个字节进行UTF-8编码。
			3.第二个字符क्षि由四个码点（共12字节）组成，每个码点需要三个字节进行UTF-8编码。
			4.英语字符a由一个码点组成，此码点只需一个字节进行UTF-8编码。
			5.字符π由一个码点组成，此码点只需两个字节进行UTF-8编码。
			6.汉字囧由一个码点组成，此码点只需三个字节进行UTF-8编码。


		从这几个例子可以看出，len(s)将返回字符串s中的字节数。 len(s)的时间复杂度为O(1)。

		如何得到一个字符串中的码点数呢？
			1.使用刚介绍的for-range循环来统计一个字符串中的码点数是一种方法，
			2.使用unicode/utf8标准库包中的RuneCountInString是另一种方法。 这两种方法的效率基本一致。
			3.第三种方法为使用len([]rune(s）)来获取字符串s中码点数。
			注意，这三种方法的时间复杂度均为O(n)。

		标准编译器从1.11版本开始，对第三种方法的表达式做了优化以避免一个不必要的深复制，从而使得它的效率和前两种方法一致。

*/
func TestEg7(t *testing.T) {
	s := "éक्षिaπ囧"
	for i, rn := range s {
		fmt.Printf("%2v: 0x%x %v \n", i, rn, string(rn))
	}
	fmt.Println(len(s))
}

/*
	更多字符串衔接方法
		除了使用+运算符来衔接字符串，我们也可以用下面的方法来衔接字符串：
			1.fmt标准库包中的Sprintf/Sprint/Sprintln函数可以用来衔接各种类型的值的字符串表示，当然也包括字符串类型的值。
			2.使用strings标准库包中的Join函数。
			3.bytes标准库包提供的Buffer类型可以用来构建一个字节切片，然后我们可以将此字节切片转换为一个字符串。
			4.从Go 1.10开始，strings标准库包中的Builder类型可以用来拼接字符串。
				和bytes.Buffer类型类似，此类型内部也维护着一个字节切片，但是它在将此字节切片转换为字符串时避免了底层字节的深复制。

		标准编译器对使用+运算符的字符串衔接做了特别的优化。
		 所以，一般说来，在被衔接的字符串的数量是已知的情况下，使用+运算符进行字符串衔接是比较高效的。
*/

/*
	语法糖：将字符串当作字节切片使用
		内置函数copy和append可以用来复制和添加切片元素。
		 事实上，做为一个特例，如果这两个函数的调用中的第一个实参为一个字节切片的话，那么第二个实参可以是一个字符串。
		  （对于append函数调用，字符串实参后必须跟随三个点...。） 换句话说，在此特例中，字符串可以当作字节切片来使用。
*/
func TestEg8(t *testing.T) {
	hello := []byte("Hello ")
	world := "world!"

	// helloWorld := append(hello, []byte(world)...) // 正常的语法
	helloWorld := append(hello, world...) // 语法糖
	fmt.Println(string(helloWorld))

	helloWorld2 := make([]byte, len(hello)+len(world))
	copy(helloWorld2, hello)
	// copy(helloWorld2[len(hello):], []byte(world)) // 正常的语法
	copy(helloWorld2[len(hello):], world) // 语法糖
	fmt.Println(string(helloWorld2))

}

/*
	更多关于字符串的比较
		上面已经提到了比较两个字符串事实上逐个比较这两个字符串中的字节。 Go编译器一般会做出如下的优化：
			1.对于==和!=比较，如果这两个字符串的长度不相等，则这两个字符串肯定不相等（无需进行字节比较）。
			2.如果这两个字符串底层引用着字符串切片的指针相等，则比较结果等同于比较这两个字符串的长度。

		所以两个相等的字符串的比较的时间复杂度取决于它们底层引用着字符串切片的指针是否相等。
		 如果相等，则对它们的比较的时间复杂度为O(1)，否则时间复杂度为O(n)。

		上面已经提到了，对于标准编译器，一个字符串赋值完成之后，目标字符串和源字符串将共享同一个底层字节序列。 所以比较这两个字符串的代价很小。
*/
func TestEg9(t *testing.T) {
	bs := make([]byte, 1<<28)
	s0 := string(bs)
	s1 := string(bs)
	s2 := s1

	// s0、s1和s2是三个相等的字符串。

	// s0的底层字节序列是bs的一个深复制。
	// s1的底层字节序列也是bs的一个深复制。
	// s0和s1底层字节序列为两个不同的字节序列。
	// s2和s1共享同一个底层字节序列。

	startTime := time.Now()
	_ = s0 == s1
	duration := time.Since(startTime)
	fmt.Println("duration for (s0 == s1):", duration)

	startTime = time.Now()
	_ = s1 == s2
	duration = time.Since(startTime)
	fmt.Println("duration for (s1 == s2):", duration)

}

/*
	所以应尽量避免比较两个很长的不共享底层字节序列的相等的（或者几乎相等的）字符串。
*/
