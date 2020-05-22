package skipLink

/*
跳跃表接口
*/

type ISkipList interface {
	Add(key int, val interface{})       //添加元素，已经存在的覆盖
	Find(key int) (interface{}, bool)   //根据key获取值
	Remove(key int) (interface{}, bool) //根据key删除指定的值
	Len() int                           //元素个数
	Clear()                             //删除所有元素
	GetValues() []interface{}           //返回所有元素的值
	GetKeys() []int                     //返回所有元素的key
	GetLevel() int                      //层数
}
