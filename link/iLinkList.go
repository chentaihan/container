package link

//链表接口

type ILinkList interface {
	PushFront(val interface{})        //首部添加元素
	PushBack(val interface{})         //尾部添加元素
	RemoveFront() (interface{}, bool) //删除首部元素，成功true，失败false
	RemoveBack() (interface{}, bool)  //删除尾部元素，成功true，失败false
	RemoveValue(val interface{}) int  ////删除指定值，返回被删除元素数量
	Front() (interface{}, bool)       //返回首部元素，存在true，不存在false
	Back() (interface{}, bool)        //返回尾部元素，存在true，不存在false
	Len() int                         //元素个数
	Exist(val interface{}) bool       //指定元素是否存在
	ToList() []interface{}            //转换成数组
	Clear()                           //删除所有元素
}
