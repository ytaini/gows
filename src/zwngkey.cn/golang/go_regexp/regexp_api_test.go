/*
 * @Author: wzmiiiiii
 * @Date: 2022-10-30 14:01:04
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-10-30 15:26:40
 * @Description:

	All类型的方法: 贪婪模式

	Submatch类型的方法: 还会返回分组的匹配项

	Find(All)?(String)?(Submatch)?(Index)?
	正则表达式提供了16个类似的查找方法，格式为Find(All)?(String)?(Submatch)?(Index)?。

		当方法名中有All的时候,它回继续查找非重叠的后续的字符串，返回slice
		当方法名中有String的时候,参数类型是字符串，否则是byte slice
		当方法名中有Submatch的时候, 还会返回子表达式/分组(capturing group)的匹配项
*/
package goregexp

import (
	"fmt"
	"regexp"
	"strconv"
	"testing"
)

// Compile 或者 MustCompile函数
// Compile() 或者 MustCompile()创建一个编译好的正则表达式对象。
// 假如正则表达式非法，那么Compile()方法回返回error,而MustCompile()编译非法正则表达式时不会返回error，而是回panic。
// 如果你想要很好的性能，不要在使用的时候才调用Compile()临时进行编译，而是预先调用Compile()编译好正则表达式对象
func Test3(t *testing.T) {
	// regexp1, _ := regexp.Compile(`^\d+.*m$`)
	// regexp2 := regexp.MustCompile(`^\d+.*m$`)
}

// MatchString与Match函数
// 如果在字符串str中有子字符串匹配这个正则表达式pattern，那么MatchString返回true
// err返回值:表示正则表达式是否合法.
func Test2(t *testing.T) {
	pattern := `^go`
	str := "go regexp"
	fmt.Println(regexp.MatchString(pattern, str))
	fmt.Println(regexp.Match(pattern, []byte(str)))

	str1 := "xigo regexp"
	fmt.Println(regexp.MatchString(pattern, str1))
	fmt.Println(regexp.Match(pattern, []byte(str1)))
}

// FindString()方法
// 用来返回第一个匹配的结果。
// 如果没有匹配的字符串，那么它回返回一个空的字符串，
// 当然如果你的正则表达式就是要匹配空字符串的话，它也会返回空字符串。
// 使用 FindStringIndex 或者 FindStringSubmatch可以区分这两种情况。
func Test4(t *testing.T) {
	str := "Golang expressions example Golang apae"

	regexp, _ := regexp.Compile("p([a-z]+)e")

	fmt.Println(regexp.FindString(str))

	// FindAllString(s string, n int):
	// n参数: 指定做几次匹配.-1表示所有的都匹配
	// 	 FindString方法的All版本，它返回指定匹配的字符串的slice。如果返回nil值代表没有匹配的字符串。
	fmt.Println(regexp.FindAllString(str, -1)) //[pre ple pae]
	fmt.Println(regexp.FindAllString(str, 1))  //[pre]
	fmt.Println(regexp.FindAllString(str, 2))  //[pre ple]
}

// FindStringIndex()
// 可以得到第一个匹配的字符串在整体字符串中的索引位置。如果没有匹配的字符串，它回返回nil值。
func Test5(t *testing.T) {
	str := "Golang regular expressions example"
	regexp, _ := regexp.Compile(`p([a-z]+)e`)
	match := regexp.FindStringIndex(str)

	// Match:  [15 18]
	fmt.Println("Match: ", match)
	// str[match[0]:match[1]] :找到子字符串
	fmt.Println("substring: ", str[match[0]:match[1]])

}

// FindStringSubmatch() 除了返回匹配的字符串外，还会返回分组的匹配项。
// 如果没有匹配项，则返回nil值。
func Test6(t *testing.T) {
	str := "Golang regular expressions example"

	regexp, err := regexp.Compile(`p([a-z]+)e`)

	match := regexp.FindStringSubmatch(str)

	fmt.Println("Match: ", match, " Error: ", err)
}

// ReplaceAllString
// 	用来替换所有匹配的字符串，返回一个源字符串的拷贝。
func Test7(t *testing.T) {
	str := "Golang regular expressions example"

	regexp, err := regexp.Compile(`p([a-z]+)e`)

	match := regexp.ReplaceAllString(str, "tutorial")

	fmt.Println("Match: ", match, " Error: ", err)

}

// ReplaceAllStringFunc(src string, repl func(string) string)
//  用来替换所有匹配的字符串，返回一个源字符串的拷贝。
// 	repl 函数: 对匹配到的子字符串进行格式化.并返回.
func Test8(t *testing.T) {
	searchIn := "John: 2578.34 William: 4567.23 Steve: 5632.18"
	pat := `\d+[.]\d+`
	re, _ := regexp.Compile(pat)

	//
	f := func(s string) string {
		v, _ := strconv.ParseFloat(s, 32)
		return strconv.FormatFloat(v*2, 'f', 2, 32)
	}

	str2 := re.ReplaceAllStringFunc(searchIn, f)
	fmt.Println(str2)
}
