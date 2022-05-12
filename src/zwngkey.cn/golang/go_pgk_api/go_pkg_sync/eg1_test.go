/*
 * @Author: zwngkey
 * @Date: 2022-05-11 23:28:26
 * @LastEditors: zwngkey 18390924907@163.com
 * @LastEditTime: 2022-05-12 14:57:09
 * @Description:
 */

package gopkgsync

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

/*
	sync包介绍
		sync包提供了基本的并发编程同步原语(concurrency primitives or synchronization primitives)，例如互斥锁sync.Mutex。
		  sync包囊括了以下数据类型：
			sync.Cond
			sync.Locker
			sync.Map
			sync.Mutex
			sync.Once
			sync.Pool
			sync.RWMutex
			sync.WaitGroup

		除了sync.Once和sync.WaitGroup这2个类型之外，其它类型主要给一些偏底层的库程序用。
			业务代码里的goroutine同步，Go设计者是建议通过channel通信来实现。
*/
/*
	sync.WaitGroup
		WaitGroup是sync包里的一个结构体类型，

		这个结构体有如下3个方法
		Add：
			func (wg *WaitGroup) Add(delta int)

		Done：Done调用会将WiatGroup的计数器减1
			func (wg *WaitGroup) Done()

		Wait：Wait调用会阻塞，直到WaitGroup的计数器为0
			func (wg *WaitGroup) Wait()

	定义一个WaitGroup变量的目的是为了等待若干个goroutine执行完成，主goroutine调用Add方法，
	  指明要等待的子goroutine数量，这些子goroutine执行完成后调用Done方法。同时，主goroutine要调用Wait方法阻塞程序，
		等WaitGroup的计数器减小到0时，Wait方法不再阻塞。

	注意:
		1.一个WaitGroup在第一次使用后不能被复制。
		2.WaitGroup不要拷贝传值，如果要显式地把WaitGroup作为函数参数，一定要传指针。
*/

/*
	sync.Once
		Once是sync包里的一个结构体类型，Once可以在并发场景下让某个操作只执行一次，
			比如设计模式里的单例只创建一个实例，比如只加载一次配置文件，
			比如对同一个channel只关闭一次（对一个已经close的channel再次close会引发panic）等。

		这个结构体只有1个方法Do，参数是要执行的函数。
		（注意：参数是函数类型，而不是函数的返回值，所以只需要把函数名作为参数给到Do即可）
			func(o *Once) Do(f func())

		可以看到Do方法的参数f这个函数类型没有参数，所以如果要执行的函数f需要传递参数就要结合Go的闭包来使用。

		如果Once.Do(f)被多次调用，只有第一次调用才会调用f，即使f在每次调用中都有不同的值。
			每个函数的执行都需要一个新的Once实例。

		注意事项:
			1.Once变量作为函数参数传递时，只能传指针，不能传值。传值给函数A的话，对于函数A而言，
				参数列表里的once形参会是一个新生成的once局部变量，和外部传入的once实参不一样。

			2.如果once.Do(f)方法调用的函数f发生了panic，那Do也会认为函数f已经return了。

			3.如果多个goroutine执行了都去调用once.Do(f)，只有某次的函数f调用返回了，所有Do方法调用才会返回，
			  否则Do方法会一直阻塞等待。如果在f里继续调用同一个once变量的Do方法，就会死锁了，因为Do在等待f返回，f又在等待Do返回。

*/
// 参考下面的例子，print函数通过Once执行，只会执行1次,同时第二个once.Do(f)中的f函数不会被执行.
func TestEg1(t *testing.T) {
	var wg sync.WaitGroup
	var once sync.Once
	size := 10
	wg.Add(size)
	/*
		启用size个goroutine，每个goroutine都调用once.Do(print)
		最终print只会执行一次
	*/
	for i := 0; i < size; i++ {
		go func() {
			defer wg.Done()
			once.Do(print)
			once.Do(func() {
				print1(1, "once")
			})
		}()
	}
	wg.Wait()

}

func print() {
	fmt.Println("test sync.Once.Do")
}
func print1(i int, s string) {
	fmt.Println("test sync.Once.Do", i, s)
}

