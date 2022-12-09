/*
 * @Author: wzmiiiiii
 * @Date: 2022-12-08 23:22:34
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-12-09 06:09:23
 */
package main

import (
	"fmt"

	. "zwngkey.cn/designpattern/creational/builder/case2"
)

func main() {
	// NewHero()
	hero := NewHeroBuilder("律师", "张三").Build()
	fmt.Println(hero)
	hero1 := NewHeroBuilder("律师", "张三").WithHairColor("白色").Build()
	fmt.Println(hero1)
	hero2 := NewHeroBuilder("律师", "张三").WithHairColor("红色").WithHairType("平头").Build()
	fmt.Println(hero2)

}
