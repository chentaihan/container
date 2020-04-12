package hashmap

import (
	"encoding/json"
	"sync"
)

type MapSync struct {
	mData map[string]interface{}
	lock  sync.RWMutex
}

func NewMapSync() *MapSync {
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

func (sm *MapSync) Copy(mData map[string]interface{}) {
	sm.lock.Lock()
	for key, value := range mData {
		sm.mData[key] = value
	}
	sm.lock.Unlock()
}

func (sm *MapSync) Clone() map[string]interface{} {
	m := make(map[string]interface{}, sm.Len())
	sm.lock.Lock()
	for key, value := range sm.mData {
		m[key] = value
	}
	sm.lock.Unlock()
	return m
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