// sync.Once实现并发安全的单例
type Singleton struct {
	member int
}

var instance *Singleton
var once sync.Once

func getInstance() *Singleton {
	/*
	   通过sync.Once实现单例，只会生成一个Singleton实例
	*/
	once.Do(func() {
		fmt.Println("once")
		instance = &Singleton{member: 100}
	})
	return instance
}
func TestEg2(t *testing.T) {
	var wg sync.WaitGroup
	size := 10
	wg.Add(size)
	/*
	   多个goroutine同时去获取Singelton实例
	*/
	for i := 0; i < size; i++ {
		go func() {
			defer wg.Done()
			instance = getInstance()
			fmt.Printf("%p\n", instance)
		}()
	}
	wg.Wait()
	fmt.Println("end")
}

/*
	sync.Mutex
		Mutex是sync包里的一个结构体类型，含义就是互斥锁。Mutex变量的默认值或者说零值是一个没有加锁的mutex，
			也就是当前mutex的状态是unlocked。

		不要对Mutex使用值传递方式进行函数调用。

		Mutex允许一个goroutine对其加锁，其它goroutine对其解锁，不要求加锁和解锁在同一个goroutine里。

		Mutex结构体类型有2个方法
			Lock()加锁。Lock()方法会把Mutex变量m锁住，如果m已经锁住了，如果再次调用Lock()就会阻塞，直到锁释放。
				func (m *Mutex) Lock()

			Unlock()解锁。Unlock()方法会把Mutex变量m解锁，如果m没有被锁，还去调用Unlock，会遇到runtime error。
				func (m *Mutex) Unlock()

*/
/*
	不加锁
	场景举例：多个 goroutine对共享变量同时执行写操作，并发是不安全的，结果和预期不符。
*/
var sum int = 0

func TestEg3(t *testing.T) {
	var wg sync.WaitGroup
	size := 100
	wg.Add(size)
	for i := 1; i <= size; i++ {
		go func(i int) {
			defer wg.Done()
			/*
				sum是多个goroutine共享的
				也就是多个goroutine同时对共享变量sum做写操作不是并发安全的
			*/
			sum += i
		}(i)
	}
	wg.Wait()
	fmt.Printf("sum of 1 to %d is: %d\n", size, sum)
}

/*
	加锁
	示例代码，通过对共享变量加互斥锁来保证并发安全，结果和预期相符。
*/
func TestEg4(t *testing.T) {
	var wg sync.WaitGroup
	var lock sync.Mutex
	size := 100
	wg.Add(size)
	for i := 1; i <= size; i++ {
		go func(i int) {
			defer wg.Done()
			/*
				sum是多个goroutine共享的
				通过加互斥锁来保证并发安全
			*/
			lock.Lock()
			defer lock.Unlock()

			sum += i
		}(i)
	}
	wg.Wait()
	fmt.Printf("sum of 1 to %d is: %d\n", size, sum)
}

/*
	sync.RWMutex
		RWMutex是sync包里的一个结构体类型，含义是读写锁。RWMutex变量的零值是一个没有加锁的mutex。

		不要对RWMutex变量使用值传递的方式进行函数调用。

		RWMutex允许一个goroutine对其加锁，其它goroutine对其解锁，不要求加锁和解锁在同一个goroutine里。

		RWMutex结构体类型有5个方法：
			Lock()，加写锁。某个goroutine加了写锁后，其它goroutine不能获取读锁，也不能获取写锁
				func (rw *RWMutex) Lock()

			Unlock()，释放写锁。
				func (rw *RWMutex) Unlock()

			RLock()，加读锁。某个goroutine加了读锁后，其它goroutine可以获取读锁，但是不能获取写锁
				func (rw *RWMutex) RLock()

			RUnlock()，释放读锁
				func (rw *RWMutex) RUnlock()

			RLocker()，获取一个类型为Locker的接口，Locker类型定义了Lock()和Unlock()方法
				func (rw *RWMutex) RLocker() Locker

		类型Locker的定义如下
			type Locker interface {
				Lock()
				Unlock()
			}

		Mutex和RWMutex这2个结构体类型实现了Locker这个interface里的所有方法，
			因此可以把Mutex和RWMutex变量或者指针赋值给Locker实例，然后通过Locker实例来加锁和解锁，
				这个在条件变量sync.Cond里会用到.

		注意事项
			Mutex和RWMutex都不是递归锁，不可重入
*/
type Counter struct {
	count int
	rw    sync.RWMutex
}

