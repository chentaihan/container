package set

type ISet interface {
	Add(val IObject) bool    //添加元素
	Exist(val IObject) bool  //判断是否存在
	Remove(val IObject) bool //删除指定的值
	Len() int                //元素个数
	Clear()                  //删除所有元素
	GetArray() []IObject     //返回所有元素（不复制）
}

type IObject interface {
	GetHashCode() int //按照这个函数排序
}
