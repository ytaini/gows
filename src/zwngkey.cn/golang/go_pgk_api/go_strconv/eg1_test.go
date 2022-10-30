/*
 * @Author: wzmiiiiii
 * @Date: 2022-10-30 15:21:54
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-10-30 17:39:50
 * @Description:

 */
package gostrconv

import (
	"fmt"
	"strconv"
	"testing"
)

// ParseFloat(s string, bitSize int)
func Test1(t *testing.T) {
	// bitSize: 32 for float32, or 64 for float64.
	// 	When bitSize=32, the result still has type float64,
	//  but it will be convertible to float32 without changing its value(但它可以转换为 float32 而不会改变它的值。)
	strconv.ParseFloat("22.3131", 32)
}

// FormatFloat(f float64, fmt byte, prec int, bitSize int) string
// 	FormatFloat 根据格式 fmt 和精度 prec 将浮点数 f 转换为字符串。
// 	假设原始值是从 bitSize 位的浮点值(32 表示 float32，64 表示 float64)获得的，它会对结果进行四舍五入。
//	格式 fmt 是'b'(-ddddp±ddd，二进制 index )、'e'(-d.dddde±dd，十进制 index )、
// 		'E'(-d.ddddE±dd，十进制 index )之一), 'f' (-ddd.dddd, 无 index ), 'g' ('e' 用于大 index , 'f' 否则),
// 		'G' ('E' 用于大 index , 'f' 否则), 'x'(-0xd.ddddp±ddd，十六进制分数和二进制 index )，
// 		或'X'(-0Xd.ddddP±ddd，十六进制分数和二进制 index )。
// 	精度 prec 控制由 'e'、'E'、'f', 'g'、'G'、'x' 和 'X' 格式打印的位数(不包括 index )。
// 		对于'e'、'E'、'f', 'x'和'X'，它是小数点后的位数。对于'g' 和'G'，它是有效数字的最大数量(删除尾随零)。
// 		特殊精度 -1 使用所需的最少位数，以便 ParseFloat 将准确返回 f。
func Test2(t *testing.T) {
	v := 3.1415926535

	s32 := strconv.FormatFloat(v, 'E', -1, 32)
	fmt.Printf("%T, %v\n", s32, s32)
}
