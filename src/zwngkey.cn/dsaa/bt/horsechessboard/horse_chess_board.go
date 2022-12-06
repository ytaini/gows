/*
 * @Author: wzmiiiiii
 * @Date: 2022-12-06 00:09:12
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-12-06 05:02:27
 */

// 题目要求:
// 马踏棋盘算法也被称为骑士周游问题
// 将马随机放在国际象棋的8×8棋盘Board[0～7][0～7]的某个方格中，马按走棋规则(马走日字)进行移动。要求每个方格只进入一次，走遍棋盘上全部64个方格

// DFS + BT + Greedy
package bt

import "sort"

// 棋盘大小
var SIZE int = 8

type ChessBoard struct {
	cb       [][]int //结果
	row, col int     //棋盘行列
	ok       bool    //是否走完
}

func NewChessBoard(size int) *ChessBoard {
	SIZE = size
	c := &ChessBoard{
		row: size,
		col: size,
	}
	c.cb = make([][]int, size)
	for i := range c.cb {
		c.cb[i] = make([]int, size)
	}
	return c
}

func (c *ChessBoard) TraversalChess(start Point) {
	if start.X > c.row-1 || start.Y > c.col-1 {
		panic("坐标非法")
	}
	cb, ok := traversalChess(c.cb, start, 1)
	c.cb = cb
	c.ok = ok
}

var flag = false //标记是否走完了整个棋盘
// 算法核心.
// 返回值: []][]int 存放每一步如何走的, bool: 是否走完整个棋盘
// 参数: p: 开始位置 (1,3),step: 当前走的步数(从1开始)
func traversalChess(chess [][]int, p Point, step int) ([][]int, bool) {

	// 标记p点已访问,同时记录了到p点时已经几步
	chess[p.X][p.Y] = step
	// 基于当前的点p得到下一步可以走的点的集合
	ps := nextPoints(p)
	// 按len(nextPoints[i])对ps进行非递减排序
	// 即: 基于ps中所有的Point的下一步的集合的数量进行排序,这样可以减少回溯的次数
	sort.Slice(ps, func(i, j int) bool {
		return len(nextPoints(ps[i])) < len(nextPoints(ps[j]))
	}) // 通过贪心优化.差距巨大!!!
	for i := 0; i < len(ps); i++ {
		p1 := ps[i]
		if chess[p1.X][p1.Y] == 0 { //说明p点没有访问过
			traversalChess(chess, p1, step+1)
		}
	}
	// 回溯
	if step < SIZE*SIZE && !flag {
		chess[p.X][p.Y] = 0
	} else {
		flag = true
	}
	return chess, flag
}

// 如果走完了整个棋盘返回true.
func (c *ChessBoard) Result() ([][]int, bool) {
	return c.cb, c.ok
}

type Point struct {
	X, Y int
}

type points []Point

// 方向
var dir = [][]int{{-2, -1}, {-2, 1}, {-1, 2}, {1, 2}, {2, 1}, {2, -1}, {1, -2}, {-1, -2}}

// 基于当前的点p得到下一步可以走的点的集合.(顺时针存储)
func nextPoints(p Point) points {
	var ps points
	for _, v := range dir {
		x := v[0] + p.X
		y := v[1] + p.Y
		if x >= 0 && y >= 0 && x < SIZE && y < SIZE {
			ps = append(ps, Point{x, y})
		}
	}
	// l := p.X - 2
	// r := p.X + 2
	// t := p.Y + 2
	// b := p.Y - 2
	// if l >= 0 {
	// 	if b+1 >= 0 {
	// 		if chess[l][b+1] == 0 {
	// 			points = append(points, &Point{l, b + 1})
	// 		}
	// 	}
	// 	if t-1 < COL {
	// 		if chess[l][t-1] == 0 {
	// 			points = append(points, &Point{l, t - 1})
	// 		}
	// 	}
	// }
	// if t < COL {
	// 	if l+1 >= 0 {
	// 		if chess[l+1][t] == 0 {
	// 			points = append(points, &Point{l + 1, t})
	// 		}
	// 	}
	// 	if r-1 < ROW {
	// 		if chess[r-1][t] == 0 {
	// 			points = append(points, &Point{r - 1, t})
	// 		}
	// 	}
	// }
	// if r < ROW {
	// 	if t-1 < COL {
	// 		if chess[r][t-1] == 0 {
	// 			points = append(points, &Point{r, t - 1})
	// 		}
	// 	}
	// 	if b+1 >= 0 {
	// 		if chess[r][b+1] == 0 {
	// 			points = append(points, &Point{r, b + 1})
	// 		}
	// 	}
	// }
	// if b >= 0 {
	// 	if r-1 < ROW {
	// 		if chess[r-1][b] != 0 {
	// 			points = append(points, &Point{r - 1, b})
	// 		}
	// 	}
	// 	if l+1 >= 0 {
	// 		if chess[l+1][b] != 0 {
	// 			points = append(points, &Point{l + 1, b})
	// 		}
	// 	}
	// }
	return ps
}
