package main

import (
	"fmt"
	"reflect"
)

type Cal struct {
	Num1, Num2 int
}

func (c Cal) GetSub(name string) {
	fmt.Printf("%s 完成了减法运算,%d - %d = %d\n", name, c.Num1, c.Num2, c.Num1-c.Num2)
}

func main() {
	v := Cal{
		Num1: 10,
		Num2: 5,
	}

	rTyp := reflect.TypeOf(v)
	rVal := reflect.ValueOf(v)

	typName := rVal.Type().Name()

	for i := 0; i < rVal.NumField(); i++ {
		fName := rTyp.Field(i).Name
		fVal := rVal.Field(i).Int()
		fmt.Printf("%s 的 第 %d 个字段为 %s,字段值为 %d\n", typName, i, fName, fVal)
	}

	for i := 0; i < rVal.NumMethod(); i++ {
		rm := rTyp.Method(i)
		funcName := rm.Name
		args := []reflect.Value{reflect.ValueOf("Tom")}
		rVal.MethodByName(funcName).Call(args)
	}

}
