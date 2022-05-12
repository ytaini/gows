/*
 * @Author: zwngkey
 * @Date: 2022-05-10 01:14:49
 * @LastEditors: zwngkey 18390924907@163.com
 * @LastEditTime: 2022-05-10 16:25:45
 * @Description: 组合式继承
 */
package gostruct

import (
	"fmt"
	"testing"
)

type People struct {
}

func (p *People) ShowA() {
	fmt.Println("show A")
	p.ShowB()
}

func (p *People) ShowB() {
	fmt.Println("show b")
}

type Teacher struct {
	People
}

func (t *Teacher) ShowB() {
	fmt.Println("teacher show b")
}

type Student struct {
	*People
}

func TestEg31(te *testing.T) {
	t := Teacher{}
	t.ShowA() //&t.People.ShowA()
}
