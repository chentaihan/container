package pool

/*
内存池
*/

import (
	"bytes"
	"sync"
)

type PoolSync struct {
	sync.Pool
}

func NewPoolSync() *PoolSync {
	sp := &PoolSync{}
	sp.New = func() interface{} {
		return &bytes.Buffer{}
	}
	return sp
}

func (ps *PoolSync) Get() *bytes.Buffer {
	return ps.Pool.Get().(*bytes.Buffer)
}

func (ps *PoolSync) Put(b *bytes.Buffer) {
	ps.Pool.Put(b)
}
