package main

import (
	"fmt"
	"reflect"
)

type sex uint8

type Person struct {
	Name string
	age  sex
}

func (p *Person) Hello() string {
	return "hello"
}
func (p Person) say() {
	fmt.Println("say")
}
func main() {
	p := Person{
		Name: "ls",
		age:  18,
	}

	typ := reflect.TypeOf(&p)
	val := reflect.ValueOf(p)
	typ = typ.Elem()

	fmt.Printf("typ.NumField(): %v\n", typ.NumField())
	sf, _ := typ.FieldByName("Name")
	fmt.Printf("sf: %v\n", sf)
	fmt.Printf("sf.Name: %v\n", sf.Name)
	fmt.Printf("sf.Type: %v\n", sf.Type)

	fmt.Printf("typ.NumMethod(): %v\n", typ.NumMethod())

	rm1, ok1 := typ.MethodByName("Hello")
	fmt.Printf("ok1: %v\n", ok1)
	res := rm1.Func.Call([]reflect.Value{
		val,
	})
	fmt.Printf("res: %v\n", res)

	rm2, ok2 := typ.MethodByName("say")
	fmt.Printf("ok2: %v\n", ok2)
	res2 := rm2.Func.Call([]reflect.Value{
		val,
	})
	fmt.Printf("res2: %v\n", res2)

}
