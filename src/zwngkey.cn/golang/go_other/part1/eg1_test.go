/*
 * @Author: zwngkey
 * @Date: 2022-05-12 20:14:00
 * @LastEditors: zwngkey 18390924907@163.com
 * @LastEditTime: 2022-05-12 20:25:51
 * @Description:
 	Go语言中fmt.Println(true)的结果一定是true么？

	true,false,nil 并不是Go语言的关键字

	Go语言的这个特性也引起了很多争议，正如Go的错误处理一样。

	我们作为使用者，需要注意：Go不允许把关键字作为标识符，其它都可以作为标识符。
*/
package gootherpartone

import (
	"fmt"
	"testing"
)

var nil = 100

var false = true

func Test51(t *testing.T) {

	true := false
	fmt.Println(true, false, nil)

}
