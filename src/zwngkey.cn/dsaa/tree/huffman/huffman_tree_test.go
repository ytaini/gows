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
	ht := CreateHuffmanTree[int](5)
	for i := 1; i < len(ht); i++ {
		fmt.Println(i, ht[i])
	}
}
