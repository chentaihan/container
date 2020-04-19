package array

//数组接口

type IArray interface {
	Add(val IObject)                       //添加元素
	Get(index int) (IObject, bool)         //根据下标获取元素
	Index(val IObject) int                 //获取指定值对应的下标，不存在就返回-1
	RemoveIndex(index int) (IObject, bool) //删除对应下标的元素
	Remove(value IObject) int              //删除指定的值
	Len() int                              //元素个数
	Clear()                                //删除所有元素
	GetArray() []IObject                   //返回所有元素（不复制）
	Copy() []IObject                       //复制所有元素
}

type IObject interface {
	GetValue() int //按照这个函数排序
}
