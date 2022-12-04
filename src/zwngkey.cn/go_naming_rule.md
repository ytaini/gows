# 通用规则

- 不要用宽泛、无意义的名字，如：
    - util
    - helper
    - info
    - common

- 缩略语要么全小写，要么全大写。

- 非缩略语则应该使用驼峰命名。

- 不要使用2/4来表达英文 to/for。

- 如无必要，不要起和包相同的名字。


## 项目命名

- 小写，如果有多个单词使用连字符分隔。



## 包命名

- 保持package的名字和目录一致.
- 尽量采取有意义,简短的包名,尽量不要和标准库冲突.
- 包名应该为小写单词,不要使用下划线或混合大小写,使用多级目录来划分层级.
- 简单明了的包命名.
- 不用复数.
- 包名谨慎使用缩写。当缩写是程序员广泛熟知的词时，可以使用缩写。
- 不要使用大而全的无意义包名

如 util、common、misc、global。package 名字应该追求清晰且越来越收敛，符合‘单一职责’原则，而不是像common一样，什么都能往里面放，越来越膨胀，让依赖关系变得复杂，不利于阅读、复用、重构。注意，xxx/utils/encryption这样的包名是允许的。

- 只有一个源文件的包，包名应该和文件名保持一致。
- 不要轻易使用别名。


## 文件命名

- 采用有意义、简短的文件名。
- 文件名应该采用小写，并且使用下划线分割各个单词。

## 结构体命名

- 采用驼峰命名方式，首字母根据访问控制采用大写或者小写。
- 结构体名应该是名词或名词短语，如 Customer、WikiPage、Account、AddressParser，它不应是动词。
- 避免使用 Data、Info 这类意义太宽泛的结构体名。
- 结构体的定义和初始化格式采用多行。


## 接口命名

- 命名规则基本保持和结构体命名规则一致。
- 单个函数的接口名以 er 作为后缀，如 Reader，Writer。
```go
// Reader 字节数组读取接口。
type Reader interface {
    // Read 读取整个给定的字节数据并返回读取的长度
    Read(p []byte) (n int, err error)
}
```

- 两个函数的接口名综合两个函数名。
- 三个以上函数的接口名，类似于结构体名。

## 量命名

### 通用

- 量名不应该以类型作为前缀/后缀。
```go
// map
filterHandlerMap -> opToHandler

// slice
uidSlice -> uids

// array
uidArray -> uids 

// 二维切片或数组。
// 比如多个班级下的学生ID。
uidSliceSlice -> classesUIDs
```

- 特有名词时，需遵循以下规则：
    - 如果变量为私有，且特有名词为首个单词，则使用小写，如 apiClient；
    - 其他情况都应该使用该名词原有的写法，如 APIClient、repoID、UserID；
    - 错误示例：UrlArray，应该写成 urlArray 或者 URLArray；

- 尽量不要用拼音命名。
- 量名遵循驼峰式，根据是否导出决定首字母大小写。
- 若量类型为 bool 类型，则名称应以 Has，Is，Can 或 Allow 等单词开头。
- 私有量和局部量规范一致，均以小写字母开头。
- 作用域较小的名字（局部变量/函数参数），尽量使用简短的名字。
- 作用域较大的名字（全局变量），不要使用缩写，要有明确的意义。
- 全局量中不要包含格式化字符，否则必然违反就近原则。
```go
// Bad
var (
    tGitHost     = "https://git.code.oa.com"
    mrCommitsUri = "/api/v3/projects/%s/merge_request/%s/commits"
)

// Good
func getMRCommitsUri() string {
    return fmt.Sprintf("/api/v3/projects/%s/merge_request/%s/commits", "foo", "bar")
}
```

## 方法接收器命名

- 推荐以类名第一个英文首字母的小写作为接收器的命名。
- 接收器的名称在函数超过 20 行的时候不要用单字符。
- 命名不能采用 me，this，self 这类易混淆名称。

## 错误命名

- 对于存储为全局变量的错误值，根据是否导出，使用前缀 Err 或 err。

- 对于自定义错误类型，请改用后缀 Error。
