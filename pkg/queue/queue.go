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
