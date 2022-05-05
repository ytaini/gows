# go官方工具链
- `go run`子命令并不推荐在正式的大项目中使用。`go run`子命令只是一种方便的方式来运行简单的Go程序。 对于正式的项目，最好使用`go build`或者`go install`子命令构建可执行程序文件来运行Go程序。

  

- 支持Go模块特性的Go项目的根目录下需要一个`go.mod`文件。此文件可以使用`go mod init`子命令来生成

  

- 上面提到的三个go子命令（go run、go build和go install） 将只会输出代码语法错误。它们不会输出可能的代码逻辑错误（即警告）。 

  go vet子命令可以用来检查可能的代码逻辑错误（即警告）

  

- 我们应该常常使用go fmt子命令来用同一种代码风格格式化Go代码。

  

- 我们可以使用go test子命令来运行单元和基准测试用例。

  

- 我们可以使用go doc子命令来（在终端中）查看Go代码库包的文档。

  

> 强烈推荐让你的Go项目支持Go模块特性来简化依赖管理。
- 对一个支持Go模块特性的项目：
  - go mod init example.com/myproject命令可以用来在当前目录中生成一个go.mod文件。 当前目录将被视为一个名为example.com/myproject的模块（即当前项目）的根目录。 此go.mod文件将被用来记录当前项目需要的依赖模块和版本信息。 我们可以手动编辑或者使用go子命令来修改此文件。
  - go mod tidy命令用来通过扫描当前项目中的所有代码来添加未被记录的依赖至go.mod文件或从go.mod文件中删除不再被使用的依赖。
  - go get命令用拉添加、升级、降级或者删除单个依赖。此命令不如go mod tidy命令常用。


- 从Go官方工具链1.16版本开始，我们可以运行go install example.com/program@latest来安装一个第三方Go程序的最新版本（至GOBIIN目录）。 在Go官方工具链1.16版本之前，对应的命令是go get -u example.com/program（现在已经被废弃而不再推荐被使用了）。
