package main

import (
	"fmt"
	"math/big"
)

func main() {
	bInt := big.NewInt(10000000000)
	fmt.Println(bInt.Bytes())
	fmt.Println(bInt.BitLen())
	fmt.Println(bInt.Bits())
	fmt.Println(bInt)
	fmt.Println(string(rune(128516)))
}