func (c *Counter) getCount() int {
	c.rw.RLock()
	defer c.rw.RUnlock()
	return c.count
}

func (c *Counter) add() {
	c.rw.Lock()
	defer c.rw.Unlock()
	c.count++
}
func TestEg5(t *testing.T) {
	var wg sync.WaitGroup

	size := 100

	var c Counter

	wg.Add(size)

	for i := 0; i < size; i++ {
		go func() {
			defer wg.Done()
			c.getCount()
			c.add()
		}()
	}
	wg.Wait()
	fmt.Println(c.count)
}

/*
	sync.Cond
		type Cond struct {
			// 根据需求初始化不同的锁，如*Mutex 和 *RWMutex
			L Locker

			notify  notifyList  // 通知列表,调用Wait()方法的goroutine会被放入list中,每次唤醒,从这里取出
			checker copyChecker // 复制检查,检查cond实例是否被复制
			....
		}

		Cond是sync包里的一个结构体类型，表示条件变量。我们知道sync.WaitGroup可以用于等待所有goroutine都执行完成，
			sync.Cond可以用于控制goroutine什么时候开始执行。

		Cond 实现了一个条件变量，在 Locker 的基础上增加的一个消息通知的功能，保存了一个通知列表，
			用来唤醒一个或所有因等待条件变量而阻塞的 Go 程，以此来实现多个 Go 程间的同步。


		不要对Cond变量使用值传递进行函数调用。

		Cond结构体类型以下几个方法与其紧密相关：
			NewCond函数，用于创建条件变量，条件变量的成员L是NewCond函数的参数l. 阻塞等待通知的操作以及通知解除阻塞的操作就是基于sync.Mutex来实现的。
				func NewCond(l Locker) *Cond

			Broadcast，发出广播，唤醒所有等待条件变量c的goroutine。注意：在调用Broadcast方法之前，要确保目标goroutine处于Wait阻塞状态，不然会出现死锁问题。
				func (c *Cond) Broadcast()

			Signal，发出信号，唤醒某一个等待条件变量c的goroutine。注意：在调用Signal方法之前，要确保目标goroutine处于Wait阻塞状态，不然会出现死锁问题。
				func (c *Cond) Signal()

			Wait，调用 Wait 会自动释放锁 c.L，并挂起调用者所在的 goroutine，因此当前协程会阻塞在 Wait 方法调用的地方。
				如果其他协程调用了 Signal 或 Broadcast 唤醒了该协程，Wait 方法结束阻塞时，并重新给 c.L 加锁，
					并且继续执行 Wait 后面的代码
				func (c *Cond) Wait()

		每个Cond变量都有一个Locker类型的成员L，L通常是*Mutex或者*RWMutex类型，调用Wait方法时要对L加锁。

*/
// 下面这个示例，先开启了10个goroutine，这10个goroutine都进入Wait阻塞状态，等待被唤醒。
func TestEg6(t *testing.T) {
	var wg sync.WaitGroup
	var mutex sync.Mutex
	cond := sync.NewCond(&mutex)
	size := 10
	wg.Add(size + 1)

	for i := 0; i < size; i++ {
		i := i
		go func() {
			defer wg.Done()
			/*调用Wait方法时，要对L加锁*/
			cond.L.Lock()
			fmt.Printf("%d ready\n", i)
			/*
			  Wait实际上是会先解锁cond.L，再阻塞当前goroutine
			  这样其它goroutine调用上面的cond.L.Lock()才能加锁成功，才能进一步执行到Wait方法，
			  等待被Broadcast或者signal唤醒。
			  Wait被Broadcast或者Signal唤醒的时候，会再次对cond.L加锁，加锁后Wait才会return
			*/
			cond.Wait()
			fmt.Printf("%d done\n", i)
			cond.L.Unlock()
		}()
	}

	/*这里sleep 2秒，确保目标goroutine都处于Wait阻塞状态
	  如果调用Broadcast之前，目标goroutine不是处于Wait状态，会死锁
	*/
	time.Sleep(2 * time.Second)
	go func() {
		defer wg.Done()
		cond.Broadcast()
	}()

	wg.Wait()
}

