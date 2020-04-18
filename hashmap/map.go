package hashmap

//对map的简单封装

import (
	"encoding/json"
)

type Map struct {
	mData map[string]interface{}
}

func NewMap() IMap {
	return &Map{
		mData: map[string]interface{}{},
	}
}

func (sm *Map) Get(key string) (interface{}, bool) {
	value, exist := sm.mData[key]
	return value, exist
}

func (sm *Map) Set(key string, value interface{}) {
	sm.mData[key] = value
}

func (sm *Map) Remove(key string) bool {
	delete(sm.mData, key)
	return true
}

func (sm *Map) Clear() {
	sm.mData = map[string]interface{}{}
}

func (sm *Map) Exist(key string) bool {
	_, exist := sm.mData[key]
	return exist
}

func (sm *Map) Marshal() ([]byte, error) {
	return json.Marshal(sm.mData)
}

func (sm *Map) Unmarshal(data []byte) error {
	m := map[string]interface{}{}
	err := json.Unmarshal(data, &m)
	sm.mData = m
	return err
}

func (sm *Map) Len() int {
	if sm == nil {
		return 0
	}
	return len(sm.mData)
}

func (as *Map) Values() []interface{} {
	list := make([]interface{}, len(as.mData))
	index := 0
	for _, val := range as.mData {
		list[index] = val
	}
	return list
}

func (as *Map) Keys() []string {
	list := make([]string, len(as.mData))
	index := 0
	for key := range as.mData {
		list[index] = key
	}
	return list
}
