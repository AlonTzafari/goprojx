package hashset

import "iter"

type HashSet[T comparable] struct {
	m map[T]struct{}
}

func New[T comparable](items ...T) *HashSet[T] {
	hs := &HashSet[T]{
		m: make(map[T]struct{}),
	}
	for _, item := range items {
		hs.Add(item)
	}
	return hs
}

func (hs *HashSet[T]) Add(item T) {
	hs.m[item] = struct{}{}
}

func (hs *HashSet[T]) Has(item T) bool {
	_, exists := hs.m[item]
	return exists
}

func (hs *HashSet[T]) Del(item T) {
	delete(hs.m, item)
}

func (hs *HashSet[T]) Len() int {
	return len(hs.m)
}

func (hs *HashSet[T]) ToSlice() []T {
	s := make([]T, 0, len(hs.m))
	for k := range hs.m {
		s = append(s, k)
	}
	return s
}

func (hs *HashSet[T]) Iter() iter.Seq[T] {
	return func(yield func(T) bool) {
		for k := range hs.m {
			if !yield(k) {
				return
			}
		}
	}
}
