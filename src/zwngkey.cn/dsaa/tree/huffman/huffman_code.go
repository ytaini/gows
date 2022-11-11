/*
 * @Author: wzmiiiiii
 * @Date: 2022-11-11 22:32:38
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-11-12 00:46:04
 * @Description:
	哈夫曼编码的算法实现
*/
package huffman

type HCNode struct {
	Code string
}

type HuffmanCode = []HCNode

func CreateHuffmanCode[T Number](ht HuffmanTree[T], n int) (hc HuffmanCode) {
	//从叶子到根逆向求每个字符的哈夫曼编码,存储在编码表hc中.
	hc = make([]HCNode, n+1)  //初始化编码表hc,0单元不使用
	for i := 1; i <= n; i++ { //逐个字符求哈夫曼编码
		cd := make([]rune, n-1) //临时存放编码的数组
		start := n - 1          //指向临时数组的指针
		c := i                  //记录当前结点
		f := ht[i].parent       //记录当前结点的双亲结点
		for f != 0 {
			start--
			if ht[f].lchild == c { //结点c是f的左孩子,则生成代码0
				cd[start] = '0'
			} else { //结点c是f的右孩子,则生成代码1
				cd[start] = '1'
			}
			c = f //继续向上回溯
			f = ht[f].parent
		}
		hc[i].Code = string(cd) //将求得的编码复制到hc中.
	}
	return
}
