/*
  - @Author: wzmiiiiii
  - @Date: 2022-11-16 21:31:25

* @LastEditors: wzmiiiiii
* @LastEditTime: 2022-12-06 02:54:24
  - @Description:
    使用回溯法得到哈夫曼编码
*/
package bt

type TreeNode struct {
	data  byte
	times int
	left  *TreeNode
	right *TreeNode
}

// 使用回溯法得到经过哈夫曼树编码以后的哈夫曼编码
// encodeMap中key为字符的ASCII码值，用byte表示，
// value为哈夫曼树从根节点到叶子节点的左右路径编码值，用string表示，
// 其中tmp记录递归过程中所走过的路径，往左走用'0'表示，往右走用'1'表示，最后再转换成string就得到了对应ASCII码的编码值。
// node: 哈夫曼树root结点
func CreateEncodingTable(node *TreeNode) map[byte]string {
	encodeTable := make(map[byte]string)
	tmp := make([]byte, 0)
	var depth func(*TreeNode)
	depth = func(tn *TreeNode) {
		if tn == nil {
			return
		}
		if tn.left == nil && tn.right == nil {
			encodeTable[tn.data] = string(tmp)
		}
		tmp = append(tmp, '0')
		depth(tn.left)
		tmp[len(tmp)-1] = '1' //回溯
		depth(tn.right)
		tmp = tmp[:len(tmp)-1] //回溯
	}
	depth(node)
	return encodeTable
}
