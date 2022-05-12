/*
 * @Author: zwngkey
 * @Date: 2022-05-04 00:05:17
 * @LastEditTime: 2022-05-04 01:03:58
 * @Description:
 */
package goset

import (
	"fmt"
	"strconv"
	"testing"
)

func Test(t *testing.T) {
	str := make([]string, 0)
	for i := 0; i < 10; i++ {
		str = append(str, "go"+strconv.Itoa(i))
	}
	set1 := NewString(str...)
	for k := range set1 {
		fmt.Println(k)
	}

	// set1.DeleteAll()
	// set1.Delete("go1", "go2")

	fmt.Println("----")
	for k := range set1 {
		fmt.Println(k)
	}
}
func Test2(t *testing.T) {
	map1 := make(map[string]Empty)

	for i := 0; i < 10; i++ {
		map1["map"+strconv.Itoa(i)] = Empty{}
	}

	set1 := StringKeySet(map1)

	_ = set1
}
