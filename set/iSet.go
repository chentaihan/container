package set

type ISet interface {
	Add(val int) bool    //添加元素
	Exist(val int) bool  //判断是否存在
	Remove(val int) bool //删除指定的值
	Len() int            //元素个数
	Clear()              //删除所有元素
	GetArray() []int     //返回所有元素（不复制）
	Copy() []int         //复制并返回所有元素
}
