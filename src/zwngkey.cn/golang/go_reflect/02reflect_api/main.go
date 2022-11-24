package main

import (
	"fmt"
	"io"
	"log"
	"reflect"
	"strconv"
	"time"
)

type Foo struct {
	name string
	age  int
	desp string
}

func (f Foo) Hello() {
	fmt.Println("hello")
}
func (f Foo) Say() {
	fmt.Println("say hello")
}
func F1() {
	//without reflection
	f := Foo{
		name: "zs",
		age:  18,
		desp: "xixi",
	}

	//with reflection
	fT := reflect.TypeOf(Foo{
		name: "li",
		age:  18,
		desp: "xixi"})

	fT2 := reflect.TypeOf(f)

	//1.如果两个Type值代表相同的类型,那么他们一定是相等的.
	i := reflect.TypeOf(int(3))
	fmt.Println(fT == fT2)              //true
	fmt.Println(i == reflect.TypeOf(4)) //true

	//2.Align():返回该类型在内存中分配时，以字节数为单位的字节数
	fmt.Printf("fT.Align(): %v\n", fT.Align())
	fmt.Printf("fT.Align(): %v\n", i.Align())

	//3.FieldAlign():返回该类型在结构中作为字段使用时，以字节数为单位的字节数
	fmt.Printf("fT.Align(): %v\n", fT.FieldAlign())
	fmt.Printf("fT.Align(): %v\n", i.FieldAlign())

	// Kind返回该接口的具体分类
	fmt.Println(fT.Kind())
	fmt.Println(i.Kind())

	// Name返回该类型在自身包内的类型名，如果是未命名类型会返回""
	fmt.Println(fT.Name()) //Foo
	fmt.Println(i.Name())  // int

	// PkgPath返回类型的包路径，即明确指定包的import路径，如"encoding/base64"
	// 如果类型为内建类型(string, error)或未命名类型(*T, struct{}, []int)，会返回""
	fmt.Println(fT.PkgPath())
	fmt.Println(i.PkgPath())

	// 返回要保存一个该类型的值需要多少字节；类似unsafe.Sizeof
	fmt.Println(fT.Size())
	fmt.Println(i.Size())

	// 检查当前类型能不能做比较运算，其实就是看这个类型底层有没有绑定 typeAlg 的 equal 方法。
	fmt.Println(fT.Comparable())
	fmt.Println(i.Comparable())

	// 返回类型方法集中可导出的方法的数量
	// 匿名字段的方法会被计算；主体类型的方法会屏蔽匿名字段的同名方法；
	// 匿名字段导致的歧义方法会滤除
	fmt.Println(fT.NumMethod())
	fmt.Println(i.NumMethod())

	// 根据方法名返回该类型方法集中的方法，使用一个布尔值说明是否发现该方法
	// 对非接口类型T或*T，返回值的Type字段和Func字段描述方法的未绑定函数状态
	// 对接口类型，返回值的Type字段描述方法的签名，Func字段为nil
	fTM, ok := fT.MethodByName("Hello")

	if !ok {
		log.Fatalln("no method")
	}
	fmt.Printf("fTM.Index: %v\n", fTM.Index) //0
	// Name是方法名
	fmt.Printf("fTM.Name: %v\n", fTM.Name) //Hello
	// PkgPath是非导出字段的包路径，对导出字段该字段为""。
	fmt.Printf("fTM.PkgPath: %v\n", fTM.PkgPath)
	//
	fmt.Printf("fTM.IsExported(): %v\n", fTM.IsExported()) //true
	fmt.Printf("fTM.Type: %v\n", fTM.Type)                 //func(main.Foo)
	fmt.Printf("fTM.Type.Name(): %q\n", fTM.Type.Name())   //""

	fV := reflect.ValueOf(f)
	fTM.Func.Call([]reflect.Value{fV}) //调用Hello 方法,接收者作为第一个参数.

	fmt.Printf("fV.Kind(): %v\n", fV.Kind()) // struct
	fmt.Printf("fV.Type(): %v\n", fV.Type()) // main.Foo
	// fT = fT.Elem()

	fVF := reflect.ValueOf(f.Hello)
	fmt.Printf("fVF.Kind(): %v\n", fVF.Kind()) //func
	fmt.Printf("fVF.Type(): %v\n", fVF.Type()) // func()
	fVF.Call(nil)

	var x float64 = 3.4
	v := reflect.ValueOf(&x)
	fmt.Printf("v.CanSet(): %v\n", v.CanSet()) //false

	// Value.Elem返回v持有的接口保管的值的Value封装，或者v持有的指针指向的值的Value封装。
	// 如果v的Kind不是Interface或Ptr会panic；如果v持有的值为nil，会返回Value零值
	p := v.Elem()
	fmt.Printf("p.CanSet(): %v\n", p.CanSet()) // true
	p.SetFloat(3.1313)

	fmt.Printf("x: %v\n", x)

	var w = new(io.Writer)

	t := reflect.TypeOf(w)
	a := t.Elem()
	fmt.Println(a.Name()) //Writer
	fmt.Println(a.Kind()) //interface
}
func main() {
	F1()
	// F2()

}
func F2() {
	var x int64 = 1
	var d time.Duration = 1 * time.Nanosecond
	f := Foo{
		name: "zs",
		age:  18,
		desp: "xixi",
	}

	fmt.Println(Any(x)) //1
	fmt.Println(Any(d)) //1
	fmt.Println(Any([]int64{int64(x)}))
	fmt.Println(Any([]time.Duration{d}))
	fmt.Println(Any("[]time.Duration{d}"))
	fmt.Println(Any(f)) //1
}

func Any(value any) string {
	return formatAtom(reflect.ValueOf(value))
}

func formatAtom(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Invalid:
		return "invalid"
	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(v.Uint(), 10)
	case reflect.Float32, reflect.Float64:
		return strconv.FormatFloat(v.Float(), 'e', 4, 64)
	case reflect.Bool:
		return strconv.FormatBool(v.Bool())
	case reflect.String:
		return strconv.Quote(v.String())
	case reflect.Chan, reflect.Func, reflect.Ptr, reflect.Slice, reflect.Map:
		return v.Type().String() + " 0x" +
			strconv.FormatUint(uint64(v.Pointer()), 16)
	default: // reflect.Array, reflect.Struct, reflect.Interface
		return v.Type().String() + " value"
	}
}
