package skipLink

import "math/rand"

type skipListNode struct {
	key   int
	val   IObject
	nexts []*skipListNode
	prev  *skipListNode
	level int
}

type SkipList struct {
	head     *skipListNode
	maxLevel int
	size     int
}

func NewSkipList(maxLevel int) *SkipList {
	if maxLevel < 0 {
		maxLevel = 32
	}
	sl := &SkipList{}
	sl.head = sl.createNode(-1, nil)
	sl.maxLevel = maxLevel
	return sl
}

func (sl *SkipList) Add(key int, val IObject) {
	curNode := sl.find(key)
	if curNode.key == key {
		curNode.val = val
		for i := 0; i < curNode.level; i++ {
			curNode.nexts[i].val = val
		}
		return
	}
	newNode := sl.createNode(key, val)
	sl.headExtend(newNode.level)
	newNode.prev = curNode
	//后一个节点的prev指向新节点
	nextNode := curNode.nexts[0]
	if nextNode != nil {
		nextNode.prev = newNode
		for i := 0; i < newNode.level; i++ {
			nextNode.nexts[i].prev = newNode
		}
	}

	start := 0
	for {
		minLevel := curNode.level
		if newNode.level < minLevel {
			minLevel = newNode.level
		}
		for i := start; i < minLevel; i++ {
			newNode.nexts[i] = curNode.nexts[i]
			newNode.nexts[i].prev = curNode
			curNode.nexts[i] = newNode.nexts[i]
		}
		if minLevel == newNode.level {
			break
		}
		start = minLevel
		curNode = curNode.prev
	}

}

func (sl *SkipList) find(key int) *skipListNode {
	var head, cur *skipListNode = sl.head, nil
	buf := make([]*skipListNode, head.level)
	for i := head.level; i >= 0; i-- {
		cur = head.nexts[i]
		for cur != nil && cur.key < key {
			cur = cur.nexts[i]
		}
		buf[i] = cur
	}
	return buf[0]
}

func (sl *SkipList) Find(key int) (IObject, bool) {
	node := sl.find(key)
	if node != nil && node.key == key {
		return node.val, true
	}
	return nil, false
}

func (sl *SkipList) Remove(key int) (IObject, bool) {
	return nil, false
}

func (sl *SkipList) Len() int {
	return sl.size
}

func (sl *SkipList) Clear() {

}

func (sl *SkipList) GetArray() []IObject {
	return nil
}

func (sl *SkipList) Copy() []IObject {
	return nil
}

func (sl *SkipList) randomLevel() int {
	level := 1
	for rand.Int31n(2) == 1 && level < sl.maxLevel {
		level++
	}
	return level
}

func (sl *SkipList) createNode(key int, val IObject) *skipListNode {
	level := sl.randomLevel()
	node := &skipListNode{
		key:   key,
		val:   val,
		level: level,
		nexts: make([]*skipListNode, level),
	}
	for i := 0; i < level; i++ {
		node.nexts[i] = &skipListNode{
			key:   node.key,
			val:   node.val,
			level: level,
		}
	}
	return node
}

func (sl *SkipList) headExtend(level int) {
	//如果新节点高度大于首节点高度，首节点高度新增到同样高度
	if level > sl.head.level {
		for i := sl.head.level; i < level; i++ {
			sl.head.nexts = append(sl.head.nexts, &skipListNode{
				key:   -1,
				val:   nil,
				level: level,
			})
		}
		sl.head.level = level
	}
}