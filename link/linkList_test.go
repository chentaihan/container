package link

import "testing"

func TestNewLinkedList(t *testing.T) {
	list := NewLinkedList()
	const count = 20
	for i := 0; i < count; i++ {
		list.PushBack(i)
	}
	if list.Len() != count {
		t.Fatal("len error")
	}
	first, _ := list.Front()
	if first != 0 {
		t.Fatal("Front error")
	}
	last, _ := list.Back()
	if last != count-1 {
		t.Fatal("Back error")
	}

	list.RemoveFront()
	first, _ = list.Front()
	if first != 1 {
		t.Fatal("Front error")
	}
	list.RemoveBack()
	last, _ = list.Back()
	if last != count-2 {
		t.Fatalf("Back error %v", last)
	}
	if list.RemoveValue(10) != 1 {
		t.Fatal("Remove error")
	}
	array := list.ToList()
	t.Log(array)
	for i := 0; i < count; i++ {
		list.PushBack(10)
	}
	array = list.ToList()
	t.Log(array)
	if list.RemoveValue(10) != count {
		t.Fatal("Remove error")
	}
	array = list.ToList()
	t.Log(array)

	if list.Exist(10) {
		t.Fatal("Exist error")
	}
	list.PushBack(10)
	if !list.Exist(10) {
		t.Fatal("Exist error")
	}
}
