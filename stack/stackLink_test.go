package stack

import (
	"fmt"
	"testing"
)

func TestNewStackLink(t *testing.T) {
	var s IStack = NewStackLink()
	size := 10
	for i := 0; i < size; i++ {
		s.Push(i)
	}
	if s.Len() != size {
		t.Fatal("len error")
	}
	t.Logf("cap=%v", s.Cap())
	for !s.Empty() {
		if val, exist := s.Top(); exist {
			fmt.Println(val)
		}
		s.Pop()
	}
	if s.Len() != 0 {
		t.Fatal("empty error")
	}
}
