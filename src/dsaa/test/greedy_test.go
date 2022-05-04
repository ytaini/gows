package test

import (
	"dsaa/greedy"
	"fmt"
	"testing"
)

func TestFindContentChildren(t *testing.T) {
	greedy.FindContentChildren([]int{3, 15, 62, 8, 27, 94}, []int{1, 4, 5, 6, 161})
}
func TestEraseOverlapIntervals(t *testing.T) {
	fmt.Printf("greedy.EraseOverlapIntervals(): %v\n", greedy.EraseOverlapIntervals2([][]int{{-52, 31}, {-73, -26}, {82, 97}, {-65, -11}, {-62, -49}, {95, 99}, {58, 95}, {-31, 49}, {66, 98}, {-63, 2}, {30, 47}, {-40, -26}}))
}
