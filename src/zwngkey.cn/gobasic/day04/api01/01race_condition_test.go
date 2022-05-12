package main

import (
	"fmt"
	"sort"
	"sync"
	"sync/atomic"
	"testing"
)

//随着构建的系统越来越复杂,正确保护对共享资源的访问以防止竞争条件变得极为重要.

func Test_HaveRaceCondition(t *testing.T) {
	var s = make([]int, 0)

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			//这里多个goroutine会尝试同时访问特定的共享变量,并且这些goroutine中的至少一个尝试修改它.
			//这里会产生特殊的竞态条件,是由数据竞争引起的.
			s = append(s, i)

		}(i)
	}

	wg.Wait()
	sort.Ints(s)
	fmt.Printf("s: %v\n", s)
	// 问题:
	// go test -v 01race_condition_test.go -race :使用-race标志执行测试,go会告诉你存在数据竞争并帮助你准确定位.
	// s: [0 2 4 5 6 7 8 9]
	// testing.go:1152: race detected during execution of test
}

//保护对这些共享资源的访问通常涉及常用的内存同步机制,比如通道或互斥锁.
func Test_NoRaceCondition(t *testing.T) {
	var s = make([]int, 0)

	var wg sync.WaitGroup
	//使用互斥锁,控制对共享变量的访问
	var mutex sync.Mutex
	//问题:对于高吞吐量的系统,性能变得非常重要.因此减少锁争用变得更加重要.
	//执行此操作的最基本的方式之一是使用读写锁(sync.WRMutex),而不是标准的sync.Mutex

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			mutex.Lock()
			defer wg.Done()
			defer mutex.Unlock()

			s = append(s, i)
		}(i)
	}

	wg.Wait()
	sort.Ints(s)
	fmt.Printf("s: %v\n", s)
}

// 原子.
// go的atomic包提供了用于实现同步算法的低级原子内存原语.

// atomic不能代替互斥锁,但是当涉及到可以使用读取-复制-更新模式管理共享资源时,它非常出色.
// 在这种技术中,我们通过引用获取当前值,当我们想要更新它时,我们不修改原始值,而是替换指针(因此没有人访问另一线程可能访问的相同资源)
//

func Test_RaceCondition_Atomic(t *testing.T) {
	var s = atomic.Value{}
	s.Store([]int{})

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			s1 := s.Load().([]int)
			s.Store(append(s1, i))
		}(i)
	}

	wg.Wait()
	s1 := s.Load().([]int)
	sort.Ints(s1)
	fmt.Printf("s1: %v\n", s1)
}

//上面的示例无法使用此模式实现,因为它应该随着时间的推移扩展现有资源而不是完全替换其内容,但在许多情况下,读取-复制-更新模式是完美的.
//在下面这个例子中,我们正在执行一个并行基准测试,比较原子和读写互斥.

type AtomicValue struct {
	value atomic.Value
}

func (a *AtomicValue) Get() bool {
	return a.value.Load().(bool)
}

func (a *AtomicValue) Set(value bool) {
	a.value.Store(value)
}

func BenchmarkAtomicValue_Get(b *testing.B) {
	atomB := AtomicValue{}
	atomB.value.Store(false)

	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			atomB.Get()
		}
	})
}

type MutexBool struct {
	mutex sync.RWMutex
	flag  bool
}

func (mb *MutexBool) Get() bool {
	mb.mutex.Lock()
	defer mb.mutex.Unlock()
	return mb.flag
}

func (mb *MutexBool) Set(value bool) {

}

func BenchmarkMutexBool_Get(b *testing.B) {
	mb := MutexBool{flag: true}

	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			mb.Get()
		}
	})
}
