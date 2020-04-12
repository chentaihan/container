package link

import "testing"

func TestNewLinkedList(t *testing.T) {
	list := NewLinkedList()
	const count = 20
	for i := 0; i < count; i++ {
		list.AddLast(i)
	}
	if list.Len() != count {
		t.Fatal("len error")
	}
	first, _ := list.GetFirst()
	if first != 0 {
		t.Fatal("GetFirst error")
	}
	last, _ := list.GetLast()
	if last != count-1 {
		t.Fatal("GetLast error")
	}

	list.RemoveFirst()
	first, _ = list.GetFirst()
	if first != 1 {
		t.Fatal("GetFirst error")
	}
	list.RemoveLast()
	last, _ = list.GetLast()
	if last != count-2 {
		t.Fatalf("GetLast error %v", last)
	}
	if list.RemoveValue(10) != 1 {
		t.Fatal("RemoveValue error")
	}
	array := list.ToList()
	t.Log(array)
	for i := 0; i < count; i++ {
		list.AddLast(10)
	}
	array = list.ToList()
	t.Log(array)
	if list.RemoveValue(10) != count {
		t.Fatal("RemoveValue error")
	}
	array = list.ToList()
	t.Log(array)

	if list.Exist(10) {
		t.Fatal("Exist error")
	}
	list.AddLast(10)
	if !list.Exist(10) {
		t.Fatal("Exist error")
	}
}
