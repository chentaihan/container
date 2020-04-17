package tree

import (
	"github.com/chentaihan/container/queue"
)

type TreeNodeInt struct {
	Val    int
	Left   *TreeNodeInt
	Right  *TreeNodeInt
	Parent *TreeNodeInt
}

type BinaryTreeInt struct {
	root  *TreeNodeInt
	count int
}

func NewBinaryTreeInt() *BinaryTreeInt {
	return &BinaryTreeInt{}
}

/**
构造一棵树
*/
func (tree *BinaryTreeInt) AddRange(nums []int) {
	for i := 0; i < len(nums); i++ {
		tree.Add(nums[i])
	}
}

func (tree *BinaryTreeInt) GetRoot() *TreeNodeInt {
	return tree.root
}

/**
树添加节点
*/
func (tree *BinaryTreeInt) Add(val int) {
	node := &TreeNodeInt{
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
func (tree *BinaryTreeInt) Find(val int) *TreeNodeInt {
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

func (tree *BinaryTreeInt) GetDepth() int {
	return treeDepthInt(tree.root)
}

/**
计算树的深度
*/
func treeDepthInt(root *TreeNodeInt) int {
	if root == nil {
		return 0
	}
	leftDepth := treeDepthInt(root.Left) + 1
	rightDepth := treeDepthInt(root.Right) + 1
	if leftDepth > rightDepth {
		return leftDepth
	}
	return rightDepth
}

func (tree *BinaryTreeInt) GetCount() int {
	return tree.count
}

/**
最小节点
*/
func (tree *BinaryTreeInt) MinNode(root *TreeNodeInt) *TreeNodeInt {
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
func (tree *BinaryTreeInt) MaxNode(root *TreeNodeInt) *TreeNodeInt {
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
func (tree *BinaryTreeInt) Remove(val int) bool {
	targetNode := tree.Find(val)
	if targetNode == nil {
		return false
	}

	var removeNode *TreeNodeInt
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
func (tree *BinaryTreeInt) ToList() []int {
	list := make([]int, 0)
	toListInt(tree.root, &list)
	return list
}

func toListInt(root *TreeNodeInt, list *[]int) {
	if root != nil {
		toListInt(root.Left, list)
		*list = append(*list, root.Val)
		toListInt(root.Right, list)
	}
}

//层序遍历
func (tree *BinaryTreeInt) FloorList(root *TreeNodeInt) []int {
	ret := make([]int, 0)
	if root == nil {
		return ret
	}
	queue := queue.NewQueue(tree.count / 2)
	queue.Enqueue(root)
	for !queue.Empty() {
		node, _ := queue.Dequeue()
		treeNode, _ := node.(*TreeNodeInt)
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
