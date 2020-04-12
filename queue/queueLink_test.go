package queue

import (
	"testing"
)

func TestNewQueueLink(t *testing.T) {
	size := 2000
	var q IQueue = NewQueueLink()
	for i := 0; i < size; i++ {
		q.Enqueue(i)
	}
	if q.Len() != size {
		t.Fatal("queue len error")
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
	if q.Len() != size*8 {
		t.Fatal("len error")
	}
}
