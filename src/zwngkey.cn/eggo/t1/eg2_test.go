/*
 * @Author: zwngkey
 * @Date: 2021-12-25 15:43:32
 * @LastEditors: zwngkey 18390924907@163.com
 * @LastEditTime: 2022-05-13 06:55:26
 * @Description:
 */
package eggot1

import (
	"fmt"
	"testing"
)

//go组合式继承.

//(p *People)
// ShowA,ShowB,ShowC,ShowD
//(p People)
// ShowC,ShowD
type People struct {
}

func (p *People) ShowA() {
	fmt.Println("show A")
	p.ShowB() //()
}

func (p *People) ShowB() {
	fmt.Println("show b")
}

// func (p People) ShowC() {
// 	fmt.Println("show c")
// }

// func (p People) ShowD() {
// 	fmt.Println("show d")
// }

//(t *Teacher)
// ShowA,ShowB,ShowC,ShowD
//(t Teacher)
// ShowC,ShowD
type Teacher struct {
	People
}

func (t *Teacher) ShowB() {
	fmt.Println("teacher show b")
}

func (t Teacher) ShowD() {
	fmt.Println("teach show d")
}

//(s *Student)
// ShowA,ShowB,ShowC,ShowD
//(s Student)
// ShowA,ShowB,ShowC,ShowD
type Student struct {
	*People
}

func Test2(te *testing.T) {
	t := Teacher{}
	t.ShowA() //底层 (&t).ShowA() --> (*Teacher).ShowA(&t) -> (*People).ShowA(&t.People)
	//(&(&t).People).ShowA()
	(&t.People).ShowA()
	(&(t.People)).ShowA()
	(&t).ShowA() //底层

	t.ShowB() //teacher show b

	// t.ShowC() //show c

	// t.ShowD() //teach show d

}
