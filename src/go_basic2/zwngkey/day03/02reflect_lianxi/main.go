package main

import (
	"fmt"
	"log"
	"reflect"
)

type Monster struct {
	Name  string `json:"name"`
	Age   int    `json:"monster_age"`
	Score float32
	Sex   string
}

func (m *Monster) Print() {
	fmt.Println("---start---")
	fmt.Println(m)
	fmt.Println("---end---")
}

func (m *Monster) GetSum(n1, n2 int) int {
	return n1 + n2
}

func (m *Monster) Set(name string, age int, score float32, sex string) {
	m.Name = name
	m.Age = age
	m.Score = score
	m.Sex = sex
}

func TestStruct(a interface{}) {
	fmt.Printf("a: %T\n", a)
	typ := reflect.TypeOf(a)
	val := reflect.ValueOf(a)
	typM := typ.Elem()
	valM := val.Elem()

	kd := valM.Kind()
	if kd != reflect.Struct {
		log.Fatalln("expect struct")
	}

	tName := typM.String()

	for i := 0; i < valM.NumField(); i++ {
		tf := typM.Field(i)
		tfName := tf.Name
		tfTag := tf.Tag.Get("json")
		vf := valM.Field(i)
		switch vf.Kind() {
		case reflect.String:
			fmt.Printf("%s的第 %d 个字段为 %s,字段值为 %v,json tag 值为 %s\n", tName, i, tfName, vf.String(), tfTag)
		case reflect.Int:
			fmt.Printf("%s的第 %d 个字段为 %s,字段值为 %v,json tag 值为 %s\n", tName, i, tfName, vf.Int(), tfTag)
		case reflect.Float32:
			fmt.Printf("%s的第 %d 个字段为 %s,字段值为 %v,json tag 值为 %s\n", tName, i, tfName, vf.Float(), tfTag)
		}
	}
	fmt.Printf("val.NumMethod(): %v\n", val.NumMethod())
	fmt.Printf("typ.NumMethod(): %v\n", typ.NumMethod())
	fmt.Printf("val.NumMethod(): %v\n", valM.NumMethod())
	fmt.Printf("typ.NumMethod(): %v\n", typM.NumMethod())

	val.Method(1).Call(nil)

	args := []reflect.Value{reflect.ValueOf(1), reflect.ValueOf(2)}
	res := val.Method(0).Call(args)
	fmt.Printf("res: %v\n", res[0].Int())

	args = []reflect.Value{reflect.ValueOf("zs"), reflect.ValueOf(16), reflect.ValueOf(float32(66.6)), reflect.ValueOf("男")}
	val.Method(2).Call(args)

	val.Method(1).Call(nil)
}

func main() {
	m := &Monster{
		Name:  "ls",
		Age:   13,
		Score: 99.0,
		Sex:   "nv",
	}
	TestStruct(m)
}
