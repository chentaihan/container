package common

import "testing"

func TestIntRemoveValue(t *testing.T) {
	tests := []struct {
		array  []int
		result []int
		value  int
	}{
		{
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0},
			[]int{2, 3, 4, 5, 6, 7, 8, 9, 0},
			1,
		},
		{
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0},
			[]int{1, 2, 4, 5, 6, 7, 8, 9, 0},
			3,
		},
		{
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			0,
		},
		{
			[]int{1, 2, 3, 4, 5, 1, 6, 7, 8, 9, 0, 1},
			[]int{2, 3, 4, 5, 6, 7, 8, 9, 0},
			1,
		},
		{
			[]int{2, 1, 3, 4, 1, 5, 1, 6, 7, 8, 9, 1, 0},
			[]int{2, 3, 4, 5, 6, 7, 8, 9, 0},
			1,
		},
		{
			[]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			[]int{},
			1,
		},
		{
			[]int{1},
			[]int{},
			1,
		},
	}

	for _, test := range tests {
		result := IntRemoveValue(test.array, test.value)
		isEqual := IntEqual(result, test.result)
		if !isEqual {
			t.Fatal("IntRemoveValue error")
		}
	}
}

func TestIntRemoveIndex(t *testing.T) {
	tests := []struct {
		array  []int
		result []int
		index  int
	}{
		{
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			-1,
		},
		{
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			0,
		},
		{
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			[]int{0, 2, 3, 4, 5, 6, 7, 8, 9},
			1,
		},
		{
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8},
			9,
		},
		{
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			10,
		},
	}

	for _, test := range tests {
		result := IntRemoveIndex(test.array, test.index)
		isEqual := IntEqual(result, test.result)
		if !isEqual {
			t.Fatal("IntRemoveValue error")
		}
	}
}

func TestIntBinarySearch(t *testing.T) {
	tests := []struct {
		array []int
		value int
		index int
	}{
		{
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			0,
			0,
		},
		{
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			1,
			1,
		},
		{
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			2,
			2,
		},
		{
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			3,
			3,
		},
		{
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			4,
			4,
		},
		{
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			5,
			5,
		},
		{
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			9,
			9,
		},
		{
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			10,
			-1,
		},
		{
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			110,
			-1,
		},
		{
			[]int{0, 2, 4, 6, 8, 10, 12, 14, 16, 18},
			1,
			-1,
		},
		{
			[]int{0, 2, 4, 6, 8, 10, 12, 14, 16, 18},
			3,
			-1,
		},
		{
			[]int{0, 2, 4, 6, 8, 10, 12, 14, 16, 18},
			100,
			-1,
		},
	}
	for index, test := range tests {
		pos := IntBinarySearch(test.array, test.value)
		if pos != test.index {
			t.Log(index, pos, test.index, test.value)
		}
	}
}

func TestIntBinarySearchPos(t *testing.T) {
	tests := []struct {
		array []int
		value int
		index int
	}{
		{
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			0,
			0,
		},
		{
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			1,
			1,
		},
		{
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			2,
			2,
		},
		{
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			3,
			3,
		},
		{
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			4,
			4,
		},
		{
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			5,
			5,
		},
		{
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			9,
			9,
		},
		{
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			10,
			10,
		},
		{
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			110,
			10,
		},
		{
			[]int{0, 2, 4, 6, 8, 10, 12, 14, 16, 18},
			1,
			1,
		},
		{
			[]int{0, 2, 4, 6, 8, 10, 12, 14, 16, 18},
			3,
			2,
		},
		{
			[]int{0, 2, 4, 6, 8, 10, 12, 14, 16, 18},
			5,
			3,
		},
		{
			[]int{0, 2, 4, 6, 8, 10, 12, 14, 16, 18},
			7,
			4,
		},
		{
			[]int{0, 2, 4, 6, 8, 10, 12, 14, 16, 18},
			17,
			9,
		},
		{
			[]int{0, 2, 4, 6, 8, 10, 12, 14, 16, 18},
			100,
			10,
		},
	}
	for index, test := range tests {
		pos, _ := IntBinarySearchPos(test.array, test.value)
		if pos != test.index {
			t.Log(index, pos, test.index, test.value)
		}
	}
}
