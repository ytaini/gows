[TOC]

# go test工具

`go test [build/test flags] [packages] [build/test flags & test binary flags]`

Go语言中的测试依赖`go test`命令。编写测试代码和编写普通的Go代码过程是类似的，并不需要学习新的语法、规则或工具。


`go test`命令是一个按照一定约定和组织的测试代码的驱动程序。
在包目录内，所有以`_test.go`为后缀名的源代码文件都是`go test`测试的一部分，不会被`go build`编译到最终的可执行文件中。

在`*_test.go`文件中有三种类型的函数，单元测试函数、基准测试函数和示例函数。

|   类型   |         格式          |              作用              |
| :------: | :-------------------: | :----------------------------: |
| 测试函数 |   函数名前缀为Test    | 测试程序的一些逻辑行为是否正确 |
| 基准函数 | 函数名前缀为Benchmark |         测试函数的性能         |
| 示例函数 |  函数名前缀为Example  |       为文档提供示例文档       |

`go test`命令会遍历所有的`*_test.go`文件中符合上述命名规则的函数，然后生成`一个临时的main包`用于调用相应的测试函数，然后构建并运行、报告测试结果，最后清理测试中生成的临时文件。

参数:
`-timeout` : 5m/5s/5h等等. 设置总超时时间
用法: `go test -timeout 5m`

`-count n` : 禁用缓存。
用法: `go test -count=1`

可以为`go test`命令添加`-v`参数，查看测试函数名称和运行时间
```shell
$ go test -v
```

可以在`go test`命令后添加`-run`参数，它对应一个正则表达式，只有函数名匹配上的测试函数才会被go test命令执行。
```shell
$ go test -v -run="regexp"
```

`-short` : 跳过某些测试用例。
用法: `go test -short`

<br>
<br>

# 单元测试

每个测试函数必须导入testing包，测试函数的基本格式（签名）如下：
```go
func TestName(t *testing.T){
    // ...
}
```

测试函数的名字必须以Test开头，可选的后缀名必须以大写字母开头，举几个例子：
```go
func TestAdd(t *testing.T){ ... }
func TestSum(t *testing.T){ ... }
func TestLog(t *testing.T){ ... }
```


其中参数t用于报告测试失败和附加的日志信息。 testing.T的拥有的方法如下：

当一个单元测试的测试函数返回时， 又或者当一个测试函数调用 FailNow 、 Fatal 、 Fatalf 、 SkipNow 、 Skip 或者 Skipf 中的任意一个时， 该测试即宣告结束。 跟 Parallel 方法一样， 以上提到的这些方法只能在运行测试函数的 goroutine 中调用。

