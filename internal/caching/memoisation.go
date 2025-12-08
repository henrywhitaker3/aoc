// Package caching
package caching

type Memoised[T any] struct {
	cache map[string]T
}

func Memoise[T any]() *Memoised[T] {
	return &Memoised[T]{
		cache: map[string]T{},
	}
}

func (m *Memoised[T]) Run(key string, f func() T) T {
	if val, ok := m.cache[key]; ok {
		return val
	}
	val := f()
	m.cache[key] = val
	return val
}

func (m *Memoised[T]) Clear(key string) {
	delete(m.cache, key)
}

func (m *Memoised[T]) Replace(key string, v T) {
	m.cache[key] = v
}
