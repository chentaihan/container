package pool

/*
同步-对象内存池
*/

import (
	"sync"
)

type PoolObjectSync struct {
	sync.Pool
	size int
}

func NewPoolObjectSync(newObject func() interface{}) IPoolObject {
	sp := &PoolObjectSync{}
	sp.New = newObject
	return sp
}

func (ps *PoolObjectSync) Get() interface{} {
	if ps.size > 0 {
		ps.size--
	}
	return ps.Pool.Get()
}

func (ps *PoolObjectSync) Put(obj interface{}) {
	if obj == nil {
		panic("PoolObjectSync.Put obj is nil")
	}
	ps.size++
	ps.Pool.Put(obj)
}

func (ps *PoolObjectSync) Len() int {
	return ps.size
}

func (ps *PoolObjectSync) Clear() {
	ps.Pool = sync.Pool{}
	ps.size = 0
}

func (ps *PoolObjectSync) Resize(size uint64) {
	l := int(size)
	if l > ps.size {
		for i := ps.size; i < l; i++ {
			ps.Put(ps.New())
		}
	} else {
		for i := l; i < ps.size; i++ {
			ps.Get()
		}
	}
}
