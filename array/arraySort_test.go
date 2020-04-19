package array

import (
	"testing"
)

type integer int

func (i integer) GetValue() int {
	return int(i)
}

func TestNewArraySort(t *testing.T) {
	tests := []struct {
		array  []integer
		result []integer
	}{
		{
			[]integer{},
			[]integer{},
		},
		{
			[]integer{1},
			[]integer{1},
		},
		{
			[]integer{1, 3},
			[]integer{1, 3},
		},
		{
			[]integer{1, 3, 2},
			[]integer{1, 2, 3},
		},
		{
			[]integer{1, 3, 2, 4, 6, 5, 9, 8, 7},
			[]integer{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}
	for index, test := range tests {
		as := NewArraySort(0)
		for i := 0; i < len(test.array); i++ {
			as.Add(test.array[i])
		}
		var result []IObject
		for i := 0; i < len(test.result); i++ {
			result = append(result, test.result[i])
		}
		list := as.GetArray()
		if !IntEqual(list, result) {
			t.Log(list, test.result)
			t.Fatal("add error ", index)
		}
	}
}

func TestRemoveValue(t *testing.T) {
	tests := []struct {
		array   []integer
		result  []integer
		rmValue integer
	}{
		{
			[]integer{},
			[]integer{},
			0,
		},
		{
			[]integer{1},
			[]integer{},
			1,
		},
		{
			[]integer{1, 3},
			[]integer{3},
			1,
		},
		{
			[]integer{1, 3, 2},
			[]integer{1, 3},
			2,
		},
		{
			[]integer{1, 3, 2},
			[]integer{1, 2},
			3,
		},
		{
			[]integer{1, 1, 1},
			[]integer{},
			1,
		},
		{
			[]integer{1, 1, 1, 1},
			[]integer{},
			1,
		},
		{
			[]integer{1, 1, 1, 1, 1},
			[]integer{},
			1,
		},
		{
			[]integer{1, 1, 1, 1, 1, 2, 2, 2},
			[]integer{2, 2, 2},
			1,
		},
		{
			[]integer{1, 1, 1, 1, 1, 2, 2, 2},
			[]integer{1, 1, 1, 1, 1},
			2,
		},
		{
			[]integer{1, 1, 1, 1, 1, 1, 2, 2, 2},
			[]integer{1, 1, 1, 1, 1, 1},
			2,
		},
		{
			[]integer{1, 1, 1, 1, 1, 1, 2, 2, 2, 2},
			[]integer{1, 1, 1, 1, 1, 1},
			2,
		},
		{
			[]integer{1, 3, 2, 4, 6, 5, 9, 8, 7},
			[]integer{1, 2, 3, 4, 5, 6, 7, 8,},
			9,
		},
	}
	for index, test := range tests {
		as := NewArraySort(0)
		for i := 0; i < len(test.array); i++ {
			as.Add(test.array[i])
		}
		var result []IObject
		for i := 0; i < len(test.result); i++ {
			result = append(result, test.result[i])
		}
		as.Remove(test.rmValue)
		list := as.GetArray()
		if !IntEqual(list, result) {
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
		array  []integer
		result []integer
		index  int
	}{
		{
			[]integer{},
			[]integer{},
			0,
		},
		{
			[]integer{1},
			[]integer{},
			0,
		},
		{
			[]integer{1, 3},
			[]integer{3},
			0,
		},
		{
			[]integer{1, 3},
			[]integer{1},
			1,
		},
		{
			[]integer{1, 3, 2},
			[]integer{1, 2},
			2,
		},
		{
			[]integer{1, 3, 2},
			[]integer{1, 3},
			1,
		},
		{
			[]integer{1, 3, 2, 4, 6, 5, 9, 8, 7},
			[]integer{1, 2, 3, 4, 5, 6, 7, 8},
			8,
		},
	}
	for index, test := range tests {
		as := NewArraySort(0)
		for i := 0; i < len(test.array); i++ {
			as.Add(test.array[i])
		}
		var result []IObject
		for i := 0; i < len(test.result); i++ {
			result = append(result, test.result[i])
		}
		as.RemoveIndex(test.index)
		list := as.GetArray()
		if !IntEqual(list, result) {
			t.Log(list, test.result)
			t.Fatal("add error ", index)
		}
		list = as.Copy()
		if !IntEqual(list, result) {
			t.Log(list, test.result)
			t.Fatal("copy error ", index)
		}
	}
}

func IntEqual(nums1, nums2 []IObject) bool {
	if len(nums1) != len(nums2) {
		return false
	}
	for i := 0; i < len(nums1); i++ {
		if nums1[i].GetValue() != nums2[i].GetValue() {
			return false
		}
	}
	return true
}
