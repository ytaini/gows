package main

import (
	"fmt"
	"reflect"
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

func f1() {
	var i = 10
	rVal := reflect.ValueOf(&i)

	fmt.Printf("rVal.Type(): %v\n", rVal.Type())
	fmt.Printf("rVal.Kind(): %v\n", rVal.Kind())
	fmt.Printf("rVal.CanSet(): %v\n", rVal.CanSet())

	fmt.Printf("rVal.Elem().CanSet(): %v\n", rVal.Elem().CanSet())
	rVal.Elem().SetInt(64)
	fmt.Printf("i: %v\n", i)

}
func f2() {
	p := Person{
		Name: "ls",
		Age:  18,
	}
	rVal := reflect.ValueOf(p)
	t := rVal.FieldByName("name")
	fmt.Printf("t: %v\n", t)
	fmt.Printf("t.Type(): %v\n", t.Type())
	fmt.Printf("t.Kind(): %v\n", t.Kind())
	t = rVal.FieldByName("age")
	fmt.Printf("t: %v\n", t)
	fmt.Printf("t.Type(): %v\n", t.Type())
	fmt.Printf("t.Kind(): %v\n", t.Kind())
	a := rVal.MethodByName("Hello")
	fmt.Printf("a: %v\n", a)
	fmt.Printf("a.Type(): %v\n", a.Type())
	fmt.Printf("a.Kind(): %v\n", a.Kind())
	ret := a.Call(nil)
	fmt.Printf("ret: %v\n", ret)
	fmt.Printf("rVal.NumMethod(): %v\n", rVal.NumMethod())

	//这种方式不可以调用私有方法
	b := rVal.MethodByName("say")
	fmt.Printf("b: %v\n", b)
	fmt.Printf("b: %v\n", b.Kind())
	// fmt.Printf("b.Type(): %v\n", b.Type())
	// b.Call(nil)
}

func main() {
	p := Person{
		Name: "ls",
		Age:  18,
	}
	p1 := test1(p).(Person)
	fmt.Printf("p1: %v\n", p1)
	fmt.Printf("p1: %T\n", p1)

	rVal := reflect.ValueOf(&p)
	fmt.Printf("rVal: %v\n", rVal)
	fmt.Printf("rVal.Type(): %v\n", rVal.Type())
	fmt.Printf("rVal.Kind(): %v\n", rVal.Kind())

	re1 := rVal.MethodByName("Hello")
	res2 := re1.Call(nil)
	fmt.Printf("res2: %v\n", res2)

	a := rVal.Elem()
	fmt.Printf("a: %v\n", a)
	fmt.Printf("a.Type(): %v\n", a.Type())
	fmt.Printf("a.Kind(): %v\n", a.Kind())

	c := a.FieldByName("Name")
	fmt.Printf("c.CanSet(): %v\n", c.CanSet())
	fmt.Printf("c.Type(): %v\n", c.Type())
	c.SetString("zs")
	fmt.Printf("p.Name: %v\n", p.Name)

	a.FieldByName("Age").SetUint(19)
	fmt.Printf("p.Age: %v\n", p.Age)

	ret := a.Addr().MethodByName("Hello")
	res := ret.Call(nil)
	fmt.Printf("res: %v\n", res)
}

func test1(a any) any {
	var temp reflect.Value = reflect.ValueOf(a)
	var i any = temp.Interface()
	return i
}
