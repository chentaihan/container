package tree

import (
	"github.com/chentaihan/container/queue"
)

type ITreeObject interface {
	GetHashCode() int
}
type TreeNode struct {
	Val    ITreeObject
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
func (tree *BinaryTree) AddRange(nums []ITreeObject) {
	for i := 0; i < len(nums); i++ {
		tree.Add(nums[i])
	}
}

/**
树添加节点
*/
func (tree *BinaryTree) Add(val ITreeObject) {
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
		if curNode.Val.GetHashCode() < val.GetHashCode() {
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
func (tree *BinaryTree) Find(val ITreeObject) *TreeNode {
	root := tree.root
	for root != nil {
		if root.Val.GetHashCode() == val.GetHashCode() {
			return root
		}
		if root.Val.GetHashCode() < val.GetHashCode() {
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
func (tree *BinaryTree) Remove(val ITreeObject) bool {
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
func (tree *BinaryTree) ToList() []ITreeObject {
	list := make([]ITreeObject, 0)
	toList(tree.root, &list)
	return list
}

func toList(root *TreeNode, list *[]ITreeObject) {
	if root != nil {
		toList(root.Left, list)
		*list = append(*list, root.Val)
		toList(root.Right, list)
	}
}

//层序遍历
func (tree *BinaryTree) FloorList(root *TreeNode) []ITreeObject {
	ret := make([]ITreeObject, 0)
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
