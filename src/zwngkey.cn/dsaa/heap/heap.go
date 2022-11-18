/*
 * @Author: wzmiiiiii
 * @Date: 2022-11-18 06:00:36
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-11-18 07:34:39
 * @Description:
	堆
		构建堆的过程叫堆化。

	go 中已经存在heap接口. container/heap包.

	heap 包为任何实现 heap.Interface 的类型提供堆操作。
		堆是一棵树，每个节点都是其子树中的最小值节点。
		树中的最小元素是根，索引为 0。
		同时heap包也是实现优先队列.
			A heap is a common way to implement a priority queue

	package heap

	type Interface interface {
		sort.Interface
		Push(x any) // add x as element Len()
		Pop() any   // remove and return element Len() - 1.
	}
*/
package heap

// Heap 通过数组切片存储二叉树节点
type Heap []int

func NewHeap() *Heap {
	return new(Heap)
}

func NewHeapByArray(seq []int) *Heap {
	n := len(seq)
	h := (*Heap)(&seq)
	for i := n/2 - 1; i >= 0; i-- {
		h.down(i, n)
	}
	return h
}

func (h *Heap) Len() int {
	return len(*h)
}

func (h *Heap) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
}

func (h *Heap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *Heap) Push(v any) {
	*h = append(*h, v.(int))
	i := h.Len() - 1 //新增元素位置
	h.up(i)
}

func (h *Heap) up(i int) {
	for {
		j := (i - 1) / 2 //父节点位置
		// 如果是根节点或者父节点值大于子节点值，则退出循环
		if i == j || !h.Less(i, j) {
			break
		}
		h.Swap(i, j) // 否则交换子节点与父节点，直到父节点值大于子节点
		i = j
	}
}

func (h *Heap) Pop() any {
	n := h.Len() - 1
	h.Swap(0, n)
	h.down(0, n)
	v := (*h)[n]
	*h = (*h)[:n]
	return v
}

func (h *Heap) down(i0, n int) bool {
	i := i0
	for {
		j1 := i*2 + 1 // 左孩子节点索引
		if j1 >= n {
			break
		}
		if j1+1 < n && h.Less(j1+1, j1) {
			j1++
		}
		if h.Less(i, j1) {
			break
		}
		h.Swap(i, j1)
		i = j1
	}
	return i > i0
}

// Remove removes and returns the element at index i from the heap.
// The complexity is O(log n) where n = h.Len().
func (h *Heap) Remove(i int) any {
	n := h.Len() - 1
	if n != i {
		h.Swap(i, n)
		if !h.down(i, n) {
			h.up(i)
		}
	}
	return h.Pop()
}

// Fix 在索引 i 处的元素更改其值后重新建立堆排序。
// 改变索引 i 处元素的值然后调用 Fix 等同于,
// 但比调用Remeve(h,i),然后调用Push(h,i)花销少.
// The complexity is O(log n) where n = h.Len().
func (h *Heap) Fix(i int) {
	if !h.down(i, h.Len()) {
		h.up(i)
	}
}
