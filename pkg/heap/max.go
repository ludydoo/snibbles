package heap

import (
	"errors"
	"golang.org/x/exp/constraints"
)

type heap[V constraints.Ordered] struct {
	elements   []V
	capacity   int
	size       int
	bubbleUp   func(heap *heap[V], i int)
	bubbleDown func(heap *heap[V], i int)
}

func (h heap[V]) All() []V {
	return h.elements
}

func (h *heap[V]) init(capacity int) {
	h.elements = make([]V, capacity)
	h.capacity = capacity
	h.size = 0
}

func (h *heap[V]) Push(v V) bool {
	if h.size == h.capacity {
		return false
	}
	h.elements[h.size] = v
	h.size++
	h.bubbleUp(h, h.size-1)
	return true
}

func (h *heap[V]) Pop() (V, error) {
	if h.size == 0 {
		return *new(V), errors.New("heap is empty")
	}
	v := h.elements[0]
	h.elements[0] = h.elements[h.size-1]
	h.elements[h.size-1] = *new(V)
	h.size--
	h.bubbleDown(h, 0)
	return v, nil
}

func (h *heap[V]) Min() (V, error) {
	if h.size == 0 {
		return *new(V), ErrHeapEmpty
	}
	return h.elements[0], nil
}

func (h *heap[V]) Max() (V, error) {
	if h.size == 0 {
		return *new(V), ErrHeapEmpty
	}
	return h.elements[h.size-1], nil
}

func (h *heap[V]) Empty() bool {
	return h.size == 0
}

func (h *heap[V]) Len() int {
	return h.size
}

func (h *heap[V]) Peek() (V, error) {
	if h.size == 0 {
		return *new(V), ErrHeapEmpty
	}
	return h.elements[0], nil
}

func (h *heap[V]) Remove(v V) bool {
	if h.size == 0 {
		return false
	}
	for i := 0; i < h.size; i++ {
		if h.elements[i] == v {
			h.elements[i] = h.elements[h.size-1]
			h.elements[h.size-1] = *new(V)
			h.size--
			h.bubbleDown(h, i)
			return true
		}
	}
	return false
}