至于其他报告方法， 比如 Log 以及 Error 的变种， 则可以在多个 goroutine 中同时进行调用。
```go
func (t *T) Deadline() (deadline time.Time, ok bool)
// The ok result is false if the -timeout flag indicates(表示) “no timeout” (0).
// 返回超时的时间点.

func (t *T) Helper()
// Helper 的作用是标记一个函数为测试辅助函数。打印文件和行信息时，将跳过该功能。助手可以从多个程序同时调用。
// 这样的话，该函数将不会在测试日志输出文件名和行号信息时出现。
// 当 go testing 系统在查找调用栈帧的时候，通过 Helper 标记过的函数将被略过，报错时将输出帮助函数调用者的信息，而不是帮助函数的内部信息.
// 这个函数的用途在于削减日志输出中（尤其是在打印调用栈帧信息时）的杂音。


func (t *T) Cleanup(f func())
// Cleanup注册一个函数，以便在测试（或子测试）及其所有子测试完成后调用。
// Cleanup将按照最后添加、最先调用的顺序调用。

func (t *T) Fail()  // 将当前测试标识为失败，但是仍继续执行该测试。
func (t *T) FailNow()
// 将当前测试标识为失败并停止执行该测试，在此之后，测试过程将在下一个测试或者下一个基准测试中继续。
// FailNow 必须在运行测试函数或者基准测试函数的 goroutine 中调用，而不能在测试期间创建的 goroutine 中调用。调用 FailNow 不会导致其他 goroutine 停止。


func (t *T) Failed() bool //Failed 用于报告测试函数是否已失败。


func (t *T) Error(args ...interface{})
func (t *T) Errorf(format string, args ...interface{}) // 调用 Error 相当于在调用 Log 之后调用 Fail 。

func (t *T) Fatal(args ...interface{})
func (t *T) Fatalf(format string, args ...interface{}) // 调用 Fatal 相当于在调用 Log 之后调用 FailNow 。


func (t *T) Log(args ...interface{})
func (t *T) Logf(format string, args ...interface{})
// Log 使用与 Printf 相同的格式化语法对它的参数进行格式化，然后将格式化后的文本记录到错误日志里面。 
// 如果输入的格式化文本最末尾没有出现新行，那么将一个新行添加到格式化后的文本末尾。
// 1）对于测试来说，Logf 产生的格式化文本只会在测试失败或者设置了 -test.v 标志的情况下被打印出来； 
// 2）对于基准测试来说，为了避免 -test.v 标志的值对测试的性能产生影响，Logf 产生的格式化文本总会被打印出来。


func (t *T) Name() string // 返回正在运行的测试或基准测试的名字。

func (t *T) Parallel() // Parallel 用于表示当前测试只会与其他带有 Parallel 方法的测试并行进行测试。

func (t *T) Run(name string, f func(t *T)) bool
// 执行名字为 name 的子测试 f ，并报告 f 在执行过程中是否出现了任何失败。Run 将一直阻塞直到 f 的所有并行测试执行完毕。
// Run 可以在多个 goroutine 里面同时进行调用，但这些调用必须发生在 t 的外层测试函数（outer test function）返回之前。

func (t *T) Setenv(key, value string) 
// 这个方法不能用于并行测试。
// Setenv调用os.Setenv(key.value），并在测试后使用Cleanup()将环境变量恢复到其原始值。

func (t *T) TempDir() string 
// TempDir返回一个临时目录供测试使用。
// 当测试及其所有子测试完成后，Cleanup()会自动删除目录。
// 每次后续调用t.TempDir都会返回一个唯一的目录；如果目录创建失败，TempDir将通过调用Fatal来终止测试。


func (t *T) Skip(args ...interface{})
func (t *T) Skipf(format string, args ...interface{}) //调用 Skipf 相当于在调用 Logf 之后调用 SkipNow 。

func (t *T) SkipNow()
// 将当前测试标识为“被跳过”并停止执行该测试。 
// 如果一个测试在失败（参考 Error、Errorf 和 Fail）之后被跳过了， 那么它还是会被判断为是“失败的”。
// 在停止当前测试之后，测试过程将在下一个测试或者下一个基准测试中继续，具体请参考 FailNow 。
// SkipNow 必须在运行测试的 goroutine 中进行调用，而不能在测试期间创建的 goroutine 中调用。 调用 SkipNow 不会导致其他 goroutine 停止。

func (t *T) Skipped() bool //Skipped 用于报告测试函数是否已被跳过。
```

<br>
<br>

### 测试组

```go
func TestSplit(t *testing.T) {
   // 定义一个测试用例类型
	type test struct {
		input string
		sep   string
		want  []string // 预期值
	}
	// 定义一个存储测试用例的切片
	tests := []test{
		{input: "a:b:c", sep: ":", want: []string{"a", "b", "c"}},
		{input: "a:b:c", sep: ",", want: []string{"a:b:c"}},
		{input: "abcd", sep: "bc", want: []string{"a", "d"}},
		{input: "沙河有沙又有河", sep: "沙", want: []string{"河有", "又有河"}},
	}
	// 遍历切片，逐一执行测试用例
	for _, tc := range tests {
		got := Split(tc.input, tc.sep)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("expected:%v, got:%v", tc.want, got)
		}
	}
}
```

<br>
<br>

### 跳过某些测试用例
为了节省时间支持在单元测试时跳过某些耗时的测试用例。
```go
func Test(t *testing.T) {
	if testing.Short() {   // Short reports whether the -test.short flag is set.
		t.Skip("short模式下会跳过该测试用例")
	}
	fmt.Println(t.Name())
}
```
当执行`go test -v -short`时就不会执行上面的`Test`测试用例。

输出:
```
=== RUN   Test
    eg2_test.go:16: short模式下会跳过该测试用例
--- SKIP: Test (0.00s)
```

