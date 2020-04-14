package set

import (
	"github.com/chentaihan/container/common"
	"testing"
)

func TestSet_Add(t *testing.T) {
	tests := []struct {
		list     []int
		sortList []int
		size     int
	}{
		{
			[]int{},
			[]int{},
			0,
		},
		{
			[]int{1},
			[]int{1},
			1,
		},
		{
			[]int{1, 4},
			[]int{1, 4},
			2,
		},
		{
			[]int{1, 4, 6},
			[]int{1, 4, 6},
			3,
		},
		{
			[]int{1, 4, 6, 6},
			[]int{1, 4, 6},
			3,
		},
		{
			[]int{1, 4, 6, 6, 4},
			[]int{1, 4, 6},
			3,
		},
		{
			[]int{1, 4, 6, 6, 4, 1},
			[]int{1, 4, 6},
			3,
		},
		{
			[]int{1, 1, 1, 1, 1, 1, 1},
			[]int{1},
			1,
		},
		{
			[]int{1, 3, 5, 7, 9, 2, 4, 6, 8, 0},
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			10,
		},
	}

	for index, test := range tests {
		s := NewSet()
		for i := 0; i < len(test.list); i++ {
			s.Add(test.list[i])
		}
		if s.Len() != test.size {
			t.Fatal("size error", index)
		}
		if !common.IntEqualSort(s.GetArray(), test.sortList) {
			t.Fatal("add error", index)
		}
	}
}

func TestSet_Remove(t *testing.T) {
	tests := []struct {
		list     []int
		sortList []int
		rmValue  int
	}{
		{
			[]int{},
			[]int{},
			0,
		},
		{
			[]int{1},
			[]int{},
			1,
		},
		{
			[]int{1, 4},
			[]int{4},
			1,
		},
		{
			[]int{1, 4},
			[]int{1},
			4,
		},
		{
			[]int{1, 4, 6, 6},
			[]int{4, 6},
			1,
		},
		{
			[]int{1, 4, 6, 6, 4},
			[]int{1, 4},
			6,
		},
		{
			[]int{1, 4, 6, 6, 4, 1},
			[]int{1, 6},
			4,
		},
		{
			[]int{1, 1, 1, 1, 1, 1, 1},
			[]int{},
			1,
		},
		{
			[]int{1, 3, 5, 7, 9, 2, 4, 6, 8, 0},
			[]int{0, 1, 2, 3, 4, 5, 6, 8, 9},
			7,
		},
	}

	for index, test := range tests {
		s := NewSet()
		for i := 0; i < len(test.list); i++ {
			s.Add(test.list[i])
		}
		s.Remove(test.rmValue)
		if !common.IntEqualSort(s.GetArray(), test.sortList) {
			t.Fatal("add error", index)
		}
	}
}
