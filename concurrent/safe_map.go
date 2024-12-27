package concurrent

import "sync"

// SafeMap 可以在并发环境下安全地使用 map
type SafeMap[K comparable, V any] struct {
	data  map[K]V
	mutex sync.RWMutex
}

// Put 向 map 安全地添加 key-value
func (s *SafeMap[K, V]) Put(key K, val V) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.data[key] = val
}

// Get 从 map 中根据 key 获取 value
func (s *SafeMap[K, V]) Get(key K) (any, bool) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	res, ok := s.data[key]
	return res, ok
}

// LoadOrStore 使用 double-check 方式，获取或者装入一个 key-value
func (s *SafeMap[K, V]) LoadOrStore(key K, newVal V) (val V, loaded bool) {
	s.mutex.RLock() //  先加读锁检查一遍
	res, ok := s.data[key]
	s.mutex.RUnlock() // 释放读锁
	if ok {
		return res, true
	}
	//	可以在这加 time.Sleep(time.Second) 进行调试，或者在 s.mutex.RUnlock() 打断点，
	//	让多个 goroutine 同时执行到 s.mutex.RUnlock()，没有 double-check 就会发生数据混乱

	s.mutex.Lock() // 加写锁
	defer s.mutex.Unlock()
	res, ok = s.data[key] // 再检查一遍
	if ok {
		return res, true
	}
	s.data[key] = newVal
	return newVal, false
}
