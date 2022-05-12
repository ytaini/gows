/*
 * @Author: zwngkey
 * @Date: 2021-11-21 23:21:06
 * @LastEditTime: 2022-05-02 16:50:19
 * @Description:
	题目描述:
		给定一些计算机课程，每个课程都有前置课程，只有完成了前置课程才可以开始当前课程的学习；
		我们的目标是选择出一组课程，这组课程必须确保按顺序学习时，能全部被完成。

	拓扑排序 深度优先遍历
*/
package dfs

import (
	"fmt"
	"sort"
	"testing"
)

func Test1(t *testing.T) {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items []string)

	visitAll = func(items []string) {
		fmt.Printf("items: %v\n", items)
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)

			}
		}
	}
	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	visitAll(keys)
	return order
}
