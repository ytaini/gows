/*
 * @Author: wzmiiiiii
 * @Date: 2022-10-31 14:29:08
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-10-31 22:06:25
 * @Description:
	最短迷宫问题:
			迷宫由n行m列的单元格组成(n,m都小于等于50),每个单元格要么是空地,要么是障碍物.
			请找出一条从起点到终点的最短路径长度

	通过bfs+队列求解最短路径长度
		算法:
			1.将起点入队.
			2.队首结点可拓展的点入队.
			如果没有可扩展的点,将队首结点出队.
			重复该步骤,直到到达目标位置或队列为空.

		思想:

*/
package main

import (
	"fmt"

	queue "zwngkey.cn/dsaa/queue/circleQueue"
)

// 表示一个坐标.
type point struct {
	x, y int
}

func (p point) equals(p1 point) bool {
	return p.x == p1.x && p.y == p1.y
}

// 定义迷宫maze,
var maze [100][100]int

// 标记某点(坐标)是否访问过.
// 也就是标记某点是否入队了
var visited [100][100]bool

// 当前迷宫行列
var n, m int

// 起点,终点坐标
var startp, endp point

// 记录某个点以及到达这个点所需的步数
type elem struct {
	p    point
	step int
}

// 定义方向为:右下左上
// 定义四个方向偏移的坐标
var dx = []int{0, 1, 0, -1}
var dy = []int{1, 0, -1, 0}

// 判断坐标是否越界,true:未越界
func In(x, y int) bool {
	return x >= 1 && x <= n && y >= 1 && y <= m
}

// 通过队列实现bfs
var q = queue.NewQueue()

// 1.将起点入队.
// 2.队首结点可拓展的点入队.
// 如果没有可扩展的点,将队首结点出队.
// 重复该步骤,直到到达目标位置或队列为空.
func bfs() {
	start := elem{startp, 0}

	q.Push(start) // 起点入队

	visited[startp.x][startp.y] = true //将起点设为已访问.

	// 标记是否有解
	var flag bool

	// 直到到达目标位置或队列为空.
	for !q.IsEmpty() { //队列为空.

		// 得到队首元素
		n := q.Peek().(elem)
		// 得到队首元素坐标
		p := n.p
		// 判断是否到达目标位置
		if p.equals(endp) {
			fmt.Println(n.step)
			flag = true //有解
			break
		}

		// 将队首元素可拓展的点入队.
		// 向4个方向试探,右下左上
		for i := 0; i < 4; i++ {
			tx := p.x + dx[i]
			ty := p.y + dy[i]
			if maze[tx][ty] == 1 && !visited[tx][ty] { //如果该点为空地且未访问
				//将该点入队.
				q.Push(elem{point{tx, ty}, n.step + 1})
				// 入队后,将该点设为已访问(已入队)
				visited[tx][ty] = true
			}
		}
		// 拓展完了,将队首元素出队.
		q.Pop()
	}

	if !flag {
		fmt.Println("无解")
	}
}

/*
程序输入:将这个粘贴到终端.
5 4
1 1 2 1
1 1 1 1
1 1 2 1
1 2 1 1
1 1 1 2
1 1 4 3
*/
func main() {
	// 当前迷宫行列
	fmt.Scanf("%d%d", &n, &m)

	// 初始化迷宫
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			fmt.Scanf("%d", &maze[i][j])
		}
	}
	//  起始和终点坐标
	fmt.Scanf("%d%d%d%d", &startp.x, &startp.y, &endp.x, &endp.y)

	bfs()
}
