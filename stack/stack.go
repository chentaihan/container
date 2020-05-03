package stack

type Stack struct {
	buf  []interface{}
	size int
}

/** Initialize your data structure here. */
func NewStack(size int) IStack {
	if size < 4 {
		size = 4
	}
	return &Stack{
		buf: make([]interface{}, 0, size),
	}
}

/** Push element x onto stack. */
func (sk *Stack) Push(x interface{}) {
	if len(sk.buf) == sk.size {
		sk.buf = append(sk.buf, x)
	} else {
		sk.buf[sk.size] = x
	}
	sk.size++
}

/** Removes the element on top of the stack and returns that element. */
func (sk *Stack) Pop() (interface{}, bool) {
	if sk.size == 0 {
		return nil, false
	}
	val := sk.buf[sk.size-1]
	sk.buf[sk.size-1] = nil //释放对变量的引用
	sk.size--
	return val, true
}

/** Find the top element. */
func (sk *Stack) Top() (interface{}, bool) {
	if sk.size == 0 {
		return nil, false
	}
	return sk.buf[sk.size-1], true
}

/** Returns whether the stack is empty. */
func (sk *Stack) Empty() bool {
	return sk.size == 0
}

/** Returns the stack 's len */
func (sk *Stack) Len() int {
	return sk.size
}

/** Returns the stack 's cap */
func (sk *Stack) Cap() int {
	return cap(sk.buf)
}
