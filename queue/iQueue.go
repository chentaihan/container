package queue

type IQueue interface {
	Enqueue(val interface{})
	Dequeue() (interface{}, bool)
	Empty() bool
	Len() int
	Cap() int
	Peek() (interface{}, bool)
	ToList() []interface{}
}
