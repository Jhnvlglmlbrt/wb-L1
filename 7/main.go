package main

import (
	"fmt"
	"sync"
)

type ConcurrentMap[V any] struct {
	sync.RWMutex
	mapa map[string]V
}

func NewMap[V any]() *ConcurrentMap[V] {
	return &ConcurrentMap[V]{
		mapa: make(map[string]V),
	}
}

func (m *ConcurrentMap[V]) Get(key string) (V, bool) {
	m.RLock()
	defer m.RUnlock()
	v, ok := m.mapa[key]
	if !ok {
		fmt.Printf("map does not have %v key\n", key)
	}
	return v, ok
}

func (m *ConcurrentMap[V]) Set(key string, val V) {
	m.Lock()
	defer m.Unlock()
	m.mapa[key] = val
}

func main() {
	m := NewMap[int]()
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			m.Set(fmt.Sprintf("%d", i), i)
		}(i)
	}

	wg.Wait()

	for i := 0; i < 10; i++ {
		val, _ := m.Get(fmt.Sprintf("%d", i))
		fmt.Printf("map[%d] = %d\n", i, val)
	}
}
