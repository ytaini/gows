/*
 * @Author: wzmiiiiii
 * @Date: 2022-11-21 19:39:27
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-11-24 18:57:52
 * @Description:
 */
package main

import "fmt"

func main() {
	data := mirrorQuery()
	fmt.Printf("data: %v\n", data)
}

func mirrorQuery() string {

	responses := make(chan string, 3)

	go func() { responses <- request("a") }()
	go func() { responses <- request("b") }()
	go func() { responses <- request("c") }()

	return <-responses
}

func request(url string) string {
	return url
}
