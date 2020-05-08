package hashmap

/*
map接口
 */

type IMap interface {
	Set(key string, value interface{})  //添加元素，已存在就覆盖
	Get(key string) (interface{}, bool) //根据key获取元素，bool：true存在，false不存在
	Exist(key string) bool              //判断key是否存在
	Remove(key string) bool             //删除指定的key
	Len() int                           //元素个数
	Clear()                             //清除所有元素
	Values() []interface{}              //获取所有值
	Keys() []string                     //获取所有key
	Marshal() ([]byte, error)           //序列化
	Unmarshal(data []byte) error        //反序列化
}

type entity struct {
	key string
	value interface{}
}