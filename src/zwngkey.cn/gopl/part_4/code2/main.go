package main

import "fmt"

type tree struct {
	value       int
	left, right *tree
}

//用一个二叉树来实现一个插入排序：
func Sort(values []int) {
	var root *tree

	for _, v := range values {
		root = add(root, v)
	}
	fmt.Printf("%#v\n", *root)

	appendValues(values[:0], root)

}

func add(t *tree, value int) *tree {
	if t == nil {
		// Equivalent to return &tree{value: value}.
		t = new(tree)
		t.value = value
		return t
	}

	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}

	return values
}

func main() {
	// var root *tree = &tree{value: 1}
	// fmt.Printf("root.value: %T\n", root.value)
	// fmt.Printf("root: %T\n", root)
	// fmt.Println(root)
	s := []int{143, 23, 5, 81, 72, 89, 49, 22, 71, 151, 13}
	// fmt.Printf("s: %v\n", s[:0])
	fmt.Printf("s: %v\n", s)
	Sort(s)
	fmt.Printf("s: %v\n", s)

}
