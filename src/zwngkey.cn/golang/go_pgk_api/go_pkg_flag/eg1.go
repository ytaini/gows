/*
 * @Author: zwngkey
 * @Date: 2022-05-02 17:30:26
 * @LastEditTime: 2022-05-02 19:12:18
 * @Description:	flag包API
 */

package gopkgflag

/*
	flag 包实现了命令行标签解析.

		定义标签需要使用flag.String(),Bool(),Int()等方法。

		还可以选择使用XXXVar()函数将标签绑定到指定变量中

		也可以传入自定义类型的标签，只要标签满足对应的值接口（接收指针指向的接收者）

		所有的标签都定义好了，就可以调用flag.Parse()来解析命令行参数并传入到定义好的标签了。

		在解析之后，标签对应的参数可以从flag.Args()获取到，它返回的slice，也可以使用flag.Arg(i)来获取单个参数。
			 参数列的索引是从0到flag.NArg()-1。

		命令行标签格式：
			-flag
			-flag=x
			-flag x  // 只有非boolean标签能这么用
		减号可以使用一个或者两个，效果是一样的。
		必须使用-flag=false的方式来解析boolean标签


		一个标签的解析会在下次出现第一个非标签参数（“-”就是一个非标签参数）的时候停止，或者是在终止符号“--”的时候停止。
*/
