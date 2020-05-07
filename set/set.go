package set

type Set struct {
	m map[int]IObject
}

func NewSet() ISet {
	return &Set{
		m: make(map[int]IObject),
	}
}

func (as *Set) Add(val IObject) bool {
	if _, exist := as.m[val.GetHashCode()]; exist {
		return false
	}
	as.m[val.GetHashCode()] = val
	return true
}

func (as *Set) Exist(val IObject) bool {
	_, exist := as.m[val.GetHashCode()]
	return exist
}

func (as *Set) Remove(val IObject) bool {
	delete(as.m, val.GetHashCode())
	return true
}

func (as *Set) Len() int {
	return len(as.m)
}

func (as *Set) Clear() {
	as.m = make(map[int]IObject)
}

func (as *Set) GetArray() []IObject {
	list := make([]IObject, len(as.m))
	index := 0
	for _, val := range as.m {
		list[index] = val
		index++
	}
	return list
}
