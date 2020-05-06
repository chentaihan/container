package pool

/**
对象内存池
*/

type IPoolObject interface {
	Get() interface{}    //从内存次中获取一个对象
	Put(obj interface{}) //将对象存入内存池
	Len() int            //内存池中对象个数
	Clear()              //删除内存池中所有对象
	Resize(size uint64)  //调整内存池中对象到指定数量
}
