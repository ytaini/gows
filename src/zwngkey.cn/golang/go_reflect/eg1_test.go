/*
 * @Author: zwngkey
 * @Date: 2022-05-04 14:56:35
 * @LastEditors: zwngkey 18390924907@163.com
 * @LastEditTime: 2022-05-11 05:11:25
 * @Description:
 */
package goreflect

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

/*
	我们可以通过reflect库包中Type和Value两个类型提供的功能来观察不同的Go值。

	Go反射机制设计的目标之一是任何非反射操作都可以通过反射机制来完成。 由于各种各样的原因，此目标并没有得到100%的实现。
	  但是，目前大部分的非反射操作都可以通过反射机制来完成。 另一方面，通过反射，我们也可以完成一些使用非反射操作不可能完成的操作。


	reflect.Type类型和值
		通过调用reflect.TypeOf函数，我们可以从一个任何非接口类型的值创建一个reflect.Type值。
		  此reflect.Type值表示着此非接口值的类型。通过此值，我们可以得到很多此非接口类型的信息。
		    同时，我们也可以将一个接口值传递给一个reflect.TypeOf函数调用，
			  但是此调用将返回一个表示着此接口值的动态类型的reflect.Type值。
			    实际上，reflect.TypeOf函数的唯一参数的类型为interface{}，
				  reflect.TypeOf函数将总是返回一个表示着此唯一接口参数值的动态类型的reflect.Type值。

		那如何得到一个表示着某个接口类型的reflect.Type值呢?  需要间接途径来达到这一目的。

		类型reflect.Type为一个接口类型，它指定了若干方法。 通过这些方法，我们能够观察到一个reflect.Type值所表示的Go类型的各种信息。
			 这些方法中的有些适用于所有种类的类型，有些只适用于一种或几种类型。 通过不合适的reflect.Type属主值调用某个方法将在运行时产生一个恐慌。
*/
func TestEg1(t *testing.T) {

	type A = [16]int16
	var c <-chan map[A][]byte
	tc := reflect.TypeOf(c)
	fmt.Println(tc.Kind())
	fmt.Println(tc.ChanDir())

	tm := tc.Elem()
	ta, tb := tm.Key(), tm.Elem()
	fmt.Println(tm.Kind(), ta.Kind(), tb.Kind())

	tx, ty := ta.Elem(), tb.Elem()
	fmt.Println(tx.Kind(), ty.Kind())
	fmt.Println(tx.Bits(), ty.Bits())
	fmt.Println(tx.ConvertibleTo(ty))
	fmt.Println(ta.ConvertibleTo(tb))

	fmt.Println(tc.Comparable())
	fmt.Println(tm.Comparable())
	fmt.Println(ta.Comparable())
	fmt.Println(tb.Comparable())
	fmt.Println(tx.Comparable())
	fmt.Println(ty.Comparable())

}

/*
	我们使用方法Elem来得到某些类型的元素类型。 实际上，此方法也可以用来得到一个指针类型的基类型.

	下面这个例子同时也展示了如何通过间接的途径得到一个表示一个接口类型的reflect.Type值。
*/
type T []interface{ m() }

func (T) m() {}
func TestEg2(t *testing.T) {
	tp := reflect.TypeOf(new(interface{}))
	tt := reflect.TypeOf(T{})
	fmt.Println(tp.Kind(), tt.Kind()) // ptr slice

	// 使用间接的方法得到表示两个接口类型的reflect.Type值。
	ti, tim := tp.Elem(), tt.Elem()
	fmt.Println(ti.Kind(), tim.Kind()) // interface interface

	fmt.Println(tt.Implements(tim))  // true
	fmt.Println(tp.Implements(tim))  // false
	fmt.Println(tim.Implements(tim)) // true

	// 所有的类型都实现了任何空接口类型。
	fmt.Println(tp.Implements(ti))  // true
	fmt.Println(tt.Implements(ti))  // true
	fmt.Println(tim.Implements(ti)) // true
	fmt.Println(ti.Implements(ti))  // true
}

/*
	我们可以通过反射列出一个类型的所有方法和一个结构体类型的所有（导出和非导出）字段的类型。
		 我们也可以通过反射列出一个函数类型的各个输入参数和返回结果类型。

*/
type F func(string, int) bool

func (f F) m(s string) bool {
	return f(s, 32)
}
func (f F) M() {}

type I interface {
	m(s string) bool
	M()
}

