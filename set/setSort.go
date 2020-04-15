package set

import "github.com/chentaihan/container/array"

type SetSort struct {
	as *array.ArraySort
}

func NewSetSort() ISet {
	return &SetSort{
		as: array.NewArraySort(0),
	}
}

func (ss *SetSort) Add(val int) bool {
	index := ss.as.Index(val)
	if index >= 0 {
		return false
	}
	ss.as.Add(val)
	return true
}

func (ss *SetSort) Exist(val int) bool {
	return ss.as.Index(val) >= 0
}

func (ss *SetSort) Remove(val int) bool {
	return ss.as.RemoveValue(val) > 0
}

func (ss *SetSort) Len() int {
	return ss.as.Len()
}

func (ss *SetSort) Clear() {
	ss.as.Clear()
}

func (ss *SetSort) GetArray() []int {
	return ss.as.GetArray()
}

func (ss *SetSort) Copy() []int {
	return ss.Copy()
}