//sync.Cond基本使用
//下述代码实现了主线程对多个goroutine的通知的功能。
func TestEg7(t *testing.T) {
	var locker sync.Mutex
	var cond = sync.NewCond(&locker)
	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func(x int) {
			cond.L.Lock() // 获取锁
			defer wg.Done()
			defer cond.L.Unlock() // 释放锁
			fmt.Println("ready:", x)
			cond.Wait() // 等待通知，阻塞当前 goroutine
			// 通知到来的时候, cond.Wait()就会结束阻塞, do something. 这里仅打印
			fmt.Println("done:", x)
		}(i)
	}
	time.Sleep(time.Second * 1) // 睡眠 1 秒，等待所有 goroutine 进入 Wait 阻塞状态
	fmt.Println("Signal...")
	cond.Signal() // 1 秒后下发一个通知给已经获取锁的 goroutine

	time.Sleep(time.Second * 1)
	fmt.Println("Signal...")
	cond.Signal() // 1 秒后下发下一个通知给已经获取锁的 goroutine

	time.Sleep(time.Second * 1)
	cond.Broadcast() // 1 秒后下发广播给所有等待的goroutine
	fmt.Println("Broadcast...")

	wg.Wait()
}

/*
	抛出一个问题：
		主线程执行的时候，如果并不想触发所有的协程，想让不同的协程可以有自己的触发条件，应该怎么实现？

	下面就是一个具体的需求：
		有四个worker和一个master，worker等待master去分配指令，
			master一直在计数，计数到5的时候通知第一个worker，计数到10的时候通知第二个和第三个worker。

	首先列出几种解决方式
		1、所有worker循环去查看master的计数值，计数值满足自己条件的时候，触发操作 >>>>>>>>>弊端：无谓的消耗资源
		2、用channel来实现，几个worker几个channel，eg:worker1的协程里<-channel(worker1)进行阻塞，计数值到5的时候，
			给worker1的channel放入值，阻塞解除，worker1开始工作。
			 >>>>>>>弊端：channel还是比较适用于一对一的场景，一对多的时候，需要起很多的channel，不是很美观
		3、用条件变量sync.Cond，针对多个worker的话，用broadcast，就会通知到所有的worker。
*/
func TestEg8(t *testing.T) {
	var lock sync.Mutex
	cond := sync.NewCond(&lock)
	wg := sync.WaitGroup{}
	wg.Add(4)

	mail := 1

	go func() {
		defer wg.Done()
		for count := 0; count <= 15; count++ {
			time.Sleep(time.Second)
			mail = count
			cond.Broadcast()
		}
	}()

	// worker1
	go func() {
		defer wg.Done()
		// 触发的条件，如果不等于5，就会进入cond.Wait()等待，此时cond.Broadcast()通知进来的时候，wait阻塞解除，
		// 进入下一个循环，	当发现mail == 5，跳出循环，开始工作。
		for mail != 5 {
			cond.L.Lock()
			cond.Wait()
			cond.L.Unlock()
		}
		fmt.Println("worker1 started to work")
		time.Sleep(3 * time.Second)
		fmt.Println("worker1 work end")
	}()
	// worker2
	go func() {
		defer wg.Done()
		for mail != 10 {
			cond.L.Lock()
			cond.Wait()
			cond.L.Unlock()
		}
		fmt.Println("worker2 started to work")
		time.Sleep(3 * time.Second)
		fmt.Println("worker2 work end")
	}()
	// worker3
	go func() {
		defer wg.Done()
		for mail != 10 {
			cond.L.Lock()
			cond.Wait()
			cond.L.Unlock()
		}
		fmt.Println("worker3 started to work")
		time.Sleep(3 * time.Second)
		fmt.Println("worker3 work end")
	}()
	wg.Wait()
}