func TestEg3(t *testing.T) {
	var x struct {
		F
		i I
	}
	tx := reflect.TypeOf(x)
	fmt.Println(tx.Kind())               //struct
	fmt.Println(tx.NumField())           //2
	fmt.Println(tx.Field(0).Anonymous)   //true
	fmt.Println(tx.Field(0).Index)       //[0]
	fmt.Println(tx.Field(0).Name)        //F
	fmt.Println(tx.Field(0).Offset)      //0
	fmt.Println(tx.Field(0).PkgPath)     //""
	fmt.Println(tx.Field(0).Tag)         //""
	fmt.Println(tx.Field(0).Type.Kind()) //func

	fmt.Println(tx.Field(1).Anonymous) //false
	fmt.Println(tx.Field(1).Index)     //[1]
	fmt.Println(tx.Field(1).Name)      //i
	fmt.Println(tx.Field(1).Offset)    //8
	// 包路径（PkgPath）是非导出字段（或者方法）的内在属性。
	fmt.Println(tx.Field(1).PkgPath)     //"github.com/zwngkey/go_reflect"
	fmt.Println(tx.Field(1).Tag)         //""
	fmt.Println(tx.Field(1).Type.Kind()) //interface

	tf := tx.Field(0).Type
	fmt.Println(tf.Kind())
	//IsVariadic 判断函数最后一个参数是否为可变参数
	fmt.Println(tf.IsVariadic())
	fmt.Println(tf.NumIn())
	fmt.Println(tf.NumOut())
	fmt.Println(tf.In(0).Kind())
	fmt.Println(tf.In(1).Kind())
	fmt.Println(tf.Out(0).Kind())

	ti := tx.Field(1).Type

	fmt.Println(tf.NumMethod())    //1
	fmt.Println(ti.NumMethod())    //2
	fmt.Println(tf.Method(0).Name) //M
	fmt.Println(ti.Method(0).Name) //M
	fmt.Println(ti.Method(1).Name) //m

	_, ok1 := tf.MethodByName("m")
	_, ok2 := ti.MethodByName("m")
	fmt.Println(ok1, ok2) // false true
}

/*
	从上面这个例子我们可以看出：
		对于非接口类型，reflect.Type.NumMethod方法只返回一个类型的所有导出的方法（包括通过内嵌得来的隐式方法）的个数，
			并且 方法reflect.Type.MethodByName不能用来获取一个类型的非导出方法；

		虽然reflect.Type.NumField方法返回一个结构体类型的所有字段（包括非导出字段）的数目，
			但是不推荐使用方法reflect.Type.FieldByName来获取非导出字段
*/
/*
	我们可以通过反射来检视结构体字段的标签信息。 结构体字段标签的类型为reflect.StructTag，
		它的方法Get和Lookup用来检视字段标签中的键值对。 一个例子：
*/
type A struct {
	X    int  `max:"99" min:"0" default:"0"`
	Y, Z bool `optional:"yes"`
}

func TestEg4(te *testing.T) {
	t := reflect.TypeOf(A{})
	x := t.Field(0).Tag
	y := t.Field(1).Tag
	z := t.Field(2).Tag
	fmt.Println(reflect.TypeOf(x)) // reflect.StructTag
	// v的类型为string
	v, present := x.Lookup("max")
	fmt.Println(len(v), present)      // 2 true
	fmt.Println(x.Get("max"))         // 99
	fmt.Println(x.Lookup("optional")) //  false
	fmt.Println(y.Lookup("optional")) // yes true
	fmt.Println(z.Lookup("optional")) // yes true
}

/*
	注意：
		1.键值对中的键不能包含空格（Unicode值为32）、双引号（Unicode值为34）和冒号（Unicode值为58）。
		2.为了形成键值对，所设想的键值对形式中的冒号的后面不能紧跟着空格字符。所以`optional: "yes"`不形成键值对。
		3.键值对中的值中的空格不会被忽略。所以
			`json:"author, omitempty"`,
			`json:" author,omitempty"`以及
			`json:"author,omitempty"`各不相同。
		4.每个字段标签应该呈现为单行才能使它的整个部分都对键值对的形成有贡献。
*/

