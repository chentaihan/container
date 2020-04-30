package main

import (
	"fmt"
	"github.com/chentaihan/container/binaryTree"
)

type integer int

func (i integer) GetHashCode() int {
	return int(i)
}

func intEqual(nums1, nums2 []binaryTree.IObject) bool {
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


func main() {
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
		root := binaryTree.NewBinaryTree()
		for i := 0; i < len(test.nums); i++ {
			root.Add(test.nums[i])
		}
		if root.GetDepth() != test.depth {
			fmt.Println("GetDepth error")
		}
		count := root.GetCount()
		if count != test.count {
			fmt.Println("GetCount error")
		}
		if test.minVal != root.MinNode(root.GetRoot()).Val.GetHashCode() {
			fmt.Println("MinNode error")
		}
		if test.maxVal != root.MaxNode(root.GetRoot()).Val.GetHashCode() {
			fmt.Println("MaxNode error")
		}
		if test.findResult != (root.Find(integer(test.findVal)) != nil) {
			fmt.Println("Find error")
		}
		var list []binaryTree.IObject
		for i := 0; i < len(test.list); i++ {
			list = append(list, test.list[i])
		}
		if !intEqual(list, root.ToList()) {
			fmt.Println("ToList error")
		}
	}
}
