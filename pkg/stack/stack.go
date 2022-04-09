package stack

type Stack[V any] struct {
	Items []V
}

func New[v any](items ...v) *Stack[v] {
	s := &Stack[v]{}
	for _, item := range items {
		s.Push(item)
	}
	return s
}

func (s *Stack[V]) Push(item V) {
	s.Items = append(s.Items, item)
}

func (s *Stack[V]) IsEmpty() bool {
	return len(s.Items) == 0
}

func (s *Stack[V]) Pop() (V, bool) {
	if s.IsEmpty() {
		return *new(V), false
	}
	index := len(s.Items) - 1
	item := s.Items[index]
	s.Items = s.Items[:index]
	return item, true
}

func (s *Stack[V]) Peek() V {
	if s.IsEmpty() {
		return *new(V)
	}
	index := len(s.Items) - 1
	return s.Items[index]
}

func (s *Stack[V]) Len() int {
	return len(s.Items)
}
