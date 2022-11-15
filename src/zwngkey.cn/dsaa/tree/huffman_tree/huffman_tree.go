/*
 * @Author: wzmiiiiii
 * @Date: 2022-11-11 19:22:03
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-11-15 00:12:31
 * @Description:
	哈夫曼树构造算法的实现
		通过一维结构数组.
*/
package huffman

import (
	"math"
)

type Number interface {
	int | float64
}
type HTNode[T Number] struct {
	weight                 T
	parent, lchild, rchild int
}

type HuffmanTree[T Number] []HTNode[T]

// 构造哈夫曼树--哈夫曼算法
func CreateHuffmanTree[T Number](input []T) (ht HuffmanTree[T]) {
	n := len(input)
	if n <= 1 {
		return nil
	}
	m := 2*n - 1 //数组共2n-1个元素

	ht = make([]HTNode[T], m+1) //0号单元未使用,ht[m]表达根节点.
	// var w T = 0
	for i := 1; i <= n; i++ { //输入前n个元素的weight值.
		// fmt.Scan(&w)
		// ht[i].weight = ww
		ht[i].weight = input[i-1]
	}
	// 初始化结束,下面建立哈夫曼树.
	for i := n + 1; i <= m; i++ { //合并产生n-1个结点--构造huffman树.
		s1, s2 := Select(ht, i-1) //在ht[k](1≤k≤i-1)中选择两个其双亲域为0,且权值最小的结点,并返回他们的下标.
		ht[s1].parent = i         //表示从F中删除s1,s2.
		ht[s2].parent = i
		ht[i].weight = ht[s1].weight + ht[s2].weight //i的权值为左右孩子权值之和.
		ht[i].lchild = s1                            //s1,s2分别作为i的左右孩子
		ht[i].rchild = s2
	}
	return ht
}

// 找出最小的两个值的下标.
func Select[T Number](ht HuffmanTree[T], n int) (m1, m2 int) {
	var firstmin T = (T)(math.MaxInt)
	var secondmin T = firstmin
	for i := 1; i <= n; i++ {
		if ht[i].parent == 0 && ht[i].weight < firstmin {
			firstmin = ht[i].weight
			m1 = i
		}
	}
	ht[m1].parent = 1

	for i := 1; i <= n; i++ {
		if ht[i].parent == 0 && ht[i].weight < secondmin {
			secondmin = ht[i].weight
			m2 = i
		}
	}
	ht[m1].parent = 0
	return
}
