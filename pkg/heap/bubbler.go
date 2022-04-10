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

func minBubbleUp[V constraints.Ordered](h *heap[V], i int) {
	for i > 0 {
		parent := (i - 1) / 2
		if h.elements[parent] > (h.elements[i]) {
			h.elements[parent], h.elements[i] = h.elements[i], h.elements[parent]
			i = parent
		} else {
			break
		}
	}
}

func minBubbleDown[V constraints.Ordered](h *heap[V], i int) {
	for i < h.size/2 {
		left := 2*i + 1
		right := 2*i + 2
		smallest := left
		if right < h.size && h.elements[right] < (h.elements[left]) {
			smallest = right
		}
		if h.elements[i] < (h.elements[smallest]) {
			break
		}
		h.elements[i], h.elements[smallest] = h.elements[smallest], h.elements[i]
		i = smallest
	}
}

func maxBubbleUp[V constraints.Ordered](h *heap[V], i int) {
	for i > 0 {
		parent := (i - 1) / 2
		if h.elements[parent] < h.elements[i] {
			h.elements[parent], h.elements[i] = h.elements[i], h.elements[parent]
			i = parent
		} else {
			break
		}
	}
}

func maxBubbleDown[V constraints.Ordered](h *heap[V], i int) {
	if i == h.size-1 {
		return
	}
	left := 2*i + 1
	right := 2*i + 2
	if left >= h.size {
		return
	}
	if right >= h.size {
		if h.elements[i] < h.elements[left] {
			h.elements[i], h.elements[left] = h.elements[left], h.elements[i]
			h.bubbleDown(h, left)
		}
		return
	}
	if h.elements[left] < h.elements[right] {
		if h.elements[i] < h.elements[right] {
			h.elements[i], h.elements[right] = h.elements[right], h.elements[i]
			h.bubbleDown(h, right)
		}
		return
	}
	if h.elements[i] < h.elements[left] {
		h.elements[i], h.elements[left] = h.elements[left], h.elements[i]
		h.bubbleDown(h, left)
	}
}
