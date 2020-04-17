# Container

## 常用集合实现

* <a href="## 1：队列">队列</a>
* <a href="## 2：栈">栈</a>
* <a href="## 2：栈">排序数组</a>
* <a href="## 2：栈">单链表</a>
* <a href="## 2：栈">Map</a>
* <a href="## 2：栈">Set</a>
* <a href="## 2：栈">SetSort</a>
* <a href="## 2：栈">搜索二叉树</a>
* <a href="## 2：栈">LRU</a>
* <a href="## 2：栈">内存池</a>
* <a href="## 2：栈">常用函数</a>

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


## 2：栈

```go
type IStack interface {
	Push(x interface{})       //入栈
	Pop() (interface{}, bool) //出栈
	Top() (interface{}, bool) //栈顶元素
	Empty() bool              //栈是否为空
	Len() int                 //栈元素个数
	Cap() int                 //栈容量
}
```

```go
package main

import (
	"fmt"
	"github.com/chentaihan/container/stack"
)

func main() {
	var s stack.IStack = stack.NewStack(100)
	//var s stack.IStack = stack.NewStackLink()
	size := 10
	for i := 0; i < size; i++ {
		s.Push(i)
	}
	if s.Len() != size {
		fmt.Println("len error")
	}
	fmt.Println("cap=", s.Cap())
	for !s.Empty() {
		if val, exist := s.Top(); exist {
			fmt.Println(val)
		}
		s.Pop()
	}
	if s.Len() != 0 {
		fmt.Println("empty error")
	}
}

```








