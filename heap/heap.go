package heap

type IObject interface {
	GetHashCode() int //按照这个函数排序
}

type IHeap interface {
	Push(h IObject)        //添加元素
	Pop() IObject          //删除顶部元素（第0个元素）
	Peek() IObject         //获取顶部元素（第0个元素）
	Remove(h IObject) bool //删除指定下标元素
	Len() int              //元素个数
	Cap() int              //容量
	Empty() bool           //堆是否为空
	Exist(h IObject) bool  //是否包含指定元素
	Clear()                //删除所有元素
	GetArray() []IObject   //获取元素数组（不复制）
	Copy() []IObject       //复制所有元素
}
