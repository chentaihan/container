package heap

import "testing"

func TestNewSmallHeap(t *testing.T) {
	count := 100
	heap := NewSmallHeap(count)
	for i := 0; i < count; i++ {
		heap.Push(integer(i))
	}
	if !heap.Exist(integer(20)) {
		t.Fatal("Contain error")
	}
	if !heap.Remove(integer(20)) {
		t.Fatal("remove error")
	}
	heap.Push(integer(20))
	for heap.Len() > 0 {
		l := heap.Len()
		val := heap.Pop()
		if val.GetValue() != count-l {
			t.Fatal("pop error", val.GetValue(), count-l)
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
