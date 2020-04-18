package queue

import (
	"testing"
)

type integer int

func (i integer) GetPriority() int {
	return int(i)
}

func (i integer) GetHashCode() int {
	return int(i)
}

func TestNewPriorityQueue(t *testing.T) {
	count := 100
	heap := NewPriorityQueue(count)
	for i := 0; i < count; i++ {
		heap.Push(integer(i))
	}
	if !heap.Contains(integer(20)) {
		t.Fatal("Contain error")
	}
	if !heap.Remove(integer(20)) {
		t.Fatal("remove error")
	}
	heap.Push(integer(20))
	for heap.Len() > 0 {
		l := heap.Len()
		val := heap.Pop()
		if val.GetHashCode() != count-l {
			t.Fatal("pop error", val.GetHashCode(), count-l)
		}
	}
	for i := 0; i < count; i++ {
		heap.Push(integer(i))
	}
	list := heap.GetArray()
	t.Log(list)
	heap.Pop()
	list = heap.GetArray()
	t.Log(list)
	heap.Push(integer(20))
	list = heap.GetArray()
	t.Log(list)
	for heap.Len() > 0 {
		heap.Pop()
	}
	if !heap.Empty() {
		t.Fatal("clear error")
	}
	heap.Clear()
	if !heap.Empty() {
		t.Fatal("clear error")
	}
}
