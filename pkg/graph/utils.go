package graph

import "container/list"

func mapKeys[T comparable, V any](m map[T]V) []T {
	r := make([]T, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}

func elementFromList[T comparable](l list.List, v T) *list.Element {
	for e := l.Front(); e != nil; e = e.Next() {
        if v == e.Value {
			return e
		}
    }
	return nil
}
