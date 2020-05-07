package set

import "github.com/chentaihan/container/binaryTree"

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
	result := make([]IObject, len(list))
	for i := 0; i < len(list); i++ {
		result[i] = list[i]
	}
	return result
}
