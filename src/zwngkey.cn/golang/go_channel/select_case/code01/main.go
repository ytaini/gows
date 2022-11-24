/*
 * @Author: wzmiiiiii
 * @Date: 2022-11-24 19:11:00
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-11-24 19:13:58
 * @Description:
	如何在select中实现优先级?
*/

// 已知，当select 存在多个 case时会随机选择一个满足条件的case执行。
// 现在我们有一个需求：我们有一个函数会持续不间断地从ch1和ch2中分别接收任务1和任务2，
// 如何确保当ch1和ch2同时达到就绪状态时，优先执行任务1，在没有任务1的时候再去执行任务2呢？
package main

import "fmt"

func Worker(ch1, ch2 <-chan int, stopCh chan struct{}) {
	for {
		select {
		case <-stopCh:
			return
		case job1 := <-ch1:
			fmt.Println(job1)
		case job2 := <-ch2:
		priority:
			for {
				select {
				case job1 := <-ch1:
					fmt.Println(job1)
				default:
					break priority
				}
			}
			fmt.Println(job2)
		}
	}
}
