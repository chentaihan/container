package queue

import (
	"testing"
)

func TestNewQueue(t *testing.T) {
	size := 2000
	var q IQueue = NewQueue(size)
	for i := 0; i < size; i++ {
		q.Enqueue(i)
	}
	if q.Len() != size {
		t.Fatal("queue len error")
	}
	if q.Cap() != size {
		t.Fatalf("queue cap error,cap=%v", q.Cap())
	}
	value, isOK := q.Peek()
	if isOK {
		if value.(int) != 0 {
			t.Fatal("queue peek error")
		}
	}

	for i := 0; i < size; i++ {
		q.Enqueue(i)
	}
	if q.Cap() != size*2 {
		t.Fatalf("queue cap error,cap=%v", q.Cap())
	}
	count := q.Len()
	index := 0
	for !q.Empty() {
		q.Dequeue()
		index++
	}
	if count != index {
		t.Fatalf("len = %v != count = %v", index, count)
	}
	for i := 0; i < size*8; i++ {
		q.Enqueue(i)
	}
	if q.Cap() != size*8 {
		t.Fatalf("queue cap error,cap=%v,want cap=%v", q.Cap(), size*8)
	}
}
