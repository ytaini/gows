/*
 * @Author: zwngkey
 * @Date: 2022-05-12 14:56:34
 * @LastEditors: zwngkey 18390924907@163.com
 * @LastEditTime: 2022-05-12 16:57:48
 * @Description: sync包API
 */
package gopkgsync

import (
	"container/list"
	"fmt"
	"log"
	"math/rand"
	"sync"
	"testing"
	"time"
)

/*
	sync.Cond 可以用来干什么？
		Go 的 sync 包中的 Cond 实现了一种条件变量，可以使用多个 Reader 等待公共资源。

		每个 Cond 都会关联一个 Lock ,当修改条件或者调用 Wait 方法，必须加锁，保护 Condition。 有点类似 Java 中的 Wait 和 NotifyAll。

		sync.Cond 条件变量是用来协调想要共享资源的那些 goroutine, 当共享资源的状态发生变化时，可以被用来通知被互斥锁阻塞的 gorountine。


	sync.Mutex 通常用来保护临界区和共享资源，条件变量 sync.Cond 用来协调想要访问的共享资源。


	sync.Cond 使用场景
		有一个协程正在修改某个数据数据，其他协程必须等待这个协程接改完数据，才能读取数据。

		上述情形下，如果单纯的使用 channel 或者互斥锁，只能有一个协程可以等待，并读取到数据，没办法通知其他协程也读取数据。

		这个时候怎么办？
			1.可以用一个全局变量标识第一个协程是否接收数据完毕，剩下的协程反复检查该变量的值，直到读取到数据。
			2.也可创建多个 channel, 每个协程阻塞在一个 Channel 上，由接收数据的协程在数据接收完毕后，挨个通知。
			3.通过Go 中其实内置来一个 sync.Cond 来解决这个问题。
*/
var done = false

func read(name string, data *string, c *sync.Cond) {
	c.L.Lock()
	for !done {
		c.Wait()
	}
	log.Println(name, "starts reading")
	time.Sleep(time.Second / 2)
	fmt.Println(*data)
	c.L.Unlock()
}

func write(name string, data *string, c *sync.Cond) {
	log.Println(name, "starts writing")
	time.Sleep(time.Second)
	c.L.Lock()
	done = true
	*data = "golang"
	c.L.Unlock()
	log.Println(name, "wakes all")
	c.Broadcast()
}

func TestSyncCond(t *testing.T) {
	cond := sync.NewCond(&sync.Mutex{})
	var data = "go"

	go read("reader1", &data, cond)
	go read("reader2", &data, cond)
	go read("reader3", &data, cond)
	write("writer", &data, cond)

	time.Sleep(time.Second * 3)
}

/*
	案例：Redis连接池
		可以看一下下面的代码，使用了 Cond 实现一个 Redis 的连接池，最关键的代码就是在链表为空的时候需要调用 Cond 的 Wait 方法，
			将 gorutine 进行阻塞。然后 goruntine 在使用完连接后，将连接返回池子后，需要通知其他阻塞的 goruntine 来获取连接。
*/

// 连接池
type Pool struct {
	lock    sync.Mutex // 锁
	clients list.List  // 连接
	cond    *sync.Cond // cond实例
	close   bool       // 是否关闭
}

// Redis Client
type Client struct {
	id int32
}

// 创建Redis Client
func NewClient() *Client {
	return &Client{
		id: rand.Int31n(100000),
	}
}

// 关闭Redis Client
func (c *Client) CloseClient() {
	fmt.Printf("Client:%d 正在关闭", c.id)
}

// 创建连接池
func NewPool(maxConnNum int) *Pool {
	pool := new(Pool)
	pool.cond = sync.NewCond(&pool.lock)

	// 创建连接
	for i := 0; i < maxConnNum; i++ {
		client := NewClient()
		pool.clients.PushBack(client)
	}

	return pool
}

// 从池子中获取连接
func (p *Pool) Pull() *Client {
	p.lock.Lock()
	defer p.lock.Unlock()

	// 已关闭
	if p.close {
		fmt.Println("Pool is closed")
		return nil
	}

	// 如果连接池没有连接 需要阻塞
	for p.clients.Len() <= 0 {
		p.cond.Wait()
	}

	// 从链表中取出头节点，删除并返回
	ele := p.clients.Remove(p.clients.Front())
	return ele.(*Client)
}

// 将连接放回池子
func (p *Pool) Push(client *Client) {
	p.lock.Lock()
	defer p.lock.Unlock()

	if p.close {
		fmt.Println("Pool is closed")
		return
	}

	// 向链表尾部插入一个连接
	p.clients.PushBack(client)

	// 唤醒一个正在等待的goruntine
	p.cond.Signal()
}

