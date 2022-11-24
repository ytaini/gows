/*
 * @Author: wzmiiiiii
 * @Date: 2022-11-21 19:34:03
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-11-22 08:12:02
 * @Description:
 */
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

	rtyp := reflect.TypeOf(v)
	rval := reflect.ValueOf(v)

	typName := rtyp.Name()

	for i := 0; i < rtyp.NumField(); i++ {
		fName := rtyp.Field(i).Name
		fVal := rval.Field(i).Int()
		fmt.Printf("%s 的 第 %d 个字段为 %s,字段值为 %d\n", typName, i, fName, fVal)
	}

	for i := 0; i < rtyp.NumMethod(); i++ {
		rm := rtyp.Method(i)
		funcName := rm.Name
		args := []reflect.Value{reflect.ValueOf("Tom")}
		rval.MethodByName(funcName).Call(args)
	}

}
