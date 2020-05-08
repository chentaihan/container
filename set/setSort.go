package set

import (
	"github.com/chentaihan/container/array"
	"unsafe"
)

type SetSort struct {
	as array.IArray
}

func NewSetSort() ISet {
	return &SetSort{
		as: array.NewArraySort(0),
	}
}

func (ss *SetSort) Add(val IObject) bool {
	index := ss.as.Index(val)
	if index >= 0 {
		return false
	}
	ss.as.Add(val)
	return true
}

func (ss *SetSort) Exist(val IObject) bool {
	return ss.as.Index(val) >= 0
}

func (ss *SetSort) Remove(val IObject) bool {
	return ss.as.Remove(val) > 0
}

func (ss *SetSort) Len() int {
	return ss.as.Len()
}

func (ss *SetSort) Clear() {
	ss.as.Clear()
}

func (ss *SetSort) GetArray() []IObject {
	list := ss.as.GetArray()
	return *(*[]IObject)(unsafe.Pointer(&list))
}
