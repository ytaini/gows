/*
 * @Author: zwngkey
 * @Date: 2022-05-12 18:39:07
 * @LastEditors: zwngkey 18390924907@163.com
 * @LastEditTime: 2022-05-12 20:11:37
 * @Description: 匿名结构体使用场景
 */
package goother

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"testing"
)

/*
	全局变量组合
		有时候我们会在程序里定义若干全局变量，有些全局变量的含义是互相关联的，这个时候我们可以使用匿名结构体把这些关联的全局变量组合在一起。

		对全局匿名结构体变量完成赋值后，后续代码都可以使用这个匿名结构体变量。

		注意：如果你的程序对于某个全局的结构体要创建多个变量，就不要用匿名结构体了。


	局部变量组合
		如果在局部作用域(比如函数或者方法体内)里，某些变量的含义互相关联，就可以组合到一个结构体里。

		同时这个结构体只是临时一次性使用，不需要创建这个结构体的多个变量，那就可以使用匿名结构体。


	构建测试数据
		匿名结构体可以和切片结合起来使用，通常用于创建一组测试数据。

*/
// DBConfig 声明全局匿名结构体变量
var DBConfig struct {
	user string
	pwd  string
	host string
	port int
	db   string
}

// SysConfig 全局匿名结构体变量也可以在声明的时候直接初始化赋值
var SysConfig = struct {
	sysName string
	mode    string
}{"tutorial", "debug"}

// 测试数据
var IndexRuneTests = []struct {
	s    string
	rune rune
	out  int
}{
	{"a A x", 'A', 2},
	{"some_text=some_value", '=', 9},
	{"☺a", 'a', 3},
	{"a☻☺b", '☺', 4},
}

/*
	嵌套锁(embed lock)
		我们经常遇到多个goroutine要操作共享变量，为了并发安全，需要对共享变量的读写加锁。

		这个时候通常需要定义一个和共享变量配套的锁来保护共享变量。

		匿名结构体和匿名字段相结合，可以写出更优雅的代码来保护匿名结构体里的共享变量，实现并发安全
*/
// hits 匿名结构体变量
// 这里同时用到了匿名结构体和匿名字段, sync.Mutex是匿名字段
// 因为匿名结构体嵌套了sync.Mutex，所以hits就有了sync.Mutex的Lock和Unlock方法
var hits struct {
	sync.Mutex
	n int
}

func Test41(t *testing.T) {
	var wg sync.WaitGroup
	N := 100
	// 启动100个goroutine对匿名结构体的成员n同时做读写操作
	wg.Add(N)
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			hits.Lock()
			defer hits.Unlock()
			hits.n++
		}()
	}
	wg.Wait()
	fmt.Println(hits.n) // 100

}

/*
	HTTP处理函数中JSON序列化和反序列化
		我们在处理http请求时，通常会和JSON数据打交道。

		比如post请求的content-type使用application/json时，服务器接收过来的json数据是key:value格式，
			不同key的value的类型可以不一样，可能是数字、字符串、数组等，
			因此会遇到使用json.Unmarshal和map[string]interface{}来接收JSON反序列化后的数据。

		使用map[string]interface{}有几个问题：
			1.没有类型检查：比如json的某个value本来预期是string类型，但是请求传过来的是bool类型，
				使用json.Unmarshal解析到map[string]interface{}是不会报错的，因为空interface可以接受任何类型数据。
			2.map是模糊的：Unmarshal后得到了map，我们还得判断这个key在map里是否存在。
				否则拿不存在的key的value，得到的可能是给nil值，如果不做检查，直接对nil指针做*操作，会引发panic。

			3.代码比较冗长：得先判断key是否存在，如果存在，要显示转换成对应的数据类型，并且还得判断转换是否成功。代码会比较冗长。

		这个时候我们就可以使用匿名结构体来接收反序列化后的数据，代码会更简洁。参见如下代码示例：
*/
// 请求命令：curl -X POST -H "content-type: application/json" http://localhost:4000/user -d '{"name":"John", "age":111}'

func AddUser(w http.ResponseWriter, r *http.Request) {
	// data 匿名结构体变量，用来接收http请求发送过来的json数据
	data := struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}{}
	// 把json数据反序列化到data变量里
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&data)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(data)
	fmt.Fprint(w, "Hello!")
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "index")
}

func Test42(t *testing.T) {
	http.HandleFunc("/", index)
	http.HandleFunc("/user", AddUser)
	log.Fatal(http.ListenAndServe("localhost:4000", nil))
}

/*
	匿名结构体可以让我们不用先定义结构体类型，再定义结构体变量。让结构体的定义和变量的定义可以结合在一起，一次性完成。
*/
