/*
 * @Author: wzmiiiiii
 * @Date: 2022-12-09 06:57:42
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-12-09 07:19:07
 */
package static

import (
	"fmt"
	"math/rand"
	"time"
)

type Movable interface {
	Move()
}

type Car struct{}

func (c *Car) Move() {
	fmt.Println("car move ...")
	randNum := rand.Intn(5)
	time.Sleep(time.Duration(randNum) * time.Second)
}

// 日志代理
type LogProxy struct {
	m Movable
}

func NewLogProxy(m Movable) *LogProxy {
	return &LogProxy{
		m: m,
	}
}

func (p *LogProxy) Move() {
	fmt.Println("start...")
	p.m.Move()
	fmt.Println("end...")
}

// 计时代理
type TimerProxy struct {
	m Movable
}

func NewTimerProxy(m Movable) *TimerProxy {
	return &TimerProxy{
		m: m,
	}
}

func (p *TimerProxy) Move() {
	start := time.Now()
	p.m.Move()
	end := time.Since(start)
	fmt.Println(end)
}
