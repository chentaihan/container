package hashmap

//同步map，读写锁map

import (
	"encoding/json"
	"sync"
)

type MapSync struct {
	mData map[string]interface{}
	lock  sync.RWMutex
}

func NewMapSync() IMap {
	return &MapSync{
		mData: map[string]interface{}{},
	}
}

func (sm *MapSync) Get(key string) (interface{}, bool) {
	sm.lock.RLock()
	value, exist := sm.mData[key]
	sm.lock.RUnlock()
	return value, exist
}

func (sm *MapSync) Set(key string, value interface{}) {
	sm.lock.Lock()
	sm.mData[key] = value
	sm.lock.Unlock()
}

func (sm *MapSync) Remove(key string) bool {
	sm.lock.Lock()
	delete(sm.mData, key)
	sm.lock.Unlock()
	return true
}

func (sm *MapSync) Clear() {
	sm.lock.Lock()
	sm.mData = map[string]interface{}{}
	sm.lock.Unlock()
}

func (sm *MapSync) Exist(key string) bool {
	sm.lock.RLock()
	_, exist := sm.mData[key]
	sm.lock.RUnlock()
	return exist
}

func (sm *MapSync) Marshal() ([]byte, error) {
	sm.lock.RLock()
	data, err := json.Marshal(sm.mData)
	sm.lock.RUnlock()
	return data, err
}

func (sm *MapSync) Unmarshal(data []byte) error {
	m := map[string]interface{}{}
	err := json.Unmarshal(data, &m)
	sm.lock.Lock()
	sm.mData = m
	sm.lock.Unlock()
	return err
}

func (sm *MapSync) Len() int {
	if sm == nil {
		return 0
	}
	sm.lock.RLock()
	l := len(sm.mData)
	sm.lock.RUnlock()
	return l
}

func (sm *MapSync) Values() []interface{} {
	list := make([]interface{}, len(sm.mData))
	index := 0
	sm.lock.RLock()
	for _, val := range sm.mData {
		list[index] = val
	}
	sm.lock.RUnlock()
	return list
}

func (sm *MapSync) Keys() []string {
	list := make([]string, len(sm.mData))
	index := 0
	sm.lock.RLock()
	for key := range sm.mData {
		list[index] = key
	}
	sm.lock.RUnlock()
	return list
}
