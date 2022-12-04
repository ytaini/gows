/*
  - @Author: wzmiiiiii
  - @Date: 2022-12-03 22:45:01

* @LastEditors: wzmiiiiii
* @LastEditTime: 2022-12-04 22:33:06
*/
package btree

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

type position int

const (
	left  = position(-1)
	none  = position(0)
	right = position(1)
)

type Int int64

func (i Int) Less(j Item) bool {
	return i < j.(Int)
}

type Item interface {
	Less(than Item) bool
}

type items []Item

// @bool : 如果返回true,表示items中存在key值.
// @int : 如果items存在key值.返回其下标.
// - 否则返回key插入的位置.
func (s items) find(key Item) (int, bool) {
	// sort.Search会利用二分查找在items中找到第一个满足 > key的Item的下标，
	// 这也就说明(key < s[i-1])是不成立的（值为 false），
	// 这时再进行一次(s[i-1] < item)判断，如果也不成立，那么就可以认为 item 与 s[i-1] 是相等的。
	index := sort.Search(len(s), func(i int) bool {
		return key.Less(s[i])
	})
	if index > 0 && !s[index-1].Less(key) {
		return index - 1, true
	}
	return index, false
}

type node struct {
	items    items   //关键字 [K1,K2,K3]
	children []*node //子节点 [A1,A2,A3]
	parent   *node
}

func createNode(m int) *node {
	return &node{}
}

// 在整个B树中搜索key,存在返回true
func (n *node) isExist(key Item) bool {
	i, found := n.items.find(key)
	if found {
		return found
	} else if len(n.children) > 0 {
		return n.children[i].isExist(key)
	}
	return found
}

// 在整颗B树中搜索key需要被插入到哪个节点.返回这个节点,和key在这个节点的索引
func (n *node) getKey(key Item) (*node, int) {
	i, found := n.items.find(key)
	if found {
		return n, i
	} else if len(n.children) > 0 {
		return n.children[i].getKey(key)
	}
	return nil, 0
}

// 获取n节点的左右兄弟节点
func (n *node) getBrotherNodes() (l, r *node) {
	if n.parent == nil {
		return nil, nil
	}
	index := 0
	for i, pn := range n.parent.children {
		if pn == n {
			index = i
			break
		}
	}
	if index == 0 {
		return nil, n.parent.children[1]
	} else {
		if len(n.parent.children) == index+1 {
			return n.parent.children[index-1], nil
		} else {
			return n.parent.children[index-1], n.parent.children[index+1]
		}
	}
}

// 要删除一个结点的数据元素Ki（1≤i≤n），
// 首先寻找该结点Ai所指子树中的最小元素Kmin（Ai所指子树中的最小数据元素Kmin一定在叶子结点上）
// 返回该叶子节点
func (n *node) getSuccessorNode(index int) *node {
	ln := n.children[index+1]
	if len(ln.children) == 0 {
		return ln
	}

	for {
		if ln.children == nil {
			break
		}
		ln = ln.children[0]
	}
	return ln
}

// 往bt中插入key.
func (n *node) insert(bt *BTree, key Item) {
	// 找到key要在n.items中插入的位置i.
	i := sort.Search(len(n.items), func(i int) bool {
		return key.Less(n.items[i])
	})
	// 将key插入到n.items指定位置
	newItems := append(n.items[0:i], append(items{key}, n.items[i:]...)...)
	n.items = newItems
	// 若加入key后,bt不满足B树属性,split方法使其平衡.
	bt.split(n)
}

// 删除n节点的n.items的index位置的值. 返回该值.
func (n *node) delete(index int) Item {
	key := n.items[index]
	if index == 0 {
		n.items = append(items{}, n.items[1:]...)
	} else {
		n.items = append(n.items[0:index], n.items[index+1:]...)
	}
	return key
}

// 移除一个孩子节点
// TODO 下标越界可能
func (n *node) deleteChild(index int) {
	if index == 0 { //删除第一个孩子结点
		n.children = append([]*node{}, n.children[1:]...)
	} else {
		n.children = append(n.children[0:index], n.children[index+1:]...)
	}
}

// 最前后插入一个子节点
func (n *node) addChild(move *node, p position) {
	if p == left {
		n.children = append([]*node{move}, n.children...)
	}
	if p == right {
		n.children = append(n.children, move)
	}
	move.parent = n
}

