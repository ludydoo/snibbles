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
