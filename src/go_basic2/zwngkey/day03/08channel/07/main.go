/*
 * @Author: zwngkey
 * @Date: 2021-12-19 16:41:00
 * @LastEditTime: 2022-04-30 20:44:35
 * @Description:
 */
package main

import (
	"fmt"
	"go_basic2/zwngkey/day03/08channel/07/links"
	"log"
	"os"
)

func main() {
	worklist := make(chan []string)

	// Start with the command-line arguments.
	go func() { worklist <- os.Args[1:] }()

	// Crawl the web concurrently.
	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				go func(link string) {
					worklist <- crawl(link)
				}(link)
			}
		}
	}
}

var tokens = make(chan struct{}, 20)

func crawl(url string) []string {
	fmt.Println(url)

	tokens <- struct{}{}
	list, err := links.Extract(url)
	<-tokens

	if err != nil {
		log.Print(err)
	}
	return list
}
