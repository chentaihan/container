package link

type ILinkList interface {
	PushFront(val interface{})
	PushBack(val interface{})
	RemoveFront() (interface{}, bool)
	RemoveBack() (interface{}, bool)
	RemoveValue(val interface{}) int
	Front() (interface{}, bool)
	Back() (interface{}, bool)
	Len() int
	Exist(val interface{}) bool
	ToList() []interface{}
	Clear()
}
