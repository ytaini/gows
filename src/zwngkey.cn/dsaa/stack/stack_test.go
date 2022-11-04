/*
 * @Author: wzmiiiiii
 * @Date: 2022-10-30 00:13:39
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-10-30 00:13:39
 * @Description:
 */

package stack

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	stack := New[int]()
	fmt.Println(stack.Pop())
	fmt.Printf("stack.IsEmpty(): %v\n", stack.IsEmpty())
	stack.Push(12)
	stack.Push(13)
	stack.Print()
	fmt.Println("----")
	fmt.Printf("stack.IsEmpty(): %v\n", stack.IsEmpty())
	fmt.Println(stack.Pop())

	fmt.Println("----")
	fmt.Println(stack.Peek())
	fmt.Println(stack.Pop())

	fmt.Println("----")
	stack.Print()
	fmt.Println(stack.Peek())
}