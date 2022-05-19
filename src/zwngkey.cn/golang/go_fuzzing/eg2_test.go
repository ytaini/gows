/*
 * @Author: zwngkey
 * @Date: 2022-05-18 18:04:39
 * @LastEditors: zwngkey 18390924907@163.com
 * @LastEditTime: 2022-05-18 20:29:19
 * @Description:
	go 模糊测试
*/
package gofuzzing

import (
	"testing"
)

/*
	Fuzzing 是一种自动化的测试技术， 它不断的创建输入用来测试程序的bug。
		Go fuzzing使用覆盖率智能指导遍历被模糊化测试的代码，发现缺陷并报告给用户。
		由于模糊测试可以达到人类经常忽略的边缘场景，因此它对于发现安全漏洞和缺陷特别有价值。
*/

/*
	下面是模糊测试必须遵循的规则：
		1.模糊测试必须是一个名称类似FuzzXxx的函数，仅仅接收一个*testing.F类型的参数,没有返回值

		2.模糊测试必须在*_test.go文件中才能运行

		3.Fuzz target(模糊目标)必须是对(*testing.F).Fuzz的方法调用，参数是一个函数，并且此函数的第一个参数是*testing.T,然后是模糊参数(fuzzing argument)，没有返回值

		4.一个模糊测试中必须只有一个模糊目标

		5.所有的种子语料库(seed corpus)必须具有与模糊参数相同的类型,顺序相同。对(*testing.F).Add的调用也是如此, 同样也适用模糊测试中的testdata/fuzz中的语料文件

	模糊参数只能是下面的类型
		string, []byte
		int, int8, int16, int32/rune, int64
		uint, uint8/byte, uint16, uint32, uint64
		float32, float64
		bool

*/

/*
	https://colobu.com/2022/01/03/go-fuzzing/
*/
func Test21(t *testing.T) {

}
