package set

import (
	"github.com/chentaihan/container/binaryTree"
	"unsafe"
)

type TreeSet struct {
	tree binaryTree.ITree
}

func NewTreeSet() ISet {
	return &TreeSet{
		tree: binaryTree.NewBinaryTree(),
	}
}

func (as *TreeSet) Add(val IObject) bool {
	if as.tree.Find(val) != nil {
		return false
	}
	as.tree.Add(val)
	return true
}

func (as *TreeSet) Exist(val IObject) bool {
	return as.tree.Find(val) != nil
}

func (as *TreeSet) Remove(val IObject) bool {
	return as.tree.Remove(val)
}

func (as *TreeSet) Len() int {
	return as.tree.GetCount()
}

func (as *TreeSet) Clear() {
	as.tree = binaryTree.NewBinaryTree()
}

func (as *TreeSet) GetArray() []IObject {
	list := as.tree.ToList()
	return *(*[]IObject)(unsafe.Pointer(&list))
}
