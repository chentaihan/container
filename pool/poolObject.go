package pool

/**
对象内存池
*/
type IObject interface {
	New() interface{}
}

type PoolObject struct {
	list      []interface{}
	newObject func() interface{}
}

func NewPoolObject(newObject func() interface{}) IPoolObject {
	po := &PoolObject{
		newObject: newObject,
	}
	return po
}

func (po *PoolObject) Get() interface{} {
	if len(po.list) == 0 {
		return po.newObject()
	}
	obj := po.list[len(po.list)-1]
	po.list[len(po.list)-1] = nil
	po.list = po.list[:len(po.list)-1]
	return obj
}

func (po *PoolObject) Put(obj interface{}) {
	if obj == nil {
		panic("PoolObject.Put obj is nil")
	}
	po.list = append(po.list, obj)
}

func (po *PoolObject) Len() int {
	return len(po.list)
}

func (po *PoolObject) Clear() {
	po.list = nil
}

func (po *PoolObject) Resize(size uint64) {
	l := int(size)
	if l > len(po.list) {
		for i := len(po.list); i < l; i++ {
			po.list = append(po.list, po.newObject())
		}
	} else {
		for i := l; i < len(po.list); i++ {
			po.list[i] = nil
		}
		po.list = po.list[:size]
	}
}
