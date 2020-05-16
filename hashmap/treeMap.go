package hashmap

//二叉树实现的map

import (
	"encoding/json"
	"github.com/chentaihan/container/common"
	"github.com/chentaihan/container/binaryTree"
)

type treeEntity struct {
	entity
	hashCode int
}

func (et *treeEntity) GetHashCode() int {
	if et.hashCode < 0 {
		et.hashCode = int(common.GetHashCode([]byte(et.key)))
	}
	return et.hashCode
}

func newTreeEntity(key string, value interface{}) *treeEntity {
	return &treeEntity{
		entity: entity{
			key:   key,
			value: value,
		},
		hashCode: -1,
	}
}

type TreeMap struct {
	tree binaryTree.ITree
}

func NewTreeMap() IMap {
	return &TreeMap{
		tree: binaryTree.NewBinaryTree(),
	}
}

func (tm *TreeMap) Set(key string, value interface{}) {
	et := newTreeEntity(key, value)
	if node := tm.tree.Find(et); node != nil {
		node.Val = et
	} else {
		tm.tree.Add(et)
	}
}

func (tm *TreeMap) Get(key string) (interface{}, bool) {
	et := newTreeEntity(key, nil)
	if node := tm.tree.Find(et); node != nil {
		return node.Val.(*treeEntity).value, true
	} else {
		return nil, false
	}
}

func (tm *TreeMap) Exist(key string) bool {
	et := newTreeEntity(key, nil)
	return tm.tree.Find(et) != nil
}

func (tm *TreeMap) Remove(key string) bool {
	et := newTreeEntity(key, nil)
	return tm.tree.Remove(et)
}

func (tm *TreeMap) Len() int {
	return tm.tree.GetCount()
}

func (tm *TreeMap) Clear() {
	tm.tree = binaryTree.NewBinaryTree()
}

func (tm *TreeMap) Values() []interface{} {
	list := tm.tree.ToList()
	result := make([]interface{}, len(list))
	for i := 0; i < len(list); i++ {
		result[i] = list[i].(*treeEntity).value
	}
	return result
}

func (tm *TreeMap) Keys() []string {
	list := tm.tree.ToList()
	result := make([]string, len(list))
	for i := 0; i < len(list); i++ {
		result[i] = list[i].(*treeEntity).key
	}
	return result
}

func (tm *TreeMap) Marshal() ([]byte, error) {
	list := tm.tree.ToList()
	m := make(map[string]interface{}, len(list))
	for i := 0; i < len(list); i++ {
		item := list[i].(*treeEntity)
		m[item.key] = item.value
	}
	return json.Marshal(m)
}

func (tm *TreeMap) Unmarshal(data []byte) error {
	m := map[string]interface{}{}
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for key, value := range m {
		tm.Set(key, value)
	}
	return err
}