<br>
<br>

### 子测试

Go1.7+中新增了子测试，支持在测试函数中使用t.Run执行一组测试用例，这样就不需要为不同的测试数据定义多个测试函数了。

我们可以按照如下方式使用`t.Run`执行子测试：

注意: 我们都知道可以通过`-run=RegExp`来指定运行的测试用例，还可以通过`/`来指定要运行的子测试用例，例如：`go test -v -run=Split/simple`只会运行`simple`对应的子测试用例。
```go
func TestSplit(t *testing.T) {
	type test struct { // 定义test结构体
		input string
		sep   string
		want  []string
	}
	tests := map[string]test{ // 测试用例使用map存储
		"simple":      {input: "a:b:c", sep: ":", want: []string{"a", "b", "c"}},
		"wrong sep":   {input: "a:b:c", sep: ",", want: []string{"a:b:c"}},
		"more sep":    {input: "abcd", sep: "bc", want: []string{"a", "d"}},
		"leading sep": {input: "沙河有沙又有河", sep: "沙", want: []string{"河有", "又有河"}},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := Split(tc.input, tc.sep)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("expected:%#v, got:%#v", tc.want, got)
			}
		})
	}
}
```



<br>
<br>


### 帮助函数

对一些重复的逻辑，抽取出来作为公共的帮助函数(helpers)，可以增加测试代码的可读性和可维护性。 借助帮助函数，可以让测试用例的主逻辑看起来更清晰。

例如，我们可以将创建子测试的逻辑抽取出来：

```go
// calc_test.go
package main

import "testing"

type calcCase struct{ A, B, Expected int }

func createMulTestCase(t *testing.T, c *calcCase) {
	// t.Helper()
	if ans := Mul(c.A, c.B); ans != c.Expected {
		t.Fatalf("%d * %d expected %d, but %d got",
			c.A, c.B, c.Expected, ans)
	}

}

func TestMul(t *testing.T) {
	createMulTestCase(t, &calcCase{2, 3, 6})
	createMulTestCase(t, &calcCase{2, -3, -6})
	createMulTestCase(t, &calcCase{2, 0, 1}) // wrong case
}
```

在这里，我们故意创建了一个错误的测试用例，运行 go test，用例失败，会报告错误发生的文件和行号信息：

```shell
$ go test
--- FAIL: TestMul (0.00s)
    calc_test.go:11: 2 * 0 expected 1, but 0 got
FAIL
exit status 1
FAIL    example 0.007s
```

可以看到，错误发生在第11行，也就是帮助函数 createMulTestCase 内部。18, 19, 20行都调用了该方法，我们第一时间并不能够确定是哪一行发生了错误。有些帮助函数还可能在不同的函数中被调用，报错信息都在同一处，不方便问题定位。因此，Go 语言在 1.9 版本中引入了 t.Helper()，用于标注该函数是帮助函数，报错时将输出帮助函数调用者的信息，而不是帮助函数的内部信息。

修改 createMulTestCase，调用 t.Helper()

```go
func createMulTestCase(c *calcCase, t *testing.T) {
    t.Helper()
	t.Run(c.Name, func(t *testing.T) {
		if ans := Mul(c.A, c.B); ans != c.Expected {
			t.Fatalf("%d * %d expected %d, but %d got",
				c.A, c.B, c.Expected, ans)
		}
	})
}
```
运行 go test，报错信息如下，可以非常清晰地知道，错误发生在第 20 行。

```shell
$ go test
--- FAIL: TestMul (0.00s)
    calc_test.go:20: 2 * 0 expected 1, but 0 got
FAIL
exit status 1
FAIL    example 0.006s
```

关于 helper 函数的 2 个建议：

- 不要返回错误， 帮助函数内部直接使用 t.Error 或 t.Fatal 即可，在用例主逻辑中不会因为太多的错误处理代码，影响可读性。
- 调用 t.Helper() 让报错信息更准确，有助于定位。


<br>
<br>

### setup 和 teardown

如果在同一个测试文件中，每一个测试用例运行前后的逻辑是相同的，一般会写在 setup 和 teardown 函数中。例如执行前需要实例化待测试的对象，如果这个对象比较复杂，很适合将这一部分逻辑提取出来；执行后，可能会做一些资源回收类的工作，例如关闭网络连接，释放文件等。标准库 testing 提供了这样的机制：

