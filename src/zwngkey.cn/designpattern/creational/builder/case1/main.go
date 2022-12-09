/*
 * @Author: wzmiiiiii
 * @Date: 2022-12-09 00:10:29
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-12-09 00:15:40
 */

package main

import "fmt"

func main() {

	// 不同具体建造者,得到不同口味的汉堡(不同形式的同一类型的对象)
	// builder2 := NewConcreteBurgerBuilder2()
	// director := NewDirector(builder2) //

	builder1 := NewConcreteBurgerBuilder1()
	director := NewDirector(builder1) //
	director.Construct()              //构建一个汉堡
	burger := director.GetBurger()    //得到汉堡
	fmt.Println(burger)

}
