package main

import (
	"fmt"
	"github.com/chentaihan/container/btree"
)

// Integer implements the IObject interface for integers.
type Integer int

// Less returns true if int(a) < int(b).
func (a Integer) Less(b btree.IObject) bool {
	return a < b.(Integer)
}

func main() {
	tree := btree.NewBTree(10)
	for i := 0; i < 200; i += 2 {
		tree.Add(Integer(i))
	}
	fmt.Println("len", tree.Len())
	tree.Add(Integer(1))
	ii := tree.Find(Integer(1))
	if ii != nil {
		fmt.Println("find", ii)
	}
	tree.Remove(Integer(1))
	ii = tree.Find(Integer(1))
	if ii != nil {
		fmt.Println(ii)
	} else {
		fmt.Println("not found")
	}
	fmt.Println("min:", tree.Min())
	fmt.Println("max:", tree.Max())
	tree.RemoveMin()
	tree.RemoveMax()
	fmt.Println("min:", tree.Min())
	fmt.Println("max:", tree.Max())
	fmt.Println("Find 1=", tree.Find(Integer(1)))
	fmt.Println("Find 20=", tree.Find(Integer(20)))
	tree.Clear(false)
	fmt.Println("len", tree.Len())

}
