package main

import "fmt"

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	go counter(naturals)
	go squaresFunc(squares, naturals)
	print(squares)

}
func print(in <-chan int) {
	for x := range in {
		fmt.Println(x)
	}
}

func squaresFunc(out chan<- int, in <-chan int) {
	for x := range in {
		out <- x * x
	}
	close(out)
}

func counter(out chan<- int) {
	for x := 1; x <= 100; x++ {
		out <- x
	}
	close(out)
}
