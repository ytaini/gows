/*
 * @Author: wzmiiiiii
 * @Date: 2022-11-21 09:14:45
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-11-21 09:25:19
 * @Description:
 */
package main

import "fmt"

func main() {
	t := Teacher{}
	t.ShowA()
}

// 57e1 T main.(*People).ShowA
// 5851 T main.(*People).ShowB
// 66e4 T main.(*People).ShowC
// 6754 T main.(*People).ShowD
// 58b1 T main.People.ShowC
// 5911 T main.People.ShowD
type People struct {
}

func (p *People) ShowA() {
	fmt.Println("show A")
	p.ShowB() //()
}

func (p *People) ShowB() {
	fmt.Println("show b")
}

func (p People) ShowC() {
	fmt.Println("show c")
}

func (p People) ShowD() {
	fmt.Println("show d")
}

// 67c4 T main.(*Teacher).ShowA
// 5971 T main.(*Teacher).ShowB
// 67d4 T main.(*Teacher).ShowC
// 6844 T main.(*Teacher).ShowD
// 68b4 T main.Teacher.ShowC
// 59d1 T main.Teacher.ShowD
type Teacher struct {
	People
}

// 生成的(*Teacher).ShowA 方法如下所示:
// func (t *Teacher) ShowA() {
// 	(&(t.People)).ShowA()
// }

func (t *Teacher) ShowB() {
	fmt.Println("teacher show b")
}

func (t Teacher) ShowD() {
	fmt.Println("teach show d")
}

// 9591 T main.(*Student).ShowA
// 95a1 T main.(*Student).ShowB
// 95b1 T main.(*Student).ShowC
// 9621 T main.(*Student).ShowD
// 9691 T main.Student.ShowA
// 96f1 T main.Student.ShowB
// 9751 T main.Student.ShowC
// 97c1 T main.Student.ShowD
type Student struct {
	*People
}
