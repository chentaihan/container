package cachel

import (
	"github.com/chentaihan/container/common"
	"strconv"
	"testing"
)

func TestNewLru(t *testing.T) {
	tests := []struct {
		list []int
		cap  int
	}{

		{
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			0,
		},
		{
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			1,
		},
		{
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			3,
		},
		{
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			4,
		},
		{
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			5,
		},
		{
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			6,
		},
		{
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			7,
		},
		{
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			8,
		},
		{
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			9,
		},
		{
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			10,
		},
		{
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			11,
		},
	}

	for index, test := range tests {
		lru := NewLru(test.cap)
		for i := 0; i < len(test.list); i++ {
			lru.Add(strconv.Itoa(test.list[i]), test.list[i])
			list := toIntArray(lru.Values())
			var array []int
			if i < test.cap {
				array = test.list[:i+1]
			} else {
				ii := i - test.cap + 1
				array = test.list[ii : ii+test.cap]
			}
			if !common.IntEqualSort(list, array) {
				t.Log(list, array)
				t.Fatal("add error", index, i, test.cap)
			}
			t.Log(index, i, "success")
		}
	}
}

func toIntArray(array []interface{}) []int {
	list := make([]int, len(array))
	for i := 0; i < len(list); i++ {
		list[i] = common.ToInt(array[i])
	}
	return list
}

func TestLru_Remove(t *testing.T) {
	tests := []struct {
		list    []int
		rmValue int
		result  []int
	}{
		{
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			0,
			[]int{9, 8, 7, 6, 5, 4, 3, 2, 1},
		},
		{
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			1,
			[]int{9, 8, 7, 6, 5, 4, 3, 2, 0},
		},
		{
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			3,
			[]int{9, 8, 7, 6, 5, 4, 2, 1, 0},
		},
		{
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			4,
			[]int{9, 8, 7, 6, 5, 3, 2, 1, 0},
		},
		{
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			5,
			[]int{9, 8, 7, 6, 4, 3, 2, 1, 0},
		},
		{
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			6,
			[]int{9, 8, 7, 5, 4, 3, 2, 1, 0},
		},
		{
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			7,
			[]int{9, 8, 6, 5, 4, 3, 2, 1, 0},
		},
		{
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			8,
			[]int{9, 7, 6, 5, 4, 3, 2, 1, 0},
		},
		{
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			9,
			[]int{8, 7, 6, 5, 4, 3, 2, 1, 0},
		},
		{
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			10,
			[]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
		},
		{
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			11,
			[]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
		},
	}

	for index, test := range tests {
		lru := NewLru(100)
		for i := 0; i < len(test.list); i++ {
			lru.Add(strconv.Itoa(test.list[i]), test.list[i])
		}
		lru.Remove(strconv.Itoa(test.rmValue))
		list := toIntArray(lru.Values())
		if !common.IntEqualSort(list, test.result) {
			t.Log(list, test.rmValue, test.result)
			t.Fatal("add error", index)
		}
	}
}

func TestLru_Get(t *testing.T) {
	tests := []struct {
		list   []int
		val    int
		result []int
	}{
		{
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			0,
			[]int{0, 9, 8, 7, 6, 5, 4, 3, 2, 1},
		},
		{
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			1,
			[]int{1, 9, 8, 7, 6, 5, 4, 3, 2, 0},
		},
		{
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			2,
			[]int{2, 9, 8, 7, 6, 5, 4, 3, 1, 0},
		},
		{
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			3,
			[]int{3, 9, 8, 7, 6, 5, 4, 2, 1, 0},
		},
		{
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			4,
			[]int{4, 9, 8, 7, 6, 5, 3, 2, 1, 0},
		},
		{
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			5,
			[]int{5, 9, 8, 7, 6, 4, 3, 2, 1, 0},
		},
		{
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			6,
			[]int{6, 9, 8, 7, 5, 4, 3, 2, 1, 0},
		},
		{
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			7,
			[]int{7, 9, 8, 6, 5, 4, 3, 2, 1, 0},
		},
		{
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			8,
			[]int{8, 9, 7, 6, 5, 4, 3, 2, 1, 0},
		},
		{
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			9,
			[]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
		},
		{
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			10,
			[]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
		},
		{
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			11,
			[]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
		},
	}

	for index, test := range tests {
		lru := NewLru(100)
		for i := 0; i < len(test.list); i++ {
			lru.Add(strconv.Itoa(test.list[i]), test.list[i])
		}
		lru.Get(strconv.Itoa(test.val))
		list := toIntArray(lru.Values())
		if !common.IntEqualSort(list, test.result) {
			t.Log(list, test.val, test.result)
			t.Fatal("add error", index)
		}
	}
}
