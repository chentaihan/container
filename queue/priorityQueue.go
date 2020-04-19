package queue

import "container/heap"

type IPriorityQueue interface {
	GetPriority() int //按照这个函数排序，越小优先级越高
	GetHashCode() int //唯一标识一个对象
}

type smallHeapList []IPriorityQueue

func (h *smallHeapList) Less(i, j int) bool {
	return (*h)[i].GetPriority() < (*h)[j].GetPriority()
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
	*h = append(*h, v.(IPriorityQueue))
}

func (h *smallHeapList) Remove(index int) bool {
	if index < 0 || index >= h.Len() {
		return false
	}
	*h = append((*h)[:index], (*h)[index+1:]...)
	return true
}

func NewPriorityQueue(cap int) *PriorityQueue {
	if cap < 0 {
		cap = 0
	}
	objectList := smallHeapList(make([]IPriorityQueue, 0, cap))
	return &PriorityQueue{
		list: &objectList,
	}
}

type PriorityQueue struct {
	list *smallHeapList
}

func (bh *PriorityQueue) Push(h IPriorityQueue) {
	heap.Push(bh.list, h)
}

func (bh *PriorityQueue) Pop() IPriorityQueue {
	return heap.Pop(bh.list).(IPriorityQueue)
}

func (bh *PriorityQueue) Peek() IPriorityQueue {
	if bh.Len() == 0 {
		return nil
	}
	return (*bh.list)[0]
}

func (bh *PriorityQueue) Len() int {
	return bh.list.Len()
}

func (bh *PriorityQueue) Cap() int {
	return cap(*bh.list)
}

func (bh *PriorityQueue) Empty() bool {
	return bh.list.Len() == 0
}

func (bh *PriorityQueue) GetArray() []IPriorityQueue {
	return *bh.list
}

func (bh *PriorityQueue) Copy() []IPriorityQueue {
	result := make([]IPriorityQueue, bh.list.Len())
	for i := 0; i < len(result); i++ {
		result[i] = (*bh.list)[i]
	}
	return result
}

func (bh *PriorityQueue) Remove(h IPriorityQueue) bool {
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

func (bh *PriorityQueue) Exist(h IPriorityQueue) bool {
	if bh.Len() == 0 || bh.Peek().GetPriority() > h.GetHashCode() {
		return false
	}
	for i := 0; i < bh.Len(); i++ {
		if (*bh.list)[i].GetHashCode() == h.GetHashCode() {
			return true
		}
	}
	return false
}

func (bh *PriorityQueue) Clear() {
	bh.list = new(smallHeapList)
}
