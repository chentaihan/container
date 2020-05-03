package btree

type IBTree interface {
	Clone() IBTree               //复制
	Add(item IObject) IObject    //添加元素，已经存在的更新
	Remove(item IObject) IObject //删除指定元素
	RemoveMin() IObject          //删除最小值
	RemoveMax() IObject          //删除最大值
	Find(key IObject) IObject    //查找元素
	Min() IObject                //返回最小值
	Max() IObject                //返回最大值
	Len() int                    //元素总数
	Clear(toFreeList bool)       //删除所有元素
}

type IObject interface {
	Less(than IObject) bool
}
