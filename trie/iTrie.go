package trie

/*
前缀树接口
 */

type ITrie interface {
	Add(word string)            //添加元素
	Find(word string) bool      //查找元素是否存在（完全匹配）
	Count(word string) int      //元素出现的次数
	StartWith(word string) bool //查找元素是否存在（前缀匹配）
	Remove(word string) bool    //删除元素
	GetWordCount() int          //获取元素综述
	GetTopN(n int) []string     //出现最多次数的N个元素
}
