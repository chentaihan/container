package set

type Set struct {
	m map[IObject]struct{}
}

func NewSet() ISet {
	return &Set{
		m: make(map[IObject]struct{}),
	}
}

func (as *Set) Add(val IObject) bool {
	if _, exist := as.m[val]; exist {
		return false
	}
	as.m[val] = struct{}{}
	return true
}

func (as *Set) Exist(val IObject) bool {
	_, exist := as.m[val]
	return exist
}

func (as *Set) Remove(val IObject) bool {
	delete(as.m, val)
	return true
}

func (as *Set) Len() int {
	return len(as.m)
}

func (as *Set) Clear() {
	as.m = make(map[IObject]struct{})
}

func (as *Set) GetArray() []IObject {
	list := make([]IObject, len(as.m))
	index := 0
	for key := range as.m {
		list[index] = key
		index++
	}
	return list
}
