/*
 * @Author: wzmiiiiii
 * @Date: 2022-12-05 03:06:48
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-12-05 03:06:59
 */
/*
 * @Author: wzmiiiiii
 * @Date: 2022-12-05 00:31:32
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-12-05 01:52:58
 */
package floyd

import (
	"testing"
)

func Test(t *testing.T) {
	graph := NewAMGraph()
	graph.Vertexes([]rune{'A', 'B', 'C', 'D', 'E', 'F', 'G'})
	edges := []*Edge{
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
	graph.Arcs(edges)
	graph.PrintAMGraph()
	graph.Floyd()
}
