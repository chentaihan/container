package binaryTree

//树接口

type ITree interface {
	Add(val IObject)                  //添加元素
	Find(val IObject) *TreeNode       //查找元素
	GetRoot() *TreeNode               //获取根节点
	Remove(val IObject) bool          //删除元素
	GetDepth() int                    //树深度
	GetCount() int                    //节点个数
	MinNode(root *TreeNode) *TreeNode //获取最小子节点（从当前节点开始查找）
	MaxNode(root *TreeNode) *TreeNode //获取最大子节点（从当前节点开始查找）
	ToList() []IObject                //获取所有节点值
}

type IObject interface {
	GetHashCode() int
}