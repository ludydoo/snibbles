package heap

import "golang.org/x/exp/constraints"

type ErrHeapEmptyType struct {
	msg string
}

func (e ErrHeapEmptyType) Error() string {
	return e.msg
}

var ErrHeapEmpty = &ErrHeapEmptyType{msg: "heap empty"}

type Heap[V constraints.Ordered] interface {
	Push(v V) bool
	Pop() (V, error)
	Peek() (V, error)
	Len() int
	Empty() bool
	All() []V
	Remove(v V) bool
}

var _ Heap[int] = &heap[int]{}

func Max[V constraints.Ordered](capacity int) Heap[V] {
	h := &heap[V]{
		bubbleUp:   maxBubbleUp[V],
		bubbleDown: maxBubbleDown[V],
	}
	h.init(capacity)
	return h
}

func Min[V constraints.Ordered](capacity int) Heap[V] {
	h := &heap[V]{
		bubbleUp:   minBubbleUp[V],
		bubbleDown: minBubbleDown[V],
	}
	h.init(capacity)
	return h
}
