/*
 * @Author: zwngkey
 * @Date: 2022-04-30 00:27:45
 * @LastEditTime: 2022-05-02 15:19:10
 * @Description:
 */
package goother

import "fmt"

/*
	表达式、语句和简单语句
		简单说来，一个表达式表示一个值，而一条语句表示一个操作。
			但是在实际中，有些个表达式可能同时表示多个值，有些语句可能是由很多更基本的语句组成的。
			另外，根据场合不同，某些语句也可以被视为表达式。

		Go中，某些语句被称为简单语句。Go中各种流程控制语句的某些部分可能会被要求必须为简单语句或者表达式。

	表达式:
		字面量、变量和有名常量等均属于单值表达式
		运算符操作（不包括赋值部分）也都属于单值表达式。
		如果一个函数至少返回一个值，则它的调用属于表达式。
			特别的，如果此函数返回两个或两个以上的值，则对它的调用称为多值表达式。
			不返回任何结果的函数的调用不属于表达式。
		自定义函数（包括方法）本身都属于函数类型的值，所以它们都是单值表达式。
		通道的接收数据操作（不包括赋值部分）也属于表达式

	一些表达式的例子：
		123
		true
		B
		B + " language"
		a - 789
		a > 0 // 一个类型不确定布尔值
		f     // 一个类型为“func ()”的表达式

		// 下面这些即可以被视为简单语句，也可以被视为表达式。
		f() // 函数调用
		<-c // 通道接收操作

	Go中的一些表达式，包括刚提及的通道的接收数据操作，可能会表示可变数量的值。
		根据不同的场景，这样的表达式可能呈现为单值表达式，也可能呈现为多值表达式。


	简单语句:
		Go中有六种简单语句类型：
			变量短声明语句。
			纯赋值语句，包括x op= y这种运算形式。
			有返回结果的函数或方法调用，以及通道的接收数据操作。
			通道的发送数据操作。
			空语句.
			自增（x++）和自减（x--）语句。
		在Go中，自增和自减语句不能被当作表达式使用。

	一些简单语句的例子：
		c := make(chan bool) // 通道将在以后讲解
		a = 789
		a += 5
		a = f() // 这是一个纯赋值语句
		a++
		a--
		c <- true // 一个通道发送操作
		z := <-c  // 一个使用通道接收操作
				  // 做为源值的变量短声明语句

	一些非简单语句:
		标准变量声明语句。
		（有名）常量声明语句。
		类型声明语句。
		（代码）包引入语句。
		显式代码块。一个显式代码块起始于一个左大括号{，终止于一个右大括号}。 一个显式代码块中可以包含若干子语句。
		函数声明。 一个函数声明中可以包含若干子语句。
		流程控制跳转语句。
		函数返回（return）语句。
		延迟函数调用和协程创建语句。

	一些非简单语句：
		import "time"
		var a = 123
		const B = "Go"
		type Choice bool
		func f() int {
			for a < 10 {
				break
			}

			defer f()

			go fun(){}()

			// 这是一个显式代码块。
			{
				// ...
			}
			return 567
		}
*/
func Testeg1() {

}

/*
	Go中的流程控制
		Go中几种和特定种类的类型相关的流程控制代码块：
			容器类型相关的for-range循环代码块。
			接口类型相关的type-switch多条件分支代码块。
			通道类型相关的select-case多分支代码块。

		Go也支持break、continue和goto等跳转语句。 另外，Go还支持一个特有的fallthrough跳转语句。

		Go所支持的六种流程控制代码块中，除了if-else条件分支代码块，其它五种称为可跳出代码块。
			我们可以在一个可跳出代码块中使用break语句以跳出此代码块。

		我们可以在for和for-range两种循环代码块中使用continue语句提前结束一个循环步。
			除了这两种循环代码块，其它四种代码块称为分支代码块。

		每种流程控制的一个分支都属于一条语句。这样的语句常常会包含很多子语句。

		上面所提及的流程控制语句都属于狭义上的流程控制语句。
			协程、延迟函数调用、以及恐慌和恢复，以及并发同步技术属于广义上的流程控制语句。

*/
/*
	if-else条件分支控制代码块
		if InitSimpleStatement; Condition {
			// do something
		} else {
			// do something
		}
		InitSimpleStatement部分是可选的，如果它没被省略掉，则它必须为一条简单语句。
		 如果它被省略掉，它可以被视为一条空语句（简单语句的一种）。
		 在实际编程中，InitSimpleStatement常常为一条变量短声明语句。

		Condition必须为一个结果为布尔值的表达式（它被称为条件表达式）。
		 Condition部分可以用一对小括号括起来，但大多数情况下不需要。

		每个if-else流程控制包含一个隐式代码块，一个if分支显式代码块和一个可选的else分支代码块。
			这两个分支代码块内嵌在这个隐式代码块中。

		如果InitSimpleStatement语句是一个变量短声明语句，则在此语句中声明的变量被声明在外层的隐式代码块中。
*/

