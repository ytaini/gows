#  go 模糊测试

https://colobu.com/2022/01/03/go-fuzzing/

Go 自 1.18开始在标准工具链中开始支持模糊测试(fuzzing)。

`Fuzzing` 是一种自动化的测试技术， 它不断的创建输入用来测试程序的bug。`Go fuzzing`使用覆盖率智能指导遍历被模糊化测试的代码，发现缺陷并报告给用户。由于模糊测试可以达到人类经常忽略的边缘场景，因此它对于发现安全漏洞和缺陷特别有价值。


下面是一个模糊测试的示例，突出标识了它的主要组件。
![](../assets/example1.png)

##  必备条件
下面是模糊测试必须遵循的规则：

1. 模糊测试必须是一个名称类似FuzzXxx的函数，仅仅接收一个*testing.F类型的参数,没有返回值

2. 模糊测试必须在*_test.go文件中才能运行

3. Fuzz target(模糊目标)必须是对(*testing.F).Fuzz的方法调用，参数是一个函数，并且此函数的第一个参数是*testing.T,然后是模糊参数(fuzzing argument)，没有返回值

4. 一个模糊测试中必须只有一个模糊目标

5. 所有的种子语料库(seed corpus)必须具有与模糊参数相同的类型,顺序相同。对(*testing.F).Add的调用也是如此, 同样也适用模糊测试中的testdata/fuzz中的语料文件

6. 模糊参数只能是下面的类型
    string, []byte
    int, int8, int16, int32/rune, int64
    uint, uint8/byte, uint16, uint32, uint64
    float32, float64
    bool

## 建议
下面的建议可以帮助你应付大部分的模糊测试：

- Fuzzing应该在支持覆盖率检测的平台上运行（目前是AMD64和ARM64），这样语料库可以在运行时有意义地增长，并且在进行Fuzzing时可以覆盖更多的代码。

- 模糊目标应该是快速的和确定性的，这样模糊引擎可以有效地工作，并且新的故障和代码覆盖率可以很容易地重现。

- 由于模糊目标是在多个工作进程之间以不确定的顺序并行调用的，因此模糊目标的状态不应持续到每次调用结束后，并且模糊目标的行为不应依赖于全局状态。

## 使用

运行模糊测试
1.执行如下命令来运行模糊测试:

这个方式只会使用种子语料库，而不会生成随机测试数据。通过这种方式可以用来验证种子语料库的测试数据是否可以测试通过。(fuzz test without fuzzing)
`$ go test`

如果reverse_test.go文件里有其它单元测试函数或者模糊测试函数，但是只想运行FuzzReverse模糊测试函数，
    我们可以执行`go test -run=FuzzReverse`命令。

注意：go test默认会执行所有以TestXxx开头的单元测试函数和以FuzzXxx开头的模糊测试函数，
    默认不运行以BenchmarkXxx开头的性能测试函数，如果我们想运行 benchmark用例，则需要加上 `-bench` 参数。

2.如果要基于种子语料库生成随机测试数据用于模糊测试，需要给go test命令增加 -fuzz参数。
    `go test -fuzz={FuzzTestName}`

## 定制

默认的go命令满足大部分模糊化测试场景，所以典型的一个模糊化运行命令应该如下:
`$ go test -fuzz={FuzzTestName}`

但是go命令提供一些定制化的参数来运行模糊测试，它们在cmd/go文档中都有介绍。

这里重点说几个:

`-fuzztime`: 执行的模糊目标在退出的时候要执行的总时间或者迭代次数，默认是用不结束
`-fuzzminimizetime`: 模糊目标在每次最少尝试时要执行的时间或者迭代次数，默认是60秒。你可以禁用最小化尝试，只需把这个参数设置为0
`-parallel`: 同时执行的模糊化数量，默认是`$GOMAXPROCS`。当前进行模糊化测试时设置`-cpu`无效果
`-keepfuzzing`: Keep running the fuzz test if a crasher is found. (default false)
`-race`: Enable data race detection(检测) while fuzzing. (default false)


## 术语

corpus entry: 语料库条目，可以在模糊化测试时使用。它可以是一个特殊格式的文件，也可以是(*testing.F).Add方法的调用
coverage guidance: 一种模糊化方法，它使用代码覆盖范围的扩展来确定哪些语料库条目值得保留以备将来使用。
fuzz target: 模糊测试的函数，在模糊化过程中对语料库条目和生成的值执行模糊测试。通过将函数传递给(*testing.F).Fuzz来提供模糊测试
fuzz test: 一个test文件中的函数，格式为FuzzXxx(*testing.F),用来执行模糊测试
fuzzing: 一种自动测试类型，它不断地修改程序的输入，以发现代码可能易受影响的问题，如bug或漏洞。
fuzzing arguments: 被传递到模糊目标的类型，并且可以被mutator修改变异
fuzzing engine: 一种管理模糊化的工具，包括维护语料库、调用变体、识别新覆盖范围和报告失败。
generated corpus: 语料库在模糊化过程中由模糊引擎随时间进行维护，以跟踪进度。它存储在$GOCACHE/fuzz中。
mutator: 模糊处理时使用的一种工具，它在将语料库条目传递给模糊目标之前对其进行随机操作。
package: 同一个文件夹下源代码的集合，会被编译在一起。Go代码组织的方式。ation.
seed corpus: 用户提供的用于模糊测试的语料库，可用于指导模糊引擎。它由模糊测试中的f.Add调用提供的语料库条目和包中testdata/fuzz/{FuzzTestName}目录中的文件组成。
test file: 格式为xxx_test.go类型的文件，包含测试、benchmark和模糊测试代码
vulnerability: 代码中的安全敏感弱点，可被攻击者利用。

## 其他
- 教程
    - 如果想了解Go模糊化介绍性的文章，请参考[the blog post](https://go.dev/blog/fuzz-beta)

- 文档
    - testing 包描述了模糊测试中用到的testing.F类型
    - [cmd/go](https://pkg.go.dev/cmd/go) 描述了和模糊测试相关的参数
- 技术细节
    - [Design draft](https://go.googlesource.com/proposal/+/master/design/draft-fuzzing.md)