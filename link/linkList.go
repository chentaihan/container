package link

type Node struct {
	value interface{}
	next  *Node
}

type LinkList struct {
	head *Node
	tail *Node
	size int
}

func NewLinkedList() ILinkList {
	return &LinkList{}
}

func (lk *LinkList) PushFront(val interface{}) {
	lk.head = &Node{
		value: val,
		next:  lk.head,
	}
	if lk.tail == nil {
		lk.tail = lk.head
	}
	lk.size++
}

func (lk *LinkList) PushBack(val interface{}) {
	node := &Node{
		value: val,
	}
	if lk.head == nil {
		lk.head, lk.tail = node, node
	} else {
		lk.tail.next = node
		lk.tail = node
	}
	lk.size++
}

func (lk *LinkList) RemoveFront() (interface{}, bool) {
	if lk.head == nil {
		return nil, false
	}
	val := lk.head.value
	lk.head = lk.head.next
	if lk.size == 1 {
		lk.tail = nil
	}
	lk.size--
	return val, true
}

func (lk *LinkList) RemoveBack() (interface{}, bool) {
	if lk.head == nil {
		return nil, false
	}
	val := lk.tail.value
	lk.size--
	if lk.size == 1 {
		lk.head, lk.tail = nil, nil
		return val, true
	}
	node := lk.head
	for node.next.next != nil {
		node = node.next
	}
	node.next = nil
	lk.tail = node
	return val, true
}

func (lk *LinkList) RemoveValue(val interface{}) int {
	count := 0
	for cur := &lk.head; *cur != nil; {
		if (*cur).value == val {
			count++
			lk.size--
			*cur = (*cur).next
		} else {
			lk.tail = *cur
			cur = &(*cur).next
		}
	}
	return count
}

func (lk *LinkList) Front() (interface{}, bool) {
	if lk.head == nil {
		return nil, false
	}
	return lk.head.value, true
}

func (lk *LinkList) Back() (interface{}, bool) {
	if lk.head == nil {
		return nil, false
	}
	return lk.tail.value, true
}

func (lk *LinkList) Len() int {
	return lk.size
}

func (lk *LinkList) Exist(val interface{}) bool {
	if lk.head == nil {
		return false
	}
	for node := lk.head; node != nil; node = node.next {
		if node.value == val {
			return true
		}
	}
	return false
}

func (lk *LinkList) ToList() []interface{} {
	list := make([]interface{}, lk.size)
	index := 0
	for node := lk.head; node != nil; node = node.next {
		list[index] = node.value
		index++
	}
	return list
}

func (lk *LinkList) Clear() {
	lk.head = nil
	lk.tail = nil
	lk.size = 0
}
