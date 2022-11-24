/*
 * @Author: wzmiiiiii
 * @Date: 2022-11-21 19:34:16
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-11-22 10:11:37
 * @Description:
 */
package main

import (
	"fmt"
	"reflect"
)

type sex uint8

type Person struct {
	Name string
	age  sex
	Addr
}

func (p *Person) Hello() string {
	return "hello"
}
func (p Person) say() {
	fmt.Println("say")
}

type Addr struct {
	city string
}

func main() {
	p := Person{
		Name: "ls",
		age:  18,
	}

	typ := reflect.TypeOf(&p)
	tpv := reflect.TypeOf(p)
	val := reflect.ValueOf(p)
	valp := reflect.ValueOf(&p)

	sf1 := reflect.TypeOf(p).FieldByIndex([]int{2, 0})
	fmt.Println(sf1.Name) //city
	fmt.Println(sf1.Type) //string

	fmt.Printf("typ.Kind(): %v\n", typ.Kind())   //ptr
	fmt.Printf("typ.Kind(): %v\n", tpv.Kind())   //struct
	fmt.Printf("val.Kind(): %v\n", val.Kind())   //struct
	fmt.Printf("valp.Kind(): %v\n", valp.Kind()) //ptr

	fmt.Printf("typ.Name(): %v\n", typ.Name())                 //""
	fmt.Printf("typ.Name(): %v\n", tpv.Name())                 //Person
	fmt.Printf("val.Type().Name(): %v\n", val.Type().Name())   //Person
	fmt.Printf("valp.Type().Name(): %v\n", valp.Type().Name()) //""

	fmt.Printf("typ.Elem().NumField(): %v\n", typ.Elem().NumField()) // 2
	fmt.Printf("tpv.NumField(): %v\n", tpv.NumField())               // 2
	fmt.Printf("val.NumField(): %v\n", val.NumField())               // 2
	fmt.Printf("valp.NumField(): %v\n", valp.Elem().NumField())      //2

	sf, _ := typ.Elem().FieldByName("Name")
	fmt.Printf("sf: %v\n", sf)           //{Name  string  0 [0] false}
	fmt.Printf("sf.Name: %v\n", sf.Name) //Name
	fmt.Printf("sf.Type: %v\n", sf.Type) //string

	sf, _ = tpv.FieldByName("Name")
	fmt.Printf("sf: %v\n", sf)           //{Name  string  0 [0] false}
	fmt.Printf("sf.Name: %v\n", sf.Name) //Name
	fmt.Printf("sf.Type: %v\n", sf.Type) //string

	fmt.Printf("typ.NumMethod(): %v\n", typ.NumMethod())   //0
	fmt.Printf("tpv.NumMethod(): %v\n", tpv.NumMethod())   //0
	fmt.Printf("val.NumMethod(): %v\n", val.NumMethod())   //0
	fmt.Printf("valp.NumMethod(): %v\n", valp.NumMethod()) //1

	rm1, ok1 := typ.MethodByName("Hello")
	fmt.Printf("rm1: %v,ok1: %v\n", rm1, ok1) //fasle

	rm2, ok2 := typ.MethodByName("say")
	fmt.Printf("rm2: %v,ok2: %v\n", rm2, ok2) //fasle

}