```go
func setup() {
	fmt.Println("Before all tests")
}

func teardown() {
	fmt.Println("After all tests")
}

func Test1(t *testing.T) {
	fmt.Println("I'm test1")
}

func Test2(t *testing.T) {
	fmt.Println("I'm test2")
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}
```

- 在这个测试文件中，包含有2个测试用例，Test1 和 Test2。
- 如果测试文件中包含函数 `TestMain`，那么生成的测试将调用 `TestMain(m)`，而不是直接运行测试。
- 调用 `m.Run()` 触发所有测试用例的执行，并使用 `os.Exit()` 处理返回的状态码，如果不为0，说明有用例失败。
- 因此可以在调用 m.Run() 前后做一些额外的准备(setup)和回收(teardown)工作。


执行 go test，将会输出
```shell
$ go test
Before all tests
I'm test1
I'm test2
PASS
After all tests
ok      example 0.006s
```


### 子测试的Setup与Teardown
有时候我们可能需要为每个测试集设置Setup与Teardown，也有可能需要为每个子测试设置Setup与Teardown。下面我们定义两个函数工具函数如下：

```go
// 测试集的Setup与Teardown
func setupTestCase(t *testing.T) func(t *testing.T) {
	t.Log("如有需要在此执行:测试之前的setup")
	return func(t *testing.T) {
		t.Log("如有需要在此执行:测试之后的teardown")
	}
}

// 子测试的Setup与Teardown
func setupSubTest(t *testing.T) func(t *testing.T) {
	t.Log("如有需要在此执行:子测试之前的setup")
	return func(t *testing.T) {
		t.Log("如有需要在此执行:子测试之后的teardown")
	}
}
```

使用方式如下：
```go
func TestSplit(t *testing.T) {
	type test struct { // 定义test结构体
		input string
		sep   string
		want  []string
	}
	tests := map[string]test{ // 测试用例使用map存储
		"simple":      {input: "a:b:c", sep: ":", want: []string{"a", "b", "c"}},
		"wrong sep":   {input: "a:b:c", sep: ",", want: []string{"a:b:c"}},
		"more sep":    {input: "abcd", sep: "bc", want: []string{"a", "d"}},
		"leading sep": {input: "沙河有沙又有河", sep: "沙", want: []string{"", "河有", "又有河"}},
	}
	teardownTestCase := setupTestCase(t) // 测试之前执行setup操作
	defer teardownTestCase(t)            // 测试之后执行teardown操作

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			teardownSubTest := setupSubTest(t) // 子测试之前执行setup操作
			defer teardownSubTest(t)           // 测试之后执行teardown操作
			got := Split(tc.input, tc.sep)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("expected:%#v, got:%#v", tc.want, got)
			}
		})
	}
}
```

测试结果如下：
```shell
split $ go test -v
=== RUN   TestSplit
=== RUN   TestSplit/simple
=== RUN   TestSplit/wrong_sep
=== RUN   TestSplit/more_sep
=== RUN   TestSplit/leading_sep
--- PASS: TestSplit (0.00s)
    split_test.go:71: 如有需要在此执行:测试之前的setup
    --- PASS: TestSplit/simple (0.00s)
        split_test.go:79: 如有需要在此执行:子测试之前的setup
        split_test.go:81: 如有需要在此执行:子测试之后的teardown
    --- PASS: TestSplit/wrong_sep (0.00s)
        split_test.go:79: 如有需要在此执行:子测试之前的setup
        split_test.go:81: 如有需要在此执行:子测试之后的teardown
    --- PASS: TestSplit/more_sep (0.00s)
        split_test.go:79: 如有需要在此执行:子测试之前的setup
        split_test.go:81: 如有需要在此执行:子测试之后的teardown
    --- PASS: TestSplit/leading_sep (0.00s)
        split_test.go:79: 如有需要在此执行:子测试之前的setup
        split_test.go:81: 如有需要在此执行:子测试之后的teardown
    split_test.go:73: 如有需要在此执行:测试之后的teardown
=== RUN   ExampleSplit
--- PASS: ExampleSplit (0.00s)
PASS
ok      github.com/Q1mi/studygo/code_demo/test_demo/split       0.006s
```

