/*
 * @Author: wzmiiiiii
 * @Date: 2022-11-24 15:16:03
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-11-24 15:18:50
 * @Description:
 */
package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan struct{})

	go func() {
		time.Sleep(time.Second)
		close(c)
	}()
	v, ok := <-c
	fmt.Println(v, ok)
}
