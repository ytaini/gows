/*
 * @Author: wzmiiiiii
 * @Date: 2022-12-09 01:32:32
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-12-09 01:37:06
 */
package main

import (
	"fmt"

	. "zwngkey.cn/designpattern/idioms/fluentapi"
)

func main() {
	profile := NewFluentServiceProfileBuilder().
		WithId("service1").
		WithType("order").
		WithStatus("normal").
		WithEndpoint("localhost", 8080).
		WithRegion("region1", "hunan", "China").
		WithPriority(1).
		WithLoad(100).
		Build()

	fmt.Println(profile)
}
