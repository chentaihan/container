package main

import (
	"fmt"
	"github.com/chentaihan/container/array"
	"github.com/chentaihan/container/common"
)

func main() {
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
		as := array.NewArraySort(0)
		for i := 0; i < len(test.array); i++ {
			as.Add(test.array[i])
		}
		list := as.GetArray()
		if !common.IntEqual(list, test.result) {
			fmt.Println(list, test.result)
			fmt.Println("add error ", index)
		}
		as.RemoveValue(1)
	}
}
