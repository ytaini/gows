/*
 * @Author: zwngkey
 * @Date: 2022-05-13 00:54:20
 * @LastEditors: zwngkey 18390924907@163.com
 * @LastEditTime: 2022-05-18 20:36:17
 * @Description: go 模糊测试
 */
package gofuzzing

import (
	"fmt"
	"testing"
	"unicode/utf8"
)

/*
	注意：fuzzing模糊测试和Go已有的单元测试以及性能测试框架是互为补充的，并不是替代关系。

	单元测试有局限性，每个测试输入必须由开发者指定，显示加到单元测试的测试用例里。

	fuzzing的优点之一是可以基于开发者代码里指定的测试输入作为基础数据，进一步自动生成新的随机测试数据，
		用来发现指定测试输入没有覆盖到的边界情况。

	Fuzzing也有一定的局限性。

	在单元测试里，因为测试输入是固定的，你可以知道调用Reverse函数后每个输入字符串得到的反转字符串应该是什么，
		然后在单元测试的代码里判断Reverse的执行结果是否和预期相符。
			例如，对于测试用例Reverse("Hello, world")，单元测试预期的结果是 "dlrow ,olleH"。

	但是使用fuzzing时，我们没办法预期输出结果是什么，因为测试的输入除了我们代码里指定的用例之外，还有fuzzing随机生成的。
		对于随机生成的测试输入，我们当然没办法提前知道输出结果是什么。

	Go模糊测试和单元测试在语法上有如下差异：
		Go模糊测试函数以FuzzXxx开头，单元测试函数以TestXxx开头

		Go模糊测试函数以 *testing.F作为入参，单元测试函数以*testing.T作为入参

		Go模糊测试会调用f.Add函数和f.Fuzz函数。
			f.Add函数把指定输入作为模糊测试的种子语料库(seed corpus)，fuzzing基于种子语料库生成随机输入。
			f.Fuzz函数接收一个fuzz target函数作为入参。fuzz target函数有多个参数，第一个参数是*testing.T，其它参数是被模糊的类型(注意：被模糊的类型目前只支持部分内置类型, 列在 Go Fuzzing docs，未来会支持更多的内置类型)。

*/
func Reverse(s string) string {
	b := []byte(s)
	for i, j := 0, len(b)-1; i < len(b)/2; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}
func Test11(t *testing.T) {
	input := "The quick brown fox jumped over the lazy dog"
	rev := Reverse(input)
	doubleRev := Reverse(rev)
	fmt.Printf("original: %q\n", input)
	fmt.Printf("reversed: %q\n", rev)
	fmt.Printf("reversed again: %q\n", doubleRev)

}
func TestReverse(t *testing.T) {
	testcases := []struct {
		in, want string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{" ", " "},
		{"!12345", "54321!"},
	}
	for _, tc := range testcases {
		rev := Reverse(tc.in)
		if rev != tc.want {
			t.Errorf("Reverse: %q, want %q", rev, tc.want)
		}
	}
}
func FuzzReverse(f *testing.F) {
	testcases := []string{"Hello, world", " ", "!12345"}
	for _, tc := range testcases {
		f.Add(tc) // Use f.Add to provide a seed corpus
	}
	f.Fuzz(func(t *testing.T, orig string) {
		rev := Reverse(orig)
		doubleRev := Reverse(rev)
		if orig != doubleRev {
			t.Errorf("Before: %q, after: %q", orig, doubleRev)
		}
		if utf8.ValidString(orig) && !utf8.ValidString(rev) {
			t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
		}
	})
}

/*
  运行模糊测试
	1.执行如下命令来运行模糊测试。
		这个方式只会使用种子语料库，而不会生成随机测试数据。通过这种方式可以用来验证种子语料库的测试数据是否可以测试通过。(fuzz test without fuzzing)
		$ go test

		如果reverse_test.go文件里有其它单元测试函数或者模糊测试函数，但是只想运行FuzzReverse模糊测试函数，
			我们可以执行go test -run=FuzzReverse命令。

		注意：go test默认会执行所有以TestXxx开头的单元测试函数和以FuzzXxx开头的模糊测试函数，
			默认不运行以BenchmarkXxx开头的性能测试函数，如果我们想运行 benchmark用例，则需要加上 -bench 参数。

	2.如果要基于种子语料库生成随机测试数据用于模糊测试，需要给go test命令增加 -fuzz参数。
		go test -fuzz={FuzzTestName}

*/
