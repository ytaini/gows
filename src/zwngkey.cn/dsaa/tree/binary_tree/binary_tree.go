/*
 * @Author: wzmiiiiii
 * @Date: 2022-11-12 01:30:55
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-11-12 01:41:21
 * @Description:
	二叉树的遍历
*/
package binarytree

// 二叉树的二叉链表存储表示
// 二叉链表，由于缺乏父链的指引，在找回父节点时需要重新扫描树得知父节点的节点地址。
type BiTNode struct {
	data           any
	lchild, rchild *BiTNode //左右孩子指针
}
