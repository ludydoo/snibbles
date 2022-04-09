package bst

import (
	"golang.org/x/exp/constraints"
)

type InOrderIterator[V constraints.Ordered] struct {
	root *Node[V]
	ptr  *Node[V]
}

func newInOrderIterator[V constraints.Ordered](node *Node[V]) *InOrderIterator[V] {
	iterator := &InOrderIterator[V]{
		root: node,
	}
	return iterator
}

func (i *InOrderIterator[V]) init(position Direction) bool {
	if i.ptr == nil {
		switch position {
		case Forward:
			i.ptr = i.root.MinNode()
		case Backward:
			i.ptr = i.root.MaxNode()
		}
		return true
	}
	return false
}

func (i *InOrderIterator[V]) nextChild(n *Node[V], direction Direction) *Node[V] {
	if direction == Forward {
		return n.Right
	} else {
		return n.Left
	}
}

func (i *InOrderIterator[V]) next(direction Direction) bool {
	if i.root == nil {
		return false
	}
	if i.init(direction) {
		return true
	}
	nextChild := i.nextChild(i.ptr, direction)
	if nextChild != nil {
		if direction == Forward {
			i.ptr = nextChild.MinNode()
		} else {
			i.ptr = nextChild.MaxNode()
		}
		return true
	}
	for i.ptr.Parent != nil && i.ptr == i.nextChild(i.ptr.Parent, direction) {
		i.ptr = i.ptr.Parent
	}
	i.ptr = i.ptr.Parent
	return i.ptr != nil
}

func (i *InOrderIterator[V]) Next() bool {
	return i.next(Forward)
}

func (i *InOrderIterator[V]) Previous() bool {
	return i.next(Backward)
}

func (i *InOrderIterator[V]) All() []V {
	var result []V
	for i.Next() {
		result = append(result, i.ptr.Key)
	}
	return result
}

func (i *InOrderIterator[V]) Value() V {
	if i.ptr == nil {
		return *new(V)
	}
	return i.ptr.Key
}
