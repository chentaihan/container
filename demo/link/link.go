package main

import (
	"fmt"
	"github.com/chentaihan/container/link"
)

func main() {
	list := link.NewLinkedList()
	const count = 20
	for i := 0; i < count; i++ {
		list.PushBack(i)
	}
	if list.Len() != count {
		fmt.Println("len error")
	}
	first, _ := list.Front()
	if first != 0 {
		fmt.Println("Front error")
	}
	last, _ := list.Back()
	if last != count-1 {
		fmt.Println("Back error")
	}

	list.RemoveFront()
	first, _ = list.Front()
	if first != 1 {
		fmt.Println("Front error")
	}
	list.RemoveBack()
	last, _ = list.Back()
	if last != count-2 {
		fmt.Println("Back error ", last)
	}
	if list.RemoveValue(10) != 1 {
		fmt.Println("Remove error")
	}
	array := list.ToList()
	fmt.Println(array)
	for i := 0; i < count; i++ {
		list.PushBack(10)
	}
	array = list.ToList()
	fmt.Println(array)
	if list.RemoveValue(10) != count {
		fmt.Println("Remove error")
	}
	array = list.ToList()
	fmt.Println(array)

	if list.Exist(10) {
		fmt.Println("Exist error")
	}
	list.PushBack(10)
	if !list.Exist(10) {
		fmt.Println("Exist error")
	}
}
