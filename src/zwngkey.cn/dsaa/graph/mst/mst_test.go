/*
 * @Author: wzmiiiiii
 * @Date: 2022-12-02 19:59:43
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-12-03 01:06:11
 */
package mst

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	graph := NewAMGraph()
	graph.Vertexes([]rune{'A', 'B', 'C', 'D', 'E', 'F', 'G'})
	arcs := []*Edge{
		NewEdge('A', 'B', 5),
		NewEdge('A', 'C', 7),
		NewEdge('A', 'G', 2),
		NewEdge('G', 'B', 3),
		NewEdge('G', 'E', 4),
		NewEdge('G', 'F', 6),
		NewEdge('C', 'E', 8),
		NewEdge('F', 'E', 5),
		NewEdge('F', 'D', 4),
		NewEdge('B', 'D', 9),
	}
	graph.Arcs(arcs)
	graph.PrintAMGraph()

	mst := graph.MSTPrim('B')
	mst.PrintAMGraph()
	fmt.Println(mst.GetWeight())
}

func Test1(t *testing.T) {
	graph := NewAMGraph()
	graph.Vertexes([]rune{'A', 'B', 'C', 'D', 'E', 'F', 'G'})
	arcs := Edges{
		NewEdge('A', 'B', 12),
		NewEdge('A', 'F', 16),
		NewEdge('A', 'G', 14),
		NewEdge('B', 'C', 10),
		NewEdge('B', 'F', 7),
		NewEdge('C', 'F', 6),
		NewEdge('C', 'D', 3),
		NewEdge('C', 'E', 5),
		NewEdge('D', 'E', 4),
		NewEdge('E', 'F', 2),
		NewEdge('E', 'G', 8),
		NewEdge('F', 'G', 9),
	}
	graph.Arcs(arcs)
	graph.PrintAMGraph()

	mst := graph.MSTKruskal()
	mst.PrintAMGraph()
	fmt.Println(mst.GetWeight())
	for _, v := range mst.Edges {
		fmt.Println(v)
	}
}
