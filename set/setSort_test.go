package set

import (
	"github.com/chentaihan/container/common"
	"testing"
)

func TestNewSetSort(t *testing.T) {
	tests := []struct {
		array  []int
		result []int
	}{
		{
			[]int{},
			[]int{},
		},
		{
			[]int{1},
			[]int{1},
		},
		{
			[]int{1, 3},
			[]int{1, 3},
		},
		{
			[]int{1, 3, 2},
			[]int{1, 2, 3},
		},
		{
			[]int{1, 3, 2, 1, 2, 3, 3, 3, 3, 3},
			[]int{1, 2, 3},
		},
		{
			[]int{1, 3, 2, 4, 6, 5, 9, 8, 7},
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			[]int{1, 3, 2, 4, 6, 5, 9, 8, 7, 1, 3, 2, 4, 6, 5, 9, 8, 7},
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}
	for index, test := range tests {
		ss := NewSetSort()
		for i := 0; i < len(test.array); i++ {
			ss.Add(test.array[i])
		}
		list := ss.GetArray()
		if !common.IntEqual(list, test.result) {
			t.Log(list, test.result)
			t.Fatal("add error ", index)
		}
	}
}
