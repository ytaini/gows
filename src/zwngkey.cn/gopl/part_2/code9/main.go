package main

import "fmt"

func main() {
	srcSlice := []int{1, 2, 3, 4, 5}
	destSlice := []int{9, 8, 7}

	c := copy(srcSlice, destSlice)
	fmt.Printf("c: %v\n", c)
	fmt.Printf("srcSlice: %v\n", srcSlice)
	fmt.Printf("destSlice: %v\n", destSlice)

	fmt.Println("------------------------------------------")

	test_copy()
}

func test_copy() {

	const eleCount = 100

	srcSlice := make([]int, eleCount)

	for i := range srcSlice {
		srcSlice[i] = i
	}

	refSrc := srcSlice

	copySlice := make([]int, eleCount)

	copy(copySlice, srcSlice)

	srcSlice[0] = 900

	fmt.Println(refSrc[0])    //900
	fmt.Println(copySlice[0]) //0

	copy(copySlice, srcSlice[3:6])

	for i := 0; i < 10; i++ {
		fmt.Printf("%d\t", copySlice[i])
	}

}
