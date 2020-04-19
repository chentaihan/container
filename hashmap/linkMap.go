package hashmap

//对map的简单封装

import (
	"container/list"
	"encoding/json"
)

type LinkMap struct {
	list list.List
	m    map[string]*list.Element
}

func NewLinkMap() IMap {
	return &LinkMap{
		list: list.List{},
		m:    make(map[string]*list.Element),
	}
}

func (lm *LinkMap) Set(key string, val interface{}) {
	item, exist := lm.m[key]
	if !exist {
		et := &entity{
			key:   key,
			value: val,
		}
		lm.m[key] = lm.list.PushBack(et)
	} else {
		item.Value.(*entity).value = val
	}
}

func (lm *LinkMap) Get(key string) (interface{}, bool) {
	if item, exist := lm.m[key]; !exist {
		return nil, exist
	} else {
		curItem, _ := item.Value.(*entity)
		return curItem.value, exist
	}
}

func (lm *LinkMap) Remove(key string) bool {
	if item, exist := lm.m[key]; !exist {
		return exist
	} else {
		lm.list.Remove(item)
		curItem, _ := item.Value.(*entity)
		delete(lm.m, curItem.key)
		return exist
	}
}

func (lm *LinkMap) Exist(key string) bool {
	_, exist := lm.m[key]
	return exist
}

func (lm *LinkMap) Len() int {
	return len(lm.m)
}

func (lm *LinkMap) Cap() int {
	return len(lm.m)
}

func (lm *LinkMap) Clear() {
	lm.list.Init()
	lm.m = make(map[string]*list.Element)
}

func (lm *LinkMap) Marshal() ([]byte, error) {
	return json.Marshal(lm.m)
}

func (lm *LinkMap) Unmarshal(data []byte) error {
	m := map[string]interface{}{}
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for key, value := range m {
		lm.Set(key, value)
	}
	return err
}

func (lm *LinkMap) Values() []interface{} {
	list := make([]interface{}, 0, lm.list.Len())
	for item := lm.list.Front(); item != nil; item = item.Next() {
		curItem, _ := item.Value.(*entity)
		list = append(list, curItem.value)
	}
	return list
}

func (lm *LinkMap) Keys() []string {
	list := make([]string, 0, lm.list.Len())
	for item := lm.list.Front(); item != nil; item = item.Next() {
		curItem, _ := item.Value.(*entity)
		list = append(list, curItem.key)
	}
	return list
}
