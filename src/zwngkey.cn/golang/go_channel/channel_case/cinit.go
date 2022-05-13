/*
 * @Author: zwngkey
 * @Date: 2022-05-13 21:37:37
 * @LastEditors: zwngkey 18390924907@163.com
 * @LastEditTime: 2022-05-13 21:47:47
 * @Description:
 */
package channelcase

import (
	"math/rand"
	"time"
)

type ChInt = chan int
type RecvChInt = <-chan int
type SendChInt = chan<- int

type void = struct{}

type ChVoid = chan void

const s = time.Second

func init() {
	rand.Seed(time.Now().UnixNano())
}
