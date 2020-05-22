package lru

/**
1：基于过期时间的LRU，将元素保存到集合中，expireTime秒后自动删除
2：采取的是惰性删除机制，访问任何一个接口，都会删除一定数量的过期的数据
*/

import (
	"container/heap"
	"time"
)

const REMOVE_MAX_COUNT = 50 //每次最多删除过期数据数量

type IObject interface {
	GetHashCode() int
}

type lruTimeItem struct {
	value IObject
	time  int64 //单位：秒
	index int
}

type lruQueue []*lruTimeItem

func (pq lruQueue) Len() int { return len(pq) }

func (pq lruQueue) Less(i, j int) bool {
	return pq[i].time < pq[j].time
}

func (pq lruQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *lruQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*lruTimeItem)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *lruQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

type LruTime struct {
	m     map[int]*lruTimeItem
	queue lruQueue
}

func NewLruTime(cap int) ILruTime {
	if cap < 0 {
		cap = 0
	}
	return &LruTime{
		queue: make(lruQueue, 0, cap),
		m:     make(map[int]*lruTimeItem),
	}
}

// 添加元素，如果已经存在的就更新过期时间
// expireTime秒后自动删除
func (mq *LruTime) Add(val IObject, expireTime int64) {
	mq.RemoveOutOfTime()
	now := time.Now().Unix()
	now += expireTime
	var item *lruTimeItem
	ok := false
	hashCode := val.GetHashCode()
	if item, ok = mq.m[hashCode]; ok {
		item.time = now
		heap.Fix(&mq.queue, item.index)
	} else {
		item = &lruTimeItem{
			time:  now,
			value: val,
		}
		heap.Push(&mq.queue, item)
	}
	mq.m[hashCode] = item
}

func (mq *LruTime) Get(hashCode int) (IObject, bool) {
	mq.RemoveOutOfTime()
	if val, exist := mq.m[hashCode]; exist {
		return val.value, true
	}
	return nil, false
}

func (mq *LruTime) RemoveOutOfTime() int {
	now := time.Now().Unix()
	count := 0
	for len(mq.queue) > 0 && count < REMOVE_MAX_COUNT {
		if mq.queue[0].time < now {
			delete(mq.m, mq.queue[0].value.GetHashCode())
			heap.Pop(&mq.queue)
		} else {
			break
		}
		count++
	}
	return count
}

func (mq *LruTime) Peek() (IObject, bool) {
	mq.RemoveOutOfTime()
	if mq.Len() == 0 {
		return nil, false
	}
	return mq.queue[0].value, true
}

func (mq *LruTime) Len() int {
	mq.RemoveOutOfTime()
	return len(mq.queue)
}

func (mq *LruTime) Clear() {
	mq.queue = mq.queue[:0]
	mq.m = make(map[int]*lruTimeItem)
}

func (mq *LruTime) GetArray() []IObject {
	mq.RemoveOutOfTime()
	list := make([]IObject, mq.Len())
	for i := 0; i < len(list); i++ {
		list[i] = mq.queue[i].value
	}
	return list
}

func (mq *LruTime) Stop() {

}
