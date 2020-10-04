package common

import (
	"sync"
)

// ConcurrentMap type.
type ConcurrentMap []*ConcurrentMapShared

// Share count.
// Default value is 64.
var share int = 64

// ConcurrentMapShared structure.
// 'items' is the fregment map.
// 'mutex' is the read-write mutex for this fregment map.
type ConcurrentMapShared struct {
	items map[string]interface{}
	mutex sync.RWMutex
}

// NewConcurrentMap function.
// Make a new cocurrent map with default share count.
func NewConcurrentMap() *ConcurrentMap {
	return NewConcurrentMapSpecific(share)
}

// NewConcurrentMapSpecific function.
// Make a new cocurrent map with share count.
func NewConcurrentMapSpecific(count int) *ConcurrentMap {
	share = count
	m := make(ConcurrentMap, share)
	for i := 0; i < share; i++ {
		m[i] = &ConcurrentMapShared{
			items: map[string]interface{}{},
		}
	}
	return &m
}

func (m ConcurrentMap) getSharedMap(key string) *ConcurrentMapShared {
	return m[uint(hash(key))%uint(share)]
}

func hash(key string) uint32 {
	hash := uint32(2166136261)
	prime32 := uint32(16777619)
	for i := 0; i < len(key); i++ {
		hash *= prime32
		hash ^= uint32(key[i])
	}
	return hash
}

// Put function.
// Put the key and value to this concurrent map.
func (m ConcurrentMap) Put(key string, value interface{}) {
	sharedMap := m.getSharedMap(key)
	sharedMap.mutex.Lock()
	sharedMap.items[key] = value
	sharedMap.mutex.Unlock()
}

// Get function.
// Get the value with key from this concurrent map.
func (m ConcurrentMap) Get(key string) (value interface{}, ok bool) {
	sharedMap := m.getSharedMap(key)
	sharedMap.mutex.RLock()
	value, ok = sharedMap.items[key]
	sharedMap.mutex.RUnlock()
	return value, ok
}

// Remove function.
// Remove the key and value from this concurrent map.
func (m ConcurrentMap) Remove(key string) {
	sharedMap := m.getSharedMap(key)
	sharedMap.mutex.RLock()
	delete(sharedMap.items, key)
	sharedMap.mutex.RUnlock()
}

// Count function.
// Get the item count of this concurrent map.
func (m ConcurrentMap) Count() int {
	count := 0
	for i := 0; i < share; i++ {
		m[i].mutex.RLock()
		count += len(m[i].items)
		m[i].mutex.RUnlock()
	}
	return count
}

// Keys function.
// Get all the keys of this concurrent map.
func (m ConcurrentMap) Keys() []string {
	count := m.Count()
	keys := make([]string, count)

	ch := make(chan string, count)
	go func() {
		wg := sync.WaitGroup{}
		wg.Add(share)
		for i := 0; i < share; i++ {
			go func(ms *ConcurrentMapShared) {
				defer wg.Done()

				ms.mutex.RLock()
				for k := range ms.items {
					ch <- k
				}
				ms.mutex.RUnlock()
			}(m[i])
		}

		wg.Wait()
		close(ch)
	}()

	for k := range ch {
		keys = append(keys, k)
	}
	return keys
}
