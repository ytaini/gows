package gobasic

/*
通过把一个现有(非interface)的类型定义为一个新的类型时,新的类型不会继承现有类型的方法
*/
import "sync"

type MyMutex sync.Mutex

func Eg101() {
	// var mtx myMutex
	// mtx.Lock()   //error
	// mtx.Unlock() //error

	// var lock MyLocker
	// lock.Lock() //ok
	// lock.Unlock() //ok
}

/*
如果你确实需要原有类型的方法，你可以定义一个新的struct类型，用匿名方式把原有类型嵌入其中。
*/
type MyLocker struct {
	sync.Mutex
}
