package common

import (
	"sort"
)

func IntEqual(nums1, nums2 []int) bool {
	if (nums1 == nil && nums2 != nil) || (nums1 != nil && nums2 == nil) {
		return false
	}
	if len(nums1) != len(nums2) {
		return false
	}
	for i := 0; i < len(nums1); i++ {
		if nums1[i] != nums2[i] {
			return false
		}
	}
	return true
}

func IntEqualSort(nums1, nums2 []int) bool {
	if (nums1 == nil && nums2 != nil) || (nums1 != nil && nums2 == nil) {
		return false
	}
	if len(nums1) != len(nums2) {
		return false
	}
	sort.Ints(nums1)
	sort.Ints(nums2)
	for i := 0; i < len(nums1); i++ {
		if nums1[i] != nums2[i] {
			return false
		}
	}
	return true
}

func IntEqualTwoDim(nums1, nums2 [][]int) bool {
	if (nums1 == nil && nums2 != nil) || (nums1 != nil && nums2 == nil) {
		return false
	}
	if len(nums1) != len(nums2) {
		return false
	}
	for i := 0; i < len(nums1); i++ {
		if !IntEqual(nums1[i], nums2[i]) {
			return false
		}
	}
	return true
}

func IntCopy(array []int) []int {
	ret := make([]int, len(array))
	for i := 0; i < len(array); i++ {
		ret[i] = array[i]
	}
	return ret
}

func IntIn(array []int, val int) bool {
	for i := 0; i < len(array); i++ {
		if array[i] == val {
			return true
		}
	}
	return false
}

func IntRemoveValue(array []int, val int) []int {
	if len(array) == 0 {
		return array
	}
	index := 0
	count := 0
	for i := 0; i < len(array); i++ {
		if array[i] != val {
			if count > 0 {
				array[index] = array[i]
			}
			index++
		} else {
			count++
		}
	}
	return array[:len(array)-count]
}

func IntRemoveIndex(array []int, index int) []int {
	if len(array) == 0 || index < 0 || index >= len(array) {
		return array
	}
	return append(array[:index], array[index+1:]...)
}

//二分查找
func IntBinarySearch(nums []int, target int) int {
	start, middle, end := 0, 0, len(nums)-1
	for start <= end {
		middle = start + (end-start)/2
		if nums[middle] == target {
			return middle
		} else if nums[middle] > target {
			end = middle - 1
		} else {
			start = middle + 1
		}
	}
	return -1
}

//二分查找插入的位置
func IntBinarySearchPos(nums []int, target int) (int, bool) {
	if len(nums) == 0 {
		return 0, false
	}
	start, middle, end := 0, 0, len(nums)-1
	for start <= end {
		middle = start + (end-start)/2
		if nums[middle] == target {
			return middle, true
		} else if nums[middle] > target {
			end = middle - 1
		} else {
			start = middle + 1

		}
	}

	if target < nums[middle] {
		return middle, false
	} else {
		return middle + 1, false
	}
}

//在指定位置插入一个元素
func IntInsertValue(nums []int, index, value int) []int {
	nums = append(nums, value)
	if index >= len(nums) {
		return nums
	}
	if index < 0 {
		index = 0
	}
	for i := len(nums) - 1; i > index; i-- {
		nums[i] = nums[i-1]
	}
	nums[index] = value
	return nums
}
