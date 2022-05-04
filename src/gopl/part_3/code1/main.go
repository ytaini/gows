package main

import (
	"fmt"
	"sort"
)

func main() {
	// map1()
	// map2()
	// map3()
	m1 := map[string]int{
		"1": 1,
		"2": 2,
		"3": 3,
		"4": 4,
	}
	m2 := map[string]int{
		"1": 1,
		"2": 2,
		"5": 3,
		"4": 4,
	}

	fmt.Println(equal(m1, m2))

}

func equal(m1, m2 map[string]int) bool {
	if len(m1) != len(m2) {
		return false
	}

	for k, v := range m1 {
		if value, ok := m2[k]; !ok || v != value {
			return false
		}
	}
	return true
}

func map3() {
	colors := map[string]int{
		"blue":   0,
		"yellow": 1,
		"pink":   2,
		"red":    3,
	}

	code := colors["blue"]

	fmt.Printf("%v\n", code)          //输出:0
	fmt.Printf("%v\n", colors["123"]) //输出:0 这时不能区分一个已经存在的value=0,与value不存在而返回零值的0.

	code, ok := colors["blue"]
	if !ok {
		fmt.Print("blue is not a key")
		code = -1
	}
	fmt.Printf("code: %v\n", code)

	code, ok = colors["qwe"]

	if !ok {
		fmt.Print("qwe is not a key ")
		code = -1
	}
	fmt.Printf("code: %v\n", code)

	if num, yes := colors["qwe"]; !yes {
		num = -1
		fmt.Println("blue is not a key")
		fmt.Printf("num: %v\n", num)
	}

}

func map2() {
	var man map[string]string = make(map[string]string)
	man["school"] = "湖大"
	man["tel"] = "123729"
	man["type"] = "高富帅"

	fmt.Println(man["sex"]) // "":空字符串

	man["like"] = man["like"] + ",apple"
	fmt.Println(man["like"]) // ",apple" 当man["like"]中key不存在时,将返回value对应类型的零值.

	for t, v := range man {
		fmt.Printf("man [ %q , %q ] \n", t, v) //每次迭代的顺序是不确定的.
	}

	//如果要按顺序遍历:我们必须显式的对key进行排序.
	var props []string
	for k := range man {
		props = append(props, k) //将map中所有的Key添加到props切片中
	}

	sort.Strings(props) //对切片进行排序

	for _, v := range props {
		fmt.Println(v, "-", man[v]) //根据props中的元素得到map中value的值.
	}

}

func map1() {
	var m map[string]string = map[string]string{
		"name": "zhangsan",
		"age":  "18",
	}

	m["sex"] = "男"                  //ADD
	m["address"] = "湖南"             //ADD
	m["age"] = "20"                 //MODIFY
	fmt.Printf("m: %v\n", m["age"]) //CHECK
	fmt.Printf("m: %v\n", m)
	delete(m, "address") //DELETE
	fmt.Printf("m: %v\n", m)
}
