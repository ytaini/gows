/*
 * @Author: wzmiiiiii
 * @Date: 2022-10-31 03:22:25
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-10-31 23:44:18
 * @Description:
	最短迷宫问题:
		迷宫由n行m列的单元格组成(n,m都小于等于50),每个单元格要么是空地,要么是障碍物.
		请找出一条从起点到终点的最短路径长度


	通过dfs求解最短迷宫路径,或最短的步数等
s*/
package main

import (
	"fmt"
	"math"
)

// 起始坐标
var startx, starty int

// 终点坐标.
var endx, endy int

// 记录当前最小路径的长度
var min = math.MaxInt

// 记录当前迷宫行列
var n, m int

// 迷宫/地图
var maze [100][100]int //其中1为空地,2位障碍物

// 访问数组
var v [100][100]int //0:未访问,1:已访问

// x,y为当前的坐标,step为到当前坐标已经经过的步数.
func dfs(x, y, step int) {
	// 判断当前坐标是否到达终点
	if x == endx && endy == y {
		// 当达到终点时.
		// 更新min
		if step < min { //如果当前到达终点的步数比最小值还小.就更新min
			min = step
		}
		return //开始回溯
	}
	// 策略: 右->下->左->上
	// 分别试探每个方向
	// 向右试探
	if maze[x][y+1] == 1 && v[x][y+1] == 0 { // 如果右边是空地且没有被访问过.就可以试探
		//设置当前点为已访问
		v[x][y+1] = 1
		dfs(x, y+1, step+1) //再从当前点进行dfs.
		// 当前点试探完成后,回退时需要将当前点设置为未访问.
		v[x][y+1] = 0
	}
	// 当右边为障碍物时,向下试探
	if maze[x+1][y] == 1 && v[x+1][y] == 0 { // 如果下边是空地且没有被访问过.就可以试探
		//设置当前点为已访问
		v[x+1][y] = 1
		dfs(x+1, y, step+1) //再从当前点进行dfs.
		// 当当前点试探完成后,需要将当前点设置为未访问.
		v[x+1][y] = 0
	}
	// 当右边,下边为障碍物时,向左试探
	if maze[x-1][y] == 1 && v[x-1][y] == 0 { // 如果下边是空地且没有被访问过.就可以试探
		//设置当前点为已访问
		v[x-1][y] = 1
		dfs(x-1, y, step+1) //再从当前点进行dfs.
		// 当当前点试探完成后,需要将当前点设置为未访问.
		v[x-1][y] = 0
	}
	// 当右边,下边,左边为障碍物时,向上试探
	if maze[x][y-1] == 1 && v[x][y-1] == 0 { // 如果下边是空地且没有被访问过.就可以试探
		//设置当前点为已访问
		v[x][y-1] = 1
		dfs(x-1, y, step+1) //再从当前点进行dfs.
		// 当当前点试探完成后,需要将当前点设置为未访问.
		v[x][y-1] = 0
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
	// // 起始和终点坐标
	fmt.Scanf("%d%d%d%d", &startx, &starty, &endx, &endy)

	// // 将起点坐标设为已访问
	v[startx][starty] = 1
	// // 从起点开始dfs
	dfs(startx, starty, 0)
	// // 求得最短路径长度.
	fmt.Println(min)
}
