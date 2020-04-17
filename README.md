# Container

## 常用集合实现

* <a href="#队列">队列</a>
* <a href="#栈">栈</a>
* <a href="#排序数组">排序数组</a>
* <a href="#栈">单链表</a>
* <a href="#栈">Map</a>
* <a href="#栈">Set</a>
* <a href="#栈">SetSort</a>
* <a href="#栈">搜索二叉树</a>
* <a href="#栈">LRU</a>
* <a href="#栈">内存池</a>
* <a href="#栈">常用函数</a>

## <a id="队列">1：队列</a>

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


## <a id="排序数组">3：排序数组</a>

```go
package main

import (
	"fmt"
	"github.com/chentaihan/container/array"
	"github.com/chentaihan/container/common"
)

func main() {
	tests := []struct {
		array  []int
		result []int
	}{
		{
			[]int{},
			[]int{},
		},
		{
			[]int{1},
			[]int{1},
		},
		{
			[]int{1, 3},
			[]int{1, 3},
		},
		{
			[]int{1, 3, 2},
			[]int{1, 2, 3},
		},
		{
			[]int{1, 3, 2, 4, 6, 5, 9, 8, 7},
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}
	for index, test := range tests {
		as := array.NewArraySort(0)
		for i := 0; i < len(test.array); i++ {
			as.Add(test.array[i])
		}
		list := as.GetArray()
		if !common.IntEqual(list, test.result) {
			fmt.Println(list, test.result)
			fmt.Println("add error ", index)
		}
		as.RemoveValue(1)
	}
}


```














