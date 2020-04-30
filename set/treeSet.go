package set

import "github.com/chentaihan/container/binaryTree"

type TreeSet struct {
	tree *binaryTree.BinaryTreeInt
}

func NewTreeSet() ISet {
	return &TreeSet{
		tree: binaryTree.NewBinaryTreeInt(),
	}
}

func (as *TreeSet) Add(val int) bool {
	if as.tree.Find(val) != nil {
		return false
	}
	as.tree.Add(val)
	return true
}

func (as *TreeSet) Exist(val int) bool {
	return as.tree.Find(val) != nil
}

func (as *TreeSet) Remove(val int) bool {
	return as.tree.Remove(val)
}

func (as *TreeSet) Len() int {
	return as.tree.GetCount()
}

func (as *TreeSet) Clear() {
	as.tree = binaryTree.NewBinaryTreeInt()
}

func (as *TreeSet) GetArray() []int {
	return as.tree.ToList()
}

func (as *TreeSet) Copy() []int {
	list := make([]int, as.tree.GetCount())
	copy(list, as.GetArray())
	return list
}
