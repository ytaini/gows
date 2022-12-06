/*
 * @Author: wzmiiiiii
 * @Date: 2022-12-05 02:57:52
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-12-05 19:28:33
 */
package floyd

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

// 定义 边的切片
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
			if i == j {
				g.arcs[i][j] = 0
			} else {
				g.arcs[i][j] = INF
			}

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

// floyd最短路径。
// - 即，统计图中各个顶点间的最短路径。
func (g *AMGraph) Floyd() {
	// 路径。 path[i][j]=k表示，"顶点i"到"顶点j"的最短路径会经过顶点k。
	path := make([][]int, g.vexNum)

	for i := range path {
		path[i] = make([]int, g.vexNum)
	}
	// 长度数组。即，dist[i][j]=sum表示，"顶点i"到"顶点j"的最短路径的长度是sum。
	dist := make([][]int, g.vexNum)

	for i := range dist {
		dist[i] = make([]int, g.vexNum)
	}

	// 初始化
	for i := 0; i < g.vexNum; i++ {
		for j := 0; j < g.vexNum; j++ {
			dist[i][j] = g.arcs[i][j]
			// path[i][j] = j //path[i][j]表示从Vi到Vj需要经过的点，初始化D[i][j]=j
			path[i][j] = -1
		}
	}

	// 计算最短路径
	for k := 0; k < g.vexNum; k++ {
		for i := 0; i < g.vexNum; i++ {
			for j := 0; j < g.vexNum; j++ {
				if i == j || k == i || k == j {
					continue
				}
				// 如果经过下标为k顶点路径比原两点间路径更短，则更新dist[i][j]和path[i][j]
				var tmp int
				if dist[i][k] == INF || dist[k][j] == INF {
					tmp = INF
				} else {
					tmp = dist[i][k] + dist[k][j]
				}
				if dist[i][j] > tmp {
					dist[i][j] = tmp
					path[i][j] = k
				}
			}
		}
	}
	print(g.vexs, dist)
	print(g.vexs, path)
	fmt.Println(getPath(path, 0, 3)) // 0 5 5 4 4 3

}

// 求路径
func getPath(path [][]int, i, j int) string {
	if path[i][j] == -1 {
		return fmt.Sprintf("<%d,%d>", i, j)
	} else {
		k := path[i][j]
		return getPath(path, i, k) + " " + getPath(path, k, j) + " "
	}
}

func print(vexs []rune, s [][]int) {
	fmt.Printf("\t")
	for _, v := range vexs {
		fmt.Printf("%c\t", v)
	}
	fmt.Println()
	for _, vexs := range s {
		fmt.Printf("%c\t", '-')
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
