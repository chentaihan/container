package common

import (
	"sort"
)

func StringEqual(nums1, nums2 []string) bool {
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

func StringEqualSort(nums1, nums2 []string) bool {
	if (nums1 == nil && nums2 != nil) || (nums1 != nil && nums2 == nil) {
		return false
	}
	if len(nums1) != len(nums2) {
		return false
	}
	sort.Strings(nums1)
	sort.Strings(nums2)
	for i := 0; i < len(nums1); i++ {
		if nums1[i] != nums2[i] {
			return false
		}
	}
	return true
}

func StringEqualTwoDim(nums1, nums2 [][]string) bool {
	if (nums1 == nil && nums2 != nil) || (nums1 != nil && nums2 == nil) {
		return false
	}
	if len(nums1) != len(nums2) {
		return false
	}
	for i := 0; i < len(nums1); i++ {
		if !StringEqual(nums1[i], nums2[i]) {
			return false
		}
	}
	return true
}

func StringCopy(array []string) []string {
	ret := make([]string, len(array))
	for i := 0; i < len(array); i++ {
		ret[i] = array[i]
	}
	return ret
}

func StringIn(array []string, val string) bool {
	for i := 0; i < len(array); i++ {
		if array[i] == val {
			return true
		}
	}
	return false
}

func StringRemoveValue(array []string, val string) []string {
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

func StringRemoveIndex(array []string, index int) []string {
	if len(array) == 0 || index < 0 || index >= len(array) {
		return array
	}
	return append(array[:index], array[index+1:]...)
}

//二分查找
func StringBinarySearch(nums []string, target string) int {
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
func StringBinarySearchPos(nums []string, target string) (int, bool) {
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
