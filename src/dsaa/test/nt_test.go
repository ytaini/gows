package test

import (
	nt "dsaa/number_theory"
	"dsaa/util"
	"fmt"
	"testing"
)

func TestGCD(t *testing.T) {
	fmt.Printf("util.GCD(): %v\n", util.GCD(25, 225))
}

func TestNtGcd(t *testing.T) {
	fmt.Printf("nt.NTGCD([]int{3, 5, 10, 11, 15, 17, 19}): %v\n", nt.NTGCD([]int{3, 5, 10, 11, 13, 17, 19}))
}

// [7,5,6,8,3]

func TestFindGCD(t *testing.T) {
	fmt.Printf("nt.FindGCD([]int{7, 5, 6, 8, 3}): %v\n", nt.FindGCD([]int{7, 5, 6, 8, 3}))
}

func Test1(t *testing.T) {
	fmt.Printf("nt.HasGroupsSizeX([]int{1, 1, 1, 2, 2, 2, 3, 3, 3}): %v\n", nt.HasGroupsSizeX([]int{2, 2, 2, 2, 3, 3, 3, 3, 3, 3, 1, 1}))
}
