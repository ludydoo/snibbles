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
