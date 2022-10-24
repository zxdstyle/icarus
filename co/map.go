package co

import (
	"errors"
	json "github.com/json-iterator/go"
	"sync"
)

var (
	DuplicateKey = errors.New("key already exists")
)

type Map[K comparable, V any] struct {
	lock sync.RWMutex
	data map[K]V
}

func NewMap[K comparable, V any]() *Map[K, V] {
	return &Map[K, V]{
		lock: sync.RWMutex{},
		data: make(map[K]V),
	}
}

// Keys creates an array of the map keys.
func (m *Map[K, V]) Keys() []K {
	m.lock.RLock()
	defer m.lock.RUnlock()

	result := make([]K, 0, len(m.data))
	for k := range m.data {
		result = append(result, k)
	}
	return result
}

// Values creates an array of the map values.
func (m *Map[K, V]) Values() []V {
	m.lock.RLock()
	defer m.lock.RUnlock()

	result := make([]V, 0, len(m.data))
	for _, v := range m.data {
		result = append(result, v)
	}
	return result
}

// Add sets key-value to the hash map. returns an error if the key already exists
func (m *Map[K, V]) Add(key K, value V) error {
	m.lock.Lock()
	defer m.lock.Unlock()

	if m.data == nil {
		m.data = make(map[K]V)
	}
	if _, exists := m.data[key]; exists {
		return DuplicateKey
	}
	m.data[key] = value
	return nil
}

// Set sets key-value to the hash map.
func (m *Map[K, V]) Set(key K, value V) {
	m.lock.Lock()
	defer m.lock.Unlock()

	if m.data == nil {
		m.data = make(map[K]V)
	}
	m.data[key] = value
}

// Get returns the value by given `key`.
func (m *Map[K, V]) Get(key K) (value V) {
	m.lock.RLock()
	defer m.lock.RUnlock()

	if m.data != nil {
		value, _ = m.data[key]
	}
	return
}

// Gets returns all values by given `key`.
func (m *Map[K, V]) Gets(keys ...K) (values []V) {
	m.lock.RLock()
	defer m.lock.RUnlock()

	if m.data != nil {
		for _, key := range keys {
			value, ok := m.data[key]
			if ok {
				values = append(values, value)
			}
		}

	}
	return
}

// Size returns the size of the map.
func (m *Map[K, V]) Size() int {
	m.lock.RLock()
	defer m.lock.RUnlock()

	return len(m.data)
}

// IsEmpty checks whether the map is empty. It returns true if map is empty, or else false.
func (m *Map[K, V]) IsEmpty() bool {
	return m.Size() == 0
}

// Iterator iterates the hash map readonly with custom callback function `f`.  If `f` returns true, then it continues iterating; or false to stop.
func (m *Map[K, V]) Iterator(f func(key K, value V) bool) {
	m.lock.RLock()
	defer m.lock.RUnlock()
	for k, v := range m.data {
		if !f(k, v) {
			break
		}
	}
}

// Sets batch sets key-values to the hash map.
func (m *Map[K, V]) Sets(data map[K]V) {
	m.lock.Lock()
	defer m.lock.Unlock()

	if m.data == nil {
		m.data = data
	} else {
		for k, v := range data {
			m.data[k] = v
		}
	}
}

// Search searches the map with given `key`. Second return parameter `found` is true if key was found, otherwise false.
func (m *Map[K, V]) Search(key K) (value V, found bool) {
	m.lock.RLock()
	defer m.lock.RUnlock()

	if m.data != nil {
		value, found = m.data[key]
	}
	return
}

// GetOrSet returns the value by key, or sets value with given `value` if it does not exist and then returns this value.
func (m *Map[K, V]) GetOrSet(key K, value V) interface{} {
	if v, ok := m.Search(key); !ok {
		return m.doSetWithLockCheck(key, value)
	} else {
		return v
	}
}

// GetOrSetFuncLock returns the value by key,
// or sets value with returned value of callback function `f` if it does not exist
// and then returns this value.
//
// GetOrSetFuncLock differs with GetOrSetFunc function is that it executes function `f`
// with mutex.Lock of the hash map.
func (m *Map[K, V]) GetOrSetFuncLock(key K, f func() V) V {
	if v, ok := m.Search(key); !ok {
		return m.doSetWithLockCheck(key, f)
	} else {
		return v
	}
}

// Remove deletes value from map by given `key`, and return this deleted value.
func (m *Map[K, V]) Remove(key K) (value V) {
	m.lock.Lock()
	defer m.lock.Unlock()

	if m.data != nil {
		var ok bool
		if value, ok = m.data[key]; ok {
			delete(m.data, key)
		}
	}
	return
}

// Removes batch deletes values of the map by keys.
func (m *Map[K, V]) Removes(keys ...K) {
	m.lock.Lock()
	defer m.lock.Unlock()

	if m.data != nil {
		for _, key := range keys {
			delete(m.data, key)
		}
	}
}

// Contains checks whether a key exists. It returns true if the `key` exists, or else false.
func (m *Map[K, V]) Contains(key K) (ok bool) {
	m.lock.RLock()
	defer m.lock.RUnlock()

	if m.data != nil {
		_, ok = m.data[key]
	}
	return
}

// Clear deletes all data of the map, it will remake a new underlying data map.
func (m *Map[K, V]) Clear() {
	m.lock.Lock()
	defer m.lock.Unlock()

	m.data = make(map[K]V)
}

// Merge merges two hash maps.
// The `other` map will be merged into the map `m`.
func (m *Map[K, V]) Merge(others ...*Map[K, V]) {
	m.lock.Lock()
	defer m.lock.Unlock()

	if m.data == nil {
		m.data = make(map[K]V)
	}
	for _, other := range others {
		other.Iterator(func(key K, val V) bool {
			m.data[key] = val
			return true
		})
	}
}

// MapCopy returns a copy of the underlying data of the hash map.
func (m *Map[K, V]) MapCopy() map[K]V {
	m.lock.RLock()
	defer m.lock.RUnlock()

	data := make(map[K]V, len(m.data))
	for k, v := range m.data {
		data[k] = v
	}
	return data
}

// doSetWithLockCheck checks whether value of the key exists with mutex.Lock,
// if not exists, set value to the map with given `key`,
// or else just return the existing value.
//
// When setting value, if `value` is type of `func() interface {}`,
// it will be executed with mutex.Lock of the hash map,
// and its return value will be set to the map with `key`.
//
// It returns value with given `key`.
func (m *Map[K, V]) doSetWithLockCheck(key K, value interface{}) V {
	m.lock.Lock()
	defer m.lock.Unlock()

	if m.data == nil {
		m.data = make(map[K]V)
	}
	if v, ok := m.data[key]; ok {
		return v
	}
	if f, ok := value.(func() V); ok {
		value = f()
	}
	val, ok := value.(V)
	if ok {
		m.data[key] = val
		return val
	}
	return val
}

// MarshalJSON implements the interface MarshalJSON for json.Marshal.
func (m *Map[K, V]) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.data)
}

// UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.
func (m *Map[K, V]) UnmarshalJSON(b []byte) error {
	m.lock.Lock()
	defer m.lock.Unlock()

	if m.data == nil {
		m.data = make(map[K]V)
	}
	var data map[K]V
	if err := json.Unmarshal(b, &data); err != nil {
		return err
	}
	for k, v := range data {
		m.data[k] = v
	}
	return nil
}
