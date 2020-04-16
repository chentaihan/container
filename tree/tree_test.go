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
		fmt.Println(test.minVal == tree.MinNode(tree.root).Val)
		fmt.Println(test.maxVal == tree.MaxNode(tree.root).Val)
		fmt.Println(test.findResult == (tree.Find(test.findVal) != nil))
		fmt.Println(common.IntEqual(test.list, tree.ToList()))
	}
}

func TestBinaryTree_Remove(t *testing.T) {
	tests := []struct {
		nums    []int
		rmVAlue int
		list    []int
	}{
		{
			[]int{1},
			1,
			[]int{},
		},
		{
			[]int{1, 2},
			1,
			[]int{2},
		},
		{
			[]int{1, 2},
			2,
			[]int{1},
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
			[]int{1, 4, 3, 2},
			4,
			[]int{1, 3, 2},
		},
		{
			[]int{1, 4, 3, 2},
			3,
			[]int{1, 4, 2},
		},
		{
			[]int{1, 4, 3, 2},
			2,
			[]int{1, 4, 3},
		},
		{
			[]int{1, 4, 3, 2},
			1,
			[]int{2, 4, 3},
		},
		{
			[]int{1, 2, 3, 4},
			1,
			[]int{2, 3, 4},
		},
		{
			[]int{1, 2, 3, 4},
			2,
			[]int{1, 3, 4},
		},
		{
			[]int{1, 2, 3, 4},
			4,
			[]int{1, 2, 3},
		},
		{
			[]int{4, 3, 2, 1},
			4,
			[]int{3, 2, 1},
		},
		{
			[]int{4, 3, 2, 1},
			3,
			[]int{4, 2, 1},
		},
		{
			[]int{4, 3, 2, 1},
			2,
			[]int{4, 3, 1},
		},
		{
			[]int{4, 3, 2, 1},
			1,
			[]int{4, 3, 2},
		},
		{
			[]int{6, 3, 10, 2, 4, 8, 12},
			6,
			[]int{8, 3, 10, 2, 4, 12},
		},
		{
			[]int{6, 3, 10, 2, 4, 8, 12},
			3,
			[]int{6, 4, 10, 2, 8, 12},
		},
		{
			[]int{6, 3, 10, 2, 4, 8, 12},
			10,
			[]int{6, 3, 12, 2, 4, 8},
		},
		{
			[]int{6, 3, 10, 2, 4, 8, 12},
			2,
			[]int{6, 3, 10, 4, 8, 12},
		},
		{
			[]int{6, 3, 10, 2, 4, 8, 12},
			4,
			[]int{6, 3, 10, 2, 8, 12},
		},
		{
			[]int{6, 3, 10, 2, 4, 8, 12},
			8,
			[]int{6, 3, 10, 2, 4, 12},
		},
		{
			[]int{6, 3, 10, 2, 4, 8, 12},
			12,
			[]int{6, 3, 10, 2, 4, 8},
		},
		{
			[]int{6, 3, 10, 2, 4, 8, 12, 1},
			1,
			[]int{6, 3, 10, 2, 4, 8, 12},
		},
		{
			[]int{6, 3, 10, 2, 4, 8, 12, 1, 15},
			15,
			[]int{6, 3, 10, 2, 4, 8, 12, 1},
		},
		{
			[]int{6, 3, 10, 2, 4, 8, 12, 1, 11},
			11,
			[]int{6, 3, 10, 2, 4, 8, 12, 1},
		},
		{
			[]int{6, 3, 10, 2, 4, 8, 12, 1, 11, 5},
			5,
			[]int{6, 3, 10, 2, 4, 8, 12, 1, 11},
		},
	}
	for index, test := range tests {
		tree := NewBinaryTree()
		tree.AddRange(test.nums)
		tree.Remove(test.rmVAlue)
		if !common.IntEqual(test.list, tree.FloorList(tree.root)) {
			t.Log(test.list, test.rmVAlue, tree.FloorList(tree.root))
			t.Fatal("remove error", index)
		}
	}
}