/*
	reflect代码包也提供了一些其它函数来动态地创建出来一些无名组合类型。
*/
func TestEg5(t *testing.T) {
	ta := reflect.ArrayOf(5, reflect.TypeOf(123))
	fmt.Println(ta) // [5]int
	tc := reflect.ChanOf(reflect.SendDir, ta)
	fmt.Println(tc) // chan<- [5]int
	tp := reflect.PtrTo(ta)
	fmt.Println(tp) // *[5]int
	ts := reflect.SliceOf(tp)
	fmt.Println(ts) // []*[5]int
	tm := reflect.MapOf(ta, tc)
	fmt.Println(tm) // map[[5]int]chan<- [5]int
	tf := reflect.FuncOf([]reflect.Type{ta},
		[]reflect.Type{tp, tc}, false)
	fmt.Println(tf) // func([5]int) (*[5]int, chan<- [5]int)
	tt := reflect.StructOf([]reflect.StructField{
		{Name: "Age", Type: reflect.TypeOf("abc")},
	})
	fmt.Println(tt)            // struct { Age string }
	fmt.Println(tt.NumField()) // 1

}

/*
	注意，到目前为止（Go 1.18），我们无法通过反射动态创建一个接口类型。这是Go反射目前的一个限制。

	另一个限制是使用反射动态创建结构体类型的时候可能会有各种不完美的情况出现。

	第三个限制是我们无法通过反射来声明一个新的类型。
*/

/*
	reflect.Value类型和值
		类似的，我们可以通过调用reflect.ValueOf函数，从一个非接口类型的值创建一个reflect.Value值。
		  此reflect.Value值代表着此非接口值。 和reflect.TypeOf函数类似，
		    reflect.ValueOf函数也只有一个interface{}类型的参数。
			  当我们将一个接口值传递给一个reflect.ValueOf函数调用时，
			    此调用返回的是代表着此接口值的动态值的一个reflect.Value值。
				  我们必须通过间接的途径获得一个代表一个接口值的reflect.Value值。

		被一个reflect.Value值代表着的值常称为此reflect.Value值的底层值（underlying value）。

		reflect.Value类型有很多方法。 我们可以调用这些方法来观察和操纵一个reflect.Value属主值表示的Go值。
		  这些方法中的有些适用于所有种类类型的值，有些只适用于一种或几种类型的值。
		    通过不合适的reflect.Value属主值调用某个方法将在运行时产生一个恐慌。

		一个reflect.Value值的CanSet方法将返回此reflect.Value值代表的Go值是否可以被修改（可以被赋值）。
		  如果一个Go值可以被修改，则我们可以调用对应的reflect.Value值的Set方法来修改此Go值。
			注意：reflect.ValueOf函数直接返回的reflect.Value值都是不可修改的。
*/
func TestEg7(t *testing.T) {
	n := 1
	p := &n
	fmt.Printf("%p\n", p)
	vp := reflect.ValueOf(p)
	fmt.Println(vp.CanSet(), vp.CanAddr())
	fmt.Println(vp)
	vn := vp.Elem() //取得vp的底层指针值引用的值的代表值
	fmt.Println(vn.CanSet(), vn.CanAddr())
	fmt.Println(vn.Addr())
	// vn.SetInt(123)
	vn.Set(reflect.ValueOf(123))
	fmt.Println(n)
}

/*
	一个结构体值的非导出字段不能通过反射来修改。

	下例中同时也展示了如何间接地获取底层值为接口值的reflect.Value值。

	从下两例中，我们可以得知有两种方法获取一个代表着一个指针所引用着的值的reflect.Value值：
		1.通过调用代表着此指针值的reflect.Value值的Elem方法。
		2.将代表着此指针值的reflect.Value值的传递给一个reflect.Indirect函数调用。
		 （如果传递给一个reflect.Indirect函数调用的实参不代表着一个指针值，则此调用返回此实参的一个复制。）

*/
func TestEg8(t *testing.T) {
	var s struct {
		X any
		y any
	}
	vp := reflect.ValueOf(&s)
	fmt.Println(vp.CanSet(), vp.CanAddr()) //false false

	// 如果vp代表着一个指针，下一行等价于"vs := vp.Elem()"。
	vs := reflect.Indirect(vp)
	fmt.Println(vs.CanSet(), vs.CanAddr()) //true true

	// vx和vy都各自代表着一个接口值。
	vx, vy := vs.Field(0), vs.Field(1)

	fmt.Println(vx.CanSet(), vx.CanAddr()) // true true
	// vy is addressable but not modifiable.
	fmt.Println(vy.CanSet(), vy.CanAddr()) // false true

	vb := reflect.ValueOf(123)
	vx.Set(vb) // okay, 因为vx代表的值是可修改的。
	// vy.Set(vb)  // 会造成恐慌，因为vy代表的值是不可修改的。

	fmt.Printf("%#v\n", s)              //{X:123, y:interface {}(nil)}
	fmt.Println(vx.IsNil(), vy.IsNil()) // false true
}

