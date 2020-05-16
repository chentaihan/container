package heap

/**
大堆
 */

import "container/heap"

type bigHeapList []IObject

func (h *bigHeapList) Less(i, j int) bool {
	return (*h)[i].GetHashCode() > (*h)[j].GetHashCode()
}

func (h *bigHeapList) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *bigHeapList) Len() int {
	return len(*h)
}

func (h *bigHeapList) Pop() (v interface{}) {
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return
}

func (h *bigHeapList) Push(v interface{}) {
	*h = append(*h, v.(IObject))
}

func (h *bigHeapList) Remove(index int) bool {
	if index < 0 || index >= h.Len() {
		return false
	}
	*h = append((*h)[:index], (*h)[index+1:]...)
	return true
}

func NewBigHeap(cap int) IHeap {
	if cap < 0 {
		cap = 0
	}
	objectList := bigHeapList(make([]IObject, 0, cap))
	return &BigHeap{
		list: &objectList,
	}
}

type BigHeap struct {
	list *bigHeapList
}

func (bh *BigHeap) Push(h IObject) {
	heap.Push(bh.list, h)
}

func (bh *BigHeap) Pop() IObject {
	return heap.Pop(bh.list).(IObject)
}

func (bh *BigHeap) Peek() IObject {
	if bh.Len() == 0 {
		return nil
	}
	return (*bh.list)[0]
}

func (bh *BigHeap) Len() int {
	return bh.list.Len()
}

func (bh *BigHeap) Cap() int {
	return cap(*bh.list)
}

func (bh *BigHeap) Empty() bool {
	return bh.list.Len() == 0
}

func (bh *BigHeap) GetArray() []IObject {
	return *bh.list
}

func (bh *BigHeap) Copy() []IObject {
	result := make([]IObject, bh.list.Len())
	for i := 0; i < len(result); i++ {
		result[i] = (*bh.list)[i]
	}
	return result
}

func (bh *BigHeap) Remove(h IObject) bool {
	if bh.Len() == 0 || bh.Peek().GetHashCode() < h.GetHashCode() {
		return false
	}
	for i := 0; i < bh.Len(); i++ {
		if (*bh.list)[i].GetHashCode() == h.GetHashCode() {
			return bh.list.Remove(i)
		}
	}
	return false
}

func (bh *BigHeap) Exist(h IObject) bool {
	if bh.Len() == 0 || bh.Peek().GetHashCode() < h.GetHashCode() {
		return false
	}
	for i := 0; i < bh.Len(); i++ {
		if (*bh.list)[i].GetHashCode() == h.GetHashCode() {
			return true
		}
	}
	return false
}

func (bh *BigHeap) Clear() {
	bh.list = new(bigHeapList)
}
