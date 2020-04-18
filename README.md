# Container

## 常用集合实现

* <a href="#队列">队列</a>
* <a href="#队列">优先级队列</a>
* <a href="#栈">栈</a>
* <a href="#排序数组">排序数组</a>
* <a href="#单链表">单链表</a>
* <a href="#Map">Map</a>
* <a href="#Set">Set</a>
* <a href="#Set">SetSort</a>
* <a href="#二叉搜索树">二叉搜索树</a>
* <a href="#LRU">LRU</a>
* <a href="#堆">堆</a>
* <a href="https://github.com/chentaihan/container/tree/master/pool">内存池</a>
* <a href="https://github.com/chentaihan/container/tree/master/common">常用函数</a>

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

```

## <a id="栈">2：栈</a>

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

## <a id="单链表">4：单链表</a>

```go
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
		fmt.Println("RemoveValue error")
	}
	array := list.ToList()
	fmt.Println(array)
	for i := 0; i < count; i++ {
		list.PushBack(10)
	}
	array = list.ToList()
	fmt.Println(array)
	if list.RemoveValue(10) != count {
		fmt.Println("RemoveValue error")
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


```

## <a id="Map">4：Map</a>

```go
package main

import (
	"fmt"
	"github.com/chentaihan/container/hashmap"
)

func main() {
	tests := []struct {
		key   string
		value string
	}{
		{
			"key1",
			"value1",
		},
		{
			"key2",
			"value2",
		},
		{
			"key3",
			"value3",
		},
		{
			"key4",
			"value4",
		},
		{
			"key5",
			"value5",
		},
	}
	sm := hashmap.NewMapSync()
	for _, test := range tests {
		sm.Set(test.key, test.value)
	}
	for _, test := range tests {
		value, _ := sm.Get(test.key)
		if value.(string) != test.value {
			fmt.Println("equal ", test.key, value.(string), test.value)
		}
		if !sm.Exist(test.key) {
			fmt.Println("exist ", test.key, test.value)
		}
	}
	if sm.Len() != len(tests) {
		fmt.Println("len: ", sm.Len(), len(tests))
	}
	if sm.Exist("asdfghjtre") {
		fmt.Println("exist ", "asdfghjtre")
	}
	data, _ := sm.Marshal()
	fmt.Println(string(data))
	fmt.Println("success")

	dataString := `{"key1":"value1","key2":"value2","key3":"value3","key4":"value4","key5":"value5","key6":"value6"}`
	err := sm.Unmarshal([]byte(dataString))
	if err != nil {
		fmt.Println(err)
	}
	for _, test := range tests {
		if !sm.Exist(test.key) {
			fmt.Println("exist ", test.key, test.value)
		}
	}
	if !sm.Exist("key6") {
		fmt.Println("key6 not exist ")
	}
	value6, _ := sm.Get("key6")
	if value6 != "value6" {
		fmt.Println("key6 value error ", value6)
	}
	sm.Clear()
	if sm.Len() != 0 {
		fmt.Println("clear ", sm.Len())
	}
}

```


## <a id="Set">5：Set</a>

```go
type ISet interface {
	Add(val int) bool    //添加元素
	Exist(val int) bool  //判断是否存在
	Remove(val int) bool //删除指定的值
	Len() int            //元素个数
	Clear()              //删除所有元素
	GetArray() []int     //返回所有元素（不复制）
	Copy() []int         //复制并返回所有元素
}

```

```go
package main

import (
	"fmt"
	"github.com/chentaihan/container/common"
	"github.com/chentaihan/container/set"
)

func main() {
	tests := []struct {
		list     []int
		sortList []int
		size     int
	}{
		{
			[]int{1, 1, 1, 1, 1, 1, 1},
			[]int{1},
			1,
		},
		{
			[]int{1, 3, 5, 7, 9, 2, 4, 6, 8, 0},
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			10,
		},
	}

	for index, test := range tests {
		var s set.ISet = set.NewSet()
		for i := 0; i < len(test.list); i++ {
			s.Add(test.list[i])
		}
		if s.Len() != test.size {
			fmt.Println("size error", index)
		}
		if !common.IntEqualSort(s.GetArray(), test.sortList) {
			fmt.Println("add error", index)
		}
		s.Remove(1)
		s.Clear()
	}
}

```


## <a id="二叉搜索树">6：二叉搜索树</a>

```go
package main

import (
	"fmt"
	"github.com/chentaihan/container/common"
	"github.com/chentaihan/container/tree"
)

func main() {
	tests := []struct {
		nums       []int
		depth      int
		count      int
		minVal     int
		maxVal     int
		findVal    int
		findResult bool
		list       []int
	}{
		{
			[]int{5, 4, 8, 2, 3, 9, 1, 6, 7, 15, 12, 20, 14, 11, 25, 18},
			6,
			16,
			1,
			25,
			18,
			true,
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 14, 15, 18, 20, 25},
		},
		{
			[]int{5, 4, 8, 2, 3, 9, 1, 6, 7, 15, 12, 20, 14, 11, 25, 18, 30},
			7,
			17,
			1,
			30,
			13,
			false,
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 14, 15, 18, 20, 25, 30},
		},
	}
	for _, test := range tests {
		tree := tree.NewBinaryTreeInt()
		tree.AddRange(test.nums)
		if tree.GetDepth() != test.depth {
			fmt.Println("GetDepth error")
		}
		count := tree.GetCount()
		if count != test.count {
			fmt.Println("GetCount error")
		}
		if test.minVal != tree.MinNode(tree.GetRoot()).Val {
			fmt.Println("MinNode error")
		}
		if test.maxVal != tree.MaxNode(tree.GetRoot()).Val {
			fmt.Println("MaxNode error")
		}
		if test.findResult != (tree.Find(test.findVal) != nil) {
			fmt.Println("Find error")
		}
		if !common.IntEqual(test.list, tree.ToList()) {
			fmt.Println("ToList error")
		}
		tree.Remove(1)
	}
}


```

## <a id="LRU">6：LRU</a>

```go
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

```

## <a id="堆">7：堆</a>

```go
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

```















































































































































































































































































