/*
 * @Author: wzmiiiiii
 * @Date: 2022-11-11 22:43:02
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-11-12 00:52:30
 * @Description:
 */

package main

import (
	"fmt"

	"zwngkey.cn/dsaa/tree/huffman"
)

/*
输入:
0.4
0.3
0.15
0.05
0.04
0.03
0.03
*/

func main() {
	ht := huffman.CreateHuffmanTree[float64](7)
	for i := 1; i < len(ht); i++ {
		fmt.Println(ht[i])
	}
	hc := huffman.CreateHuffmanCode(ht, 7)
	for i := 1; i < len(hc); i++ {
		fmt.Printf("%s\n", hc[i].Code)
	}
}
