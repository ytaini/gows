/*
 * @Author: wzmiiiiii
 * @Date: 2022-10-31 13:10:39
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-11-01 13:05:59
 * @Description:
	最短迷宫问题:
		迷宫由n行m列的单元格组成(n,m都小于等于50),每个单元格要么是空地,要么是障碍物.
		请找出一条从起点到终点的最短路径长度

	通过dfs求解最短路径长度

	优化第一个版本
*/
package main

import (
	"fmt"
	"math"

	"zwngkey.cn/dsaa/queue"
)

type Point struct {
	x, y int //记录当前坐标
}

// 起始坐标,终点坐标.
var startp, endp Point

// 记录当前最小路径的长度
var min = math.MaxInt

// 记录当前迷宫行列
var n, m int

// 迷宫-画布
var maze [100][100]int //其中1为空地,2位障碍物

// 标记点(坐标)是否访问过.
var visited [100][100]bool //false:未访问,true:已访问

// 定义方向数组
// {0, 1} 向右
// {1, 0} 向下
// {0, -1} 向左
// {-1, 0} 向上
var dir = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

// var dx = []int{0, 1, 0, -1}
// var dy = []int{1, 0, -1, 0}

// 记录路径坐标集
var temp = queue.New[Point]()

// 最短路径坐标集
var result *queue.Queue[Point]

// p为当前的坐标,step为到当前坐标已经经过的步数.
func dfs(p Point, step int) {
	// 判断当前坐标是否到达终点
	if p.x == endp.x && p.y == endp.y {
		// 当达到终点时.
		// 更新min
		if step < min { //如果当前 到达终点的步数比最小值还小.就更新min
			min = step
			result = temp.Copy()
		}
		return //开始回溯
	}
	// 策略: 右->下->左->上
	// 通过循环分别试探每个方向
	for i := 0; i < 4; i++ {
		tx := p.x + dir[i][0]
		ty := p.y + dir[i][1]
		if maze[tx][ty] == 1 && !visited[tx][ty] {
			visited[tx][ty] = true
			p := Point{tx, ty}
			temp.Push(p)
			dfs(p, step+1)

			// 回溯部分
			// 当前点试探完成后,回溯时需要将当前点设置为未访问.
			visited[tx][ty] = false
			temp.Pop()
		}
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

	//  将起点坐标设为已访问
	visited[startp.x][startp.y] = true

	// 记录起始坐标
	temp.Push(startp)

	//  从起点开始dfs
	dfs(startp, 0)

	//  求得最短路径长度.
	fmt.Println("最短路径长度:", min)

	len := result.Len()
	for i := 1; i <= len; i++ {
		p, _ := result.Pop()
		fmt.Printf("路径坐标%d:%v\n", i, p)
	}

}
