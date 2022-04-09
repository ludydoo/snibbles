package linkedlist

type Node[V comparable] struct {
	Value V
	Next  *Node[V]
}

type LinkedList[V comparable] struct {
	Head *Node[V]
}

func (l *LinkedList[V]) Add(v V) int {
	n := &Node[V]{Value: v}
	if l.Head == nil {
		l.Head = n
		return 0
	} else {
		i := 1
		cur := l.Head
		for cur.Next != nil {
			cur = cur.Next
			i++
		}
		cur.Next = n
		return i
	}
}

func (l *LinkedList[V]) RemoveAt(index int) {
	if index == 0 {
		l.Head = l.Head.Next
		return
	}
	cur := l.Head
	for i := 0; i < index-1; i++ {
		cur = cur.Next
	}
	cur.Next = cur.Next.Next
}

func (l *LinkedList[V]) Remove(v V) int {
	if l.Head.Value == v {
		l.Head = l.Head.Next
		return 0
	}
	cur := l.Head
	i := 0
	for cur.Next != nil {
		if cur.Next.Value == v {
			cur.Next = cur.Next.Next
			return i + 1
		}
		cur = cur.Next
		i++
	}
	return -1
}

func (l *LinkedList[V]) ElementAt(i int) *Node[V] {
	cur := l.Head
	for j := 0; j < i; j++ {
		cur = cur.Next
	}
	return cur
}

func (l *LinkedList[V]) Insert(v V, index int) {
	n := &Node[V]{Value: v}
	if index == 0 {
		n.Next = l.Head
		l.Head = n
		return
	}
	cur := l.Head
	for i := 0; i < index-1; i++ {
		cur = cur.Next
	}
	n.Next = cur.Next
	cur.Next = n
}

func (l *LinkedList[V]) IsEmpty() bool {
	return l.Head == nil
}