/*
	注意：reflect.Value.Elem方法也可以用来获取一个代表着一个接口值的动态值的reflect.Value值
*/
func TestEg9(t *testing.T) {
	var z = 123
	var y = &z
	var x interface{} = y
	p := &x
	v := reflect.ValueOf(p)
	vx := v.Elem() //得到代表p引用的值的reflect.Value值.发现vx底层值为接口类型
	fmt.Println(vx.CanSet(), vx.CanAddr())

	vy := vx.Elem() //得到vx底层值的动态值的reflect.Value值.发现vy底层值为指针类型
	fmt.Println(vy.CanSet(), vy.CanAddr())

	vz := vy.Elem() //得到vy底层值的基类型的reflect.Value值.
	fmt.Println(vz.CanSet(), vz.CanAddr())

	vz.Set(reflect.ValueOf(789))
	fmt.Println(z) // 789
}

/*
	reflect标准库包中也提供了一些对应着内置函数或者各种非反射功能的函数。
		下面这个例子展示了如何利用这些函数将一个（效率不高的）自定义泛型函数绑定到不同的类型的函数值上。
*/
func InvertSlice(args []reflect.Value) []reflect.Value {
	inSlice, n := args[0], args[0].Len()
	outSlice := reflect.MakeSlice(inSlice.Type(), 0, n)
	for i := n - 1; i >= 0; i-- {
		element := inSlice.Index(i)
		outSlice = reflect.Append(outSlice, element)
	}
	return []reflect.Value{outSlice}
}

func Bind(p interface{}, f func([]reflect.Value) []reflect.Value) {
	// invert代表着一个函数值。
	invert := reflect.ValueOf(p).Elem()
	invert.Set(reflect.MakeFunc(invert.Type(), f))
}
func TestEg10(t *testing.T) {
	var invertInts func([]int) []int
	Bind(&invertInts, InvertSlice)
	fmt.Println(invertInts([]int{2, 3, 5})) // [5 3 2]

	// var invertStrs func([]string) []string
	// Bind(&invertStrs, InvertSlice)
	// fmt.Println(invertStrs([]string{"Go", "C"})) // [C Go]
}

/*
	如果一个reflect.Value值的底层值为一个函数值，则我们可以调用此reflect.Value值的Call方法来调用此函数。
		 每个Call方法调用接受一个[]reflect.Value类型的参数（表示传递给相应函数调用的各个实参）
		 	并返回一个同类型结果（表示相应函数调用返回的各个结果）。

	请注意：非导出结构体字段值不能用做反射函数调用中的实参。
		如果下例中的vt.FieldByName("A")被替换为vt.FieldByName("b")，则将产生一个恐慌。
*/
type B struct {
	A, b int
}

func (t B) AddSubThenScale(n int) (int, int) {
	return n * (t.A + t.b), n * (t.A - t.b)
}
func TestEg11(te *testing.T) {
	t := B{5, 2}
	vt := reflect.ValueOf(t)

	vm := vt.MethodByName("AddSubThenScale")
	results := vm.Call([]reflect.Value{reflect.ValueOf(3)})
	fmt.Println(results[0].Int(), results[1].Int()) // 21 9

	neg := func(x int) int {
		return -x
	}
	vf := reflect.ValueOf(neg)
	fmt.Println(vf.Call(results[:1])[0].Int()) // -21
	fmt.Println(vf.Call([]reflect.Value{
		vt.FieldByName("A"), // 如果是字段b，则造成恐慌
	})[0].Int()) // -5
}

// 下面是一个使用映射反射值的例子。
func TestEg12(t *testing.T) {
	valueOf := reflect.ValueOf
	m := map[string]int{"Unix": 1973, "Windows": 1985}
	v := valueOf(m)
	// 第二个实参为Value零值时，表示删除一个映射条目。
	v.SetMapIndex(valueOf("Windows"), reflect.Value{})
	v.SetMapIndex(valueOf("Linux"), valueOf(1991))
	for i := v.MapRange(); i.Next(); {
		fmt.Println(i.Key(), "\t:", i.Value())
	}

}

