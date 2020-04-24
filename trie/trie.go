package trie

import "github.com/chentaihan/container/heap"

/**
前缀树
*/
type TrieNode struct {
	Count    int
	Val      byte
	IsEnd    bool
	Children map[byte]*TrieNode
}

type Trie struct {
	root  *TrieNode
	count int
}

func NewTrie() ITrie {
	return &Trie{root: &TrieNode{
		Children: map[byte]*TrieNode{},
	}}
}

//添加元素
func (trie *Trie) Add(word string) {
	node := trie.root
	for i := 0; i < len(word); i++ {
		index := word[i]
		if node.Children[index] == nil {
			node.Children[index] = &TrieNode{
				Count:    1,
				Val:      word[i],
				IsEnd:    false,
				Children: map[byte]*TrieNode{},
			}
		} else {
			node.Children[index].Count++
		}
		node = node.Children[index]
	}
	if !node.IsEnd {
		trie.count++
	}
	node.IsEnd = true
}

//查找元素是否存在
func (trie *Trie) Find(word string) bool {
	node := trie.root
	for i := 0; i < len(word); i++ {
		index := word[i]
		if node.Children[index] == nil {
			return false
		}
		node = node.Children[index]
	}
	return node.IsEnd
}

//指定单词出现的次数
func (trie *Trie) Count(word string) int {
	node := trie.root
	for i := 0; i < len(word); i++ {
		index := word[i]
		if node.Children[index] == nil {
			return 0
		}
		node = node.Children[index]
	}
	return node.Count
}

//是否存在以指定的单词开头的单词
func (trie *Trie) StartWith(word string) bool {
	node := trie.root
	for i := 0; i < len(word); i++ {
		index := word[i]
		if node.Children[index] == nil {
			return false
		}
		node = node.Children[index]
	}
	return true
}

//删除指定的单词
func (trie *Trie) Remove(word string) bool {
	if !trie.Find(word) {
		return false
	}
	node := trie.root
	for i := 0; i < len(word); i++ {
		index := word[i]
		if node.Children[index] == nil {
			return false
		}
		node = node.Children[index]
		node.Count--
	}
	if node.Count == 0 {
		node.IsEnd = false
	}
	trie.count--
	return true
}

//单词数量
func (trie *Trie) GetWordCount() int {
	return trie.count
}

//出现次数最多的N个单词（倒序）
func (trie *Trie) GetTopN(n int) []string {
	if n <= 0 {
		return []string{}
	}
	h := heap.NewSmallHeap(n)
	trie.getTopN(trie.root, []byte{}, h)
	for h.Len() > n {
		h.Pop()
	}

	result := make([]string, h.Len())
	for index := h.Len() - 1; index >= 0; index-- {
		result[index] = h.Pop().(trieItem).word
	}
	return result
}

type trieItem struct {
	word  string
	count int
}

func (ti trieItem) GetValue() int {
	return ti.count
}

func (trie *Trie) getTopN(root *TrieNode, prefix []byte, h heap.IHeap) {
	if root != nil {
		prefix = append(prefix, root.Val)
		if root.IsEnd {
			h.Push(trieItem{
				string(prefix[1:]),
				root.Count,
			})
		}
		for _, value := range root.Children {
			trie.getTopN(value, prefix, h)
		}
	}
}
