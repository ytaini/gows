package main

import (
	"fmt"
	"log"
	"reflect"
)

type user struct {
	Name string
	Age  int
}

func main() {
	var model *user

	rTyp := reflect.TypeOf(model)

	if rTyp.Kind() == reflect.Ptr {
		rTyp = rTyp.Elem()
		if rTyp.Kind() != reflect.Struct {
			log.Fatalln("type error")
		}
	}

	//创建*user类型的对象实例
	rValModel := reflect.New(rTyp)

	model = rValModel.Interface().(*user)

	rValModel = rValModel.Elem()

	rValModel.FieldByName("Name").SetString("ls")
	rValModel.FieldByName("Age").SetInt(18)

	fmt.Printf("model: %v\n", model)
}
