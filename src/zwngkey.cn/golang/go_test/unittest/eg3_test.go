/*
 * @Author: wzmiiiiii
 * @Date: 2022-12-08 04:13:39
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-12-08 04:36:42
 */
package unit_test

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func setup() {
	fmt.Println("Before all tests")
}

func teardown() {
	// 先打印4 再打印3
	defer func() {
		fmt.Println(3)
	}()
	defer func() {
		fmt.Println(4)
	}()
	fmt.Println("After all tests")
}

func Test1(t *testing.T) {
	fmt.Println("I'm test1")
}

func Test2(t *testing.T) {
	// 先打印1 再打印2
	defer t.Cleanup(func() {
		fmt.Println(1)
	})
	defer t.Cleanup(func() {
		fmt.Println(2)
	})
	fmt.Println("I'm test2")
	fmt.Println(time.Now())
	fmt.Println(t.Deadline())
	time.Sleep(10 * time.Second)
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}