<br>
<br>

### 表格驱动测试

> 介绍

表格驱动测试不是工具、包或其他任何东西，它只是编写更清晰测试的一种方式和视角。

编写好的测试并非易事，但在许多情况下，表格驱动测试可以涵盖很多方面：表格里的每一个条目都是一个完整的测试用例，包含输入和预期结果，有时还包含测试名称等附加信息，以使测试输出易于阅读。

使用表格驱动测试能够很方便的维护多个测试用例，避免在编写单元测试时频繁的复制粘贴。

表格驱动测试的步骤通常是定义一个测试用例表格，然后遍历表格，并使用t.Run对每个条目执行必要的测试。

> 示例
官方标准库中有很多表格驱动测试的示例，例如fmt包中便有如下测试代码：
```go
var flagtests = []struct {
	in  string
	out string
}{
	{"%a", "[%a]"},
	{"%-a", "[%-a]"},
	{"%+a", "[%+a]"},
	{"%#a", "[%#a]"},
	{"% a", "[% a]"},
	{"%0a", "[%0a]"},
	{"%1.2a", "[%1.2a]"},
	{"%-1.2a", "[%-1.2a]"},
	{"%+1.2a", "[%+1.2a]"},
	{"%-+1.2a", "[%+-1.2a]"},
	{"%-+1.2abc", "[%+-1.2a]bc"},
	{"%-1.2abc", "[%-1.2a]bc"},
}
func TestFlagParser(t *testing.T) {
	var flagprinter flagPrinter
	for _, tt := range flagtests {
		t.Run(tt.in, func(t *testing.T) {
			s := Sprintf(tt.in, &flagprinter)
			if s != tt.out {
				t.Errorf("got %q, want %q", s, tt.out)
			}
		})
	}
}
```
通常表格是匿名结构体切片，可以定义结构体或使用已经存在的结构进行结构体数组声明。name属性用来描述特定的测试用例。
 
<br>
<br>

### 并行测试

表格驱动测试中通常会定义比较多的测试用例，而Go语言又天生支持并发，所以很容易发挥自身并发优势将表格驱动测试并行化。 想要在单元测试过程中使用并行测试，可以像下面的代码示例中那样通过添加t.Parallel()来实现。

```go
func TestSplitAll(t *testing.T) {
	t.Parallel()  // 将 TLog 标记为能够与其他测试并行运行
	// 定义测试表格
	// 这里使用匿名结构体定义了若干个测试用例
	// 并且为每个测试用例设置了一个名称
	tests := []struct {
		name  string
		input string
		sep   string
		want  []string
	}{
		{"base case", "a:b:c", ":", []string{"a", "b", "c"}},
		{"wrong sep", "a:b:c", ",", []string{"a:b:c"}},
		{"more sep", "abcd", "bc", []string{"a", "d"}},
		{"leading sep", "沙河有沙又有河", "沙", []string{"", "河有", "又有河"}},
	}
	// 遍历测试用例
	for _, tt := range tests {
		tt := tt  // 注意这里重新声明tt变量（避免多个goroutine中使用了相同的变量）
		t.Run(tt.name, func(t *testing.T) { // 使用t.Run()执行子测试
			t.Parallel()  // 将每个测试用例标记为能够彼此并行运行
			got := Split(tt.input, tt.sep)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("expected:%#v, got:%#v", tt.want, got)
			}
		})
	}
}
```
这样我们执行go test -v的时候就会看到每个测试用例并不是按照我们定义的顺序执行，而是互相并行了。

<br>
<br>


### 使用工具生成测试代码

社区里有很多自动生成表格驱动测试函数的工具，比如gotests等，很多编辑器如Goland也支持快速生成测试文件。这里简单演示一下gotests的使用。

> 安装

```
go get -u github.com/cweill/gotests/...
```
如果不确定，按 F1 或者 `Ctrl + Shift + P`唤出指令面板（Command Palette），输入 `Go: Generate Unit Tests` ，如果出现三个提示 For File ，For Function，For Package ，就说明安装好了。

如果没有，唤出指令面板，执行 `Go: Install/Update Tools`，找到 `gotests` 安装即可。

> 使用

