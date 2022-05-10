/*
 * @Author: zwngkey
 * @Date: 2022-05-10 18:05:29
 * @LastEditors: zwngkey 18390924907@163.com
 * @LastEditTime: 2022-05-10 21:21:59
 * @Description:
 */
package gounsafe

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

/*
	如何正确地使用非类型安全指针？
		使用模式一：将类型*T1的一个值转换为非类型安全指针值，然后将此非类型安全指针值转换为类型*T2。

		利用非类型安全指针相关的转换规则，我们可以将一个*T1值转换为类型*T2，其中T1和T2为两个任意类型。
			 然而，我们只有在T1的尺寸不小于T2并且此转换具有实际意义的时候才应该实施这样的转换。

		使用模式二：将一个非类型安全指针值转换为一个uintptr值，然后使用此uintptr值。

		使用模式三：将一个非类型安全指针转换为一个uintptr值，然后此uintptr值参与各种算术运算，
			再将算术运算的结果uintptr值转回非类型安全指针。

		使用模式四：将非类型安全指针值转换为uintptr值并传递给syscall.Syscall函数调用。

		使用模式五：将reflect.Value.Pointer或者reflect.Value.UnsafeAddr方法的uintptr返回值立即转换为非类型安全指针。

		使用模式六：将一个reflect.SliceHeader或者reflect.StringHeader值的Data字段转换为非类型安全指针，以及其逆转换。


*/
// 将一个[]MyString值和一个[]string值转换为对方的类型。结果切片和被转换的切片将共享底层元素。
//（这样的转换是不可能通过安全的方式来实现的。）
func TestEg1(t *testing.T) {
	type Mystring string

	ms := []Mystring{"c", "go", "java"}

	// ss := []string(ms)

	ss := *(*[]string)(unsafe.Pointer(&ms))

	ss[1] = "rust"

	ms = *(*[]Mystring)(unsafe.Pointer(&ss))

	ss = unsafe.Slice((*string)(&ms[0]), len(ms))

	fmt.Println(ss)
}

//此模式在实践中的另一个应用是将一个不再使用的字节切片转换为一个字符串（从而避免对底层字节序列的一次开辟和复制）
func ByteSlice2String(bs []byte) string {
	return *(*string)(unsafe.Pointer(&bs))
}

type T struct {
	x bool
	y [3]int16
}

const N = unsafe.Offsetof(T{}.y)
const M = unsafe.Sizeof(T{}.y[0])

func TestEg2(te *testing.T) {
	t := T{y: [3]int16{123, 456, 789}}
	p := unsafe.Pointer(&t)
	// "uintptr(p) + N + M + M"为t.y[2]的内存地址。
	// 其实，对于这样地址加减运算，更推荐使用unsafe.Add函数来完成。
	ty2 := (*int16)(unsafe.Pointer(uintptr(p) + N + M + M))

	fmt.Println(*ty2)
}

/*
	我们可以将一个字符串的指针值转换为一个*reflect.StringHeader指针值，从而可以对此字符串的内部进行修改。
		类似地，我们可以将一个切片的指针值转换为一个*reflect.SliceHeader指针值，从而可以对此切片的内部进行修改。
*/
func TestEg3(t *testing.T) {
	a := [...]byte{'G', 'o', 'l', 'a', 'n', 'g'}
	_ = a
	s := "Java"
	hdr := (*reflect.StringHeader)(unsafe.Pointer(&s))
	hdr.Data = uintptr(unsafe.Pointer(&a))
	hdr.Len = len(a)
	fmt.Println(s)
	// 现在，字符串s和切片a共享着底层的byte字节序列，
	// 从而使得此字符串中的字节变得可以修改。
	a[2], a[3], a[4], a[5] = 'o', 'g', 'l', 'e'
	fmt.Println(s) // Google
}

func TestEg4(t *testing.T) {
	a := [6]byte{'G', 'o', '1', '0', '1'}
	bs := []byte("Golang")
	hdr := (*reflect.SliceHeader)(unsafe.Pointer(&bs))
	hdr.Data = uintptr(unsafe.Pointer(&a))

	hdr.Len = 2
	hdr.Cap = len(a)
	fmt.Printf("%s\n", bs) // Go
	bs = bs[:cap(bs)]
	fmt.Printf("%s\n", bs) // Go101
}
func String2ByteSlice(str string) (bs []byte) {
	strHdr := (*reflect.StringHeader)(unsafe.Pointer(&str))
	sliceHdr := (*reflect.SliceHeader)(unsafe.Pointer(&bs))
	sliceHdr.Data = strHdr.Data
	sliceHdr.Cap = strHdr.Len
	sliceHdr.Len = strHdr.Len
	return
}
