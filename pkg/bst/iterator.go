/*
 * MIT License
 *
 * Copyright (c) 2022 Ludovic Cleroux
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

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
