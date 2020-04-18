package main

import (
	"fmt"
	heap2 "github.com/chentaihan/container/heap"
)

type integer int

func (i integer) GetHashCode() int {
	return int(i)
}

func main() {
	count := 100
	heap := heap2.NewBigHeap(count) //大堆
	//heap := heap2.NewSmallHeap(count) //小堆
	for i := 0; i < count; i++ {
		heap.Push(integer(i))
	}
	if !heap.Contains(integer(20)) {
		fmt.Println("Contain error")
	}
	if !heap.Remove(integer(20)) {
		fmt.Println("remove error")
	}
	heap.Push(integer(20))
	for heap.Len() > 0 {
		l := heap.Len()
		val := heap.Pop()
		if val.GetHashCode() != l-1 {
			fmt.Println("pop error", val.GetHashCode(), l-1)
		}
	}
	for i := 0; i < count; i++ {
		heap.Push(integer(i))
	}
	list := heap.GetArray()
	fmt.Println(list)
	heap.Pop()
	list = heap.GetArray()
	fmt.Println(list)
	heap.Push(integer(20))
	list = heap.GetArray()
	fmt.Println(list)
	for heap.Len() > 0 {
		heap.Pop()
	}
	if !heap.Empty() {
		fmt.Println("clear error")
	}
	heap.Clear()
	if !heap.Empty() {
		fmt.Println("clear error")
	}

}
