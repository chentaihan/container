package cachel

import "container/list"

type lruItem struct {
	key string
	val interface{}
}

type Lru struct {
	list list.List
	m    map[string]*list.Element
	cap  int
	size int
}

func NewLru(cap int) *Lru {
	if cap < 0 {
		cap = 0
	}
	return &Lru{
		list: list.List{},
		m:    make(map[string]*list.Element),
		cap:  cap,
	}
}

func (lru *Lru) Add(key string, val interface{}) {
	li := &lruItem{
		key: key,
		val: val,
	}
	lru.m[key] = lru.list.PushFront(li)
	lru.size++
	lru.removeOldest()
}

func (lru *Lru) removeOldest() {
	if lru.size > lru.cap {
		item := lru.list.Back()
		oldest, _ := item.Value.(*lruItem)
		lru.list.Remove(item)
		delete(lru.m, oldest.key)
		lru.size--
	}
}

func (lru *Lru) Get(key string) (interface{}, bool) {
	if item, exist := lru.m[key]; !exist {
		return nil, exist
	} else {
		curItem, _ := item.Value.(*lruItem)
		lru.list.Remove(item)
		li := &lruItem{
			key: curItem.key,
			val: curItem.val,
		}
		lru.list.PushFront(li)
		return curItem.val, exist
	}
}

func (lru *Lru) Remove(key string) (interface{}, bool) {
	if item, exist := lru.m[key]; !exist {
		return nil, exist
	} else {
		lru.list.Remove(item)
		curItem, _ := item.Value.(*lruItem)
		delete(lru.m, curItem.key)
		lru.size--
		return curItem.val, exist
	}
}

func (lru *Lru) Len() int {
	return lru.size
}

func (lru *Lru) Cap() int {
	return lru.cap
}

func (lru *Lru) SetCap(cap int) {
	if cap < 0 {
		cap = 0
	}
	lru.cap = cap
	for lru.size > lru.cap {
		lru.removeOldest()
	}
}

func (lru *Lru) Clear() {
	lru.list.Init()
	lru.size = 0
	lru.m = make(map[string]*list.Element)
}

func (lru *Lru) Values() []interface{} {
	list := make([]interface{}, 0, lru.list.Len())
	for item := lru.list.Front(); item != nil; item = item.Next() {
		curItem, _ := item.Value.(*lruItem)
		list = append(list, curItem.val)
	}
	return list
}

func (lru *Lru) Keys() []string {
	list := make([]string, 0, lru.list.Len())
	for item := lru.list.Front(); item != nil; item = item.Next() {
		curItem, _ := item.Value.(*lruItem)
		list = append(list, curItem.key)
	}
	return list
}
