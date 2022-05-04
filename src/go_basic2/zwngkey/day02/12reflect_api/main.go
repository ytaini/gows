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
func f1() {
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
	i := reflect.TypeOf(3)
	fmt.Println(fT == fT2)              //true
	fmt.Println(i == reflect.TypeOf(4)) //true

	//2.Align():返回该类型在内存中分配时，以字节数为单位的字节数
	fmt.Printf("fT.Align(): %v\n", fT.Align())
	fmt.Printf("fT.Align(): %v\n", i.Align())

	//3.FieldAlign():返回该类型在结构中作为字段使用时，以字节数为单位的字节数
	fmt.Printf("fT.Align(): %v\n", fT.FieldAlign())
	fmt.Printf("fT.Align(): %v\n", i.FieldAlign())

	//
	fmt.Println(fT.Name())
	fmt.Println(i.Name())

	//
	fmt.Println(fT.PkgPath())
	fmt.Println(i.PkgPath())

	fmt.Println(fT.Size())
	fmt.Println(i.Size())

	fmt.Println(fT.Comparable())
	fmt.Println(i.Comparable())

	fmt.Println(fT.NumMethod())
	fmt.Println(i.NumMethod())

	fmt.Println(fT.Kind())
	fmt.Println(i.Kind())

	fTM, ok := fT.MethodByName("Hello")

	if !ok {
		log.Fatalln("no method")
	}
	fmt.Printf("fTM.Index: %v\n", fTM.Index)
	fmt.Printf("fTM.Name: %v\n", fTM.Name)
	fmt.Printf("fTM.PkgPath: %v\n", fTM.PkgPath)
	fmt.Printf("fTM.IsExported(): %v\n", fTM.IsExported())
	fmt.Printf("fTM.Type: %v\n", fTM.Type)
	fmt.Printf("fTM.Type.Name(): %v\n", fTM.Type.Name())

	fV := reflect.ValueOf(f)
	fTM.Func.Call([]reflect.Value{fV})

	fmt.Printf("fV.Kind(): %v\n", fV.Kind())
	fmt.Printf("fV.Type(): %v\n", fV.Type())
	// fT = fT.Elem()

	fVF := reflect.ValueOf(f.Hello)
	fmt.Printf("fVF.Kind(): %v\n", fVF.Kind())
	fmt.Printf("fVF.Type(): %v\n", fVF.Type())
	fVF.Call(nil)

	var x float64 = 3.4
	v := reflect.ValueOf(&x)
	fmt.Printf("v.CanSet(): %v\n", v.CanSet())
	p := v.Elem()
	fmt.Printf("p.CanSet(): %v\n", p.CanSet())
	p.SetFloat(3.1313)

	fmt.Printf("x: %v\n", x)
	fmt.Printf("p: %v\n", p)

	var w = new(io.Writer)

	t := reflect.TypeOf(w)
	a := t.Elem()
	fmt.Println(a.Name())
	fmt.Println(a.Kind())
}
func main() {
	f1()
	f2()

}
func f2() {
	var x int64 = 1
	var d time.Duration = 1 * time.Nanosecond
	fmt.Println(Any(x))
	fmt.Println(Any(d))
	fmt.Println(Any([]int64{int64(x)}))
	fmt.Println(Any([]time.Duration{d}))
	fmt.Println(Any("[]time.Duration{d}"))
}

func Any(value interface{}) string {
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
