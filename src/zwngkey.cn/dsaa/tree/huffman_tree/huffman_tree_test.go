/*
 * @Author: wzmiiiiii
 * @Date: 2022-11-11 20:22:44
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-11-11 22:33:21
 * @Description:
 */
package huffman

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	weights := []int{1, 1, 1, 2, 2, 2, 4, 4, 4, 5, 5, 9}
	// weights := []float64{0.4,0.3,0.15,0.05,0.04,0.03,0.03}
	ht := CreateHuffmanTree(weights)
	for i := 1; i < len(ht); i++ {
		fmt.Println(i, ht[i])
	}
	hc := CreateHuffmanCode(ht, len(weights))
	for i := 1; i < len(hc); i++ {
		fmt.Printf("%s\n", hc[i].Code)
	}
}
