/*
 * @Author: zwngkey
 * @Date: 2022-05-18 22:37:36
 * @LastEditors: zwngkey 18390924907@163.com
 * @LastEditTime: 2022-05-18 22:55:52
 * @Description:
	使用sync.Cond实现
*/
package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// 首先，我们定义一个Locker和一个Cond,并且定义顾客等待的座位数。
// 来了一位顾客，座位数加一，Tony老师叫起一位等待的顾客开始理发时，座位数减一。
var (
	seatsLock sync.Mutex
	seats     int
	cond      = sync.NewCond(&seatsLock)
)

// 理发师的工作就是不断的检查是否有顾客等待，如果有，就叫起一位顾客开始理发，理发耗时是随机的，理完再去叫下一位顾客。
// 如果没有顾客，那么理发师就会被阻塞(开始睡觉)。

// 注意这里Cond的使用方法，Wait之后需要for循环检查条件是否满足，并且Wait上下会有Locker的使用。
func Barber() {
	for {
		// 等待一个用户
		log.Println("Tony老师尝试请求一个顾客")
		seatsLock.Lock()
		for seats == 0 {
			cond.Wait()
		}
		seats--
		seatsLock.Unlock()
		log.Println("Tony老师找到一位顾客，开始理发")
		RandomPause(2000)
		log.Println("Tony老师理完了")
	}
}

// customers模拟陆陆续续的顾客的到来
func Customers() {
	for {
		RandomPause(1000)
		go Customer()
	}
}

// 顾客到来之后，先请求seatsLock， 避免多个顾客同时进来的并发竞争。然后他会检查是否有空闲的座位，如果有则坐下并通知理发师。
// 此时理发师如果睡眠则会被唤醒，如果正在理发会忽略。如果没有空闲的座位则离开。
func Customer() {
	seatsLock.Lock()
	defer seatsLock.Unlock()
	if seats == 5 {
		log.Println("没有空闲座位了，一位顾客离开了")
		return
	}
	seats++
	cond.Broadcast()
	log.Println("一位顾客开始坐下排队理发")
}

func RandomPause(r int) {
	time.Sleep(time.Duration(rand.Intn(r)) * time.Millisecond)
}