// 下面是一个使用通道反射值的例子。
// reflect.Value类型的TrySend和TryRecv方法对应着只有一个case分支和一个default分支的select流程控制代码块。
func TestEg13(t *testing.T) {
	c := make(chan string, 2)
	vc := reflect.ValueOf(c)
	vc.Send(reflect.ValueOf("C"))

	succeeded := vc.TrySend(reflect.ValueOf("Go"))
	fmt.Println(succeeded) // true

	succeeded = vc.TrySend(reflect.ValueOf("C++"))
	fmt.Println(succeeded) // false

	fmt.Println(vc.Len(), vc.Cap()) // 2 2

	vs, succeeded := vc.TryRecv()
	fmt.Println(vs.String(), succeeded) // C true

	vs, sentBeforeClosed := vc.Recv()
	fmt.Println(vs.String(), sentBeforeClosed) // Go true

	vs, succeeded = vc.TryRecv()
	fmt.Println(vs.String()) //
	fmt.Println(succeeded)   // false

}

// 我们可以使用reflect.Select函数在运行时刻来模拟具有不定case分支数量的select流程控制代码块。
func TestEg14(t *testing.T) {
	c := make(chan int, 1)
	vc := reflect.ValueOf(c)
	succeeded := vc.TrySend(reflect.ValueOf(123))
	fmt.Println(succeeded, vc.Len(), vc.Cap()) // true 1 1

	vSend, vZero := reflect.ValueOf(789), reflect.Value{}
	branches := []reflect.SelectCase{
		{Dir: reflect.SelectDefault, Chan: vZero, Send: vZero},
		{Dir: reflect.SelectRecv, Chan: vc, Send: vZero},
		{Dir: reflect.SelectSend, Chan: vc, Send: vSend},
	}
	selIndex, vRecv, sentBeforeClosed := reflect.Select(branches)
	fmt.Println(selIndex)         // 1
	fmt.Println(sentBeforeClosed) // true
	fmt.Println(vRecv.Int())      // 123
	vc.Close()
	// 再模拟一次select流程控制代码块。因为vc已经关闭了，
	// 所以需将最后一个case分支去除，否则它可能会造成一个恐慌。
	selIndex, _, sentBeforeClosed = reflect.Select(branches[:2])
	fmt.Println(selIndex, sentBeforeClosed) // 1 false
}

// 一些reflect.Value值可能表示着不合法的Go值。 这样的值为reflect.Value类型的零值（即没有底层值的reflect.Value值）
func TestEg15(t *testing.T) {

	var z reflect.Value // 一个reflect.Value零值
	fmt.Printf("%#v\n", z)
	v := reflect.ValueOf((*int)(nil)).Elem()
	fmt.Println(v)      //
	fmt.Println(v == z) // true

	var i = reflect.ValueOf([]interface{}{nil}).Index(0)
	fmt.Println(i, i.Kind())   //
	fmt.Println(i.Elem())      //
	fmt.Println(i.Elem() == z) // true
}

/*
	使用空接口interface{}值做为中介，一个Go值可以转换为一个reflect.Value值。
		逆过程类似，通过调用一个reflect.Value值的Interface方法得到一个interface{}值，
			然后将此interface{}断言为原来的Go值。 但是，请注意，
				调用一个代表着非导出字段的reflect.Value值的Interface方法将导致一个恐慌。
*/
func TestEg16(te *testing.T) {
	vx := reflect.ValueOf(123)
	vy := reflect.ValueOf("abc")
	vz := reflect.ValueOf([]bool{false, true})
	vt := reflect.ValueOf(time.Time{})

	x := vx.Interface().(int)
	y := vy.Interface().(string)
	z := vz.Interface().([]bool)
	m := vt.MethodByName("IsZero").Interface().(func() bool)
	fmt.Println(x, y, z, m()) // 123 abc [false true] true

	type T struct{ x int }
	t := &T{3}
	v := reflect.ValueOf(t).Elem().Field(0)
	fmt.Println(v)             // 3
	fmt.Println(v.Interface()) // panic

}

/*
	从Go 1.17开始，一个切片可以被转化为一个相同元素类型的数组的指针类型。 但是如果在这样的一个转换中数组类型的长度过长，
		将导致恐慌产生。 因此Go 1.17同时引入了一个Value.CanConvert(T Type)方法，用来检查一个转换是否会成功（即不会产生恐慌）。
*/
func TestEg17(t *testing.T) {
	s := reflect.ValueOf([]int{1, 2, 3, 4, 5})
	ts := s.Type()
	t1 := reflect.TypeOf(&[5]int{})
	t2 := reflect.TypeOf(&[6]int{})
	fmt.Println(ts.ConvertibleTo(t1)) // true
	fmt.Println(ts.ConvertibleTo(t2)) // true
	fmt.Println(s.CanConvert(t1))     // true
	fmt.Println(s.CanConvert(t2))     // false
}
