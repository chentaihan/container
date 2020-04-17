package stack

type IStack interface {
	Push(x interface{})       //入栈
	Pop() (interface{}, bool) //出栈
	Top() (interface{}, bool) //栈顶元素
	Empty() bool              //栈是否为空
	Len() int                 //栈元素个数
	Cap() int                 //栈容量
}
