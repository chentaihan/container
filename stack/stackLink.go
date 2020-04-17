package stack

import "github.com/chentaihan/container/link"

type StackLink struct {
	list link.ILinkList
}

/** Initialize your data structure here. */
func NewStackLink() IStack {
	return &StackLink{
		list: link.NewLinkedList(),
	}
}

/** Push element x onto stack. */
func (sk *StackLink) Push(x interface{}) {
	sk.list.PushFront(x)
}

/** Removes the element on top of the stack and returns that element. */
func (sk *StackLink) Pop() (interface{}, bool) {
	return sk.list.RemoveFront()
}

/** Get the top element. */
func (sk *StackLink) Top() (interface{}, bool) {
	return sk.list.Front()
}

/** Returns whether the stack is empty. */
func (sk *StackLink) Empty() bool {
	return sk.list.Len() == 0
}

/** Returns the stack 's len */
func (sk *StackLink) Len() int {
	return sk.list.Len()
}

/** Returns the stack 's cap */
func (sk *StackLink) Cap() int {
	return sk.list.Len()
}
