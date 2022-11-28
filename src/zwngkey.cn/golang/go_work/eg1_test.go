/*
 * @Author: zwngkey
 * @Date: 2022-05-13 00:26:35
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-11-27 22:53:51
 * @Description:	go workspace mode
 */
package gowork

/*
	go.mod里使用replace指令的缺点:
		我们在提交某个module的代码到代码仓库时，需要删除最后的go.mod中的replace指令，
			否则其他开发者下载后会编译报错，因为他们本地可能没有util目录，或者util目录的路径和你的不一样。


	Go 1.18工作区模式
		该模式下不再需要在go.mod里使用replace指令，而是新增一个go.work文件。
			go.work文件与各模块在同一目录.

		go.work内容如下：
			go 1.18

			use (
				./main
				./util
			)


		在workspace目录下执行如下命令即可自动生成go.work
			$ go work init module_name1 module_name2 ...

		go work init后面跟的都是Module对应的目录。

		如果go命令执行的当前目录或者父目录有go.work文件，或者通过GOWORK环境变量指定了go.work的路径，那go命令就会进入工作区模式。
			在工作区模式下，go就可以通过go.work下的module路径找到并使用本地的module代码。

		注意：go.work不需要提交到代码仓库中，仅限于本地开发使用。


	总结
		为了解决多个有依赖的Module本地同时开发的问题，Go 1.18引入了工作区模式。

		工作区模式是对已有的Go Module开发模式的优化
*/
