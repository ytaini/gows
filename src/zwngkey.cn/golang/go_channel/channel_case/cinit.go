/*
 * @Author: zwngkey
 * @Date: 2022-05-13 21:37:37
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-11-24 19:16:19
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

type ChVoid = chan struct{}

const s = time.Second

func init() {
	rand.Seed(time.Now().UnixNano())
}
