package tree

import (
	"testing"
)

type integer int

func (i integer) GetHashCode() int {
	return int(i)
}

func intEqual(nums1, nums2 []ITree) bool {
	if (nums1 == nil && nums2 != nil) || (nums1 != nil && nums2 == nil) {
		return false
	}
	if len(nums1) != len(nums2) {
		return false
	}
	for i := 0; i < len(nums1); i++ {
		if nums1[i].GetHashCode() != nums2[i].GetHashCode() {
			return false
		}
	}
	return true
}

func TestNewBinaryTreeCommon(t *testing.T) {
	tests := []struct {
		nums       []integer
		depth      int
		count      int
		minVal     int
		maxVal     int
		findVal    int
		findResult bool
		list       []integer
	}{
		{
			[]integer{5, 4, 8, 2, 3, 9},
			4,
			6,
			2,
			9,
			5,
			true,
			[]integer{2, 3, 4, 5, 8, 9},
		},
		{
			[]integer{5, 4, 8, 2, 3, 9, 7},
			4,
			7,
			2,
			9,
			4,
			true,
			[]integer{2, 3, 4, 5, 7, 8, 9},
		},
		{
			[]integer{5, 4, 8, 2, 3, 9, 1},
			4,
			7,
			1,
			9,
			9,
			true,
			[]integer{1, 2, 3, 4, 5, 8, 9},
		},
		{
			[]integer{5, 4, 8, 2, 3, 9, 1, 6},
			4,
			8,
			1,
			9,
			1,
			true,
			[]integer{1, 2, 3, 4, 5, 6, 8, 9},
		},
		{
			[]integer{5, 4, 8, 2, 3, 9, 1, 6, 7},
			4,
			9,
			1,
			9,
			10,
			false,
			[]integer{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			[]integer{5, 4, 8, 2, 3, 9, 1, 6, 7, 15},
			4,
			10,
			1,
			15,
			15,
			true,
			[]integer{1, 2, 3, 4, 5, 6, 7, 8, 9, 15},
		},
		{
			[]integer{5, 4, 8, 2, 3, 9, 1, 6, 7, 15, 12},
			5,
			11,
			1,
			15,
			3,
			true,
			[]integer{1, 2, 3, 4, 5, 6, 7, 8, 9, 12, 15},
		},
		{
			[]integer{5, 4, 8, 2, 3, 9, 1, 6, 7, 15, 12, 20},
			5,
			12,
			1,
			20,
			10,
			false,
			[]integer{1, 2, 3, 4, 5, 6, 7, 8, 9, 12, 15, 20},
		},
		{
			[]integer{5, 4, 8, 2, 3, 9, 1, 6, 7, 15, 12, 20, 14},
			6,
			13,
			1,
			20,
			0,
			false,
			[]integer{1, 2, 3, 4, 5, 6, 7, 8, 9, 12, 14, 15, 20},
		},
		{
			[]integer{5, 4, 8, 2, 3, 9, 1, 6, 7, 15, 12, 20, 14, 11},
			6,
			14,
			1,
			20,
			13,
			false,
			[]integer{1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 14, 15, 20},
		},
		{
			[]integer{5, 4, 8, 2, 3, 9, 1, 6, 7, 15, 12, 20, 14, 11, 25},
			6,
			15,
			1,
			25,
			12,
			true,
			[]integer{1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 14, 15, 20, 25},
		},
		{
			[]integer{5, 4, 8, 2, 3, 9, 1, 6, 7, 15, 12, 20, 14, 11, 25, 18},
			6,
			16,
			1,
			25,
			18,
			true,
			[]integer{1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 14, 15, 18, 20, 25},
		},
		{
			[]integer{5, 4, 8, 2, 3, 9, 1, 6, 7, 15, 12, 20, 14, 11, 25, 18, 30},
			7,
			17,
			1,
			30,
			13,
			false,
			[]integer{1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 14, 15, 18, 20, 25, 30},
		},
	}
	for _, test := range tests {
		tree := NewBinaryTree()
		for i := 0; i < len(test.nums); i++ {
			tree.Add(test.nums[i])
		}
		if tree.GetDepth() != test.depth {
			t.Fatal("GetDepth error")
		}
		count := tree.GetCount()
		if count != test.count {
			t.Fatal("GetCount error")
		}
		if test.minVal != tree.MinNode(tree.root).Val.GetHashCode() {
			t.Fatal("MinNode error")
		}
		if test.maxVal != tree.MaxNode(tree.root).Val.GetHashCode() {
			t.Fatal("MaxNode error")
		}
		if test.findResult != (tree.Find(integer(test.findVal)) != nil) {
			t.Fatal("Find error")
		}
		var list []ITree
		for i := 0; i < len(test.list); i++ {
			list = append(list, test.list[i])
		}
		if !intEqual(list, tree.ToList()) {
			t.Fatal("ToList error")
		}
	}
}

func TestBinaryTreeCommon_Remove(t *testing.T) {
	tests := []struct {
		nums    []integer
		rmVAlue int
		list    []integer
	}{
		{
			[]integer{1},
			1,
			[]integer{},
		},
		{
			[]integer{1, 2},
			1,
			[]integer{2},
		},
		{
			[]integer{1, 2},
			2,
			[]integer{1},
		},
		{
			[]integer{1, 3, 2},
			3,
			[]integer{1, 2},
		},
		{
			[]integer{1, 3, 2},
			2,
			[]integer{1, 3},
		},
		{
			[]integer{1, 4, 3, 2},
			4,
			[]integer{1, 3, 2},
		},
		{
			[]integer{1, 4, 3, 2},
			3,
			[]integer{1, 4, 2},
		},
		{
			[]integer{1, 4, 3, 2},
			2,
			[]integer{1, 4, 3},
		},
		{
			[]integer{1, 4, 3, 2},
			1,
			[]integer{2, 4, 3},
		},
		{
			[]integer{1, 2, 3, 4},
			1,
			[]integer{2, 3, 4},
		},
		{
			[]integer{1, 2, 3, 4},
			2,
			[]integer{1, 3, 4},
		},
		{
			[]integer{1, 2, 3, 4},
			4,
			[]integer{1, 2, 3},
		},
		{
			[]integer{4, 3, 2, 1},
			4,
			[]integer{3, 2, 1},
		},
		{
			[]integer{4, 3, 2, 1},
			3,
			[]integer{4, 2, 1},
		},
		{
			[]integer{4, 3, 2, 1},
			2,
			[]integer{4, 3, 1},
		},
		{
			[]integer{4, 3, 2, 1},
			1,
			[]integer{4, 3, 2},
		},
		{
			[]integer{6, 3, 10, 2, 4, 8, 12},
			6,
			[]integer{8, 3, 10, 2, 4, 12},
		},
		{
			[]integer{6, 3, 10, 2, 4, 8, 12},
			3,
			[]integer{6, 4, 10, 2, 8, 12},
		},
		{
			[]integer{6, 3, 10, 2, 4, 8, 12},
			10,
			[]integer{6, 3, 12, 2, 4, 8},
		},
		{
			[]integer{6, 3, 10, 2, 4, 8, 12},
			2,
			[]integer{6, 3, 10, 4, 8, 12},
		},
		{
			[]integer{6, 3, 10, 2, 4, 8, 12},
			4,
			[]integer{6, 3, 10, 2, 8, 12},
		},
		{
			[]integer{6, 3, 10, 2, 4, 8, 12},
			8,
			[]integer{6, 3, 10, 2, 4, 12},
		},
		{
			[]integer{6, 3, 10, 2, 4, 8, 12},
			12,
			[]integer{6, 3, 10, 2, 4, 8},
		},
		{
			[]integer{6, 3, 10, 2, 4, 8, 12, 1},
			1,
			[]integer{6, 3, 10, 2, 4, 8, 12},
		},
		{
			[]integer{6, 3, 10, 2, 4, 8, 12, 1, 15},
			15,
			[]integer{6, 3, 10, 2, 4, 8, 12, 1},
		},
		{
			[]integer{6, 3, 10, 2, 4, 8, 12, 1, 11},
			11,
			[]integer{6, 3, 10, 2, 4, 8, 12, 1},
		},
		{
			[]integer{6, 3, 10, 2, 4, 8, 12, 1, 11, 5},
			5,
			[]integer{6, 3, 10, 2, 4, 8, 12, 1, 11},
		},
	}
	for index, test := range tests {
		tree := NewBinaryTree()
		for i := 0; i < len(test.nums); i++ {
			tree.Add(test.nums[i])
		}
		tree.Remove(integer(test.rmVAlue))
		list := make([]ITree, 0)
		for i := 0; i < len(test.list); i++ {
			list = append(list, test.list[i])
		}
		if !intEqual(list, tree.FloorList(tree.root)) {
			t.Log(test.list, test.rmVAlue, tree.FloorList(tree.root))
			t.Fatal("remove error", index)
		}
	}
}
