package main

import "fmt"

// 字典树节点
type TrieNode struct {
	end int
	path int
	nexts [] *TrieNode
}

// 构造字典树节点
func newTrieNode() *TrieNode {
	return &TrieNode{0,0,make([]*TrieNode, 26)}
}

// 字典树
type Trie struct {
	root *TrieNode
}

// 构造一个字典树
func newTrie() *Trie {
	return &Trie{root:newTrieNode()}
}

func (trie *Trie) insert(word string) {
	if len(word) == 0 {
		return
	}
	node := trie.root
	chs := []byte(word)
	var index byte
	index = 0
	node.path++

	for i := 0; i < len(chs); i++ {
		//fmt.Println(chs[i])
		index = chs[i] - byte('a')
		//fmt.Println(index)
		if node.nexts[index] == nil {
			node.nexts[index] = newTrieNode()
		}
		node = node.nexts[index]
		node.path++
	}
	node.end++
}


func (trie *Trie) search(word string) bool {
	if len(word) == 0 {
		return false
	}

	node := trie.root
	chs := []byte(word)
	var index byte
	index = 0

	for i := 0; i < len(chs); i++ {
		index = chs[i] - byte('a')
		if node.nexts[index] == nil {
			return false
		}
		node = node.nexts[index]
	}
	return node.end != 0
}


func (trie *Trie) delete(word string) {
	if trie.search(word) {
		node := trie.root
		chs := []byte(word)
		node.path++
		var index byte
		index = 0

		for i := 0; i < len(chs); i++ {
			index = chs[i] - byte('a')
			node.nexts[index].path--
			if node.nexts[index].path == 0 {
				node.nexts[index] = nil
				return
			}
			node = node.nexts[index]
		}
		node.end--
	}
}

func (trie *Trie) prefixNumber(pre string) int {
	if len(pre) == 0 {
		return 0
	}

	node := trie.root
	chs := []byte(pre)
	var index byte
	index = 0

	for i := 0; i < len(chs); i++ {
		index = chs[i] - byte('a')
		if node.nexts[index] == nil {
			return 0
		}
		node = node.nexts[index]
	}
	return node.path
}

func main() {
	trie := newTrie()
	trie.insert("abc")
	trie.insert("cat")
	trie.insert("dog")
	trie.insert("deer")
	trie.insert("panda")
	trie.insert("abcd")
	trie.insert("abd")
	trie.insert("bcd")
	trie.insert("b")

	fmt.Println(trie.search("abc"))
	fmt.Println(trie.prefixNumber("abc"))
	fmt.Println(trie.prefixNumber("b"))
	trie.delete("bcd")
	fmt.Println(trie.search("b"))
	fmt.Println(trie.search("bcd"))
}

