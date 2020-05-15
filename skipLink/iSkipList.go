package skipLink

/*
跳跃表接口
*/

type ISkipList interface {
	Add(val IObject)               //添加元素
	Get(index int) (IObject, bool) //根据下标获取元素
	Remove(value IObject) bool     //删除指定的值
	Len() int                      //元素个数
	Clear()                        //删除所有元素
	GetArray() []IObject           //返回所有元素（不复制）
	Copy() []IObject               //复制所有元素
}

type IObject interface {
	GetHashCode() int //按照这个函数排序
}