将焦点（也就是输入光标）停留在 某个 函数上，F1 唤出 指令面板，执行 Go: Generate Unit Tests For Function ，测试代码就自动生成了。

或者执行

```
gotests -all -w split.go
```
上面的命令表示，为split.go文件的所有函数生成测试代码至split_test.go文件（目录下如果事先存在这个文件就不再生成）

```go
package base_demo

import (
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) {
	type args struct {
		s   string
		sep string
	}
	tests := []struct {
		name       string
		args       args
		wantResult []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := Split(tt.args.s, tt.args.sep); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("Split() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
```
代码格式与我们上面的类似，只需要在TODO位置添加我们的测试逻辑就可以了。

<br>
<br>

### testify/assert

`testify`是一个社区非常流行的`Go单元测试工具包`，其中使用最多的功能就是它提供的断言工具——`testify/assert`或`testify/require`。

> 安装

`go get github.com/stretchr/testify`

> 使用示例

我们在写单元测试的时候，通常需要使用断言来校验测试结果，但是由于Go语言官方没有提供断言，所以我们会写出很多的`if...else..`.语句。而`testify/assert`为我们提供了很多`常用的断言函数，并且能够输出友好、易于阅读的错误描述信息`。

比如我们之前在TestSplit测试函数中就使用了`reflect.DeepEqual`来判断期望结果与实际结果是否一致。
```go
t.Run(tt.name, func(t *testing.T) { // 使用t.Run()执行子测试
	got := Split(tt.input, tt.sep)
	if !reflect.DeepEqual(got, tt.want) {
		t.Errorf("expected:%#v, got:%#v", tt.want, got)
	}
})
```
使用`testify/assert`之后就能将上述判断过程简化如下：
```go
t.Run(tt.name, func(t *testing.T) { // 使用t.Run()执行子测试
	got := Split(tt.input, tt.sep)
	assert.Equal(t, got, tt.want)  // 使用assert提供的断言函数
})
```

当我们有多个断言语句时，还可以使用`assert := assert.New(t)`创建一个assert对象，它拥有前面所有的断言方法，只是不需要再传入`Testing.T`参数了

```go
func TestSomething(t *testing.T) {
  assert := assert.New(t)

  // assert equality
  assert.Equal(123, 123, "they should be equal")

  // assert inequality
  assert.NotEqual(123, 456, "they should not be equal")

  // assert for nil (good for errors)
  assert.Nil(object)

  // assert for not nil (good when you expect something)
  if assert.NotNil(object) {

    // now we know that object isn't nil, we are safe to make
    // further assertions without causing any errors
    assert.Equal("Something", object.Value)
  }
}
```
`testify/assert`提供了非常多的断言函数，这里没办法一一列举出来，可以查看[官方文档](https://pkg.go.dev/github.com/stretchr/testify/assert#pkg-functions)了解。

`testify/require`拥有`testify/assert`所有断言函数，它们的唯一区别就是——`testify/require`遇到失败的用例会`立即终止本次测试`。

此外，testify包还提供了`mock、http`等其他测试工具.



<br>
<br>

### 测试覆盖率

测试覆盖率是你的代码被测试套件覆盖的百分比。通常我们使用的都是语句的覆盖率，也就是在测试中至少被运行一次的代码占总代码的比例。

Go提供内置功能来检查你的代码覆盖率。我们可以使用`go test -cover`来查看测试覆盖率。例如：
```shell
split $ go test -cover
PASS
coverage: 100.0% of statements
ok      github.com/Q1mi/studygo/code_demo/test_demo/split       0.005s
```

Go还提供了一个额外的`-coverprofile`参数，用来将覆盖率相关的记录信息输出到一个文件。例如：
```shell
split $ go test -cover -coverprofile=c.out
PASS
coverage: 100.0% of statements
ok      github.com/Q1mi/studygo/code_demo/test_demo/split       0.005s
```

上面的命令会将覆盖率相关的信息输出到当前文件夹下面的`c.out`文件中，然后我们执行`go tool cover -html=c.out`，使用`cover`工具来处理生成的记录信息，该命令会打开本地的浏览器窗口生成一个HTML报告。
![](assets/cover.png)
上图中每个用`绿色标记的语句块表示被覆盖`了，而`红色的表示没有被覆盖`。

