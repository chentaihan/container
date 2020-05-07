package array

func NewArraySort(cap int) IArray {
	if cap < 0 {
		cap = 0
	}
	return &ArraySort{
		buf:  make([]IObject, cap),
		size: 0,
	}
}

type ArraySort struct {
	buf  []IObject
	size int
}

func (as *ArraySort) Add(val IObject) {
	as.checkCap()
	index, _ := BinarySearchPos(as.buf[:as.size], val)
	for i := as.size - 1; i >= index; i-- {
		as.buf[i+1] = as.buf[i]
	}
	as.buf[index] = val
	as.size++
}

func (as *ArraySort) Get(index int) (IObject, bool) {
	if index >= as.size || index < 0 {
		return nil, false
	}
	return as.buf[index], true
}

func (as *ArraySort) Index(val IObject) int {
	return BinarySearch(as.GetArray(), val)
}

func (as *ArraySort) RemoveIndex(index int) (IObject, bool) {
	if index >= as.size || index < 0 {
		return nil, false
	}
	val := as.buf[index]
	for i := index; i < as.size-1; i++ {
		as.buf[index] = as.buf[index+1]
	}
	as.size--
	return val, true
}

func (as *ArraySort) Remove(value IObject) int {
	index := BinarySearch(as.GetArray(), value)
	if index < 0 {
		return 0
	}
	start, end := index, index
	for start >= 0 && as.buf[start].GetHashCode() == value.GetHashCode() {
		start--
	}
	for end < as.size && as.buf[end] == value {
		end++
	}
	as.buf = append(as.buf[:start+1], as.buf[end:]...)
	count := end - start - 1
	as.size -= count
	return count
}

func (as *ArraySort) Len() int {
	return as.size
}

func (as *ArraySort) checkCap() {
	if as.size < cap(as.buf) {
		return
	}
	newCap := 0
	if cap(as.buf) > 1024*1024 { //大于1M，每次扩大1M
		newCap = as.size + 1024*1024
	} else if cap(as.buf) < 4 { //小于4，就是4
		newCap = 4
	} else { //2倍扩容
		newCap = cap(as.buf) << 1
	}
	buf := make([]IObject, newCap)
	copy(buf, as.GetArray())
	as.buf = buf
}

func (as *ArraySort) Clear() {
	as.buf = []IObject{}
	as.size = 0
}

func (as *ArraySort) GetArray() []IObject {
	return as.buf[:as.size]
}

func (as *ArraySort) Copy() []IObject {
	list := make([]IObject, as.size)
	copy(list, as.GetArray())
	return list
}

//二分查找
func BinarySearch(nums []IObject, target IObject) int {
	start, middle, end := 0, 0, len(nums)-1
	for start <= end {
		middle = start + (end-start)/2
		if nums[middle].GetHashCode() == target.GetHashCode() {
			return middle
		} else if nums[middle].GetHashCode() > target.GetHashCode() {
			end = middle - 1
		} else {
			start = middle + 1
		}
	}
	return -1
}

//二分查找插入的位置
func BinarySearchPos(nums []IObject, target IObject) (int, bool) {
	if len(nums) == 0 {
		return 0, false
	}
	start, middle, end := 0, 0, len(nums)-1
	for start <= end {
		middle = start + (end-start)/2
		if nums[middle].GetHashCode() == target.GetHashCode() {
			return middle, true
		} else if nums[middle].GetHashCode() > target.GetHashCode() {
			end = middle - 1
		} else {
			start = middle + 1

		}
	}

	if target.GetHashCode() < nums[middle].GetHashCode() {
		return middle, false
	} else {
		return middle + 1, false
	}
}
