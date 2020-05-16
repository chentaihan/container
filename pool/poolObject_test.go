package pool

import "testing"

func TestNewPoolObject(t *testing.T) {
	pool := NewPoolObject(func() interface{} {
		return new(int)
	})
	var count int = 100
	for i := 0; i < count; i++ {
		obj := new(int)
		pool.Put(obj)
	}
	if pool.Len() != count {
		t.Fatal("put error")
	}
	for i := 0; i < count; i++ {
		pool.Get()
	}
	if pool.Len() != 0 {
		t.Fatal("get error")
	}
	pool.Resize(uint64(count))
	if pool.Len() != count {
		t.Fatal("Resize error")
	}
	pool.Clear()
	if pool.Len() != 0 {
		t.Fatal("clear error")
	}
}

func TestNewPoolObjectSync(t *testing.T) {
	pool := NewPoolObjectSync(func() interface{} {
		return new(int)
	})
	count := 100
	for i := 0; i < count; i++ {
		obj := new(int)
		pool.Put(obj)
	}
	if pool.Len() != count {
		t.Fatal("put error")
	}
	for i := 0; i < count; i++ {
		pool.Get()
	}
	if pool.Len() != 0 {
		t.Fatal("get error")
	}
	pool.Resize(uint64(count))
	if pool.Len() != count {
		t.Fatal("Resize error")
	}
	pool.Clear()
	if pool.Len() != 0 {
		t.Fatal("clear error")
	}
}
