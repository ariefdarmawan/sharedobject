package sharedobject

import (
	"sync"
)

// SharedData is shared data
type SharedData struct {
	lock sync.RWMutex
	data map[string]interface{}
}

// NewSharedData initiate new SharedData
func NewSharedData() *SharedData {
	d := &SharedData{}
	d.lock = new(sync.RWMutex)
	d.data = map[string]interface{}{}
	return d
}

// Get data with given key and default value if not exist
func (s *SharedData) Get(key string, def interface{}) interface{} {
	var out interface{}
	var b bool
	hasData := false

	s.lock.RLock()
	if s.data != nil {
		hasData = true
		out, b = s.data[key]
	}
	s.lock.RUnlock()

	if hasData {
		if b {
			return out
		}

		return def
	}

	return def
}

// Set data with given key and value
func (s *SharedData) Set(key string, value interface{}) {
	s.lock.Lock()
	if s.data == nil {
		s.data = map[string]interface{}{}
	}
	s.data[key] = value
	s.lock.Unlock()
}

// Remove data with given key
func (s *SharedData) Remove(key string) {
	s.lock.Lock()
	delete(s.data, key)
	s.lock.Unlock()
}

// Count data
func (s *SharedData) Count() int {
	out := 0
	s.lock.RLock()
	if s.data != nil {
		out = len(s.data)
	}
	s.lock.RUnlock()
	return out
}
