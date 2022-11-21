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
	return ""
}
