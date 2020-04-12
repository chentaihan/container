package hashmap

import (
	"encoding/json"
	"sync"
)

type MapIntSync struct {
	mData map[string]int
	lock  sync.RWMutex
}

func NewMapIntSync() *MapIntSync {
	return &MapIntSync{
		mData: map[string]int{},
	}
}

func (sm *MapIntSync) Get(key string) (int, bool) {
	sm.lock.RLock()
	value, exist := sm.mData[key]
	sm.lock.RUnlock()
	return value, exist
}

func (sm *MapIntSync) Set(key string, value int) {
	sm.lock.Lock()
	sm.mData[key] = value
	sm.lock.Unlock()
}

func (sm *MapIntSync) Copy(mData map[string]int) {
	sm.lock.Lock()
	for key, value := range mData {
		sm.mData[key] = value
	}
	sm.lock.Unlock()
}

func (sm *MapIntSync) Clone() map[string]int {
	m := make(map[string]int, sm.Len())
	sm.lock.Lock()
	for key, value := range sm.mData {
		m[key] = value
	}
	sm.lock.Unlock()
	return m
}

func (sm *MapIntSync) Clear() {
	sm.lock.Lock()
	sm.mData = map[string]int{}
	sm.lock.Unlock()
}

func (sm *MapIntSync) Exist(key string) bool {
	sm.lock.RLock()
	_, exist := sm.mData[key]
	sm.lock.RUnlock()
	return exist
}

func (sm *MapIntSync) Marshal() ([]byte, error) {
	sm.lock.RLock()
	data, err := json.Marshal(sm.mData)
	sm.lock.RUnlock()
	return data, err
}

func (sm *MapIntSync) Unmarshal(data []byte) error {
	m := map[string]int{}
	err := json.Unmarshal(data, &m)
	sm.lock.Lock()
	sm.mData = m
	sm.lock.Unlock()
	return err
}

func (sm *MapIntSync) Len() int {
	if sm == nil {
		return 0
	}
	sm.lock.RLock()
	l := len(sm.mData)
	sm.lock.RUnlock()
	return l
}
