package lru

/**
LRU是Least Recently Used的缩写，即最近最少使用，是一种常用的页面置换算法，
选择最近最久未使用的页面予以淘汰。该算法赋予每个页面一个访问字段，
用来记录一个页面自上次被访问以来所经历的时间 t，当须淘汰一个页面时，
选择现有页面中其 t 值最大的，即最近最少使用的页面予以淘汰。
 */

type ILru interface {
	Add(key string, val interface{})       //添加元素
	Get(key string) (interface{}, bool)    //获取元素，不存在返回false
	Remove(key string) (interface{}, bool) //删除元素，删除失败返回false
	Len() int                              //元素个数
	Cap() int                              //容量
	SetCap(cap int)                        //调整容量，只缩不增
	Clear()                                //删除所有元素
	Values() []interface{}                 //获取所有值
	Keys() []string                        //获取所有key
}
