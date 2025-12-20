package hashset

type HashSet[T comparable] struct {
	m map[T]struct{}
}

func New[T comparable]() *HashSet[T] {
	return &HashSet[T]{
		m: make(map[T]struct{}),
	}
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
