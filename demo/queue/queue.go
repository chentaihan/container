package main

import (
	"fmt"
	"github.com/chentaihan/container/queue"
)

func main() {
	queueTest()         //普通队列测试
	priorityQueueTest() //优先级队列测试
}

func queueTest() {
	size := 2000
	var q queue.IQueue = queue.NewQueue(size) //数组实现队列
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

type integer int

func (i integer) GetPriority() int {
	return int(i)
}

func (i integer) GetHashCode() int {
	return int(i)
}

func priorityQueueTest() {
	count := 100
	heap := queue.NewPriorityQueue(count)
	for i := 0; i < count; i++ {
		heap.Push(integer(i))
	}
	if !heap.Exist(integer(20)) {
		fmt.Println("Contain error")
	}
	if !heap.Remove(integer(20)) {
		fmt.Println("remove error")
	}
	heap.Push(integer(20))
	for heap.Len() > 0 {
		l := heap.Len()
		val := heap.Pop()
		if val.GetHashCode() != count-l {
			fmt.Println("pop error", val.GetHashCode(), count-l)
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