// 获取node在父节点中的index
func (n *node) getIndexInParent() int {
	if n.parent == nil {
		panic("n is root node")
	}
	index := 0
	for i, pn := range n.parent.children {
		if pn == n {
			index = i
			break
		}
	}
	return index
}

type BTree struct {
	root     *node
	m        int //阶
	minItems int //除根节点外单个节点最少包含的关键字数量
	maxItems int //单个节点最多包含的关键字数量
}

// NewBtree 创建m阶B树
func NewBtree(m int) *BTree {
	if m < 3 {
		panic("m needs >= 3")
	}
	bt := &BTree{}
	bt.root = createNode(bt.m)
	bt.m = m
	bt.minItems = int(math.Ceil(float64(m)/2) - 1)
	bt.maxItems = m - 1
	return bt
}

// 查找是否存在
func (bt *BTree) IsExist(key Item) bool {
	return bt.root.isExist(key)
}

// Insert 插入数据
func (bt *BTree) Insert(key Item) {
	if bt.IsExist(key) { //不能插入相同的key
		return
	}
	node := getLeafNode(bt.root, key)
	node.insert(bt, key)
}

// 批量插入
func (bt *BTree) InsertMultiple(keys []int) {
	for _, v := range keys {
		bt.Insert(Int(v))
	}
}

func (bt *BTree) DeleteAll(keys []int) {
	for _, v := range keys {
		bt.Delete(Int(v))
	}
}

// 删除
func (bt *BTree) Delete(key Item) {
	n, i := bt.root.getKey(key)
	if n == nil {
		return
	}
	if len(n.children) == 0 && n.parent == nil { //只有一个根节点
		n.delete(i)
		return
	}
	if len(n.children) == 0 && n.parent != nil && len(n.items) > bt.minItems { //叶子节点 & 节点数量充足
		n.delete(i)
		return
	}
	if len(n.children) == 0 && n.parent != nil && len(n.items) <= bt.minItems { //叶子节点、节点不足
		n.delete(i)
		bt.reBalance(n)
		return
	}
	// 在非叶结点上删除数据元素的算法思想：假设要删除一个结点的数据元素Ki（1≤i≤n），
	// 首先寻找该结点Ai所指子树中的最小元素Kmin（Ai所指子树中的最小数据元素Kmin一定为叶结点上），
	// 然后用kmin覆盖要删除的数据元素Ki，最后再以指针Ai所指结点为根结点查找并删除Kmin
	// （即再以Ai所指结点为B树的根结点，以Kmin为要删除数据元素再次调用B树上的删除算法）。
	// 这样就把非叶结点上的删除问题转化成了叶结点上的删除问题
	if len(n.children) != 0 { // 非叶子节点
		sn := n.getSuccessorNode(i)
		n.items[i] = sn.items[0] //后继替换
		sn.delete(0)             //后继key删除
		bt.reBalance(sn)
		return
	}
}

// 合并两个节点
func (bt *BTree) mergeNodes(l, r *node) {
	ix := l.getIndexInParent()
	k := l.parent.delete(ix)
	rChildren := r.children

	l.items = append(l.items, k)
	l.items = append(l.items, r.items...)
	l.parent.deleteChild(ix + 1)
	// 右子分支合并到左分支中
	if len(l.children) > 0 {
		for _, v := range rChildren {
			v.parent = l
		}
		l.children = append(l.children, rChildren...)
	}
	if l.parent.parent == nil && len(l.parent.items) == 0 { //根节点
		l.parent.children = make([]*node, 0)
		l.parent = nil
		bt.root = l
	} else {
		bt.reBalance(l.parent)
	}

}

