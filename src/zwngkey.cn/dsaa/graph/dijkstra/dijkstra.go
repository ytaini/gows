/*
 * @Author: wzmiiiiii
 * @Date: 2022-12-05 00:30:26
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-12-05 01:56:44
 */
package dijkstra

import (
	"fmt"
	"math"
)

// 定义最大的权
const INF = math.MaxInt

// 通过邻接矩阵构建无向网
type AMGraph struct {
	vexNum, arcNum int     //图的当前顶点数和边数
	vexs           []rune  //顶点表
	arcs           [][]int //邻接矩阵
}

// 定义 表示边的结构体
type Edge struct {
	v1, v2 rune
	weight int
}

func NewEdge(v1, v2 rune, weight int) *Edge {
	return &Edge{v1, v2, weight}
}

// Edges 定义 边的切片
type Edges []*Edge

func (e Edges) Len() int           { return len(e) }
func (e Edges) Swap(i, j int)      { e[i], e[j] = e[j], e[i] }
func (e Edges) Less(i, j int) bool { return e[i].weight < e[j].weight }

func NewAMGraph() *AMGraph {
	return &AMGraph{}
}

// 通过[]rune 输入 图的所有顶点
func (g *AMGraph) Vertexes(vers []rune) {
	vexNum := len(vers)
	g.vexNum = vexNum
	g.vexs = make([]rune, vexNum)
	g.arcs = make([][]int, vexNum)
	for i := 0; i < vexNum; i++ {
		g.arcs[i] = make([]int, vexNum)
	}
	for i := 0; i < vexNum; i++ {
		for j := 0; j < vexNum; j++ {
			g.arcs[i][j] = INF
		}
	}
	copy(g.vexs, vers)
}

// 通过[]*edge 输入图的所有边
func (g *AMGraph) Arcs(edges Edges) {
	g.arcNum = len(edges)
	for i := 0; i < len(edges); i++ {
		v1 := g.locateVex(edges[i].v1)
		v2 := g.locateVex(edges[i].v2)
		g.arcs[v1][v2] = edges[i].weight
		g.arcs[v2][v1] = edges[i].weight
	}
}

// 通过顶点找到其对应的下标
func (g *AMGraph) locateVex(vex rune) int {
	for i, v := range g.vexs {
		if vex == v {
			return i
		}
	}
	return -1
}

// 打印邻接矩阵表示的图
func (g *AMGraph) PrintAMGraph() {
	fmt.Printf("\t")
	for _, v := range g.vexs {
		fmt.Printf("%c\t", v)
	}
	fmt.Println()
	for i, vexs := range g.arcs {
		fmt.Printf("%c\t", g.vexs[i])
		for _, v := range vexs {
			if v == INF {
				fmt.Printf("%c\t", 'X')
			} else {
				fmt.Printf("%d\t", v)
			}
		}
		fmt.Println()
	}
}

// vertex : 起始顶点. 即计算"顶点vertex"到其它顶点的最短路径及长度。
func (g *AMGraph) Dijkstra(vertex rune) {
	sv := g.locateVex(vertex)
	// 前驱顶点数组。即，prev[i]的值是"顶点sv"到"顶点i"的最短路径所经历的全部顶点中，位于"顶点i"之前的那个顶点。
	prev := make([]int, g.vexNum)

	// 长度数组。即，dist[i]是"顶点sv"到"顶点i"的最短路径的长度。
	dist := make([]int, g.vexNum)

	// visited[i]=true表示"顶点sv"到"顶点i"的最短路径已成功获取。
	visited := make([]bool, g.vexNum)

	// 初始化
	for i := 0; i < g.vexNum; i++ {
		dist[i] = g.arcs[sv][i] //  顶点i的最短路径为"顶点sv"到"顶点i"的权。
	}
	// 对顶点sv 自身进行初始化
	visited[sv] = true
	dist[sv] = 0

	// 遍历g.vexNum-1次,每次找出一个顶点的最短路径
	for i := 1; i < g.vexNum; i++ {
		var k int
		min := INF
		// 寻找当前最短路径
		// 即，在未获取最短路径的顶点中，找到离sv最近的顶点(k)。
		for j := 0; j < g.vexNum; j++ {
			if !visited[j] && dist[j] < min {
				min = dist[j]
				k = j
			}
		}
		// 标记"顶点k"为已经获取到最短路径
		visited[k] = true

		// 修正当前最短路径和前驱顶点
		// 即，当已经得到"顶点k的最短路径"之后，更新"未获取最短路径的顶点的最短路径和前驱顶点"。
		for j := 0; j < g.vexNum; j++ {
			var tmp int
			if g.arcs[k][j] == INF { //防溢出
				tmp = INF
			} else {
				tmp = min + g.arcs[k][j]
			}
			if !visited[j] && tmp < dist[j] {
				dist[j] = tmp
				prev[j] = k
			}
		}
	}

	fmt.Printf("顶点%c到各顶点的最短路径为:\n", vertex)
	for i := 0; i < g.vexNum; i++ {
		if i == sv {
			continue
		}
		fmt.Printf("%c->%c: %d\n", vertex, g.vexs[i], dist[i])
		fmt.Printf("路径为:")
		path(g.vexs, dist, prev, i, sv)
		fmt.Printf("%c\n", g.vexs[i])
	}
}

func path(vexs []rune, dist, prev []int, end int, start int) {
	if prev[end] == 0 {
		fmt.Printf("%c--[%d]->", vexs[start], dist[end])
		return
	}
	path(vexs, dist, prev, prev[end], start)
	fmt.Printf("%c--[%d]->", vexs[prev[end]], dist[end])
}
