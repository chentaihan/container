package main

import (
	"fmt"
	"github.com/chentaihan/container/common"
	"github.com/chentaihan/container/lru"
	"strconv"
	"time"
)

type integer int

func (i integer) GetHashCode() int {
	return int(i)
}

func toIntArray(array []interface{}) []int {
	list := make([]int, len(array))
	for i := 0; i < len(list); i++ {
		list[i] = common.ToInt(array[i])
	}
	return list
}

func testLru() {
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
		lru := lru.NewLru(test.cap)
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

func testLruTime() {
	tests := []struct {
		array []integer
	}{
		{
			[]integer{},
		},
		{
			[]integer{1},
		},
		{
			[]integer{1, 2},
		},
		{
			[]integer{1, 2, 3},
		},
		{
			[]integer{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			[]integer{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
		},
	}

	for index, test := range tests {
		lru := lru.NewLruTime(10)
		for i := 0; i < len(test.array); i++ {
			lru.Add(test.array[i], int64(i+1))
		}
		result := lru.GetArray()
		fmt.Println(result)
		fmt.Println(lru.Get(3))
		fmt.Println(lru.Get(5))
		fmt.Println(lru.Peek())
		fmt.Println("len:", lru.Len())
		time.Sleep(time.Duration(index) * time.Second)
		result = lru.GetArray()
		fmt.Println(result)
		fmt.Println(lru.Get(3))
		fmt.Println(lru.Get(5))
		fmt.Println(lru.Peek())
		fmt.Println("len:", lru.Len())
		lru.Clear()
		fmt.Println(lru.Peek())
		fmt.Println("len:", lru.Len())
		fmt.Println()
	}
}

func testLruTimeSync() {
	tests := []struct {
		array []integer
	}{
		{
			[]integer{},
		},
		{
			[]integer{1},
		},
		{
			[]integer{1, 2},
		},
		{
			[]integer{1, 2, 3},
		},
		{
			[]integer{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			[]integer{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
		},
	}

	for index, test := range tests {
		fmt.Println("index: ", index)
		lru := lru.NewLruTimeSync(10)
		for i := 0; i < len(test.array); i++ {
			lru.Add(test.array[i], int64(i+1))
		}
		result := lru.GetArray()
		fmt.Println(result)
		fmt.Println(lru.Get(3))
		fmt.Println(lru.Get(5))
		fmt.Println(lru.Peek())
		fmt.Println("len:", lru.Len())
		//time.Sleep(time.Duration(index) * time.Second)
		result = lru.GetArray()
		fmt.Println(result)
		fmt.Println(lru.Get(3))
		fmt.Println(lru.Get(5))
		fmt.Println(lru.Peek())
		fmt.Println("len:", lru.Len())
		lru.Clear()
		fmt.Println(lru.Peek())
		fmt.Println("len:", lru.Len())
		lru.Stop()
		fmt.Println("stop success")
		fmt.Println()
	}
	lru := lru.NewLruTimeSync(10)
	for i := 1; i <= 30000; i++ {
		lru.Add(integer(i), int64(i%6))
	}
	fmt.Println(lru.Len())
	for lru.Len() > 0 {
		time.Sleep(time.Second)
		fmt.Println("len: ", lru.Len())
	}

}

func main() {
	testLruTimeSync()
}
