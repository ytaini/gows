package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	//portScan()
	//portScan1()
	result := make(chan string)
	stream := make(chan string, 100)
	go worker(stream, result)
	for i := 1; i < 65535; i++ {
		stream <- fmt.Sprintf("localhost:%d", i)
	}
	close(stream)
	go func() {
		wg.Wait()
		close(result)
	}()
	for addr := range result {
		fmt.Println(addr)
	}
}

// goroutine池
func worker(stream <-chan string, result chan<- string) {
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for address := range stream {
				conn, err := net.Dial("tcp", address)
				if err != nil {
					continue
				}
				conn.Close()
				result <- address
			}
		}()
	}
}

func portScan1() {
	openedAddrs := make(chan string)
	t := time.Now()
	for i := 1; i < 66535; i++ {
		wg.Add(1)
		i := i
		go func() {
			defer wg.Done()
			address := fmt.Sprintf("localhost:%d", i)
			conn, err := net.Dial("tcp", address)
			if err != nil {
				return
			}
			defer conn.Close()
			openedAddrs <- address
		}()
	}
	go func() {
		wg.Wait()
		close(openedAddrs)
	}()
	for addr := range openedAddrs {
		fmt.Println(addr)
	}
	t1 := time.Since(t)
	fmt.Println(t1)
}

func portScan() {
	var openedAddrs []string
	for i := 1; i < 65535; i++ {
		address := fmt.Sprintf("localhost:%d", i)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			continue
		}
		conn.Close()
		openedAddrs = append(openedAddrs, address)
	}
	fmt.Println(openedAddrs)
}
