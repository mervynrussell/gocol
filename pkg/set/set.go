package set

type Set[T comparable] interface {
	Add(item T)
	AddAll([]T)
	Clear()
	Contains(T) bool
	Len() int
	Remove(item T)
	All() []T
	IsEmpty() bool
	Intersection(Set[T]) Set[T]
	Union(Set[T]) Set[T]
	Difference(Set[T]) Set[T]
	Subset(Set[T]) bool
	Equals(Set[T]) bool
	Pop() (*T, bool)
}

// Set implementation based on map
type mapSet[T comparable] struct {
	data map[T]bool
}

// New Set based on mapSet
func New[T comparable]() Set[T] {
	return &mapSet[T]{data: make(map[T]bool)}
}

func NewFrom[T comparable](items []T) Set[T] {
	m := New[T]()
	m.AddAll(items)
	return m
}

func(s mapSet[T]) Subset(v Set[T]) bool {
	if v.Len() > len(s.data) {
		return false
	}

	for _, i := range v.All() {
		if !s.Contains(i) {
			return false
		}
	}
	return true
}

func(s mapSet[T]) Pop() (*T, bool) {
	if s.Len() > 0 {
		for k := range s.data {
			delete(s.data, k)
			return &k, true
		}
	}
	return nil, false
}

func(s mapSet[T]) Equals(v Set[T]) bool {
	if len(s.data) != v.Len() {
		return false
	}

	for k := range s.data {
		if !v.Contains(k) {
			return false
		}
	}
	return true
}

func(s mapSet[T]) IsEmpty() bool {
	return len(s.data) == 0
}

func(s mapSet[T]) Intersection(v Set[T]) Set[T] {
	intersection := New[T]()
	for _, i := range v.All() {
		if s.Contains(i) {
			intersection.Add(i)
		}
	}
	return intersection
}

func(s mapSet[T]) Difference(v Set[T]) Set[T] {
	difference := New[T]()
	for _, i := range s.All() {
		if !v.Contains(i) {
			difference.Add(i)
		}
	}
	return difference
}

func(s mapSet[T]) Union(v Set[T]) Set[T] {
	union := New[T]()
	union.AddAll(append(s.All(), v.All()...))
	return union
}

func(s mapSet[T]) Add(v T) {
	if _, ok := s.data[v]; !ok {
		s.data[v] = true
	}
}

func(s mapSet[T]) AddAll(items []T) {
	for _, item := range items {
		s.Add(item)
	}
}

func(s mapSet[T]) Clear() {
	clear(s.data)
}

func(s mapSet[T]) Contains(v T) bool {
	_, ok := s.data[v]
	return ok
}

func(s mapSet[T]) Len() int {
	return len(s.data)
}

func(s mapSet[T]) Remove(v T) {
	delete(s.data, v)
}

func(s mapSet[T])All() []T {
	r := make([]T, 0, len(s.data))
    for k := range s.data {
        r = append(r, k)
    }
    return r
}
