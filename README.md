# Container

## 常用集合实现

* 队列（数组和单链表）
* 栈（数组和单链表）
* 排序数组
* 单链表
* Map
* Set
* SetSort
* 搜索二叉树
* LRU
* 内存池
* 常用函数

## 1：队列

```go
type IQueue interface {
	Enqueue(val interface{})      //入队
	Dequeue() (interface{}, bool) //出队
	Empty() bool                  //队列是否为空
	Len() int                     //队列长度
	Cap() int                     //队列容量
	Peek() (interface{}, bool)    //取队首元素
	ToList() []interface{}        //转换成数组
}
```

```go
package main

import (
    "fmt"
    "github.com/chentaihan/container/queue"
)

func main() {
    size := 2000
    var q queue.IQueue = queue.NewQueue(size)
    //var q queue.IQueue = queue.NewQueueLink()
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
```



