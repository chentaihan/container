package main

import (
	"fmt"
	"github.com/chentaihan/container/queue"
)

func main() {
	size := 2000
	var q queue.IQueue = queue.NewQueue(size)  //数组实现队列
	//var q queue.IQueue = queue.NewQueueLink() //链表实现队列
	for i := 0; i < size; i++ {
		q.Enqueue(i)
	}
	if q.Len() != size {
		fmt.Println("queue len error")
	}
	if q.Cap() != size {
		fmt.Println("queue cap error,cap=", q.Cap())
	}
	value, isOK := q.Peek()
	if isOK {
		if value.(int) != 0 {
			fmt.Println("queue peek error")
		}
	}

	for i := 0; i < size; i++ {
		q.Enqueue(i)
	}
	if q.Cap() != size*2 {
		fmt.Println("queue cap error,cap=", q.Cap())
	}
	count := q.Len()
	index := 0
	for !q.Empty() {
		q.Dequeue()
		index++
	}
	if count != index {
		fmt.Println(fmt.Sprintf("len = %v != count = %v", index, count))
	}
	for i := 0; i < size*8; i++ {
		q.Enqueue(i)
	}
	if q.Cap() != size*8 {
		fmt.Println(fmt.Sprintf("queue cap error,cap=%v,want cap=%v", q.Cap(), size*8))
	}
}
