package tree

type TreeNode struct {
	Val    int
	Left   *TreeNode
	Right  *TreeNode
	Parent *TreeNode
}

type BinaryTree struct {
	root *TreeNode
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
	return getNodeCount(tree.root)
}

/**
计算树总节点数
*/
func getNodeCount(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + getNodeCount(root.Left) + getNodeCount(root.Right)
}

/**
最小节点
*/
func (tree *BinaryTree) MinNode() *TreeNode {
	root := tree.root
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
func (tree *BinaryTree) MaxNode() *TreeNode {
	root := tree.root
	if root == nil {
		return root
	}
	for root.Right != nil {
		root = root.Right
	}
	return root
}

func (tree *BinaryTree) ToList() []int {
	var list []int
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