// 关闭池子
func (p *Pool) Close() {
	p.lock.Lock()
	defer p.lock.Unlock()

	// 关闭连接
	for e := p.clients.Front(); e != nil; e = e.Next() {
		client := e.Value.(*Client)
		client.CloseClient()
	}

	// 重置数据
	p.close = true
	p.clients.Init()
}

func TestEg9(t *testing.T) {

	var wg sync.WaitGroup

	pool := NewPool(3)
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func(index int) {

			defer wg.Done()

			// 获取一个连接
			client := pool.Pull()

			fmt.Printf("Time:%s | 【goruntine#%d】获取到client[%d]\n", time.Now().Format("15:04:05"), index, client.id)
			time.Sleep(time.Second * 5)
			fmt.Printf("Time:%s | 【goruntine#%d】使用完毕，将client[%d]放回池子\n", time.Now().Format("15:04:05"), index, client.id)

			// 将连接放回池子
			pool.Push(client)
		}(i)
	}

	wg.Wait()
}

/*
	Go语言里普通map的读写不是并发安全的，sync.Map的读写是并发安全的。

	sync.Map可以理解为类似一个map[interface{}]interface{}的结构，key可以类型不一样，value也可以类型不一样，
		多个goroutine对其进行读写不需要额外加锁。

	Go官方设计sync.Map主要满足以下2个场景的用途
		1.每个key只写一次，其它对该key的操作都是读操作
		2.多个goroutine同时读写map，但是每个goroutine只读写各自的keys

	以上2种场景，相对于对普通的map加Mutex或者RWMutex来实现并发安全，使用sync.Map不用在业务代码里加锁，会大幅减少锁竞争，
		提升性能。其它更为常见的场景还是使用普通的Map，搭配Mutex或者RWMutex来使用。

	不能对sync.Map使用值传递方式进行函数调用。

	sync.Map结构体类型有如下几个方法：
		Delete: 删除map里的key，即使key不存在，执行Delete操作也没任何影响
			func (m *Map) Delete(key interface{})

		Load，从map里取出key对应的value。如果key存在map里，返回值value就是对应的值，ok就是true。如果key不在map里，返回值value就是nil，ok就是false。
			func (m* Map) Load(key interface{}) (value interface{}, ok bool)

		LoadAndDelete，删除map里的key。如果key存在map里，返回值value就是对应的值，loaded就是true。如果key不在map里，返回值value就是nil，loaded就是false。
			func (m* Map) LoadAndDelete(key interface{}) (value interface{}, loaded bool)

		LoadOrStore，从map里取出key对应的value。如果key在map里不存在，就把LoadOrStrore函数调用传入的参数<key, value>存储到map里，并返回参数里的value。
			如果key在map里，那loaded是true，如果key不在map里，那loaded是false。
			func (m* Map) LoadOrStore(key, value interface{}) (actual interface{}, loaded bool)

		Range，遍历map里的所有<key, value>对，把每个<key, value>对，都作为参数传递给f去调用，如果遍历执行过程中，f返回false，那range迭代就结束了。
			func (m* Map) Range(f func(key, value interface{}) bool)

		Store，往map里插入<key, vaue>对，即使key已经存在于map里，也没有任何影响
			func (m* Map) Store(key, value interface{})

	注意事项
		1.sync.Map不支持len和cap函数
		2.在评估要不要使用sync.Map的时候，先考察业务场景是否符合上面描述的场景1和2，符合再考虑用sync.Map，不符合就用普通map+Mutex或者RWMutex。
*/
func TestEg21(t *testing.T) {
	/*统计字符串里每个字符出现的次数*/
	m := sync.Map{}
	str := "abcabcd"
	for _, value := range str {
		if temp, ok := m.Load(value); !ok {
			m.Store(value, 1)
		} else {
			m.Store(value, temp.(int)+1)
		}
	}
	/*使用sync.Map里的Range遍历map*/
	m.Range(func(key, value any) bool {
		fmt.Println(string(key.(rune)), value)
		return true
	})
}

var m sync.Map

/*
sync.Map里每个key只写一次，属于场景1
*/
func changeMap(key int) {
	m.Store(key, 1)
}

func TestEg22(t *testing.T) {

	var wg sync.WaitGroup
	size := 10
	wg.Add(size)

	for i := 0; i < size; i++ {
		i := i
		go func() {
			defer wg.Done()
			changeMap(i)
		}()
	}
	wg.Wait()

	/*使用sync.Map里的Range遍历map*/
	m.Range(func(key, value interface{}) bool {
		fmt.Println(key, value)
		return true
	})
}
