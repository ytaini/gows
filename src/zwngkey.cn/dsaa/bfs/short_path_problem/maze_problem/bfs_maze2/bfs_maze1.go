/*
 * @Author: wzmiiiiii
 * @Date: 2022-11-01 01:04:52
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-11-01 14:22:23
 * @Description:
	迷宫问题: bfs求最短路径+输出最短路径各个坐标

	题目:下面表示一个迷宫,其中0为空地,1为障碍物,只能横着走或竖着走,不能斜着走,要求编程找出从左上角到右下角的最短路径.
	输入:
	0 1 0 0 0
	0 1 0 1 0
	0 0 0 0 0
	0 1 1 1 0
	0 0 0 1 0
	输出:
	(0,0)
	(1,0)
	(2,0)
	(2,1)
	(2,2)
	(2,3)
	(2,4)
	(3,4)
	(4,4)
*/
package main

import (
	"fmt"

	"zwngkey.cn/dsaa/queue"
)

var maze [5][5]int
var visit [5][5]bool

var dir = [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

type point struct {
	x, y int
}

type node struct {
	p    point
	path []point //某点所经过的所有点
}

var q = queue.New[node]()

var startx, starty, endx, endy = 0, 0, 4, 4

func bfs() []point {
	// 起始位置入队
	start := point{startx, starty}
	q.Push(node{p: start, path: []point{start}})
	// 标记已入队
	visit[startx][starty] = true

	for !q.IsEmpty() { //无解
		n := q.Peek()
		if n.p.x == endx && n.p.y == endy { // 最优解
			return n.path
		}
		for i := 0; i < 4; i++ {
			tx := n.p.x + dir[i][0]
			ty := n.p.y + dir[i][1]
			if in(tx, ty) && maze[tx][ty] == 0 && !visit[tx][ty] {
				p1 := point{tx, ty}
				pa1 := []point{}
				pa1 = append(pa1, n.path...)
				pa1 = append(pa1, p1)
				n1 := node{p1, pa1}
				q.Push(n1)
				visit[tx][ty] = true
			}
		}
		q.Pop()
	}
	return nil
}

// 判断越界
func in(x, y int) bool {
	return x >= 0 && y >= 0 && x <= 4 && y <= 4
}

/*
迷宫:
0 1 0 0 0
0 1 0 1 0
0 0 0 0 0
0 1 1 1 0
0 0 0 1 0
*/
func main() {
	// 输入迷宫
	for i := 0; i <= endx; i++ {
		for j := 0; j <= endy; j++ {
			fmt.Scanf("%d", &maze[i][j])
		}
	}
	// 通过bfs找到路径
	path := bfs()
	for _, v := range path {
		fmt.Println(v)
	}
}
