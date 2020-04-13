package tree

import (
	"fmt"
	"github.com/chentaihan/container/common"
	"testing"
)

func TestNewBinaryTree(t *testing.T) {
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
			[]int{5, 4, 8, 2, 3, 9},
			4,
			6,
			2,
			9,
			5,
			true,
			[]int{2, 3, 4, 5, 8, 9},
		},
		{
			[]int{5, 4, 8, 2, 3, 9, 7},
			4,
			7,
			2,
			9,
			4,
			true,
			[]int{2, 3, 4, 5, 7, 8, 9},
		},
		{
			[]int{5, 4, 8, 2, 3, 9, 1},
			4,
			7,
			1,
			9,
			9,
			true,
			[]int{1, 2, 3, 4, 5, 8, 9},
		},
		{
			[]int{5, 4, 8, 2, 3, 9, 1, 6},
			4,
			8,
			1,
			9,
			1,
			true,
			[]int{1, 2, 3, 4, 5, 6, 8, 9},
		},
		{
			[]int{5, 4, 8, 2, 3, 9, 1, 6, 7},
			4,
			9,
			1,
			9,
			10,
			false,
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			[]int{5, 4, 8, 2, 3, 9, 1, 6, 7, 15},
			4,
			10,
			1,
			15,
			15,
			true,
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 15},
		},
		{
			[]int{5, 4, 8, 2, 3, 9, 1, 6, 7, 15, 12},
			5,
			11,
			1,
			15,
			3,
			true,
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 12, 15},
		},
		{
			[]int{5, 4, 8, 2, 3, 9, 1, 6, 7, 15, 12, 20},
			5,
			12,
			1,
			20,
			10,
			false,
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 12, 15, 20},
		},
		{
			[]int{5, 4, 8, 2, 3, 9, 1, 6, 7, 15, 12, 20, 14},
			6,
			13,
			1,
			20,
			0,
			false,
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 12, 14, 15, 20},
		},
		{
			[]int{5, 4, 8, 2, 3, 9, 1, 6, 7, 15, 12, 20, 14, 11},
			6,
			14,
			1,
			20,
			13,
			false,
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 14, 15, 20},
		},
		{
			[]int{5, 4, 8, 2, 3, 9, 1, 6, 7, 15, 12, 20, 14, 11, 25},
			6,
			15,
			1,
			25,
			12,
			true,
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 14, 15, 20, 25},
		},
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
		tree := NewBinaryTree()
		tree.AddRange(test.nums)
		fmt.Println(tree.GetDepth() == test.depth)
		count := tree.GetCount()
		fmt.Println(count == test.count)
		fmt.Println(test.minVal == tree.MinNode().Val)
		fmt.Println(test.maxVal == tree.MaxNode().Val)
		fmt.Println(test.findResult == (tree.Find(test.findVal) != nil))
		fmt.Println(common.IntEqual(test.list, tree.ToList()))
	}
}