// 再平衡
func (bt *BTree) reBalance(n *node) {
	if n.parent == nil { //根节点
		return
	}
	if len(n.items) >= bt.minItems {
		return
	}
	// 判断兄弟节点是否可借
	var bn *node
	p := none
	l, r := n.getBrotherNodes()
	if l != nil && len(l.items) > bt.minItems {
		bn = l
		p = left
	}
	if bn == nil && r != nil && len(r.items) > bt.minItems {
		bn = r
		p = right
	}
	if bn != nil { //可以借
		var key, pk Item
		ix := bn.getIndexInParent()
		if p == left {
			key = bn.delete(len(bn.items) - 1)
			pk = bn.parent.items[ix]
			bn.parent.items[ix] = key
			n.items = append(items{pk}, n.items...)

			//  借完需要处理子节点的归属问题
			if len(bn.children) > 0 {
				moveChild := bn.children[len(bn.children)-1]
				bn.deleteChild(len(bn.children) - 1)
				n.addChild(moveChild, left)
			}
		} else {
			key = bn.delete(0)
			pk = bn.parent.items[ix-1]
			bn.parent.items[ix-1] = key
			n.items = append(n.items, pk)
			//借完需要处理子节点的归属问题
			if len(bn.children) > 0 {
				moveChild := bn.children[0]
				bn.deleteChild(0)
				n.addChild(moveChild, right)
			}
		}

	} else { //不可以借
		if l != nil {
			bt.mergeNodes(l, n)
		} else {
			bt.mergeNodes(n, r)
		}
	}
}

// 插入拆分
func (bt *BTree) split(n *node) {
	if len(n.items) <= bt.maxItems {
		return
	}
	middle := bt.m / 2
	l := make(items, middle)
	r := make(items, len(n.items)-middle-1)
	copy(l, n.items[:middle])
	copy(r, n.items[middle+1:])
	middleItem := n.items[middle]
	n.items = l

	// 右节点
	rNode := createNode(bt.m)
	rNode.items = r

	if n.parent != nil { //有父节点
		i := n.getIndexInParent()
		if i == 0 {
			n.parent.items = append(items{middleItem}, n.parent.items[:]...)
			n.parent.children = append(n.parent.children[:1], append([]*node{rNode}, n.parent.children[1:]...)...)
		} else {
			n.parent.items = append(n.parent.items[0:i], append(items{middleItem}, n.parent.items[i:]...)...)
			n.parent.children = append(n.parent.children[:i+1], append([]*node{rNode}, n.parent.children[i+1:]...)...)
		}
		rNode.parent = n.parent
		if len(n.children)-1 > len(n.items) {
			temp := make([]*node, len(n.children))
			copy(temp, n.children)
			n.children = n.children[:len(n.items)+1]
			rNode.children = append(rNode.children, temp[len(n.items)+1:]...)
			for _, node := range rNode.children {
				node.parent = rNode
			}
		}
		// 递归处理父节点
		bt.split(n.parent)
	} else { // 没有父节点
		newRoot := createNode(bt.m)
		newRoot.items = append(newRoot.items, middleItem)
		newRoot.children = append(newRoot.children, n, rNode)
		bt.root = newRoot
		n.parent = newRoot
		rNode.parent = bt.root

		if len(n.children) > 0 { //处理被拆分节点的子节点
			rChildren := make([]*node, bt.m-middle)
			copy(rChildren, n.children[middle+1:])
			n.children = n.children[0 : middle+1]
			rNode.children = rChildren
			for _, v := range rNode.children {
				v.parent = rNode
			}
		}
	}
}

// 以root为根节点,返回可以插入key的LeafNode
func getLeafNode(root *node, key Item) *node {
	if len(root.children) == 0 {
		return root
	}
	for index, item := range root.items {
		if key.Less(item) {
			return getLeafNode(root.children[index], key)
		}
	}
	return getLeafNode(root.children[len(root.items)], key)
}

func (bt *BTree) PrintBTree() {
	printBTree(bt.root, 0)
}

func printBTree(n *node, depth int) {
	var buf strings.Builder
	for i := 1; i < depth; i++ {
		buf.WriteString("|    ")
	}
	if depth > 0 {
		buf.WriteString("|----")
	}
	for i, v := range n.items {
		if i == 0 {
			buf.WriteString(fmt.Sprintf("[%v", v))
		} else if i == len(n.items)-1 {
			buf.WriteString(fmt.Sprintf(" %v", v))
		} else {
			buf.WriteString(fmt.Sprintf(" %v", v))
		}
	}
	buf.WriteString("]")
	fmt.Println(buf.String())
	for _, v := range n.children {
		printBTree(v, depth+1)
	}
}

func (bt *BTree) IsBTree() bool {
	return isBtree(bt, bt.root)
}

func isBtree(bt *BTree, node *node) bool {
	if len(node.items) > bt.maxItems || len(node.children) > bt.m {
		return false
	}
	flag := true
	for _, n := range node.children {
		if !isBtree(bt, n) {
			flag = false
			break
		}
	}
	return flag
}
