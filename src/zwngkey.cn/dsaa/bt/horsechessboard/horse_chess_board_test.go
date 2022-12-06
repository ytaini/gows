/*
 * @Author: wzmiiiiii
 * @Date: 2022-12-06 02:53:27
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-12-07 04:41:00
 */
package bt

import (
	"fmt"
	"testing"
	"time"
)

func Test12(t *testing.T) {
	cb := NewChessBoard(7)
	t1 := time.Now()
	cb.TraversalChess(Point{3, 3})
	t2 := time.Since(t1)
	fmt.Println(t2)
	res, ok := cb.Result()
	if ok {
		for _, v := range res {
			for _, p := range v {
				fmt.Printf("%d\t", p)
			}
			fmt.Println()
		}
	}
}

func Test11(t *testing.T) {
	// points := nextPoint(&Point{4, 4})
	points := nextPoints(Point{5, 5})
	for _, p := range points {
		fmt.Println(p.X, p.Y)
	}
	fmt.Println(len(nextPoints(points[0])))
	fmt.Println(len(nextPoints(points[1])))
}
