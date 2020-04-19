package hashmap

//二叉树实现的map

import (
	"encoding/json"
	"github.com/chentaihan/container/common"
	"github.com/chentaihan/container/tree"
)

type IObject interface {
	GetHashCode() int
}

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
	tree *tree.BinaryTree
}

func NewTreeMap() IMap {
	return &TreeMap{
		tree: tree.NewBinaryTree(),
	}
}

func (as *TreeMap) Set(key string, value interface{}) {
	et := newTreeEntity(key, value)
	if node := as.tree.Find(et); node != nil {
		node.Val = et
	} else {
		as.tree.Add(et)
	}
}

func (as *TreeMap) Get(key string) (interface{}, bool) {
	et := newTreeEntity(key, nil)
	if node := as.tree.Find(et); node != nil {
		return node.Val.(*treeEntity).value, true
	} else {
		return nil, false
	}
}

func (as *TreeMap) Exist(key string) bool {
	et := newTreeEntity(key, nil)
	return as.tree.Find(et) != nil
}

func (as *TreeMap) Remove(key string) bool {
	et := newTreeEntity(key, nil)
	return as.tree.Remove(et)
}

func (as *TreeMap) Len() int {
	return as.tree.GetCount()
}

func (as *TreeMap) Clear() {
	as.tree = tree.NewBinaryTree()
}

func (as *TreeMap) Values() []interface{} {
	list := as.tree.ToList()
	result := make([]interface{}, len(list))
	for i := 0; i < len(list); i++ {
		result[i] = list[i].(*treeEntity).value
	}
	return result
}

func (as *TreeMap) Keys() []string {
	list := as.tree.ToList()
	result := make([]string, len(list))
	for i := 0; i < len(list); i++ {
		result[i] = list[i].(*treeEntity).key
	}
	return result
}

func (sm *TreeMap) Marshal() ([]byte, error) {
	list := sm.tree.ToList()
	m := make(map[string]interface{}, len(list))
	for i := 0; i < len(list); i++ {
		item := list[i].(*treeEntity)
		m[item.key] = item.value
	}
	return json.Marshal(m)
}

func (sm *TreeMap) Unmarshal(data []byte) error {
	m := map[string]interface{}{}
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for key, value := range m {
		sm.Set(key, value)
	}
	return err
}
