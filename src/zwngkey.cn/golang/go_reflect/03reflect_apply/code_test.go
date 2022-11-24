package myreflect

import (
	"fmt"
	"reflect"
	"testing"
)

type sex uint8

type Person struct {
	Name string
	Age  sex
}

func (p *Person) Hello() string {
	return "hello"
}
func (p Person) say() {
	fmt.Println("say")
}

func Test1(t *testing.T) {
	var i = 10
	rVal := reflect.ValueOf(&i)

	fmt.Printf("rVal.Type(): %v\n", rVal.Type())     //*int
	fmt.Printf("rVal.Kind(): %v\n", rVal.Kind())     // Ptr
	fmt.Printf("rVal.CanSet(): %v\n", rVal.CanSet()) //false

	fmt.Printf("rVal.Elem().CanSet(): %v\n", rVal.Elem().CanSet()) //true
	rVal.Elem().SetInt(64)
	fmt.Printf("i: %v\n", i) //64
}
func Test2(te *testing.T) {
	p := Person{
		Name: "ls",
		Age:  18,
	}
	rVal := reflect.ValueOf(p)
	t := rVal.FieldByName("Name")
	fmt.Printf("t: %q\n", t.Interface().(string)) // "ls"
	fmt.Printf("t.Type(): %v\n", t.Type())        // string
	fmt.Printf("t.Kind(): %v\n", t.Kind())        // string
	t = rVal.FieldByName("Age")
	fmt.Printf("t: %v\n", t.Uint())        //18
	fmt.Printf("t.Type(): %v\n", t.Type()) //myreflect.sex
	fmt.Printf("t.Kind(): %v\n", t.Kind()) //uint8

	fmt.Println(rVal.CanAddr()) // false

	pVal := reflect.ValueOf(&p)
	a := pVal.MethodByName("Hello")
	fmt.Printf("a.Type(): %v\n", a.Type()) // func() string
	fmt.Printf("a.Kind(): %v\n", a.Kind()) // func
	ret := a.Call(nil)
	fmt.Printf("ret: %v\n", ret)                           //[hello]
	fmt.Printf("rVal.NumMethod(): %v\n", rVal.NumMethod()) //0
	fmt.Printf("rVal.NumMethod(): %v\n", pVal.NumMethod()) //1

	//这种方式不可以调用私有方法
	b := rVal.MethodByName("say")
	fmt.Printf("b.IsValid(): %v\n", b.IsValid()) //false
	fmt.Printf("b: %v\n", b)
	fmt.Printf("b: %v\n", b.Kind())
	// fmt.Printf("b.Type(): %v\n", b.Type())
	// b.Call(nil)
}

func Test3(t *testing.T) {
	p := Person{
		Name: "ls",
		Age:  18,
	}
	p1 := test1(p).(Person)
	fmt.Printf("p1: %v\n", p1) // {ls 18}
	fmt.Printf("p1: %T\n", p1) //myreflect.Person

	rVal := reflect.ValueOf(&p)
	fmt.Printf("rVal: %v\n", rVal)               // &{ls 18}
	fmt.Printf("rVal.Type(): %v\n", rVal.Type()) //*myreflect.Person
	fmt.Printf("rVal.Kind(): %v\n", rVal.Kind()) //ptr

	re1 := rVal.MethodByName("Hello")
	res2 := re1.Call(nil)
	fmt.Printf("res2: %v\n", res2) //[hello ]

	a := rVal.Elem()
	fmt.Printf("a: %v\n", a)               //{ls 18}
	fmt.Printf("a.Type(): %v\n", a.Type()) //myreflect.Person
	fmt.Printf("a.Kind(): %v\n", a.Kind()) //struct

	c := a.FieldByName("Name")
	fmt.Printf("c.CanSet(): %v\n", c.CanSet()) //true
	fmt.Printf("c.Type(): %v\n", c.Type())     //string
	c.SetString("zs")
	fmt.Printf("p.Name: %v\n", p.Name) // zs

	a.FieldByName("Age").SetUint(19)
	fmt.Printf("p.Age: %v\n", p.Age) //19

	fmt.Println(a.CanAddr()) //true
	ret := a.Addr().MethodByName("Hello")
	res := ret.Call(nil)
	fmt.Printf("res: %v\n", res) //[hello]
}

func test1(a any) any {
	temp := reflect.ValueOf(a)
	i := temp.Interface()
	return i
}
