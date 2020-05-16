package heap

/**
小堆
 */

import "container/heap"

type smallHeapList []IObject

func (h *smallHeapList) Less(i, j int) bool {
	return (*h)[i].GetHashCode() < (*h)[j].GetHashCode()
}

func (h *smallHeapList) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *smallHeapList) Len() int {
	return len(*h)
}

func (h *smallHeapList) Pop() (v interface{}) {
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return
}

func (h *smallHeapList) Push(v interface{}) {
	*h = append(*h, v.(IObject))
}

func (h *smallHeapList) Remove(index int) bool {
	if index < 0 || index >= h.Len() {
		return false
	}
	*h = append((*h)[:index], (*h)[index+1:]...)
	return true
}

func NewSmallHeap(cap int) IHeap {
	if cap < 0 {
		cap = 0
	}
	objectList := smallHeapList(make([]IObject, 0, cap))
	return &SmallHeap{
		list: &objectList,
	}
}

type SmallHeap struct {
	list *smallHeapList
}

func (bh *SmallHeap) Push(h IObject) {
	heap.Push(bh.list, h)
}

func (bh *SmallHeap) Pop() IObject {
	return heap.Pop(bh.list).(IObject)
}

func (bh *SmallHeap) Peek() IObject {
	if bh.Len() == 0 {
		return nil
	}
	return (*bh.list)[0]
}

func (bh *SmallHeap) Len() int {
	return bh.list.Len()
}

func (bh *SmallHeap) Cap() int {
	return cap(*bh.list)
}

func (bh *SmallHeap) Empty() bool {
	return bh.list.Len() == 0
}

func (bh *SmallHeap) GetArray() []IObject {
	return *bh.list
}

func (bh *SmallHeap) Copy() []IObject {
	result := make([]IObject, bh.list.Len())
	for i := 0; i < len(result); i++ {
		result[i] = (*bh.list)[i]
	}
	return result
}

func (bh *SmallHeap) Remove(h IObject) bool {
	if bh.Len() == 0 || h.GetHashCode() < bh.Peek().GetHashCode() {
		return false
	}
	for i := 0; i < bh.Len(); i++ {
		if (*bh.list)[i].GetHashCode() == h.GetHashCode() {
			return bh.list.Remove(i)
		}
	}
	return false
}

func (bh *SmallHeap) Exist(h IObject) bool {
	if bh.Len() == 0 || bh.Peek().GetHashCode() > h.GetHashCode() {
		return false
	}
	for i := 0; i < bh.Len(); i++ {
		if (*bh.list)[i].GetHashCode() == h.GetHashCode() {
			return true
		}
	}
	return false
}

func (bh *SmallHeap) Clear() {
	bh.list = new(smallHeapList)
}
