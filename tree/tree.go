package tree

import (
	"fmt"
	"github.com/chentaihan/container/common"
	"github.com/chentaihan/container/queue"
)

type TreeNode struct {
	Val    int
	Left   *TreeNode
	Right  *TreeNode
	Parent *TreeNode
}

type BinaryTree struct {
	root  *TreeNode
	count int
}

func NewBinaryTree() *BinaryTree {
	return &BinaryTree{}
}

/**
构造一棵树
*/
func (tree *BinaryTree) AddRange(nums []int) {
	for i := 0; i < len(nums); i++ {
		tree.Add(nums[i])
	}
}

/**
树添加节点
*/
func (tree *BinaryTree) Add(val int) {
	node := &TreeNode{
		Val: val,
	}
	tree.count++
	if tree.root == nil {
		tree.root = node
		return
	}
	curNode := tree.root
	for curNode != nil {
		if curNode.Val < val {
			if curNode.Right == nil {
				curNode.Right = node
				node.Parent = curNode
				break
			}
			curNode = curNode.Right
		} else {
			if curNode.Left == nil {
				curNode.Left = node
				node.Parent = curNode
				break
			}
			curNode = curNode.Left
		}
	}
}

/**
查找节点
*/
func (tree *BinaryTree) Find(val int) *TreeNode {
	root := tree.root
	for root != nil {
		if root.Val == val {
			return root
		}
		if root.Val < val {
			root = root.Right
		} else {
			root = root.Left
		}
	}
	return nil
}

func (tree *BinaryTree) GetDepth() int {
	return treeDepth(tree.root)
}

/**
计算树的深度
*/
func treeDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := treeDepth(root.Left) + 1
	rightDepth := treeDepth(root.Right) + 1
	if leftDepth > rightDepth {
		return leftDepth
	}
	return rightDepth
}

func (tree *BinaryTree) GetCount() int {
	return tree.count
}

/**
最小节点
*/
func (tree *BinaryTree) MinNode(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	for root.Left != nil {
		root = root.Left
	}
	return root
}

/**
最大节点
*/
func (tree *BinaryTree) MaxNode(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	for root.Right != nil {
		root = root.Right
	}
	return root
}

/**
删除节点
rootNode用双指针是为了能删除根节点
*/
func (tree *BinaryTree) Remove(val int) bool {
	targetNode := tree.Find(val)
	if targetNode == nil {
		return false
	}

	var removeNode *TreeNode
	if targetNode.Right != nil || targetNode.Left != nil {
		//右子树最小值代替当节点
		if targetNode.Right != nil {
			removeNode = tree.MinNode(targetNode.Right)
		} else {
			removeNode = tree.MaxNode(targetNode.Left)
		}

		targetNode.Val = removeNode.Val
		//删除右子树最小值对应的节点
		parentNode := removeNode.Parent
		childNode := removeNode.Left
		if removeNode.Right != nil {
			childNode = removeNode.Right
		}
		if parentNode.Left == removeNode {
			parentNode.Left = childNode
		} else {
			parentNode.Right = childNode
		}
	} else {
		parentNode := targetNode.Parent
		if parentNode == nil { //删除的是根节点，且没有子节点
			tree.root = nil
		} else if parentNode.Left == targetNode { //删除代替要被删除节点的叶子节点
			parentNode.Left = nil
		} else {
			parentNode.Right = nil
		}
	}
	tree.count--
	return true
}

//中序遍历得到排序数组
func (tree *BinaryTree) ToList() []int {
	list := make([]int, 0)
	toList(tree.root, &list)
	return list
}

func toList(root *TreeNode, list *[]int) {
	if root != nil {
		toList(root.Left, list)
		*list = append(*list, root.Val)
		toList(root.Right, list)
	}
}

//层序遍历
func (tree *BinaryTree) FloorList(root *TreeNode) []int {
	ret := make([]int, 0)
	if root == nil {
		return ret
	}
	queue := queue.NewQueue(tree.count / 2)
	queue.Enqueue(root)
	for !queue.Empty() {
		node, _ := queue.Dequeue()
		treeNode, _ := node.(*TreeNode)
		if treeNode.Left != nil {
			queue.Enqueue(treeNode.Left)
		}
		if treeNode.Right != nil {
			queue.Enqueue(treeNode.Right)
		}
		ret = append(ret, treeNode.Val)
	}
	return ret
}

func TestBinaryTree_Remove1() {
	tests := []struct {
		nums    []int
		rmVAlue int
		list    []int
	}{

		{
			[]int{1, 3, 2},
			3,
			[]int{1, 2},
		},
		{
			[]int{1, 3, 2},
			2,
			[]int{1, 3},
		},
		{
			[]int{1, 4, 3, 2},
			4,
			[]int{1, 3, 2},
		},
		{
			[]int{1, 4, 3, 2},
			3,
			[]int{1, 4, 2},
		},
		{
			[]int{1, 4, 3, 2},
			2,
			[]int{1, 4, 3},
		},
		{
			[]int{6, 3, 10, 2, 4, 8, 12},
			6,
			[]int{8, 3, 10, 2, 4, 12},
		},
		{
			[]int{6, 3, 10, 2, 4, 8, 12},
			3,
			[]int{6, 4, 10, 2, 8, 12},
		},
		{
			[]int{6, 3, 10, 2, 4, 8, 12},
			10,
			[]int{6, 3, 12, 2, 4, 8},
		},
		{
			[]int{6, 3, 10, 2, 4, 8, 12},
			2,
			[]int{6, 3, 10, 4, 8, 12},
		},
		{
			[]int{6, 3, 10, 2, 4, 8, 12},
			4,
			[]int{6, 3, 10, 2, 8, 12},
		},
		{
			[]int{6, 3, 10, 2, 4, 8, 12},
			8,
			[]int{6, 3, 10, 2, 4, 12},
		},
		{
			[]int{6, 3, 10, 2, 4, 8, 12},
			12,
			[]int{6, 3, 10, 2, 4, 8},
		},
		{
			[]int{6, 3, 10, 2, 4, 8, 12, 1},
			1,
			[]int{6, 3, 10, 2, 4, 8, 12},
		},
		{
			[]int{6, 3, 10, 2, 4, 8, 12, 1, 15},
			15,
			[]int{6, 3, 10, 2, 4, 8, 12, 1},
		},
		{
			[]int{6, 3, 10, 2, 4, 8, 12, 1, 11},
			11,
			[]int{6, 3, 10, 2, 4, 8, 12, 1},
		},
		{
			[]int{6, 3, 10, 2, 4, 8, 12, 1, 11, 5},
			5,
			[]int{6, 3, 10, 2, 4, 8, 12, 1, 11},
		},
	}
	for index, test := range tests {
		tree := NewBinaryTree()
		tree.AddRange(test.nums)
		tree.Remove(test.rmVAlue)
		if !common.IntEqual(test.list, tree.FloorList(tree.root)) {
			fmt.Println(test.list, test.rmVAlue, tree.FloorList(tree.root))
			fmt.Println("remove error", index)
		}
	}
}
