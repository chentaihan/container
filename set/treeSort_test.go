package set

import (
	"github.com/chentaihan/container/common"
	"testing"
)

func TestNewTreeSet(t *testing.T) {
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
		ss := NewTreeSet()
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

func TestTreeSet_Remove(t *testing.T) {
	tests := []struct {
		array   []int
		rmValue int
		result  []int
	}{
		{
			[]int{},
			0,
			[]int{},
		},
		{
			[]int{1},
			1,
			[]int{},
		},
		{
			[]int{1, 3},
			1,
			[]int{3},
		},
		{
			[]int{1, 3, 2},
			3,
			[]int{1, 2},
		},
		{
			[]int{1, 3, 2},
			2,
			[]int{1, 3},
		},
		{
			[]int{1, 3, 2, 1, 2, 3, 3, 3, 3, 3},
			2,
			[]int{1, 3},
		},
		{
			[]int{1, 3, 2, 4, 6, 5, 9, 8, 7},
			4,
			[]int{1, 2, 3, 5, 6, 7, 8, 9},
		},
		{
			[]int{1, 3, 2, 4, 6, 5, 9, 8, 7, 1, 3, 2, 4, 6, 5, 9, 8, 7},
			9,
			[]int{1, 2, 3, 4, 5, 6, 7, 8},
		},
	}
	for index, test := range tests {
		ss := NewTreeSet()
		for i := 0; i < len(test.array); i++ {
			ss.Add(test.array[i])
		}
		ss.Remove(test.rmValue)
		list := ss.GetArray()
		if !common.IntEqual(list, test.result) {
			t.Log(list, test.rmValue, test.result)
			t.Fatal("add error ", index)
		}
	}
}
