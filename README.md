# Container

## <a href="https://github.com/chentaihan/container/wiki/golang实现常用集合原理介绍">golang实现常用集合原理介绍</a>  &emsp;&emsp;&emsp;&emsp; <a href="https://github.com/chentaihan/container/tree/master/demo">demo</a>

## 常用集合实现

* <a href="#队列">队列</a>
* <a href="#队列">优先级队列</a>
* <a href="#栈">栈</a>
* <a href="#排序数组">排序数组</a>
* <a href="#单链表">单链表</a>
* <a href="https://github.com/chentaihan/container/tree/master/skipLink">跳跃表</a>
* <a href="#Map">Map</a>
* <a href="#Set">Set</a>
* <a href="#Set">SetSort</a>
* <a href="#二叉搜索树">二叉搜索树</a>
* <a href="https://github.com/chentaihan/container/tree/master/trie">前缀树</a>
* <a href="https://github.com/chentaihan/container/tree/master/btree">btree</a>
* <a href="#LRU">LRU</a>
* <a href="#LRU">LRU过期时间</a>
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
//栈接口

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




## <a id="排序数组">3：排序数组</a>


```go
//数组接口

type IArray interface {
	Add(val IObject)                       //添加元素
	Get(index int) (IObject, bool)         //根据下标获取元素
	Index(val IObject) int                 //获取指定值对应的下标，不存在就返回-1
	RemoveIndex(index int) (IObject, bool) //删除对应下标的元素
	Remove(value IObject) int              //删除指定的值
	Len() int                              //元素个数
	Clear()                                //删除所有元素
	GetArray() []IObject                   //返回所有元素（不复制）
	Copy() []IObject                       //复制所有元素
}

type IObject interface {
	GetHashCode() int //按照这个函数排序
}
```



```go
package main

import (
	"fmt"
	"github.com/chentaihan/container/array"
)

type integer int

func (i integer) GetHashCode() int {
	return int(i)
}

func main() {
	tests := []struct {
		array  []integer
		result []integer
	}{
		{
			[]integer{1, 3, 2, 4, 6, 5, 9, 8, 7},
			[]integer{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}
	for index, test := range tests {
		as := array.NewArraySort(0)
		for i := 0; i < len(test.array); i++ {
			as.Add(test.array[i])
		}
		var result []array.IObject
		for i := 0; i < len(test.result); i++ {
			result = append(result, test.result[i])
		}
		list := as.GetArray()
		if !IntEqual(list, result) {
			fmt.Println(list, test.result)
			fmt.Println("add error ", index)
		}
		as.Remove(integer(1))
	}
}

func IntEqual(nums1, nums2 []array.IObject) bool {
	if len(nums1) != len(nums2) {
		return false
	}
	for i := 0; i < len(nums1); i++ {
		if nums1[i].GetHashCode() != nums2[i].GetHashCode() {
			return false
		}
	}
	return true
}

```

## <a id="单链表">4：单链表</a>


```go
//链表接口
type ILinkList interface {
	PushFront(val interface{})        //首部添加元素
	PushBack(val interface{})         //尾部添加元素
	RemoveFront() (interface{}, bool) //删除首部元素，成功true，失败false
	RemoveBack() (interface{}, bool)  //删除尾部元素，成功true，失败false
	RemoveValue(val interface{}) int  ////删除指定值，返回被删除元素数量
	Front() (interface{}, bool)       //返回首部元素，存在true，不存在false
	Back() (interface{}, bool)        //返回尾部元素，存在true，不存在false
	Len() int                         //元素个数
	Exist(val interface{}) bool       //指定元素是否存在
	ToList() []interface{}            //转换成数组
	Clear()                           //删除所有元素
}
```



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
type IMap interface {
	Set(key string, value interface{})  //添加元素，已存在就覆盖
	Get(key string) (interface{}, bool) //根据key获取元素，bool：true存在，false不存在
	Exist(key string) bool              //判断key是否存在
	Remove(key string) bool             //删除指定的key
	Len() int                           //元素个数
	Clear()                             //清除所有元素
	Values() []interface{}              //获取所有值
	Keys() []string                     //获取所有key
	Marshal() ([]byte, error)           //序列化
	Unmarshal(data []byte) error        //反序列化
}
```



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
	sm := hashmap.NewMap()       //map简单封装
	//sm := hashmap.NewMapSync() //同步map
	//sm := hashmap.NewTreeMap() //二叉树map
	//sm := hashmap.NewLinkMap() //顺序map
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
//树接口

type ITree interface {
	Add(val IObject)                  //添加元素
	Find(val IObject) *TreeNode       //查找元素
	GetRoot() *TreeNode               //获取根节点
	Remove(val IObject) bool          //删除元素
	GetDepth() int                    //树深度
	GetCount() int                    //节点个数
	MinNode(root *TreeNode) *TreeNode //获取最小子节点（从当前节点开始查找）
	MaxNode(root *TreeNode) *TreeNode //获取最大子节点（从当前节点开始查找）
	ToList() []IObject                //获取所有节点值
}
```



