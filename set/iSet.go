package set

type ISet interface {
	Add(val int) bool
	Exist(val int) bool
	Remove(val int) bool
	Len() int
	Clear()
	GetArray() []int
	Copy() []int
}
