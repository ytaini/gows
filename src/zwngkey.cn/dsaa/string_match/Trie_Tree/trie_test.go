/*
 * @Author: wzmiiiiii
 * @Date: 2022-11-18 04:50:19
 * @LastEditors: wzmiiiiii
 * @LastEditTime: 2022-11-18 05:46:41
 * @Description:
	trie树的实现
*/
package trie

import (
	"fmt"
	"testing"
)

// Trie树节点
type TrieNode struct {
	char     rune               //Unicode字符
	isEnding bool               //是否是单词结尾
	children map[rune]*TrieNode //该结点的子节点字典.
}

// 初始化Trie树节点
func NewTrieNode(char rune) *TrieNode {
	return &TrieNode{
		char:     char,
		isEnding: false,
		children: make(map[rune]*TrieNode),
	}
}

// Trie树结构
type Trie struct {
	root *TrieNode
}

// 初始化Trie树
func NewTrie() *Trie {
	return &Trie{NewTrieNode('/')}
}

// 往Trie树中插入一个单词
func (t *Trie) Insert(word string) {
	node := t.root              // 获取根节点
	for _, code := range word { // 以 Unicode 字符遍历该单词
		childNode, ok := node.children[code] // 获取 code 编码对应子节点
		if !ok {
			childNode = NewTrieNode(code)   // 不存在则初始化该节点
			node.children[code] = childNode // 然后将其添加到子节点字典
		}
		node = childNode // 当前节点指针指向当前子节点
	}
	node.isEnding = true // 一个单词遍历完所有字符后将结尾字符打上标记
}

// 在Trie树中查找一个单词
func (t *Trie) Search(word string) bool {
	node := t.root
	for _, code := range word {
		childNode, ok := node.children[code]
		if !ok {
			return false
		}
		node = childNode
	}
	// 如果isEngind=True,表示完全匹配到了
	// 否则说明不能完全匹配，只是前缀
	return node.isEnding
}

func Test(t *testing.T) {
	words := []string{"hello", "he", "her", "hi", "how", "see", "so", "张三", "张三丰"}
	trie := NewTrie()
	for _, v := range words {
		trie.Insert(v)
	}
	target := "张三丰"
	if trie.Search(target) {
		fmt.Printf("包含词\"%s\"\n", target)
	} else {
		fmt.Printf("不包含词\"%s\"\n", target)
	}
}