/*
	switch-case流程控制代码块
		switch InitSimpleStatement; CompareOperand0 {
			case CompareOperandList1:
				// do something
			case CompareOperandList2:
				// do something
			...
			case CompareOperandListN:
				// do something
			default:
				// do something
		}
		InitSimpleStatement部分必须为一条简单语句，它是可选的。

		CompareOperand0部分必须为一个表达式（如果它没被省略的话，见下）。
			此表达式的估值结果总是被视为一个类型确定值。
			如果它是一个类型不确定值，则它被视为类型为它的默认类型的类型确定值。
			因为这个原因，此表达式不能为类型不确定的nil值。 CompareOperand0常被称为switch表达式。

		每个CompareOperandListX部分（X表示1到N）必须为一个用（英文）逗号分隔开来的表达式列表。
		 其中每个表达式都必须能和CompareOperand0表达式进行比较。
		  每个这样的表达式常被称为case表达式。
		  如果其中case表达式是一个类型不确定值，则它必须能够自动隐式转化为对应的switch表达式的类型，否则编译将失败。

		当一个switch-case流程控制被执行到的时候，其中的简单语句InitSimpleStatement将率先被执行。
		随后switch表达式CompareOperand0将被估值（仅一次）。上面已经提到，此估值结果一定为一个类型确定值。
		 然后此结果值将从上到下从左到右和各个CompareOperandListX表达式列表中的各个case表达式逐个依次比较（使用==运算符）。
		  一旦发现某个表达式和CompareOperand0相等，比较过程停止并且此表达式对应的分支代码块将得到执行。
		  如果没有任何一个表达式和CompareOperand0相等，则default默认分支将得到执行（如果此分支存在的话）。

		注意，编译器可能会不允许一个switch-case流程控制中有任何两个case表达式可以在编译时刻确定相等。
			比如，当前的官方标准编译器（1.17版本）认为上例中的case 6, 7, 8一行是不合法的（如果此行未被注释掉）。
			但是其它编译器未必这么认为。
			事实上，当前的官方标准编译器允许重复的布尔case表达式在同一个switch-case流程控制中出现，
			 而gccgo（v8.2）允许重复的布尔和字符串类型的case表达式在同一个switch-case流程控制中出现。

		一条fallthrough语句必须为一个分支代码块中的最后一条语句。

		一条fallthrough语句不能出现在一个switch-case流程控制中的最后一个分支代码块中。

		一个switch-case流程控制中的InitSimpleStatement语句和CompareOperand0表达式都是可选的。
			如果CompareOperand0表达式被省略，则它被认为类型为bool类型的true值。

		Go中另外一个和其它语言的显著不同点是default分支不必一定为最后一个分支。
*/

/*

	goto :
		一条跳转标签声明之后必须立即跟随一条语句。
			如果此声明的跳转标签使用在一条goto语句中，则当此条goto语句被执行的时候，执行将跳转到此跳转标签声明后跟随的语句。

		一个跳转标签必须声明在一个函数体内，此跳转标签的使用可以在此跳转标签的声明之后或者之前，
			但是此跳转标签的使用不能出现在此跳转标签声明所处的最内层代码块之外。
*/
func Testeg2() {
	i := 0
	//Next标签声明在变量的作用域内.

Next: // 跳转标签声明
	fmt.Println(i)
	i++
	if i < 5 {
		goto Next // 跳转
	}
}

//如果一个跳转标签声明在某个变量的作用域内，则此跳转标签的使用不能出现在此变量的声明之前。
func Testeg3() {
	i := 0
Next:
	k := i + i //2. 放大变量k的作用域
	if i >= 5 {
		// error: goto Exit jumps over declaration of k
		goto Exit
	}
	// k := i + i //如果k变量声明在这,程序报错.
	// 1.创建一个显式代码块以缩小k的作用域。
	// {
	// 	k := i + i
	// 	fmt.Println(k)
	// }
	fmt.Println(k)
	i++
	goto Next
Exit: // 此标签声明在k的作用域内，但它的使用在k的作用域之外。
	fmt.Println("over")

}
