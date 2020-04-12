package common

type IStack interface {
	Push(x interface{})
	Pop() (interface{}, bool)
	Top() (interface{}, bool)
	Empty() bool
	Len() int
	Cap() int
}
