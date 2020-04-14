package set

type Set struct {
	list []int
	size int
	m    map[int]int
}

func NewSet() *Set {
	return &Set{
		list: []int{},
		m:    make(map[int]int),
	}
}

func (as *Set) Add(val int) bool {
	if _, exist := as.m[val]; exist {
		return false
	}
	as.m[val] = as.size
	if as.size < len(as.list) {
		as.list[as.size] = val
	} else {
		as.list = append(as.list, val)
	}
	as.size++
	return true
}

func (as *Set) Exist(val int) bool {
	_, exist := as.m[val]
	return exist
}

func (as *Set) Remove(val int) bool {
	if index, exist := as.m[val]; !exist {
		return false
	} else {
		as.list[index] = as.list[as.size-1]
		as.size--
		delete(as.m, val)
		return true
	}
}

func (as *Set) Len() int {
	return as.size
}

func (as *Set) Clear() {
	as.list = []int{}
	as.size = 0
	as.m = make(map[int]int)
}

func (as *Set) GetArray() []int {
	return as.list[:as.size]
}

func (as *Set) Copy() []int {
	list := make([]int, as.size)
	copy(list, as.GetArray())
	return list
}
