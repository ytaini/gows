/*
 * @Author: zwngkey
 * @Date: 2022-04-30 14:05:25
 * @LastEditTime: 2022-04-30 14:51:38
 * @Description:
 */

package gotest

/*
	test 包的工作原理：
		在执行 go test 时会编译每个包和所有后缀匹配 *_test.go 命名的文件（这些测试文件包括一些单元测试和基准测试），
		 链接和执行生成的二进制程序, 然后打印每一个测试函数的输出日志。

	go test
		当go test以包列表模式运行时，go test会缓存成功的包的测试结果以避免不必要的重复测试。
			当然，有时候我们测试的时候并不喜欢有缓存，我们可以手动禁用缓存。可以通过下列方式禁用缓存：

			1. 带上-count=1参数禁用缓存。
				如，执行下面命令测试，便会禁用缓存测试结果
				go test -v -count=1 filename_test.go

			2. 手动清除测试缓存
				除了在执行测试命令的时候加上禁用缓存参数，我们还可以执行下面的命令手动清除缓存，
					需要注意的是，每次都得清除，不然下次执行的还是上次的结果。
				go clean -testcache

		vscode设置方法:
			在settings.json中 写入 "go.testFlags": ["-v","-count=1"] 项.
*/