```go
package main

import (
	"fmt"
	"github.com/chentaihan/container/tree"
)

type integer int

func (i integer) GetHashCode() int {
	return int(i)
}

func intEqual(nums1, nums2 []tree.IObject) bool {
	if (nums1 == nil && nums2 != nil) || (nums1 != nil && nums2 == nil) {
		return false
	}
	if len(nums1) != len(nums2) {
		return false
	}
	for i := 0; i < len(nums1); i++ {
		if nums1[i].GetHashCode() != nums2[i].GetHashCode() {
			return false
		}
	}
	return true
}


func main() {
	tests := []struct {
		nums       []integer
		depth      int
		count      int
		minVal     int
		maxVal     int
		findVal    int
		findResult bool
		list       []integer
	}{
		{
			[]integer{5, 4, 8, 2, 3, 9, 1, 6, 7, 15, 12, 20, 14, 11, 25, 18, 30},
			7,
			17,
			1,
			30,
			13,
			false,
			[]integer{1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 14, 15, 18, 20, 25, 30},
		},
	}
	for _, test := range tests {
		root := tree.NewBinaryTree()
		for i := 0; i < len(test.nums); i++ {
			root.Add(test.nums[i])
		}
		if root.GetDepth() != test.depth {
			fmt.Println("GetDepth error")
		}
		count := root.GetCount()
		if count != test.count {
			fmt.Println("GetCount error")
		}
		if test.minVal != root.MinNode(root.GetRoot()).Val.GetHashCode() {
			fmt.Println("MinNode error")
		}
		if test.maxVal != root.MaxNode(root.GetRoot()).Val.GetHashCode() {
			fmt.Println("MaxNode error")
		}
		if test.findResult != (root.Find(integer(test.findVal)) != nil) {
			fmt.Println("Find error")
		}
		var list []tree.IObject
		for i := 0; i < len(test.list); i++ {
			list = append(list, test.list[i])
		}
		if !intEqual(list, root.ToList()) {
			fmt.Println("ToList error")
		}
	}
}


```

## <a id="LRU">6：LRU</a>
```go
/**
LRU是Least Recently Used的缩写，即最近最少使用，是一种常用的页面置换算法，
选择最近最久未使用的页面予以淘汰。该算法赋予每个页面一个访问字段，
用来记录一个页面自上次被访问以来所经历的时间 t，当须淘汰一个页面时，
选择现有页面中其 t 值最大的，即最近最少使用的页面予以淘汰。
*/

type ILru interface {
	Add(key string, val interface{})       //添加元素
	Get(key string) (interface{}, bool)    //获取元素，不存在返回false
	Remove(key string) (interface{}, bool) //删除元素，删除失败返回false
	Len() int                              //元素个数
	Cap() int                              //容量
	SetCap(cap int)                        //调整容量，只缩不增
	Clear()                                //删除所有元素
	Values() []interface{}                 //获取所有值
	Keys() []string                        //获取所有key
}

/**
基于过期时间的LRU，将元素保存到集合中，expireTime秒后自动删除
采取的是惰性删除机制，访问任何一个接口，都会删除一定数量的过期的数据
*/
type ILruTime interface {
	// 添加元素，如果已经存在的就更新过期时间
	// expireTime秒后自动删除
	Add(val IObject, expireTime int64)
	//获取元素，不存在返回false
	Get(hashCode int) (IObject, bool)
	//获取队首元素，即将被删除的元素
	Peek() (IObject, bool)
	//元素个数
	Len() int
	//删除所有元素
	Clear()
	//获取所有值
	GetArray() []IObject
	//删除一定数量的过期数据，返回删除数量
	RemoveOutOfTime() int
	//停止处理
	Stop()
}

```

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

```

## <a id="堆">7：堆</a>


```go
//堆接口

type IObject interface {
	GetHashCode() int //按照这个函数排序
}

type IHeap interface {
	Push(h IObject)          //添加元素
	Pop() IObject            //删除顶部元素（第0个元素）
	Peek() IObject           //获取顶部元素（第0个元素）
	Remove(h IObject) bool   //删除指定下标元素
	Len() int                //元素个数
	Cap() int                //容量
	Empty() bool             //堆是否为空
	Exist(h IObject) bool    //是否包含指定元素
	Clear()                  //删除所有元素
	GetArray() []IObject     //获取元素数组（不复制）
	Copy() []IObject         //复制所有元素
}
```



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















































































































































































































































































