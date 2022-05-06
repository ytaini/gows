/*
 * @Author: zwngkey
 * @Date: 2022-05-07 03:32:33
 * @LastEditors: zwngkey 18390924907@163.com
 * @LastEditTime: 2022-05-07 03:55:05
 * @Description: go interface
 */
package gointerface

import "testing"

/*
	接口
		接口类型是Go中的一种很特别的类型。接口类型在Go中扮演着重要的角色。
			首先，在Go中，接口值可以用来包裹非接口值；然后，通过值包裹，反射和多态得以实现。

	什么是接口类型？
		一个接口类型指定了一个方法原型的集合。 换句话说，一个接口类型定义了一个方法集。
			事实上，我们可以把一个接口类型看作是一个方法集。 接口类型中指定的任何方法原型中的方法名称都不能为空标识符_。

		我们也常说一个接口类型规定了一个（用此接口类型指定的方法集表示的）行为规范。
*/

// ReadWriter是一个定义的接口类型。
type ReadWriter interface {
	Read(buf []byte) (n int, err error)
	Write(buf []byte) (n int, err error)
}

// Runnable是一个非定义接口类型的别名。
type Runnable = interface {
	Run()
}

/*
	特别地，一个没有指定任何方法原型的接口类型称为一个空接口类型。 下面是两个空接口类型：
		// 一个非定义空接口类型。
		interface{}

		// 类型I是一个定义的空接口类型。
		type I interface{}


	类型的方法集
		每个类型都有一个方法集。
			1.对于一个非接口类型，它的方法集由所有为它声明的（包括显式和隐式的）方法的原型组成。
			2.对于一个接口类型，它的方法集就是它所指定的方法集。

		为了解释方便，一个类型的方法集常常也可称为它的任何值的方法集。

		如果两个非定义接口类型指定的方法集是等价的，则这两个接口类型为同一个类型。
			但是请注意：不同代码包中的同名非导出方法名将总被认为是不同名的。


	什么是实现（implementation）？
		如果一个任意类型T的方法集为一个接口类型的方法集的超集，则我们说类型T实现了此接口类型。
			T可以是一个非接口类型，也可以是一个接口类型。

		实现关系在Go中是隐式的。两个类型之间的实现关系不需要在代码中显式地表示出来。Go中没有类似于implements的关键字。
			 Go编译器将自动在需要的时候检查两个类型之间的实现关系。

		一个接口类型总是实现了它自己。两个指定了一个相同的方法集的接口类型相互实现了对方。
*/
/*
	在下面的例子中，类型*Book、MyInt和*MyInt都拥有一个原型为About() string的方法，
		所以它们都实现了接口类型interface {About() string}。
*/
type Book struct {
	name string
	// 更多其它字段……
}

func (book *Book) About() string {
	return "Book: " + book.name
}

type MyInt int

func (MyInt) About() string {
	return "我是一个自定义整数类型"
}

func TestEg1(t *testing.T) {
	var a interface{ About() string } = MyInt(1)
	a.About()
}

/*
	注意：因为任何方法集都是一个空方法集的超集，所以任何类型都实现了任何空接口类型。 这是Go中的一个重要事实。

	隐式实现关系的设计使得一个声明在另一个代码包（包括标准库包）中的类型可以被动地实现一个用户代码包中的接口类型。
		比如，如果我们声明一个像下面这样的接口类型，则database/sql标准库包中声明的DB和Tx类型都实现了这个接口类型，
			因为它们都拥有此接口类型指定的三个方法。
		type DatabaseStorer interface {
			Exec(query string, args ...interface{}) (sql.Result, error)
			Prepare(query string) (*sql.Stmt, error)
			Query(query string, args ...interface{}) (*sql.Rows, error)
		}
*/
