package queue

type Node[V any] struct {
	Value V
	Next  *Node[V]
}

type Queue[V any] struct {
	head *Node[V]
	tail *Node[V]
	size int
}

func (q *Queue[V]) Enqueue(v V) {
	n := &Node[V]{Value: v}
	if q.tail == nil {
		q.head = n
	} else {
		q.tail.Next = n
	}
	q.tail = n
	q.size++
}

func (q *Queue[V]) Dequeue() V {
	if q.head == nil {
		return *new(V)
	}
	v := q.head.Value
	q.head = q.head.Next
	if q.head == nil {
		q.tail = nil
	}
	q.size--
	return v
}

func (q *Queue[V]) Len() int {
	return q.size
}

func (q *Queue[V]) IsEmpty() bool {
	return q.Len() == 0
}
