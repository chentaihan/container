package main

import (
	"fmt"
	"github.com/chentaihan/container/pool"
)

func main() {
	testPoolObject()
	testPoolObjectSync()
}

func testPoolObject() {
	pool := pool.NewPoolObject(func() interface{} {
		return new(int)
	})
	var count int = 100
	for i := 0; i < count; i++ {
		obj := new(int)
		pool.Put(obj)
	}
	if pool.Len() != count {
		fmt.Println("put error")
	}
	for i := 0; i < count; i++ {
		pool.Get()
	}
	if pool.Len() != 0 {
		fmt.Println("get error")
	}
	pool.Resize(uint64(count))
	if pool.Len() != count {
		fmt.Println("Resize error")
	}
	pool.Clear()
	if pool.Len() != 0 {
		fmt.Println("clear error")
	}
}

func testPoolObjectSync() {
	pool := pool.NewPoolObjectSync(func() interface{} {
		return new(int)
	})
	count := 100
	for i := 0; i < count; i++ {
		obj := new(int)
		pool.Put(obj)
	}
	if pool.Len() != count {
		fmt.Println("put error")
	}
	for i := 0; i < count; i++ {
		pool.Get()
	}
	if pool.Len() != 0 {
		fmt.Println("get error")
	}
	pool.Resize(uint64(count))
	if pool.Len() != count {
		fmt.Println("Resize error")
	}
	pool.Clear()
	if pool.Len() != 0 {
		fmt.Println("clear error")
	}
}
