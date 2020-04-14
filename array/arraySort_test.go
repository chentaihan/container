package array

import (
	"github.com/chentaihan/container/common"
	"testing"
)

func TestNewArraySort(t *testing.T) {
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
			[]int{1, 3, 2, 4, 6, 5, 9, 8, 7},
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}
	for index, test := range tests {
		as := NewArraySort(0)
		for i := 0; i < len(test.array); i++ {
			as.Add(test.array[i])
		}
		list := as.GetArray()
		if !common.IntEqual(list, test.result) {
			t.Log(list, test.result)
			t.Fatal("add error ", index)
		}
	}
}

func TestRemoveValue(t *testing.T) {
	tests := []struct {
		array   []int
		result  []int
		rmValue int
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
			[]int{1, 3},
			[]int{3},
			1,
		},
		{
			[]int{1, 3, 2},
			[]int{1, 3},
			2,
		},
		{
			[]int{1, 3, 2},
			[]int{1, 2},
			3,
		},
		{
			[]int{1, 1, 1},
			[]int{},
			1,
		},
		{
			[]int{1, 1, 1, 1},
			[]int{},
			1,
		},
		{
			[]int{1, 1, 1, 1, 1},
			[]int{},
			1,
		},
		{
			[]int{1, 1, 1, 1, 1, 2, 2, 2},
			[]int{2, 2, 2},
			1,
		},
		{
			[]int{1, 1, 1, 1, 1, 2, 2, 2},
			[]int{1, 1, 1, 1, 1},
			2,
		},
		{
			[]int{1, 1, 1, 1, 1, 1, 2, 2, 2},
			[]int{1, 1, 1, 1, 1, 1},
			2,
		},
		{
			[]int{1, 1, 1, 1, 1, 1, 2, 2, 2, 2},
			[]int{1, 1, 1, 1, 1, 1},
			2,
		},
		{
			[]int{1, 3, 2, 4, 6, 5, 9, 8, 7},
			[]int{1, 2, 3, 4, 5, 6, 7, 8,},
			9,
		},
	}
	for index, test := range tests {
		as := NewArraySort(0)
		for i := 0; i < len(test.array); i++ {
			as.Add(test.array[i])
		}
		as.RemoveValue(test.rmValue)
		list := as.GetArray()
		if !common.IntEqual(list, test.result) {
			t.Log(list, test.result)
			t.Fatal("add error ", index)
		}
		if len(test.result) != as.Len() {
			t.Fatal("len error ", index)
		}
	}
}

func TestRemoveIndex(t *testing.T) {
	tests := []struct {
		array  []int
		result []int
		index  int
	}{
		{
			[]int{},
			[]int{},
			0,
		},
		{
			[]int{1},
			[]int{},
			0,
		},
		{
			[]int{1, 3},
			[]int{3},
			0,
		},
		{
			[]int{1, 3},
			[]int{1},
			1,
		},
		{
			[]int{1, 3, 2},
			[]int{1, 2},
			2,
		},
		{
			[]int{1, 3, 2},
			[]int{1, 3},
			1,
		},
		{
			[]int{1, 3, 2, 4, 6, 5, 9, 8, 7},
			[]int{1, 2, 3, 4, 5, 6, 7, 8},
			8,
		},
	}
	for index, test := range tests {
		as := NewArraySort(0)
		for i := 0; i < len(test.array); i++ {
			as.Add(test.array[i])
		}
		as.RemoveIndex(test.index)
		list := as.GetArray()
		if !common.IntEqual(list, test.result) {
			t.Log(list, test.result)
			t.Fatal("add error ", index)
		}
		list = as.Copy()
		if !common.IntEqual(list, test.result) {
			t.Log(list, test.result)
			t.Fatal("copy error ", index)
		}
	}
}
