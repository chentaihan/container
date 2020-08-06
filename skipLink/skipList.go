package skipLink

/**
跳跃表实现的是一个多层的单链表
一个元素一个节点，每个节点可能有多个指向后面节点的指针，每层的指针指向离他最近且高度大于等于他的节点
每个指针都是指向一个节点，而不是节点的某一层
*/

import (
	"math/rand"
)

const maxLevel = 32

type skipListNode struct {
	key   int
	val   interface{}
	nexts []*skipListNode
}

type SkipList struct {
	head  *skipListNode
	level int
	size  int
}

func NewSkipList() ISkipList {
	sl := &SkipList{}
	sl.head = sl.createNode(-1, nil, maxLevel)
	sl.level = 1
	return sl
}

func (sl *SkipList) Add(key int, val interface{}) {
	nodeList := sl.find(key)
	if nodeList[0].nexts[0] != nil && nodeList[0].nexts[0].key == key {
		nodeList[0].nexts[0].val = val
		return
	}
	level := sl.randomLevel()
	newNode := sl.createNode(key, val, level)
	minLevel := level
	if len(nodeList) < minLevel {
		minLevel = len(nodeList)
	}
	for i := 0; i < minLevel; i++ {
		newNode.nexts[i] = nodeList[i].nexts[i]
		nodeList[i].nexts[i] = newNode
	}
	if level > sl.level {
		for i := sl.level; i < level; i++ {
			sl.head.nexts[i] = newNode
		}
		sl.level = level
	}
	sl.size++
}

func (sl *SkipList) find(key int) []*skipListNode {
	buf := make([]*skipListNode, sl.level)
	cur := sl.head
	for i := sl.level - 1; i >= 0; i-- {
		for cur.nexts[i] != nil && cur.nexts[i].key < key {
			cur = cur.nexts[i]
		}
		buf[i] = cur
	}
	return buf
}

func (sl *SkipList) Find(key int) (interface{}, bool) {
	nodeList := sl.find(key)
	if nodeList[0].nexts[0] != nil && nodeList[0].nexts[0].key == key {
		return nodeList[0].nexts[0].val, true
	}
	return nil, false
}

func (sl *SkipList) Remove(key int) (interface{}, bool) {
	nodeList := sl.find(key)
	var value interface{}
	exist := false
	for i := 0; i < len(nodeList); i++ {
		if nodeList[i].nexts[i] == nil {
			continue
		}
		next := nodeList[i].nexts[i]
		if next.key == key {
			value = next.val
			exist = true
			nodeList[i].nexts[i] = next.nexts[i]
		}
	}
	if exist {
		sl.size--
	}
	return value, exist
}

func (sl *SkipList) Len() int {
	return sl.size
}

func (sl *SkipList) Clear() {
	for i := 0; i < sl.level; i++ {
		sl.head.nexts[i] = nil
	}
	sl.level = 1
	sl.size = 0
}

func (sl *SkipList) GetKeys() []int {
	list := make([]int, sl.size)
	node := sl.head.nexts[0]
	for i := 0; i < sl.size; i++ {
		list[i] = node.key
		node = node.nexts[0]
	}
	return list
}

func (sl *SkipList) GetValues() []interface{} {
	list := make([]interface{}, sl.size)
	node := sl.head.nexts[0]
	for i := 0; i < sl.size; i++ {
		list[i] = node.val
		node = node.nexts[0]
	}
	return list
}

func (sl *SkipList) randomLevel() int {
	level := 1
	for rand.Int31n(2) == 1 && level < maxLevel {
		level++
	}
	return level
}

func (sl *SkipList) createNode(key int, val interface{}, level int) *skipListNode {
	return &skipListNode{
		key:   key,
		val:   val,
		nexts: make([]*skipListNode, level),
	}
}

func (sl *SkipList) GetLevel() int {
	return sl.level
}
