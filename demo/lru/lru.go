package main

import (
	"fmt"
	"github.com/chentaihan/container/cache"
	"github.com/chentaihan/container/common"
	"strconv"
)

func toIntArray(array []interface{}) []int {
	list := make([]int, len(array))
	for i := 0; i < len(list); i++ {
		list[i] = common.ToInt(array[i])
	}
	return list
}

func main() {
	tests := []struct {
		list []int
		cap  int
	}{
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
		lru := cache.NewLru(test.cap)
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
				fmt.Println(list, array)
				fmt.Println("add error", index, i, test.cap)
			}
			fmt.Println(index, i, "success")
		}
	}
}
