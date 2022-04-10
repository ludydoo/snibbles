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
