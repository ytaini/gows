/*
 * @Author: zwngkey
 * @Date: 2021-12-19 16:03:30
 * @LastEditTime: 2022-05-13 06:27:50
 * @Description:
 */
package main

import (
	"fmt"
	"log"
	"os"
	"sync"

	"zwngkey.cn/gobasic/day03/08channel/06/thumbnail"
)

func main() {
	filenames := make(chan string, 3)
	filenames <- "./1.JPG"
	fmt.Printf("makeThumbnails6(filenames): %v\n", makeThumbnails6(filenames))
}

func makeThumbnails6(filename <-chan string) int64 {
	sizes := make(chan int64)
	var wg sync.WaitGroup
	for f := range filename {
		wg.Add(1)
		go func(f string) {
			defer wg.Done()
			thumb, err := thumbnail.ImageFile(f)
			if err != nil {
				log.Println(err)
				return
			}
			info, _ := os.Stat(thumb)
			sizes <- info.Size()
		}(f)
	}
	go func() {
		wg.Wait()
		close(sizes)
	}()

	var total int64

	for size := range sizes {
		total += size
	}
	return total
}
