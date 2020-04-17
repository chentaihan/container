package queue

type IQueue interface {
	Enqueue(val interface{})      //入队
	Dequeue() (interface{}, bool) //出队
	Empty() bool                  //队列是否为空
	Len() int                     //队列长度
	Cap() int                     //队列容量
	Peek() (interface{}, bool)    //取队首元素
	ToList() []interface{}        //转换成数组
}
