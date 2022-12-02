/*
 * @Author: wzmiiiiii
 * @Date: 2022-12-02 20:00:05
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-12-03 02:04:57
 */

// 基于尚硅谷java数据结构与算法p173
package mst

import (
	"fmt"
	"math"
	"sort"
)

// 定义最大的权
const INF = math.MaxInt

// 通过邻接矩阵构建图
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

// 定义 最小生成树结构体
type MST struct {
	*AMGraph       // 通过邻接矩阵表示最小生成树
	weight   int   // 最小生成树的权
	Edges    Edges // 构成最小生成树的边的集合
}

func (m MST) GetWeight() int {
	return m.weight
}

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
	es = edges
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

// 通过Prim算法生成最小生成树
func (g *AMGraph) MSTPrim(start rune) *MST {
	graph := NewAMGraph()
	mst := &MST{AMGraph: graph}
	mst.vexNum = g.vexNum
	mst.arcNum = g.vexNum - 1
	mst.Vertexes(g.vexs)

	visited := make([]bool, g.vexNum) //标记顶点是否访问

	startV := g.locateVex(start)

	visited[startV] = true // 标记开始顶点已访问

	minWgt := INF        //记录子图中试探到的最小权重
	var minV1, minV2 int //记录最小权重关联的顶点
	var edges Edges      //记录最小生成树的边v
	var vexs []int       //记录所有访问过的顶点
	var minTotalWgt int  //记录生成树总权重
	vexs = append(vexs, startV)

	for len(edges) != g.vexNum-1 { // 最小生成树的边为顶点数-1,
		for _, i := range vexs { // 循环所有已访问过的顶点
			for j := 0; j < g.vexNum; j++ { // j 表示没被访问过的顶点的下标
				// 如果g.vexs[i] 与 g.vexs[j] 不相邻直接跳过
				// 基于已访问过的顶点,寻找没被访问过且权重最小的顶点
				if g.arcs[i][j] != INF && !visited[j] && g.arcs[i][j] < minWgt {
					minWgt = g.arcs[i][j] // 记录当前最小权重
					// 记录当前最小权重相应的顶点
					minV1 = i
					minV2 = j
				}
			}
		}
		minTotalWgt += minWgt
		vexs = append(vexs, minV2) // 将找到的新顶点加入vexs
		edge := NewEdge(g.vexs[minV1], g.vexs[minV2], minWgt)
		edges = append(edges, edge) // 将新找到的边加入edges
		visited[minV2] = true       // 将找到的新顶点标记为已访问
		minWgt = INF                // 重置minWgt
	}
	mst.weight = minTotalWgt
	mst.Arcs(edges)
	mst.Edges = edges
	return mst
}

var es Edges // 记录输入的边切片

// 思路:
// 加入一个边之前,判断边的两个顶点是否能到达同一个终点.
// 如果不能到达同一个终点,则可以加入该边.否则不能加入.

// 通过Kruskal算法生成最小生成树
func (g *AMGraph) MSTKruskal() *MST {
	graph := NewAMGraph()
	mst := &MST{AMGraph: graph}
	mst.vexNum = g.vexNum
	mst.arcNum = g.vexNum - 1
	mst.Vertexes(g.vexs)

	sort.Sort(es)                    //对 边切片进行排序.
	ends := make([]int, g.vexNum)    //使用ends本身的下标对应一个顶点的下标,ends[i]记录的值为顶点i第一次终点的下标.
	for i := 0; i < len(ends); i++ { //初始化ends
		ends[i] = INF
	}
	var edges Edges     //记录最小生成树的边
	var totalWeight int //记录最小生成树的总权重

	// 对边切片排序后,每次循环都是得到当前状态下权重最小的边
	for i := 0; len(edges) != mst.arcNum; i++ { //生成树中有顶点数-1条边时,结束循环
		v1Index := g.locateVex(es[i].v1) // 获取第i条边的第一个顶点的下标
		v2Index := g.locateVex(es[i].v2) // 获取第i条边的第二个顶点的下标

		// 获取v1Index,v2Index这两个顶点在已有最小生成树中的最终终点.
		v1End := getEnd(ends, v1Index)
		v2End := getEnd(ends, v2Index)

		// 判断是否构成回路
		if v1End != v2End {
			ends[v1End] = v2End // Kruskal算法的重难点. 妙!!
			edges = append(edges, es[i])
			totalWeight += es[i].weight
		}

	}
	mst.Arcs(edges)
	mst.Edges = edges
	mst.weight = totalWeight
	return mst
}

// Kruskal算法的重难点. 妙!!
// 获取下标为i的顶点的最终终点的下标
// ends: 使用ends本身的下标对应一个顶点的下标,ends[i]记录的值为顶点i第一次终点的下标.
// ends的形式为: ends[6,5,3,5,5,6,INF]
func getEnd(ends []int, i int) int {
	for ends[i] != INF { //通过循环找到i顶点的最终终点
		i = ends[i]
	}
	return i
}
