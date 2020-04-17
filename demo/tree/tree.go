package main

import (
	"fmt"
	"github.com/chentaihan/container/common"
	"github.com/chentaihan/container/tree"
)

func main() {
	tests := []struct {
		nums       []int
		depth      int
		count      int
		minVal     int
		maxVal     int
		findVal    int
		findResult bool
		list       []int
	}{
		{
			[]int{5, 4, 8, 2, 3, 9, 1, 6, 7, 15, 12, 20, 14, 11, 25, 18},
			6,
			16,
			1,
			25,
			18,
			true,
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 14, 15, 18, 20, 25},
		},
		{
			[]int{5, 4, 8, 2, 3, 9, 1, 6, 7, 15, 12, 20, 14, 11, 25, 18, 30},
			7,
			17,
			1,
			30,
			13,
			false,
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 14, 15, 18, 20, 25, 30},
		},
	}
	for _, test := range tests {
		tree := tree.NewBinaryTreeInt()
		tree.AddRange(test.nums)
		if tree.GetDepth() != test.depth {
			fmt.Println("GetDepth error")
		}
		count := tree.GetCount()
		if count != test.count {
			fmt.Println("GetCount error")
		}
		if test.minVal != tree.MinNode(tree.GetRoot()).Val {
			fmt.Println("MinNode error")
		}
		if test.maxVal != tree.MaxNode(tree.GetRoot()).Val {
			fmt.Println("MaxNode error")
		}
		if test.findResult != (tree.Find(test.findVal) != nil) {
			fmt.Println("Find error")
		}
		if !common.IntEqual(test.list, tree.ToList()) {
			fmt.Println("ToList error")
		}
		tree.Remove(1)
	}
}
