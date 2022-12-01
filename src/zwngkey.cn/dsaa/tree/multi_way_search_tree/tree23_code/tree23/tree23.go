/*
 * @Author: wzmiiiiii
 * @Date: 2022-12-02 02:20:28
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-12-02 02:53:29
 */
package tree23

const ROOT_IS_BIGGER = 1
const ROOT_IS_SMALLER = -1

type Tree23[T comparable] struct {
	root     *node[T]
	size     int  // Number of elements inside of the tree
	addition bool // A flag to know if the last element has been added correctly or not
}

func NewTree23[T comparable]() *Tree23[T] {
	return &Tree23[T]{
		root: newEmptyNode[T](),
		size: 0,
	}
}

func (tr *Tree23[T]) Add(element T) bool {
	var t T
	tr.size++
	tr.addition = false
	if tr.root == nil || tr.root.leftElement == t {
		if tr.root == nil {
			tr.root = newEmptyNode[T]()
		}
		tr.root.leftElement = element
		tr.addition = true
	} else {
		newRoot := tr.addElementI(tr.root, element)
		if newRoot != nil {
			tr.root = newRoot
		}
	}
	if !tr.addition {
		tr.size--
	}
	return tr.addition
}

func (tr *Tree23[T]) AddAll(elements []T) bool {
	ok := true
	for _, e := range elements {
		if !tr.Add(e) {
			ok = false
		}
	}
	return ok
}

func (tr *Tree23[T]) AddAllSafe(elements []T) bool {
	inserted, i := 0, 0
	for _, e := range elements {
		if !tr.Add(e) {
			for _, a := range elements {
				if i >= inserted {
					return false
				} else {
					tr.Remove(a)
				}
			}
		} else {
			inserted++
		}
	}
	return true
}

func (tr *Tree23[T]) addElementI(cur *node[T], element T) *node[T] {
	var newParent *node[T]
	if !cur.isLeaf() {
		var sonAscended *node[T]
		_ = sonAscended
		if cur.leftElement == element || (cur.is3Node() && cur.rightElement == element) {
			// ...
		}
	}
	return newParent
}

func (tr *Tree23[T]) Remove(element T) {

}
