package main

import (
	"fmt"
	"github.com/chentaihan/container/link"
	"github.com/chentaihan/container/skipLink"
	"time"
)

func main() {
	sl := skipLink.NewSkipList()
	for i := 0; i < 10; i += 2 {
		sl.Add(i, i)
	}
	fmt.Println("level:", sl.GetLevel())
	for i := 1; i < 10; i++ {
		sl.Add(i, i)
	}
	fmt.Println("level:", sl.GetLevel())
	sl.Remove(5)
	fmt.Println(sl.GetValues())
	sl.Add(5,5)
	sl.Remove(0)
	sl.Remove(9)
	fmt.Println(sl.GetValues())

	sl.Add(-1, -1)
	sl.Add(20, 20)
	fmt.Println(sl.GetValues())
	for i := -1; i < 22; i++ {
		val, exist := sl.Find(i)
		fmt.Println(i, val, exist)
	}
	fmt.Println()
	for i := -1; i < 22; i++ {
		val, exist := sl.Remove(i)
		fmt.Println(i, val, exist)
	}
	fmt.Println(sl.GetValues())
	for i := -1; i < 22; i++ {
		val, exist := sl.Remove(i)
		fmt.Println(i, val, exist)
	}
	fmt.Println(sl.GetValues())
	for i := 0; i < 10; i++ {
		sl.Add(i, i)
	}
	fmt.Println(sl.Len())
	sl.Clear()
	fmt.Println(sl.Len())
	for i := 0; i < 10; i++ {
		sl.Add(i, i)
	}
	fmt.Println(sl.GetValues())
	sl.Clear()
	for i := 0; i < 3000; i += 2 {
		sl.Add(i, i)
		if i%10 == 0 {
			fmt.Println(sl.GetLevel())
		}
	}
	for i := 1; i < 3000; i += 2 {
		sl.Add(i, i)
		if i%10 == 0 {
			fmt.Println(sl.GetLevel())
		}
	}
	fmt.Println("start find...")
	for i := 0; i < 3000; i++ {
		val, exist := sl.Find(i)
		if !exist && val != i {
			fmt.Println(i, val)
		}
	}

	startTime := time.Now()
	for i := 0; i < 3000; i++ {
		val, exist := sl.Remove(i)
		if !exist && val != i {
			fmt.Println(i, val)
		}
	}
	useTime := time.Now().Sub(startTime)
	fmt.Println(useTime)
	fmt.Println(sl.GetValues())
}

func testCompareTime() {
	sl := skipLink.NewSkipList()
	count := 10000
	for i := 0; i < count; i++ {
		sl.Add(i, i)
	}
	startTime := time.Now()
	for i := 0; i < 3000; i++ {
		val, exist := sl.Find(i)
		if !exist && val != i {
			fmt.Println(i, val)
		}
	}
	useTime := time.Now().Sub(startTime)
	fmt.Println("skipList use time=", useTime)

	list := link.NewLinkedList()
	for i := 0; i < count; i++ {
		list.PushBack(i)
	}

	startTime = time.Now()
	for i := 0; i < count; i++ {
		exist := list.Exist(i)
		if !exist {
			fmt.Println(i)
		}
	}
	useTime = time.Now().Sub(startTime)
	fmt.Println("simple link use time =", useTime)
}
