package main

import (
	"fmt"
	"github.com/chentaihan/container/array"
)

type integer int

func (i integer) GetValue() int {
	return int(i)
}

func main() {
	tests := []struct {
		array  []integer
		result []integer
	}{
		{
			[]integer{1, 3, 2, 4, 6, 5, 9, 8, 7},
			[]integer{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}
	for index, test := range tests {
		as := array.NewArraySort(0)
		for i := 0; i < len(test.array); i++ {
			as.Add(test.array[i])
		}
		var result []array.IObject
		for i := 0; i < len(test.result); i++ {
			result = append(result, test.result[i])
		}
		list := as.GetArray()
		if !IntEqual(list, result) {
			fmt.Println(list, test.result)
			fmt.Println("add error ", index)
		}
		as.Remove(integer(1))
	}
}

func IntEqual(nums1, nums2 []array.IObject) bool {
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